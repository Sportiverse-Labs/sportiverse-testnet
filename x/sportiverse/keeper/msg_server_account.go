package keeper

import (
	"context"
	"fmt"

	"sportiverse/x/sportiverse/types"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateAccount(goCtx context.Context, msg *types.MsgCreateAccount) (*types.MsgCreateAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	var account = types.Account{
		Creator:     msg.Creator,
		HashAccount: msg.HashAccount,
	}

	id := k.AppendAccount(
		ctx,
		account,
	)

	return &types.MsgCreateAccountResponse{
		Id: id,
	}, nil
}

func (k msgServer) DeleteAccount(goCtx context.Context, msg *types.MsgDeleteAccount) (*types.MsgDeleteAccountResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Checks that the element exists
	val, found := k.GetAccount(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	// Checks if the msg creator is the same as the current owner
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}

	k.RemoveAccount(ctx, msg.Id)

	return &types.MsgDeleteAccountResponse{}, nil
}
