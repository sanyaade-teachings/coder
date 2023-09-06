package spice_test

import (
	"context"
	"testing"

	"github.com/coder/coder/v2/coderd/database/spice"
	"github.com/stretchr/testify/require"
)

func TestSpiceDB(t *testing.T) {
	// Output colors: https://authzed.com/docs/guides/debugging#displaying-explanations-via-zed
	err := spice.RunExample(context.Background())
	require.NoError(t, err)
	//time.Sleep(time.Second * 1000)
}
