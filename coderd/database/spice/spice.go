package spice

import (
	"context"
	_ "embed"
	"log"

	"google.golang.org/protobuf/encoding/protojson"

	"golang.org/x/xerrors"

	"google.golang.org/grpc/metadata"

	"google.golang.org/grpc"

	"github.com/authzed/authzed-go/pkg/responsemeta"

	"github.com/authzed/authzed-go/pkg/requestmeta"

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
		Schema: `definition user {}
                 definition resource {
                   relation viewer: user
                   permission view = viewer
                 }`,
	})
	if err != nil {
		return err
	}

	resp, err := permSrv.WriteRelationships(ctx, &v1.WriteRelationshipsRequest{Updates: []*v1.RelationshipUpdate{
		{
			Operation: v1.RelationshipUpdate_OPERATION_TOUCH,
			Relationship: &v1.Relationship{
				Resource: &v1.ObjectReference{
					ObjectId:   "my_book",
					ObjectType: "resource",
				},
				Relation: "viewer",
				Subject: &v1.SubjectReference{
					Object: &v1.ObjectReference{
						ObjectId:   "john_doe",
						ObjectType: "user",
					},
				},
			},
		},
	}})
	if err != nil {
		return err
	}

	token := resp.GetWrittenAt()

	for i := 0; i < 10; i++ {
		var trailerMD metadata.MD
		ctx = requestmeta.AddRequestHeaders(ctx, requestmeta.RequestDebugInformation)
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
		}, grpc.Trailer(&trailerMD))
		if err != nil {
			log.Fatal("unable to issue PermissionCheck: %w", err)
		} else {

			// All this debug stuff just shows the trace of the check
			// with information like cache hits.
			found, err := responsemeta.GetResponseTrailerMetadata(trailerMD, responsemeta.DebugInformation)
			if err != nil {
				return xerrors.Errorf("unable to get response metadata: %w", err)
			}

			debugInfo := &v1.DebugInformation{}
			err = protojson.Unmarshal([]byte(found), debugInfo)
			if err != nil {
				return err
			}

			if debugInfo.Check == nil {
				log.Println("No trace found for the check")
			} else {
				tp := NewTreePrinter()
				DisplayCheckTrace(debugInfo.Check, tp, false)
				tp.Print()
			}

		}
		log.Printf("check result: %s", checkResp.Permissionship.String())
	}

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
	ds = &DatastoreWrapper{
		Datastore: ds,
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
		server.WithDispatchCacheConfig(server.CacheConfig{Enabled: true, Metrics: true}),
		server.WithNamespaceCacheConfig(server.CacheConfig{Enabled: true, Metrics: true}),
		server.WithClusterDispatchCacheConfig(server.CacheConfig{Enabled: true, Metrics: true}),
		server.WithDatastore(ds),
		server.WithDispatchClientMetricsPrefix("coder_client"),
		server.WithDispatchClientMetricsEnabled(true),
		server.WithDispatchClusterMetricsPrefix("cluster"),
		server.WithDispatchClusterMetricsEnabled(true),
	}

	return server.NewConfigWithOptionsAndDefaults(configOpts...).Complete(ctx)
}
