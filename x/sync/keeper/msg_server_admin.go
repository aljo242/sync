package keeper

import (
	"context"
	"github.com/aljo242/sync/x/sync/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Admin(goCtx context.Context, msg *types.MsgAdmin) (*types.MsgAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	k.SetAdmin(ctx, msg.Admin)

	return &types.MsgAdminResponse{}, nil
}
