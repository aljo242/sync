package keeper_test

import (
	"testing"

	testkeeper "github.com/aljo242/sync/testutil/keeper"
	"github.com/aljo242/sync/x/sync/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.SyncKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
