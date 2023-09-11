package keeper

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/aljo242/sync/x/sync/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) Admin(goCtx context.Context, msg *types.MsgAdmin) (*types.MsgAdminResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Authority != k.authority {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect authority")
	}

	k.SetAdmin(ctx, msg.Admin)

	return &types.MsgAdminResponse{}, nil
}
