package sync_test

import (
	"testing"

	keepertest "github.com/aljo242/sync/testutil/keeper"
	"github.com/aljo242/sync/testutil/nullify"
	"github.com/aljo242/sync/x/sync"
	"github.com/aljo242/sync/x/sync/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		HeaderList: []types.Header{
			{
				BlockID: 0,
			},
			{
				BlockID: 1,
			},
		},
		HeaderCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.SyncKeeper(t)
	sync.InitGenesis(ctx, *k, genesisState)
	got := sync.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.HeaderList, got.HeaderList)
	require.Equal(t, genesisState.HeaderCount, got.HeaderCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
