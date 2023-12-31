package keeper

import (
	"context"
	"fmt"
	"github.com/aljo242/sync/x/sync/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateHeader(goCtx context.Context, msg *types.MsgCreateHeader) (*types.MsgCreateHeaderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Admin != k.GetAdmin(ctx) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, fmt.Sprintf("incorrect admin: got %s, expected %s", msg.Admin, k.GetAdmin(ctx)))
	}

	var header = types.Header{
		BlockID:     msg.BlockNumber,
		ParentHash:  msg.ParentHash,
		UncleHash:   msg.UncleHash,
		RootHash:    msg.RootHash,
		TxHash:      msg.TxHash,
		ReceiptHash: msg.ReceiptHash,
		Hash:        msg.Hash,
		BlockNumber: msg.BlockNumber,
	}

	k.SetHeader(ctx, header)
	k.SetHeaderHashMapping(ctx, msg.BlockNumber, header.Hash)
	k.SetHeaderHashMapping(ctx, msg.BlockNumber, header.TxHash)
	k.SetHeaderHashMapping(ctx, msg.BlockNumber, header.ReceiptHash)
	k.SetHeaderHashMapping(ctx, msg.BlockNumber, header.RootHash)

	return &types.MsgCreateHeaderResponse{
		BlockID: header.BlockID,
	}, nil
}

func (k msgServer) UpdateHeader(goCtx context.Context, msg *types.MsgUpdateHeader) (*types.MsgUpdateHeaderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks if the msg admin is correct
	if msg.Admin != k.GetAdmin(ctx) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect admin")
	}

	var header = types.Header{
		BlockID: msg.BlockID,
	}

	k.SetHeader(ctx, header)

	return &types.MsgUpdateHeaderResponse{}, nil
}

func (k msgServer) DeleteHeader(goCtx context.Context, msg *types.MsgDeleteHeader) (*types.MsgDeleteHeaderResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if msg.Admin != k.GetAdmin(ctx) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "incorrect admin")
	}

	k.RemoveHeader(ctx, msg.BlockID)

	return &types.MsgDeleteHeaderResponse{}, nil
}
