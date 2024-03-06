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

// TODO: Params to this function?
func (s *SpiceDB) WriteRelationship(ctx context.Context) (delete func() error, _ error) {
	opts := []grpc.CallOption{}
	if s.debug {
		opt, callback := debugSpiceDBRPC(ctx, s.logger)
		opts = append(opts, opt)
		defer callback()
	}

	// A relationship can be written like this:
	//	group:hr#member@user:camilla
	// And parsed with:
	// 	tup := tuple.Parse(rel)
	// 	v1Rel := tuple.ToRelationship(tup)
	resp, err := s.permCli.WriteRelationships(ctx, &v1.WriteRelationshipsRequest{
		Updates: []*v1.RelationshipUpdate{
			{
				Operation: v1.RelationshipUpdate_OPERATION_TOUCH,
				Relationship: &v1.Relationship{
					Resource:       nil,
					Relation:       "",
					Subject:        nil,
					OptionalCaveat: nil,
				},
			},
		},
		OptionalPreconditions: nil,
	}, opts...)
	if err != nil {
		return nil, xerrors.Errorf("write relationship: %w", err)
	}
	s.zedToken.Store(resp.WrittenAt)

	// delete is an optional callback the caller can use to delete the relationship
	// if it's no longer needed. This is helpful if their tx fails.
	delete = func() error {
		// TODO: Fill this out correctly
		resp, err := s.permCli.DeleteRelationships(ctx, &v1.DeleteRelationshipsRequest{
			RelationshipFilter:            nil,
			OptionalPreconditions:         nil,
			OptionalLimit:                 0,
			OptionalAllowPartialDeletions: false,
		})
		if err != nil {
			return err
		}

		if resp.DeletionProgress == v1.DeleteRelationshipsResponse_DELETION_PROGRESS_PARTIAL {
			// This should not happen if OptionalAllowPartialDeletions=false
			return xerrors.Errorf("partial deletion occurred")
		}
		if resp.DeletionProgress == v1.DeleteRelationshipsResponse_DELETION_PROGRESS_COMPLETE {
			return xerrors.Errorf("deletion failed")
		}
		return nil
	}
	return delete, nil
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
