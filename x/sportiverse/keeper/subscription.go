package keeper

import (
	"context"
	"encoding/binary"

	"sportiverse/x/sportiverse/types"

	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/runtime"
)

// GetSubscriptionCount get the total number of subscription
func (k Keeper) GetSubscriptionCount(ctx context.Context) uint64 {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.SubscriptionCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetSubscriptionCount set the total number of subscription
func (k Keeper) SetSubscriptionCount(ctx context.Context, count uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, []byte{})
	byteKey := types.KeyPrefix(types.SubscriptionCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendSubscription appends a subscription in the store with a new id and update the count
func (k Keeper) AppendSubscription(
	ctx context.Context,
	subscription types.Subscription,
) uint64 {
	// Create the subscription
	count := k.GetSubscriptionCount(ctx)

	// Set the ID of the appended value
	subscription.Id = count

	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SubscriptionKey))
	appendedValue := k.cdc.MustMarshal(&subscription)
	store.Set(GetSubscriptionIDBytes(subscription.Id), appendedValue)

	// Update subscription count
	k.SetSubscriptionCount(ctx, count+1)

	return count
}

// SetSubscription set a specific subscription in the store
func (k Keeper) SetSubscription(ctx context.Context, subscription types.Subscription) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SubscriptionKey))
	b := k.cdc.MustMarshal(&subscription)
	store.Set(GetSubscriptionIDBytes(subscription.Id), b)
}

// GetSubscription returns a subscription from its id
func (k Keeper) GetSubscription(ctx context.Context, id uint64) (val types.Subscription, found bool) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SubscriptionKey))
	b := store.Get(GetSubscriptionIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSubscription removes a subscription from the store
func (k Keeper) RemoveSubscription(ctx context.Context, id uint64) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SubscriptionKey))
	store.Delete(GetSubscriptionIDBytes(id))
}

// GetAllSubscription returns all subscription
func (k Keeper) GetAllSubscription(ctx context.Context) (list []types.Subscription) {
	storeAdapter := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	store := prefix.NewStore(storeAdapter, types.KeyPrefix(types.SubscriptionKey))
	iterator := storetypes.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Subscription
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetSubscriptionIDBytes returns the byte representation of the ID
func GetSubscriptionIDBytes(id uint64) []byte {
	bz := types.KeyPrefix(types.SubscriptionKey)
	bz = append(bz, []byte("/")...)
	bz = binary.BigEndian.AppendUint64(bz, id)
	return bz
}
