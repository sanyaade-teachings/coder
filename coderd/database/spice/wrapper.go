package spice

import (
	"context"

	"github.com/google/uuid"

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
	builder := policy.New()
	// New user
	user := builder.User(arg.ID)

	// Ensure we can create this user
	err := s.Check(builder.SitePlatform().CanCreate_user(ctx))
	if err != nil {
		return database.User{}, err
	}

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
	site := builder.SitePlatform()
	builder.Organization(arg.ID).
		Platform(site)

	err := s.Check(site.CanCreate_organization(ctx))
	if err != nil {
		return database.Organization{}, err
	}

	return WithRelations(ctx, s, builder.Relationships, s.Store.InsertOrganization, arg)
}

func (s *SpiceDB) InsertOrganizationMember(ctx context.Context, arg database.InsertOrganizationMemberParams) (database.OrganizationMember, error) {
	builder := policy.New()
	user := builder.User(arg.UserID)
	org := builder.Organization(arg.OrganizationID).
		MemberUser(user).
		// TODO: This is assuming a user is only an org member
		Default_permissionsUser(user).
		// We add this because a migration creates the original default org.
		// If a user is added to this org, we want to ensure the org is tracked in
		// the authz graph.
		Platform(builder.SitePlatform())

	err := s.Check(org.CanCreate_org_member(ctx))
	if err != nil {
		return database.OrganizationMember{}, err
	}

	return WithRelations(ctx, s, builder.Relationships, s.Store.InsertOrganizationMember, arg)
}

func (s *SpiceDB) GetWorkspaceByID(ctx context.Context, id uuid.UUID) (database.Workspace, error) {
	wrk := policy.New().Workspace(id)
	err := s.Check(wrk.CanView(ctx))
	if err != nil {
		return database.Workspace{}, err
	}

	return database.Workspace{}, nil
}
