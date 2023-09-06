package spice

import (
	"context"
	_ "embed"
	"log"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/spicedb/pkg/cmd/datastore"
	"github.com/authzed/spicedb/pkg/cmd/server"
	"github.com/authzed/spicedb/pkg/cmd/util"
)

var _ = v1.NewSchemaServiceClient

//go:embed schema.zed
var schema string

func DB(ctx context.Context) error {
	srv, err := newServer(ctx)
	if err != nil {
		return err
	}

	conn, err := srv.GRPCDialContext(ctx)
	if err != nil {
		return err
	}

	schemaSrv := v1.NewSchemaServiceClient(conn)
	permSrv := v1.NewPermissionsServiceClient(conn)
	go func() {
		if err := srv.Run(ctx); err != nil {
			log.Print("error while shutting down server: %w", err)
		}
	}()

	_, err = schemaSrv.WriteSchema(ctx, &v1.WriteSchemaRequest{
		Schema: schema,
	})
	if err != nil {
		return err
	}

	resp, err := permSrv.WriteRelationships(ctx, &v1.WriteRelationshipsRequest{})
	if err != nil {
		return err
	}

	token := resp.GetWrittenAt()
	// This will not work yet!
	checkResp, err := permSrv.CheckPermission(ctx, &v1.CheckPermissionRequest{
		Permission:  "view",
		Consistency: &v1.Consistency{Requirement: &v1.Consistency_AtLeastAsFresh{AtLeastAsFresh: token}},
		Resource: &v1.ObjectReference{
			ObjectId:   "my_book",
			ObjectType: "resource",
		},
		Subject: &v1.SubjectReference{
			Object: &v1.ObjectReference{
				ObjectId:   "john_doe",
				ObjectType: "user",
			},
		},
	})
	if err != nil {
		log.Fatal("unable to issue PermissionCheck: %w", err)
	}

	log.Printf("check result: %s", checkResp.Permissionship.String())

	return nil
}

func newServer(ctx context.Context) (server.RunnableServer, error) {
	ds, err := datastore.NewDatastore(ctx,
		datastore.DefaultDatastoreConfig().ToOption(),
		datastore.WithEngine(datastore.PostgresEngine),
		datastore.WithRequestHedgingEnabled(false),
		// must run migrations first
		// spicedb migrate --datastore-engine=postgres --datastore-conn-uri "postgres://postgres:postgres@localhost:5432/spicedb?sslmode=disable" head
		datastore.WithURI(`postgres://postgres:postgres@localhost:5432/spicedb?sslmode=disable`),
	)
	if err != nil {
		log.Fatalf("unable to start postgres datastore: %s", err)
	}

	configOpts := []server.ConfigOption{
		server.WithGRPCServer(util.GRPCServerConfig{
			Network: util.BufferedNetwork,
			Enabled: true,
		}),
		server.WithGRPCAuthFunc(func(ctx context.Context) (context.Context, error) {
			return ctx, nil
		}),
		server.WithHTTPGateway(util.HTTPServerConfig{
			HTTPAddress: "localhost:50001",
			HTTPEnabled: false}),
		//server.WithDashboardAPI(util.HTTPServerConfig{HTTPEnabled: false}),
		server.WithMetricsAPI(util.HTTPServerConfig{
			HTTPAddress: "localhost:50000",
			HTTPEnabled: true}),
		// disable caching since it's all in memory
		server.WithDispatchCacheConfig(server.CacheConfig{Enabled: false, Metrics: false}),
		server.WithNamespaceCacheConfig(server.CacheConfig{Enabled: false, Metrics: false}),
		server.WithClusterDispatchCacheConfig(server.CacheConfig{Enabled: false, Metrics: false}),
		server.WithDatastore(ds),
	}

	return server.NewConfigWithOptionsAndDefaults(configOpts...).Complete(ctx)
}
