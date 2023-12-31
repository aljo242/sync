package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/aljo242/sync/x/sync/types"
)

// GetAdmin gets the admin
func (k Keeper) GetAdmin(ctx sdk.Context) string {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	bz := store.Get(types.KeyPrefixAdmin)

	// Count doesn't exist: no element
	if bz == nil {
		return ""
	}

	// Parse bytes
	return string(bz)
}

// SetAdmin sets the admin
func (k Keeper) SetAdmin(ctx sdk.Context, admin string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	store.Set(types.KeyPrefixAdmin, []byte(admin))
}
