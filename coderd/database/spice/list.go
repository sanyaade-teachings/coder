package spice

import (
	"context"
	"io"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/google/uuid"
	"golang.org/x/xerrors"
	"google.golang.org/grpc"
)

// TODO: Lookup is probably not a cheap operation at present. This is an MVP.
//
//	It is split into a different struct because in the future it is likely we
//	will need to make multiple calls. At present it fetches all IDs the user
//	can access, and returns the entire slice.
func (s *SpiceDB) Lookup(ctx context.Context, permission string, resource *v1.ObjectReference) ([]uuid.UUID, error) {
	l, err := s.lookupContext(ctx, permission, resource)
	if err != nil {
		return nil, xerrors.Errorf("lookup context: %w", err)
	}
	all, err := l.Lookup()
	if err != nil {
		return nil, xerrors.Errorf("lookup: %w", err)
	}

	// TODO: this feels annoying
	ids := make([]uuid.UUID, 0, len(all))
	for _, id := range all {
		u, err := uuid.Parse(id)
		if err != nil {
			return nil, xerrors.Errorf("parse uuid: %w", err)
		}
		ids = append(ids, u)
	}
	return ids, nil
}

type lookupContext struct {
	ctx      context.Context
	sdb      *SpiceDB
	cli      v1.PermissionsService_LookupResourcesClient
	callback func()
}

func (s *SpiceDB) lookupContext(ctx context.Context, permission string, resource *v1.ObjectReference) (*lookupContext, error) {
	actor, ok := ActorFromContext(ctx)
	if !ok {
		return nil, NoActorError
	}

	debugCallback := func() {}
	opts := []grpc.CallOption{}
	if s.debug {
		debugCtx, opt, callback := debugSpiceDBRPC(ctx, s.logger)
		opts = append(opts, opt)
		debugCallback = callback
		ctx = debugCtx
	}

	resp, err := s.permCli.LookupResources(ctx, &v1.LookupResourcesRequest{
		Consistency:        &v1.Consistency{Requirement: &v1.Consistency_AtLeastAsFresh{s.zedToken.Load()}},
		ResourceObjectType: resource.ObjectType,
		Permission:         permission,
		Subject:            actor,
		//Context:            nil,
		OptionalLimit:  0,
		OptionalCursor: nil,
	}, opts...)
	if err != nil {
		return nil, err
	}

	return &lookupContext{
		ctx:      ctx,
		sdb:      s,
		cli:      resp,
		callback: debugCallback,
	}, nil
}

func (c *lookupContext) Lookup() ([]string, error) {
	ids := make([]string, 0)
	defer c.callback()
	for {
		resp, err := c.cli.Recv()
		if xerrors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			return nil, err
		}
		ids = append(ids, resp.ResourceObjectId)
	}
	return ids, nil
}
