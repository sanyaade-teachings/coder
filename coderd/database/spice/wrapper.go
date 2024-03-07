package spice

import (
	"context"

	"github.com/coder/coder/v2/coderd/database"
	"github.com/coder/coder/v2/coderd/database/spice/policy"
	"github.com/coder/coder/v2/coderd/rbac"
)

func (s *SpiceDB) InsertWorkspace(ctx context.Context, arg database.InsertWorkspaceParams) (database.Workspace, error) {
	builder := policy.New()
	org := builder.Organization(arg.OrganizationID)
	owner := builder.User(arg.OwnerID)
	builder.Workspace(arg.ID).
		For_userUser(owner).
		Organization(org)

	return WithRelations(ctx, s, builder.Relationships, s.Store.InsertWorkspace, arg)
}

func (s *SpiceDB) InsertUser(ctx context.Context, arg database.InsertUserParams) (database.User, error) {
	// Create a new user
	builder := policy.New()
	user := builder.User(arg.ID)
	// This kinda sucks, mapping coder roles to authzed roles.
	// There needs to be a better way to do this.
	for _, role := range arg.RBACRoles {
		switch role {
		case rbac.RoleOwner():
			builder.SitePlatform().Administrator(user)
			// TODO: The rest of the roles
		}
	}

	return WithRelations(ctx, s, builder.Relationships, s.Store.InsertUser, arg)
}

func (s *SpiceDB) InsertOrganization(ctx context.Context, arg database.InsertOrganizationParams) (database.Organization, error) {
	builder := policy.New()
	builder.Organization(arg.ID).
		Platform(builder.SitePlatform())

	return WithRelations(ctx, s, builder.Relationships, s.Store.InsertOrganization, arg)
}

func (s *SpiceDB) InsertOrganizationMember(ctx context.Context, arg database.InsertOrganizationMemberParams) (database.OrganizationMember, error) {
	builder := policy.New()
	user := builder.User(arg.UserID)
	builder.Organization(arg.OrganizationID).
		// TODO: This is assuming a user is only an org member
		Default_permissiosUser(user).
		// We add this because a migration creates the original default org.
		// If a user is added to this org, we want to ensure the org is tracked in
		// the authz graph.
		Platform(builder.SitePlatform())

	return WithRelations(ctx, s, builder.Relationships, s.Store.InsertOrganizationMember, arg)
}

//func (s *SpiceDB) GetWorkspaceByID(ctx context.Context, id uuid.UUID) (database.Workspace, error) {
//	wrk := policy.Workspace(id)
//	return database.Workspace{}, nil
//}
