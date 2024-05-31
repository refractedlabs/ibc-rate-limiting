package types

import (
	"github.com/cosmos/cosmos-sdk/types/bech32"
	"regexp"

	errorsmod "cosmossdk.io/errors"
	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/migrations/legacytx"
)

const (
	TypeMsgAddRateLimit                 = "AddRateLimit"
	TypeMsgUpdateRateLimit              = "UpdateRateLimit"
	TypeMsgRemoveRateLimit              = "RemoveRateLimit"
	TypeMsgResetRateLimit               = "ResetRateLimit"
	TypeMsgUpdateParams                 = "UpdateParams"
	TypeMsgSetWhitelistedAddressPair    = "SetWhitelistedAddressPair"
	TypeMsgRemoveWhitelistedAddressPair = "RemoveWhitelistedAddressPair"
)

var (
	_ sdk.Msg = &MsgAddRateLimit{}
	_ sdk.Msg = &MsgUpdateRateLimit{}
	_ sdk.Msg = &MsgRemoveRateLimit{}
	_ sdk.Msg = &MsgResetRateLimit{}
	_ sdk.Msg = &MsgUpdateParams{}

	// Implement legacy interface for ledger support
	_ legacytx.LegacyMsg = &MsgAddRateLimit{}
	_ legacytx.LegacyMsg = &MsgUpdateRateLimit{}
	_ legacytx.LegacyMsg = &MsgRemoveRateLimit{}
	_ legacytx.LegacyMsg = &MsgResetRateLimit{}
	_ legacytx.LegacyMsg = &MsgUpdateParams{}
)

// ----------------------------------------------
//               MsgAddRateLimit
// ----------------------------------------------

func NewMsgAddRateLimit(denom, channelId string, maxPercentSend sdkmath.Int, maxPercentRecv sdkmath.Int, durationHours uint64) *MsgAddRateLimit {
	return &MsgAddRateLimit{
		Denom:          denom,
		ChannelId:      channelId,
		MaxPercentSend: maxPercentSend,
		MaxPercentRecv: maxPercentRecv,
		DurationHours:  durationHours,
	}
}

func (msg MsgAddRateLimit) Type() string {
	return TypeMsgAddRateLimit
}

func (msg MsgAddRateLimit) Route() string {
	return RouterKey
}

func (msg *MsgAddRateLimit) GetSigners() []sdk.AccAddress {
	staker, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{staker}
}

func (msg *MsgAddRateLimit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgAddRateLimit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address (%s)", err)
	}

	if msg.Denom == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid denom (%s)", msg.Denom)
	}

	matched, err := regexp.MatchString(`^channel-\d+$`, msg.ChannelId)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "unable to verify channel-id (%s)", msg.ChannelId)
	}
	if !matched {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest,
			"invalid channel-id (%s), must be of the format 'channel-{N}'", msg.ChannelId)
	}

	if msg.MaxPercentSend.GT(sdkmath.NewInt(100)) || msg.MaxPercentSend.LT(sdkmath.ZeroInt()) {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest,
			"max-percent-send percent must be between 0 and 100 (inclusively), Provided: %v", msg.MaxPercentSend)
	}

	if msg.MaxPercentRecv.GT(sdkmath.NewInt(100)) || msg.MaxPercentRecv.LT(sdkmath.ZeroInt()) {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest,
			"max-percent-recv percent must be between 0 and 100 (inclusively), Provided: %v", msg.MaxPercentRecv)
	}

	if msg.MaxPercentRecv.IsZero() && msg.MaxPercentSend.IsZero() {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest,
			"either the max send or max receive threshold must be greater than 0")
	}

	if msg.DurationHours == 0 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "duration can not be zero")
	}

	return nil
}

// ----------------------------------------------
//               MsgUpdateRateLimit
// ----------------------------------------------

func NewMsgUpdateRateLimit(denom, channelId string, maxPercentSend sdkmath.Int, maxPercentRecv sdkmath.Int, durationHours uint64) *MsgUpdateRateLimit {
	return &MsgUpdateRateLimit{
		Denom:          denom,
		ChannelId:      channelId,
		MaxPercentSend: maxPercentSend,
		MaxPercentRecv: maxPercentRecv,
		DurationHours:  durationHours,
	}
}

func (msg MsgUpdateRateLimit) Type() string {
	return TypeMsgUpdateRateLimit
}

func (msg MsgUpdateRateLimit) Route() string {
	return RouterKey
}

func (msg *MsgUpdateRateLimit) GetSigners() []sdk.AccAddress {
	staker, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{staker}
}

func (msg *MsgUpdateRateLimit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateRateLimit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address (%s)", err)
	}

	if msg.Denom == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid denom (%s)", msg.Denom)
	}

	matched, err := regexp.MatchString(`^channel-\d+$`, msg.ChannelId)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "unable to verify channel-id (%s)", msg.ChannelId)
	}
	if !matched {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest,
			"invalid channel-id (%s), must be of the format 'channel-{N}'", msg.ChannelId)
	}

	if msg.MaxPercentSend.GT(sdkmath.NewInt(100)) || msg.MaxPercentSend.LT(sdkmath.ZeroInt()) {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest,
			"max-percent-send percent must be between 0 and 100 (inclusively), Provided: %v", msg.MaxPercentSend)
	}

	if msg.MaxPercentRecv.GT(sdkmath.NewInt(100)) || msg.MaxPercentRecv.LT(sdkmath.ZeroInt()) {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest,
			"max-percent-recv percent must be between 0 and 100 (inclusively), Provided: %v", msg.MaxPercentRecv)
	}

	if msg.MaxPercentRecv.IsZero() && msg.MaxPercentSend.IsZero() {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest,
			"either the max send or max receive threshold must be greater than 0")
	}

	if msg.DurationHours == 0 {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "duration can not be zero")
	}

	return nil
}

// ----------------------------------------------
//               MsgRemoveRateLimit
// ----------------------------------------------

func NewMsgRemoveRateLimit(denom, channelId string) *MsgRemoveRateLimit {
	return &MsgRemoveRateLimit{
		Denom:     denom,
		ChannelId: channelId,
	}
}

func (msg MsgRemoveRateLimit) Type() string {
	return TypeMsgRemoveRateLimit
}

func (msg MsgRemoveRateLimit) Route() string {
	return RouterKey
}

func (msg *MsgRemoveRateLimit) GetSigners() []sdk.AccAddress {
	staker, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{staker}
}

func (msg *MsgRemoveRateLimit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveRateLimit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address (%s)", err)
	}

	if msg.Denom == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid denom (%s)", msg.Denom)
	}

	matched, err := regexp.MatchString(`^channel-\d+$`, msg.ChannelId)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "unable to verify channel-id (%s)", msg.ChannelId)
	}
	if !matched {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest,
			"invalid channel-id (%s), must be of the format 'channel-{N}'", msg.ChannelId)
	}

	return nil
}

// ----------------------------------------------
//               MsgResetRateLimit
// ----------------------------------------------

func NewMsgResetRateLimit(denom, channelId string) *MsgResetRateLimit {
	return &MsgResetRateLimit{
		Denom:     denom,
		ChannelId: channelId,
	}
}

func (msg MsgResetRateLimit) Type() string {
	return TypeMsgResetRateLimit
}

func (msg MsgResetRateLimit) Route() string {
	return RouterKey
}

func (msg *MsgResetRateLimit) GetSigners() []sdk.AccAddress {
	staker, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{staker}
}

func (msg *MsgResetRateLimit) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgResetRateLimit) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address (%s)", err)
	}

	if msg.Denom == "" {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "invalid denom (%s)", msg.Denom)
	}

	matched, err := regexp.MatchString(`^channel-\d+$`, msg.ChannelId)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest, "unable to verify channel-id (%s)", msg.ChannelId)
	}
	if !matched {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidRequest,
			"invalid channel-id (%s), must be of the format 'channel-{N}'", msg.ChannelId)
	}

	return nil
}

// ----------------------------------------------
//               MsgUpdateParams
// ----------------------------------------------

func (msg *MsgUpdateParams) Type() string {
	return TypeMsgUpdateParams
}

func (msg *MsgUpdateParams) Route() string {
	return RouterKey
}

func (msg *MsgUpdateParams) GetSigners() []sdk.AccAddress {
	authority, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{authority}
}

func (msg *MsgUpdateParams) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateParams) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address (%s)", err)
	}

	return msg.Params.Validate()
}

// ----------------------------------------------
//          MsgSetWhitelistedAddressPair
// ----------------------------------------------

func (msg *MsgSetWhitelistedAddressPair) Type() string {
	return TypeMsgSetWhitelistedAddressPair
}

func (msg *MsgSetWhitelistedAddressPair) Route() string {
	return RouterKey
}

func (msg *MsgSetWhitelistedAddressPair) GetSigners() []sdk.AccAddress {
	authority, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{authority}
}

func (msg *MsgSetWhitelistedAddressPair) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSetWhitelistedAddressPair) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address (%s)", err)
	}
	if _, err := getAccountAddress(msg.Sender); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid sender address %s: %s", msg.Sender, err.Error())
	}
	if _, err := getAccountAddress(msg.Receiver); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid receiver address %s: %s", msg.Receiver, err.Error())
	}
	return nil
}

// ----------------------------------------------
//        MsgRemoveWhitelistedAddressPair
// ----------------------------------------------

func (msg *MsgRemoveWhitelistedAddressPair) Type() string {
	return TypeMsgRemoveWhitelistedAddressPair
}

func (msg *MsgRemoveWhitelistedAddressPair) Route() string {
	return RouterKey
}

func (msg *MsgRemoveWhitelistedAddressPair) GetSigners() []sdk.AccAddress {
	authority, err := sdk.AccAddressFromBech32(msg.Authority)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{authority}
}

func (msg *MsgRemoveWhitelistedAddressPair) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgRemoveWhitelistedAddressPair) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Authority); err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid authority address (%s)", err)
	}
	if _, err := getAccountAddress(msg.Sender); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid sender address %s: %s", msg.Sender, err.Error())
	}
	if _, err := getAccountAddress(msg.Receiver); err != nil {
		return sdkerrors.ErrInvalidAddress.Wrapf("invalid receiver address %s: %s", msg.Receiver, err.Error())
	}
	return nil
}

func getAccountAddress(bech32Address string) (sdk.AccAddress, error) {
	_, bz, err := bech32.DecodeAndConvert(bech32Address)
	if err != nil {
		return nil, err
	}

	if len(bz) == 0 {
		return nil, sdkerrors.ErrInvalidAddress.Wrap("address cannot be empty")
	}

	return bz, nil
}
