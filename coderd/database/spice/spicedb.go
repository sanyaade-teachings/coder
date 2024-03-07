package spice

import (
	"context"

	"github.com/authzed/authzed-go/pkg/responsemeta"
	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/spicedb/pkg/cmd/server"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"tailscale.com/syncs"

	"cdr.dev/slog"
	"github.com/coder/coder/v2/coderd/database"
	"github.com/coder/coder/v2/coderd/database/spice/policy"
)

const wrapname = "spicedb.querier"

type SpiceDB struct {
	// TODO: Do not embed anonymously. This is just a lazy way to skip
	// 		having to implement all interface methods for now.
	database.Store
	logger slog.Logger

	srv       server.RunnableServer
	schemaCli v1.SchemaServiceClient
	permCli   v1.PermissionsServiceClient
	// experimental client has bulk operations.
	expCli v1.ExperimentalServiceClient

	// zedToken is required to enforce consistency. When making a request, passing
	// this token says "I want to see the world as if it was at least after this time".
	// For a 100% consistent view, we should update this on any write.
	// In a world of HA, we have an issue that a different Coder might have done
	// a write.
	// TODO: A way of doing this across HA is storing the Zedtoken in the DB on
	//		each resource. So if you fetch a workspace, it has the Zedtoken required
	//		to get it's updated state.
	zedToken syncs.AtomicValue[*v1.ZedToken]

	ctx    context.Context
	cancel context.CancelFunc

	// debug will print extra debug information on all checks.
	debug bool
}

func (s *SpiceDB) Wrappers() []string {
	return append(s.Store.Wrappers(), wrapname)
}

func (s *SpiceDB) Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	s.ctx = ctx
	s.cancel = cancel

	// Start the server
	go func() {
		if err := s.srv.Run(ctx); err != nil {
			s.logger.Error(ctx, "spicedb run server", slog.Error(err))
		}
	}()

	// Setup the clients
	conn, err := s.srv.GRPCDialContext(ctx)
	if err != nil {
		return xerrors.Errorf("spicedb grpc dial failed: %w", err)
	}

	s.expCli = v1.NewExperimentalServiceClient(conn)
	s.schemaCli = v1.NewSchemaServiceClient(conn)
	s.permCli = v1.NewPermissionsServiceClient(conn)

	// TODO: If the server isn't running yet because it's async, will this fail?
	resp, err := s.schemaCli.WriteSchema(ctx, &v1.WriteSchemaRequest{
		Schema: policy.Schema,
	})
	if err != nil {
		return xerrors.Errorf("write schema: %w", err)
	}
	s.zedToken.Store(resp.WrittenAt)

	return nil
}

func (s *SpiceDB) Debugging(set bool) {
	s.debug = set
}

func (s *SpiceDB) Close() {
	s.cancel()
}

// WriteRelationships returns a revert function that will delete all the relationships that
// were written.
func (s *SpiceDB) WriteRelationships(ctx context.Context, relationships ...v1.Relationship) (revert func() error, _ error) {
	opts := []grpc.CallOption{}
	if s.debug {
		opt, callback := debugSpiceDBRPC(ctx, s.logger)
		opts = append(opts, opt)
		defer callback()
	}

	updates := make([]*v1.RelationshipUpdate, 0, len(relationships))
	for i := range relationships {
		// Make a copy so to ensure the delete function has the correct data.
		// We could definitely improve the memory allocations here.
		cpy := relationships[i]
		updates = append(updates, &v1.RelationshipUpdate{
			Operation:    v1.RelationshipUpdate_OPERATION_TOUCH,
			Relationship: &cpy,
		})
	}

	// A relationship can be written like this:
	//	group:hr#member@user:camilla
	// And parsed with:
	// 	tup := tuple.Parse(rel)
	// 	v1Rel := tuple.ToRelationship(tup)
	resp, err := s.permCli.WriteRelationships(ctx, &v1.WriteRelationshipsRequest{
		Updates:               updates,
		OptionalPreconditions: nil,
	}, opts...)
	if err != nil {
		return nil, xerrors.Errorf("write relationship: %w", err)
	}
	// TODO: We should probably return this? Allow it to be stored on the object or something?
	s.zedToken.Store(resp.WrittenAt)

	// revert is an optional callback the caller can use to delete the relationship
	// if it's no longer needed. This is helpful if their tx fails.
	revert = func() error {
		for i := range updates {
			updates[i].Operation = v1.RelationshipUpdate_OPERATION_DELETE
		}

		// The delete api might be quicker, but this an atomic operation.
		resp, err := s.permCli.WriteRelationships(ctx, &v1.WriteRelationshipsRequest{
			Updates:               updates,
			OptionalPreconditions: nil,
		}, opts...)
		if err != nil {
			return xerrors.Errorf("revert relationships: %w", err)
		}
		s.zedToken.Store(resp.WrittenAt)

		return nil
	}
	return revert, nil
}

// TODO: Params to this function?
func (s *SpiceDB) Check(ctx context.Context) (bool, error) {
	opts := []grpc.CallOption{}
	if s.debug {
		opt, callback := debugSpiceDBRPC(ctx, s.logger)
		opts = append(opts, opt)
		defer callback()
	}

	// A permission can be written like:
	//	"workspace:dogfood#view@user:root"
	// And parsed with:
	//	tup := tuple.Parse(perm)
	//	r := tuple.ToRelationship(tup)
	resp, err := s.permCli.CheckPermission(ctx, &v1.CheckPermissionRequest{
		Consistency: &v1.Consistency{Requirement: &v1.Consistency_AtLeastAsFresh{s.zedToken.Load()}},
		Resource:    nil,
		Permission:  "",
		Subject:     nil,
		// Context for caveats
		Context: nil,
	}, opts...)
	if err != nil {
		return false, xerrors.Errorf("check permission: %w", err)
	}

	return resp.Permissionship == v1.CheckPermissionResponse_PERMISSIONSHIP_HAS_PERMISSION, nil
}

func debugSpiceDBRPC(ctx context.Context, logger slog.Logger) (opt grpc.CallOption, debugString func()) {
	var trailerMD metadata.MD
	debugString = func() {
		if trailerMD.Len() == 0 {
			return
		}
		// All this debug stuff just shows the trace of the check
		// with information like cache hits.
		found, err := responsemeta.GetResponseTrailerMetadata(trailerMD, responsemeta.DebugInformation)
		if err != nil {
			logger.Debug(ctx, "debug rpc failed: unable to get response metadata", slog.Error(err))
			return
		}

		debugInfo := &v1.DebugInformation{}
		err = protojson.Unmarshal([]byte(found), debugInfo)
		if err != nil {
			logger.Debug(ctx, "debug rpc failed: unable to debug proto", slog.Error(err))
			return
		}

		if debugInfo.Check == nil {
			logger.Debug(ctx, "debug rpc: no trace found for the check")
			return
		}
		tp := NewTreePrinter()
		DisplayCheckTrace(debugInfo.Check, tp, false)
		logger.Debug(ctx, tp.String())
	}

	return grpc.Trailer(&trailerMD), debugString
}
