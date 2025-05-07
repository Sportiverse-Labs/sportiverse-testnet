package keeper

import (
	"context"
	"encoding/binary"

	"sportiverse/x/sportiverse/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetCommentCount get the total number of comment
func (k Keeper) GetCommentCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.CommentCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetCommentCount set the total number of comment
func (k Keeper) SetCommentCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.CommentCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendComment appends a comment in the store with a new id and update the count
func (k Keeper) AppendComment(
	ctx context.Context,
	comment types.Comment,
) uint64 {
	// Create the comment
	count := k.GetCommentCount(ctx)

	// Set the ID of the appended value
	comment.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CommentKey))
	appendedValue := k.cdc.MustMarshal(&comment)
	store.Set(GetCommentIDBytes(comment.Id), appendedValue)

	// Update comment count
	k.SetCommentCount(ctx, count+1)

	return count
}

// SetComment set a specific comment in the store
func (k Keeper) SetComment(ctx context.Context, comment types.Comment) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CommentKey))
	b := k.cdc.MustMarshal(&comment)
	store.Set(GetCommentIDBytes(comment.Id), b)
}

// GetComment returns a comment from its id
func (k Keeper) GetComment(ctx context.Context, id uint64) (val types.Comment, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CommentKey))
	b := store.Get(GetCommentIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveComment removes a comment from the store
func (k Keeper) RemoveComment(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CommentKey))
	store.Delete(GetCommentIDBytes(id))
}

// GetAllComment returns all comment
func (k Keeper) GetAllComment(ctx context.Context) (list []types.Comment) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.CommentKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Comment
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetCommentIDBytes returns the byte representation of the ID
func GetCommentIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.CommentKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
