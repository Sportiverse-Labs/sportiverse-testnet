package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateSubscription{}

func NewMsgCreateSubscription(creator string, userId string) *MsgCreateSubscription {
	return &MsgCreateSubscription{
		Creator: creator,
		UserId:  userId,
	}
}

func (msg *MsgCreateSubscription) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteSubscription{}

func NewMsgDeleteSubscription(creator string, id uint64) *MsgDeleteSubscription {
	return &MsgDeleteSubscription{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteSubscription) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
