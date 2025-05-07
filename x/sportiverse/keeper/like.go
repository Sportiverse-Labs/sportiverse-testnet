package keeper

import (
	"context"
	"encoding/binary"

	"sportiverse/x/sportiverse/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetLikeCount get the total number of like
func (k Keeper) GetLikeCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.LikeCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetLikeCount set the total number of like
func (k Keeper) SetLikeCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.LikeCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendLike appends a like in the store with a new id and update the count
func (k Keeper) AppendLike(
	ctx context.Context,
	like types.Like,
) uint64 {
	// Create the like
	count := k.GetLikeCount(ctx)

	// Set the ID of the appended value
	like.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LikeKey))
	appendedValue := k.cdc.MustMarshal(&like)
	store.Set(GetLikeIDBytes(like.Id), appendedValue)

	// Update like count
	k.SetLikeCount(ctx, count+1)

	return count
}

// SetLike set a specific like in the store
func (k Keeper) SetLike(ctx context.Context, like types.Like) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LikeKey))
	b := k.cdc.MustMarshal(&like)
	store.Set(GetLikeIDBytes(like.Id), b)
}

// GetLike returns a like from its id
func (k Keeper) GetLike(ctx context.Context, id uint64) (val types.Like, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LikeKey))
	b := store.Get(GetLikeIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLike removes a like from the store
func (k Keeper) RemoveLike(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LikeKey))
	store.Delete(GetLikeIDBytes(id))
}

// GetAllLike returns all like
func (k Keeper) GetAllLike(ctx context.Context) (list []types.Like) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.LikeKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Like
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetLikeIDBytes returns the byte representation of the ID
func GetLikeIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.LikeKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
