package keeper

import (
	"encoding/binary"

	"github.com/aljo242/sync/x/sync/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetHeaderCount get the total number of header
func (k Keeper) GetHeaderCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	bz := store.Get(types.KeyPrefixHeaderCount)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetHeaderCount set the total number of header
func (k Keeper) SetHeaderCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(types.KeyPrefixHeaderCount, bz)
}

// AppendHeader appends a header in the store with a new id and update the count
func (k Keeper) AppendHeader(
	ctx sdk.Context,
	header types.Header,
) uint64 {
	// Create the header
	count := k.GetHeaderCount(ctx)

	// Set the ID of the appended value
	header.BlockID = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixHeader)
	appendedValue := k.cdc.MustMarshal(&header)
	store.Set(GetHeaderIDBytes(header.BlockID), appendedValue)

	// Update header count
	k.SetHeaderCount(ctx, count+1)

	return count
}

// SetHeader set a specific header in the store
func (k Keeper) SetHeader(ctx sdk.Context, header types.Header) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixHeader)
	b := k.cdc.MustMarshal(&header)
	store.Set(GetHeaderIDBytes(header.BlockID), b)
}

// GetHeader returns a header from its id
func (k Keeper) GetHeader(ctx sdk.Context, id uint64) (val types.Header, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixHeader)
	b := store.Get(GetHeaderIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// SetHeaderHashMapping set a specific header blockID to hash mapping
func (k Keeper) SetHeaderHashMapping(ctx sdk.Context, blockID uint64, hash string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixHeaderHashMapping)
	store.Set([]byte(hash), GetHeaderIDBytes(blockID))
}

// GetHeaderHashMapping returns a header block ID from its hash
func (k Keeper) GetHeaderHashMapping(ctx sdk.Context, hash string) (val uint64, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixHeader)
	b := store.Get([]byte(hash))
	if b == nil {
		return val, false
	}

	return GetHeaderIDFromBytes(b), true
}

// GetHeaderFromHash returns a header block ID from its hash
func (k Keeper) GetHeaderFromHash(ctx sdk.Context, hash string) (val types.Header, found bool) {
	id, found := k.GetHeaderHashMapping(ctx, hash)
	if !found {
		return val, false
	}

	return k.GetHeader(ctx, id)
}

// RemoveHeader removes a header from the store
func (k Keeper) RemoveHeader(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixHeader)
	store.Delete(GetHeaderIDBytes(id))
}

// GetAllHeader returns all header
func (k Keeper) GetAllHeader(ctx sdk.Context) (list []types.Header) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixHeader)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Header
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetHeaderIDBytes returns the byte representation of the ID
func GetHeaderIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetHeaderIDFromBytes returns ID in uint64 format from a byte array
func GetHeaderIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
