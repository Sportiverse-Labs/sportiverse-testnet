package keeper

import (
	"context"
	"fmt"

	"sportiverse/x/sportiverse/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateSubscription(goCtx context.Context, msg *types.MsgCreateSubscription) (*types.MsgCreateSubscriptionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var subscription = types.Subscription{
		Creator: msg.Creator,
		UserId:  msg.UserId,
	}

	id := k.AppendSubscription(
		ctx,
		subscription,
	)

	return &types.MsgCreateSubscriptionResponse{
		Id: id,
	}, nil
}

func (k msgServer) DeleteSubscription(goCtx context.Context, msg *types.MsgDeleteSubscription) (*types.MsgDeleteSubscriptionResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetSubscription(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveSubscription(ctx, msg.Id)

	return &types.MsgDeleteSubscriptionResponse{}, nil
}
