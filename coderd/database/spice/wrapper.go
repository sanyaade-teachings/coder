package spice

import (
	"context"

	"golang.org/x/xerrors"

	"cdr.dev/slog"

	"github.com/coder/coder/v2/coderd/database"
	"github.com/coder/coder/v2/coderd/database/spice/policy"
)

func (s *SpiceDB) InsertWorkspace(ctx context.Context, arg database.InsertWorkspaceParams) (database.Workspace, error) {
	org := policy.Organization(arg.OrganizationID)
	owner := policy.User(arg.OwnerID)
	wrk := policy.Workspace(arg.ID).
		For_userUser(owner).
		Organization(org)

	revert, err := s.WriteRelationships(ctx, wrk.Relationships...)
	if err != nil {
		return database.Workspace{}, xerrors.Errorf("write relationships: %w", err)
	}

	workspace, err := s.Store.InsertWorkspace(ctx, arg)
	if err != nil {
		revertError := revert()
		if revertError != nil {
			s.logger.Error(ctx, "revert relationships",
				slog.F("workspace", arg.ID),
				slog.Error(revertError),
			)
		}
		return database.Workspace{}, err
	}

	return workspace, nil
}
