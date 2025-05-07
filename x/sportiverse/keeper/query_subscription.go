package keeper

import (
	"context"

	"sportiverse/x/sportiverse/types"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) SubscriptionAll(ctx context.Context, req *types.QueryAllSubscriptionRequest) (*types.QueryAllSubscriptionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var subscriptions []types.Subscription

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	subscriptionStore := prefix.NewStore(store, types.KeyPrefix(types.SubscriptionKey))

	pageRes, err := query.Paginate(subscriptionStore, req.Pagination, func(key []byte, value []byte) error {
		var subscription types.Subscription
		if err := k.cdc.Unmarshal(value, &subscription); err != nil {
			return err
		}

		subscriptions = append(subscriptions, subscription)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSubscriptionResponse{Subscription: subscriptions, Pagination: pageRes}, nil
}

func (k Keeper) Subscription(ctx context.Context, req *types.QueryGetSubscriptionRequest) (*types.QueryGetSubscriptionResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	subscription, found := k.GetSubscription(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetSubscriptionResponse{Subscription: subscription}, nil
}
