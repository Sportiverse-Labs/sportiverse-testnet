package keeper

import (
	"context"
	"fmt"

	"sportiverse/x/sportiverse/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateLike(goCtx context.Context, msg *types.MsgCreateLike) (*types.MsgCreateLikeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var like = types.Like{
		Creator: msg.Creator,
		PostId:  msg.PostId,
	}

	id := k.AppendLike(
		ctx,
		like,
	)

	return &types.MsgCreateLikeResponse{
		Id: id,
	}, nil
}

func (k msgServer) DeleteLike(goCtx context.Context, msg *types.MsgDeleteLike) (*types.MsgDeleteLikeResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetLike(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveLike(ctx, msg.Id)

	return &types.MsgDeleteLikeResponse{}, nil
}
