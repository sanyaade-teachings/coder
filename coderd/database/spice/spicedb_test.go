package spice_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"cdr.dev/slog"
	"github.com/coder/coder/v2/coderd/database"
	"github.com/coder/coder/v2/coderd/database/dbauthz"
	"github.com/coder/coder/v2/coderd/database/dbgen"
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
