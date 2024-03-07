package spice

import (
	"context"
	"log"

	"github.com/authzed/spicedb/pkg/cmd/datastore"
	"github.com/authzed/spicedb/pkg/cmd/server"
	"github.com/authzed/spicedb/pkg/cmd/util"
	"golang.org/x/exp/slices"
	"golang.org/x/xerrors"

	"cdr.dev/slog"
	"github.com/coder/coder/v2/coderd/database"
)

type SpiceServerOpts struct {
	// If 'PostgresURI' is empty, will default to in memory
	PostgresURI string
	Logger      slog.Logger
	// Store is the application database store.
	Store database.Store
}

// TODO: Handle PG vs Memory
func New(ctx context.Context, opts *SpiceServerOpts) (database.Store, error) {
	if opts.Store == nil {
		return nil, xerrors.Errorf("store is required")
	}

	// Do not double wrap
	if slices.Contains(opts.Store.Wrappers(), wrapname) {
		return opts.Store, nil
	}

	engine := datastore.MemoryEngine
	if opts.PostgresURI != "" {
		engine = datastore.PostgresEngine
	}

	defaults := datastore.DefaultDatastoreConfig()
	// Default prom metrics get registered to the default global registry.
	// This is unfortunate, just going to disable for now.
	defaults.EnableDatastoreMetrics = false
	// TODO: Look into this, we might be able to make this migrate itself without the cli.
	defaults.MigrationPhase = ""
	configOptions := []datastore.ConfigOption{
		defaults.ToOption(),
		datastore.WithEngine(engine),
		datastore.WithRequestHedgingEnabled(false),
	}

	if engine == datastore.PostgresEngine {
		// must run migrations first
		// To get cli: go install github.com/authzed/spicedb/cmd/spicedb@latest
		// 	spicedb migrate --skip-release-check --datastore-engine=postgres --datastore-conn-uri "postgres://postgres:postgres@localhost:5432/spicedb?sslmode=disable" head
		// Example URI:
		//	postgres://postgres:postgres@localhost:5432/spicedb?sslmode=disable`
		configOptions = append(configOptions, datastore.WithURI(opts.PostgresURI))
	}

	ds, err := datastore.NewDatastore(ctx,
		configOptions...,
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
			// Since this is embedded, we can assume no external actors can auth this.
			// TODO: Ensure the opened ports cannot be used?
			return ctx, nil
		}),
		server.WithHTTPGateway(util.HTTPServerConfig{
			HTTPAddress: "localhost:50001",
			HTTPEnabled: false}),
		//server.WithDashboardAPI(util.HTTPServerConfig{HTTPEnabled: false}),
		server.WithMetricsAPI(util.HTTPServerConfig{
			HTTPAddress: "localhost:50000",
			HTTPEnabled: true}),
		server.WithDispatchCacheConfig(server.CacheConfig{
			Name:        "DispatchCache",
			MaxCost:     "70%",
			NumCounters: 100_000,
			Metrics:     true,
			Enabled:     true,
		}),
		server.WithNamespaceCacheConfig(server.CacheConfig{
			Name:        "NamespaceCache",
			MaxCost:     "32MiB",
			NumCounters: 1_000,
			Metrics:     true,
			Enabled:     true,
		}),
		server.WithClusterDispatchCacheConfig(server.CacheConfig{
			Name:        "ClusterCache",
			MaxCost:     "70%",
			NumCounters: 100_000,
			Metrics:     true,
			Enabled:     true,
		}),
		server.WithDatastore(ds),
		server.WithDispatchClientMetricsPrefix("coder_client"),
		server.WithDispatchClientMetricsEnabled(true),
		server.WithDispatchClusterMetricsPrefix("cluster"),
		server.WithDispatchClusterMetricsEnabled(true),
	}

	runnable, err := server.NewConfigWithOptionsAndDefaults(configOpts...).Complete(ctx)
	if err != nil {
		return nil, xerrors.Errorf("spicedb config failed: %w", err)
	}

	return &SpiceDB{
		Store:  opts.Store,
		logger: opts.Logger,
		srv:    runnable,
	}, nil
}
