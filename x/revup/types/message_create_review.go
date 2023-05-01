package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateReview = "create_review"

var _ sdk.Msg = &MsgCreateReview{}

func NewMsgCreateReview(creator string, product string, rate string, comment string) *MsgCreateReview {
	return &MsgCreateReview{
		Creator: creator,
		Product: product,
		Rate:    rate,
		Comment: comment,
	}
}

func (msg *MsgCreateReview) Route() string {
	return RouterKey
}

func (msg *MsgCreateReview) Type() string {
	return TypeMsgCreateReview
}

func (msg *MsgCreateReview) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateReview) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateReview) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
