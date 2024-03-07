package spice

import (
	"context"

	"github.com/google/uuid"

	"github.com/coder/coder/v2/coderd/database"
	"github.com/coder/coder/v2/coderd/database/spice/policy"
)

func (s *SpiceDB) InsertWorkspace(ctx context.Context, arg database.InsertWorkspaceParams) (database.Workspace, error) {
	org := policy.Organization(arg.OrganizationID)
	owner := policy.User(arg.OwnerID)
	wrk := policy.Workspace(arg.ID).
		For_userUser(owner).
		Organization(org)

	return WithRelations(ctx, s, wrk.Relationships, s.Store.InsertWorkspace, arg)
}

func (s *SpiceDB) GetUserByID(ctx context.Context, id uuid.UUID) (database.User, error) {
	return s.Store.GetUserByID(ctx, id)
}
