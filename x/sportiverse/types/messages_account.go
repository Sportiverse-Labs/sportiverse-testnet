package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateAccount{}

func NewMsgCreateAccount(creator string, hashAccount string) *MsgCreateAccount {
	return &MsgCreateAccount{
		Creator:     creator,
		HashAccount: hashAccount,
	}
}

func (msg *MsgCreateAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteAccount{}

func NewMsgDeleteAccount(creator string, id uint64) *MsgDeleteAccount {
	return &MsgDeleteAccount{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteAccount) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
