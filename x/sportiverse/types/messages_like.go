package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateLike{}

func NewMsgCreateLike(creator string, postId string) *MsgCreateLike {
	return &MsgCreateLike{
		Creator: creator,
		PostId:  postId,
	}
}

func (msg *MsgCreateLike) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteLike{}

func NewMsgDeleteLike(creator string, id uint64) *MsgDeleteLike {
	return &MsgDeleteLike{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteLike) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
