package keeper_test

import (
	"testing"

	keepertest "github.com/aljo242/sync/testutil/keeper"
	"github.com/aljo242/sync/testutil/nullify"
	"github.com/aljo242/sync/x/sync/keeper"
	"github.com/aljo242/sync/x/sync/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func createNHeader(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Header {
	items := make([]types.Header, n)
	for i := range items {
		items[i].BlockID = keeper.AppendHeader(ctx, items[i])
	}
	return items
}

func TestHeaderGet(t *testing.T) {
	keeper, ctx := keepertest.SyncKeeper(t)
	items := createNHeader(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetHeader(ctx, item.BlockID)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestHeaderRemove(t *testing.T) {
	keeper, ctx := keepertest.SyncKeeper(t)
	items := createNHeader(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveHeader(ctx, item.BlockID)
		_, found := keeper.GetHeader(ctx, item.BlockID)
		require.False(t, found)
	}
}

func TestHeaderGetAll(t *testing.T) {
	keeper, ctx := keepertest.SyncKeeper(t)
	items := createNHeader(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllHeader(ctx)),
	)
}

func TestHeaderCount(t *testing.T) {
	keeper, ctx := keepertest.SyncKeeper(t)
	items := createNHeader(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetHeaderCount(ctx))
}
