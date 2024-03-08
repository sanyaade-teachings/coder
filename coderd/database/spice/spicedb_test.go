package spice_test

import (
	"database/sql"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/moby/moby/pkg/namesgenerator"
	"github.com/stretchr/testify/require"

	"cdr.dev/slog"
	"cdr.dev/slog/sloggers/slogtest"
	"github.com/coder/coder/v2/coderd/database"
	"github.com/coder/coder/v2/coderd/database/dbauthz"
	"github.com/coder/coder/v2/coderd/database/dbgen"
	"github.com/coder/coder/v2/coderd/database/dbmem"
	"github.com/coder/coder/v2/coderd/database/dbtestutil"
	"github.com/coder/coder/v2/coderd/database/dbtime"
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

func TestExampleDatabaseLayer(t *testing.T) {
	t.Parallel()

	logger := slogtest.Make(t, &slogtest.Options{
		IgnoreErrors:   false,
		SkipCleanup:    false,
		IgnoredErrorIs: nil,
	}).Leveled(slog.LevelDebug)

	ctx := testutil.Context(t, testutil.WaitLong)

	db, err := spice.New(ctx, &spice.SpiceServerOpts{
		PostgresURI: "",
		Logger:      logger,
		Store:       dbmem.New(),
		Debug:       false,
	})
	require.NoError(t, err)
	err = db.Run(ctx)
	require.NoError(t, err)

	god := spice.AsGod(ctx)
	def, err := db.GetDefaultOrganization(god)
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

	memberCtx := spice.AsUser(ctx, member.ID)
	otherMemberCtx := spice.AsUser(ctx, otherMember.ID)
	ownerCtx := spice.AsUser(ctx, owner.ID)

	// Workspace owner by member
	workspace, err := db.InsertWorkspace(memberCtx, database.InsertWorkspaceParams{
		ID:                uuid.New(),
		CreatedAt:         dbtime.Now(),
		UpdatedAt:         dbtime.Now(),
		OwnerID:           member.ID,
		OrganizationID:    def.ID,
		TemplateID:        uuid.New(),
		Name:              namesgenerator.GetRandomName(1),
		AutostartSchedule: sql.NullString{},
		Ttl:               sql.NullInt64{},
		LastUsedAt:        time.Time{},
		AutomaticUpdates:  database.AutomaticUpdatesNever,
	})
	require.NoError(t, err)

	// Member can view
	_, err = db.GetWorkspaceByID(memberCtx, workspace.ID)
	require.NoErrorf(t, err, "Member: %s", member.ID.String())

	// Owner can view
	_, err = db.GetWorkspaceByID(ownerCtx, workspace.ID)
	require.NoErrorf(t, err, "Owner: %s", owner.ID.String())

	// Other member cannot
	_, err = db.GetWorkspaceByID(otherMemberCtx, workspace.ID)
	require.Errorf(t, err, "OtherMember: %s", otherMember.ID.String())
}
