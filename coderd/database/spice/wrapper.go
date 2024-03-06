package spice

import (
	"context"

	"github.com/coder/coder/v2/coderd/database"
	"github.com/coder/coder/v2/coderd/database/spice/policy"
)

func (s *SpiceDB) InsertWorkspace(ctx context.Context, arg database.InsertWorkspaceParams) (database.Workspace, error) {
	s.Store.InTx(func(store database.Store) error {
		workspace, err := s.Store.InsertWorkspace(ctx, arg)
		if err != nil {
			return err
		}

		policy.Workspace(workspace.ID.String()).ViewerUser()

		// Insert relationships
		//workspace.
		//s.WriteRelationship(ctx)
		return nil
	}, nil)

	return s.Store.InsertWorkspace(ctx, arg)
}
