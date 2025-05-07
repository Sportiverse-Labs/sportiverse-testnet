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

func (k Keeper) LikeAll(ctx context.Context, req *types.QueryAllLikeRequest) (*types.QueryAllLikeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var likes []types.Like

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	likeStore := prefix.NewStore(store, types.KeyPrefix(types.LikeKey))

	pageRes, err := query.Paginate(likeStore, req.Pagination, func(key []byte, value []byte) error {
		var like types.Like
		if err := k.cdc.Unmarshal(value, &like); err != nil {
			return err
		}

		likes = append(likes, like)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllLikeResponse{Like: likes, Pagination: pageRes}, nil
}

func (k Keeper) Like(ctx context.Context, req *types.QueryGetLikeRequest) (*types.QueryGetLikeResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	like, found := k.GetLike(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetLikeResponse{Like: like}, nil
}
