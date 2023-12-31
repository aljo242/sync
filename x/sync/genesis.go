package sync

import (
	"github.com/aljo242/sync/x/sync/keeper"
	"github.com/aljo242/sync/x/sync/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the header
	for _, elem := range genState.HeaderList {
		k.SetHeader(ctx, elem)
	}

	// Set header count
	k.SetHeaderCount(ctx, genState.HeaderCount)

	k.SetAdmin(ctx, genState.Admin)

	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)
	genesis.HeaderList = k.GetAllHeader(ctx)
	genesis.HeaderCount = k.GetHeaderCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
