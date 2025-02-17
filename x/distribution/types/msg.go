package types

import (
	errorsmod "cosmossdk.io/errors"
	"github.com/cosmos/cosmos-sdk/codec/legacy"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// distribution message types
const (
	TypeMsgSetWithdrawAddress                   = "set_withdraw_address"
	TypeMsgWithdrawDelegatorReward              = "withdraw_delegator_reward"
	TypeMsgWithdrawValidatorCommission          = "withdraw_validator_commission"
	TypeMsgFundCommunityPool                    = "fund_community_pool"
	TypeMsgWithdrawTokenizeShareRecordReward    = "withdraw_tokenize_share_record_reward"
	TypeMsgWithdrawAllTokenizeShareRecordReward = "withdraw_all_tokenize_share_record_reward"
)

// Verify interface at compile time
var (
	_, _, _ sdk.Msg = &MsgSetWithdrawAddress{}, &MsgWithdrawDelegatorReward{}, &MsgWithdrawValidatorCommission{}
	_       sdk.Msg = &MsgWithdrawTokenizeShareRecordReward{}
	_       sdk.Msg = &MsgWithdrawAllTokenizeShareRecordReward{}
)

func NewMsgSetWithdrawAddress(delAddr, withdrawAddr sdk.AccAddress) *MsgSetWithdrawAddress {
	return &MsgSetWithdrawAddress{
		DelegatorAddress: delAddr.String(),
		WithdrawAddress:  withdrawAddr.String(),
	}
}

func (msg MsgSetWithdrawAddress) Route() string { return ModuleName }
func (msg MsgSetWithdrawAddress) Type() string  { return TypeMsgSetWithdrawAddress }

// Return address that must sign over msg.GetSignBytes()
func (msg MsgSetWithdrawAddress) GetSigners() []sdk.AccAddress {
	delegator, err := sdk.AccAddressFromBech32(msg.DelegatorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{delegator}
}

// get the bytes for the message signer to sign on
func (msg MsgSetWithdrawAddress) GetSignBytes() []byte {
	bz := legacy.Cdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// quick validity check
func (msg MsgSetWithdrawAddress) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.DelegatorAddress); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid delegator address: %s", err)
	}
	if _, err := sdk.AccAddressFromBech32(msg.WithdrawAddress); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid withdraw address: %s", err)
	}

	return nil
}

func NewMsgWithdrawDelegatorReward(delAddr sdk.AccAddress, valAddr sdk.ValAddress) *MsgWithdrawDelegatorReward {
	return &MsgWithdrawDelegatorReward{
		DelegatorAddress: delAddr.String(),
		ValidatorAddress: valAddr.String(),
	}
}

func (msg MsgWithdrawDelegatorReward) Route() string { return ModuleName }
func (msg MsgWithdrawDelegatorReward) Type() string  { return TypeMsgWithdrawDelegatorReward }

// Return address that must sign over msg.GetSignBytes()
func (msg MsgWithdrawDelegatorReward) GetSigners() []sdk.AccAddress {
	delegator, err := sdk.AccAddressFromBech32(msg.DelegatorAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{delegator}
}

// get the bytes for the message signer to sign on
func (msg MsgWithdrawDelegatorReward) GetSignBytes() []byte {
	bz := legacy.Cdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// quick validity check
func (msg MsgWithdrawDelegatorReward) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.DelegatorAddress); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid delegator address: %s", err)
	}
	if _, err := sdk.ValAddressFromBech32(msg.ValidatorAddress); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid validator address: %s", err)
	}
	return nil
}

func NewMsgWithdrawValidatorCommission(valAddr sdk.ValAddress) *MsgWithdrawValidatorCommission {
	return &MsgWithdrawValidatorCommission{
		ValidatorAddress: valAddr.String(),
	}
}

func (msg MsgWithdrawValidatorCommission) Route() string { return ModuleName }
func (msg MsgWithdrawValidatorCommission) Type() string  { return TypeMsgWithdrawValidatorCommission }

// Return address that must sign over msg.GetSignBytes()
func (msg MsgWithdrawValidatorCommission) GetSigners() []sdk.AccAddress {
	valAddr, _ := sdk.ValAddressFromBech32(msg.ValidatorAddress)
	return []sdk.AccAddress{sdk.AccAddress(valAddr)}
}

// get the bytes for the message signer to sign on
func (msg MsgWithdrawValidatorCommission) GetSignBytes() []byte {
	bz := legacy.Cdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// quick validity check
func (msg MsgWithdrawValidatorCommission) ValidateBasic() error {
	if _, err := sdk.ValAddressFromBech32(msg.ValidatorAddress); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid validator address: %s", err)
	}
	return nil
}

// NewMsgFundCommunityPool returns a new MsgFundCommunityPool with a sender and
// a funding amount.
func NewMsgFundCommunityPool(amount sdk.Coins, depositor sdk.AccAddress) *MsgFundCommunityPool {
	return &MsgFundCommunityPool{
		Amount:    amount,
		Depositor: depositor.String(),
	}
}

// Route returns the MsgFundCommunityPool message route.
func (msg MsgFundCommunityPool) Route() string { return ModuleName }

// Type returns the MsgFundCommunityPool message type.
func (msg MsgFundCommunityPool) Type() string { return TypeMsgFundCommunityPool }

// GetSigners returns the signer addresses that are expected to sign the result
// of GetSignBytes.
func (msg MsgFundCommunityPool) GetSigners() []sdk.AccAddress {
	depositor, err := sdk.AccAddressFromBech32(msg.Depositor)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{depositor}
}

// GetSignBytes returns the raw bytes for a MsgFundCommunityPool message that
// the expected signer needs to sign.
func (msg MsgFundCommunityPool) GetSignBytes() []byte {
	bz := legacy.Cdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic performs basic MsgFundCommunityPool message validation.
func (msg MsgFundCommunityPool) ValidateBasic() error {
	if !msg.Amount.IsValid() {
		return errorsmod.Wrap(sdkerrors.ErrInvalidCoins, msg.Amount.String())
	}
	if _, err := sdk.AccAddressFromBech32(msg.Depositor); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid depositor address: %s", err)
	}
	return nil
}

func NewMsgWithdrawTokenizeShareRecordReward(ownerAddr sdk.AccAddress, recordId uint64) *MsgWithdrawTokenizeShareRecordReward {
	return &MsgWithdrawTokenizeShareRecordReward{
		OwnerAddress: ownerAddr.String(),
		RecordId:     recordId,
	}
}

func (msg MsgWithdrawTokenizeShareRecordReward) Route() string { return ModuleName }
func (msg MsgWithdrawTokenizeShareRecordReward) Type() string {
	return TypeMsgWithdrawTokenizeShareRecordReward
}

// Return address that must sign over msg.GetSignBytes()
func (msg MsgWithdrawTokenizeShareRecordReward) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

// get the bytes for the message signer to sign on
func (msg MsgWithdrawTokenizeShareRecordReward) GetSignBytes() []byte {
	bz := legacy.Cdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// quick validity check
func (msg MsgWithdrawTokenizeShareRecordReward) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.OwnerAddress); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid owner address: %s", err)
	}
	return nil
}

func NewMsgWithdrawAllTokenizeShareRecordReward(ownerAddr sdk.AccAddress) *MsgWithdrawAllTokenizeShareRecordReward {
	return &MsgWithdrawAllTokenizeShareRecordReward{
		OwnerAddress: ownerAddr.String(),
	}
}

func (msg MsgWithdrawAllTokenizeShareRecordReward) Route() string { return ModuleName }
func (msg MsgWithdrawAllTokenizeShareRecordReward) Type() string {
	return TypeMsgWithdrawAllTokenizeShareRecordReward
}

// Return address that must sign over msg.GetSignBytes()
func (msg MsgWithdrawAllTokenizeShareRecordReward) GetSigners() []sdk.AccAddress {
	owner, err := sdk.AccAddressFromBech32(msg.OwnerAddress)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{owner}
}

// get the bytes for the message signer to sign on
func (msg MsgWithdrawAllTokenizeShareRecordReward) GetSignBytes() []byte {
	bz := legacy.Cdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// quick validity check
func (msg MsgWithdrawAllTokenizeShareRecordReward) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.OwnerAddress); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid owner address: %s", err)
	}
	return nil
}
