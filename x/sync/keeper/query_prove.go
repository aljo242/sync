package keeper

import (
	"context"
	"fmt"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/status"

	"github.com/aljo242/sync/x/sync/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
)

func (k Keeper) Prove(goCtx context.Context, req *types.QueryProveRequest) (*types.QueryProveResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)

	header, found := k.GetHeader(ctx, req.BlockID)
	if !found {
		return nil, fmt.Errorf("unable to find header for blockID %d: %w", req.BlockID, sdkerrors.ErrKeyNotFound)
	}

	proven := header.TxHash == req.Proof
	if !proven {
		return &types.QueryProveResponse{Valid: proven}, fmt.Errorf("mismatch: header: %s, proof %s", header.Hash, req.Proof)
	}

	return &types.QueryProveResponse{Valid: proven}, nil
}
