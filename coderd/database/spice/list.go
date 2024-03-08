package spice

import (
	"context"

	v1 "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

func (s *SpiceDB) Lookup(ctx context.Context, permission string, resource *v1.ObjectReference) ([]uuid.UUID, error) {
	actor, ok := ActorFromContext(ctx)
	if !ok {
		return nil, NoActorError
	}

	//if actor.Object.ObjectType == god && actor.Object.ObjectId == "god" {
	//	return nil
	//}

	opts := []grpc.CallOption{}
	if s.debug {
		debugCtx, opt, callback := debugSpiceDBRPC(ctx, s.logger)
		opts = append(opts, opt)
		defer callback()
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

	return nil, nil
}
