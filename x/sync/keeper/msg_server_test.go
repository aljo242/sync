package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/aljo242/sync/testutil/keeper"
	"github.com/aljo242/sync/x/sync/keeper"
	"github.com/aljo242/sync/x/sync/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.SyncKeeper(t)
	k.SetParams(ctx, types.DefaultParams())
	k.SetAdmin(ctx, types.DefaultAdmin)

	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

func TestMsgServer(t *testing.T) {
	ms, ctx := setupMsgServer(t)
	require.NotNil(t, ms)
	require.NotNil(t, ctx)
}
