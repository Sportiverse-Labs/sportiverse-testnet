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

func (k Keeper) CommentAll(ctx context.Context, req *types.QueryAllCommentRequest) (*types.QueryAllCommentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var comments []types.Comment

	store := runtime.KVStoreAdapter(k.storeService.OpenKVStore(ctx))
	commentStore := prefix.NewStore(store, types.KeyPrefix(types.CommentKey))

	pageRes, err := query.Paginate(commentStore, req.Pagination, func(key []byte, value []byte) error {
		var comment types.Comment
		if err := k.cdc.Unmarshal(value, &comment); err != nil {
			return err
		}

		comments = append(comments, comment)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllCommentResponse{Comment: comments, Pagination: pageRes}, nil
}

func (k Keeper) Comment(ctx context.Context, req *types.QueryGetCommentRequest) (*types.QueryGetCommentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	comment, found := k.GetComment(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetCommentResponse{Comment: comment}, nil
}
