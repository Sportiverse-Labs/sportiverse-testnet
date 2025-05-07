package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateComment{}

func NewMsgCreateComment(creator string, postId string, body string, timestamp string) *MsgCreateComment {
	return &MsgCreateComment{
		Creator:   creator,
		PostId:    postId,
		Body:      body,
		Timestamp: timestamp,
	}
}

func (msg *MsgCreateComment) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateComment{}

func NewMsgUpdateComment(creator string, id uint64, postId string, body string, timestamp string) *MsgUpdateComment {
	return &MsgUpdateComment{
		Id:        id,
		Creator:   creator,
		PostId:    postId,
		Body:      body,
		Timestamp: timestamp,
	}
}

func (msg *MsgUpdateComment) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteComment{}

func NewMsgDeleteComment(creator string, id uint64) *MsgDeleteComment {
	return &MsgDeleteComment{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteComment) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
