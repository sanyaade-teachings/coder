package spice

import (
	"context"
	"log"

	"github.com/authzed/authzed-go/pkg/responsemeta"
	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/spicedb/pkg/cmd/server"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/encoding/protojson"
	"tailscale.com/syncs"

	"cdr.dev/slog"
	"github.com/coder/coder/v2/coderd/database/spice/policy"
)

type SpiceDB struct {
	logger slog.Logger

	srv       server.RunnableServer
	schemaSrv v1.SchemaServiceClient
	permSrv   v1.PermissionsServiceClient

	// This might not be the most efficient, but we should update
	// this on any write.
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
	s.schemaSrv = v1.NewSchemaServiceClient(conn)
	s.permSrv = v1.NewPermissionsServiceClient(conn)

	// TODO: If the server isn't running yet because it's async, will this fail?
	resp, err := s.schemaSrv.WriteSchema(ctx, &v1.WriteSchemaRequest{
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

func (s *SpiceDB) Check(ctx context.Context) (bool, error) {
	opts := []grpc.CallOption{}
	if s.debug {
		// Add debug information to the request so we can see the trace of the check.
		var trailerMD metadata.MD
		opts = append(opts, grpc.Trailer(&trailerMD))
		defer func() {
			if trailerMD.Len() > 0 {
				// All this debug stuff just shows the trace of the check
				// with information like cache hits.
				found, err := responsemeta.GetResponseTrailerMetadata(trailerMD, responsemeta.DebugInformation)
				if err != nil {
					s.logger.Debug(ctx, "unable to get response metadata", slog.Error(err))
					return
				}

				debugInfo := &v1.DebugInformation{}
				err = protojson.Unmarshal([]byte(found), debugInfo)
				if err != nil {
					s.logger.Debug(ctx, "unable to unmarshal debug proto data", slog.Error(err))
					return
				}

				if debugInfo.Check == nil {
					log.Println("No trace found for the check")
				} else {
					tp := NewTreePrinter()
					DisplayCheckTrace(debugInfo.Check, tp, false)
					s.logger.Debug(ctx, tp.String())
				}
			}
		}()
	}

	resp, err := s.permSrv.CheckPermission(ctx, &v1.CheckPermissionRequest{
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
