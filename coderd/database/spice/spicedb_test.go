package spice_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"cdr.dev/slog"
	"cdr.dev/slog/sloggers/slogtest"
	"github.com/coder/coder/v2/coderd/database"
	"github.com/coder/coder/v2/coderd/database/dbauthz"
	"github.com/coder/coder/v2/coderd/database/dbgen"
	"github.com/coder/coder/v2/coderd/database/dbmem"
	"github.com/coder/coder/v2/coderd/database/dbtestutil"
	"github.com/coder/coder/v2/coderd/database/spice"
	"github.com/coder/coder/v2/coderd/rbac"
	"github.com/coder/coder/v2/testutil"
)

func TestInNestedTX(t *testing.T) {
	t.Parallel()

	ctx := testutil.Context(t, testutil.WaitLong)
	db, _ := dbtestutil.NewDB(t)
	s, err := spice.New(ctx, &spice.SpiceServerOpts{
		PostgresURI: "",
		Logger:      slog.Logger{},
		Store:       db,
	})
	require.NoError(t, err, "must not error")

	u := dbgen.User(t, db, database.User{})
	actor := rbac.Subject{
		ID:     u.ID.String(),
		Roles:  rbac.RoleNames{rbac.RoleOwner()},
		Groups: []string{},
		Scope:  rbac.ScopeAll,
	}

	ctx = dbauthz.As(ctx, actor)
	// A double nested tx
	err = s.InTx(func(tx database.Store) error {
		return tx.InTx(func(tx database.Store) error {
			_, err := tx.GetUserByID(ctx, u.ID)
			return err
		}, nil)
	}, nil)
	require.NoError(t, err, "must not error")
}

func TestSimplePermissionCheck(t *testing.T) {
	t.Parallel()

	db := dbmem.New()
	logger := slogtest.Make(t, nil)

	ctx := testutil.Context(t, testutil.WaitLong)

	def, err := db.GetDefaultOrganization(ctx)
	require.NoError(t, err)

	owner := dbgen.User(t, db, database.User{
		RBACRoles: []string{rbac.RoleOwner()},
	})
	member := dbgen.User(t, db, database.User{})
	otherMember := dbgen.User(t, db, database.User{})

	dbgen.OrganizationMember(t, db, database.OrganizationMember{
		UserID:         owner.ID,
		OrganizationID: def.ID,
	})
	dbgen.OrganizationMember(t, db, database.OrganizationMember{
		UserID:         member.ID,
		OrganizationID: def.ID,
	})
	dbgen.OrganizationMember(t, db, database.OrganizationMember{
		UserID:         otherMember.ID,
		OrganizationID: def.ID,
	})

	sdb, err := spice.New(testutil.Context(t, testutil.WaitLong), &spice.SpiceServerOpts{
		PostgresURI: "",
		Logger:      logger,
		Store:       db,
	})
	require.NoError(t, err)

	memberCtx := spice.AsUser(ctx, member.ID)
	otherMemberCtx := spice.AsUser(ctx, otherMember.ID)
	ownerCtx := spice.AsUser(ctx, owner.ID)

	// Workspace owner by member
	workspace, err := sdb.InsertWorkspace(memberCtx, database.InsertWorkspaceParams{
		OwnerID:        member.ID,
		OrganizationID: def.ID,
	})
	require.NoError(t, err)

	// Member can view
	_, err = sdb.GetWorkspaceByID(memberCtx, workspace.ID)
	require.NoError(t, err)

	// Owner can view
	_, err = sdb.GetWorkspaceByID(ownerCtx, workspace.ID)
	require.NoError(t, err)

	// Other member cannot
	_, err = sdb.GetWorkspaceByID(otherMemberCtx, workspace.ID)
	require.Error(t, err)
}
