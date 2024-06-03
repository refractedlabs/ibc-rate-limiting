package keeper

import (
	"context"
	"slices"

	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/Stride-Labs/ibc-rate-limiting/ratelimit/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the ratelimit MsgServer interface
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

// Adds a new rate limit. Fails if the rate limit already exists or the channel value is 0
func (k msgServer) AddRateLimit(goCtx context.Context, msg *types.MsgAddRateLimit) (*types.MsgAddRateLimitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := k.assertGovOrAdmin(ctx, msg.Authority); err != nil {
		return nil, err
	}

	if err := k.Keeper.AddRateLimit(ctx, msg); err != nil {
		return nil, err
	}

	return &types.MsgAddRateLimitResponse{}, nil
}

// Updates an existing rate limit. Fails if the rate limit doesn't exist
func (k msgServer) UpdateRateLimit(goCtx context.Context, msg *types.MsgUpdateRateLimit) (*types.MsgUpdateRateLimitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := k.assertGovOrAdmin(ctx, msg.Authority); err != nil {
		return nil, err
	}

	if err := k.Keeper.UpdateRateLimit(ctx, msg); err != nil {
		return nil, err
	}

	return &types.MsgUpdateRateLimitResponse{}, nil
}

// Removes a rate limit. Fails if the rate limit doesn't exist
func (k msgServer) RemoveRateLimit(goCtx context.Context, msg *types.MsgRemoveRateLimit) (*types.MsgRemoveRateLimitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := k.assertGovOrAdmin(ctx, msg.Authority); err != nil {
		return nil, err
	}

	_, found := k.Keeper.GetRateLimit(ctx, msg.Denom, msg.ChannelId)
	if !found {
		return nil, types.ErrRateLimitNotFound
	}

	k.Keeper.RemoveRateLimit(ctx, msg.Denom, msg.ChannelId)
	return &types.MsgRemoveRateLimitResponse{}, nil
}

// Resets the flow on a rate limit. Fails if the rate limit doesn't exist
func (k msgServer) ResetRateLimit(goCtx context.Context, msg *types.MsgResetRateLimit) (*types.MsgResetRateLimitResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := k.assertGovOrAdmin(ctx, msg.Authority); err != nil {
		return nil, err
	}

	if err := k.Keeper.ResetRateLimit(ctx, msg.Denom, msg.ChannelId); err != nil {
		return nil, err
	}

	return &types.MsgResetRateLimitResponse{}, nil
}

// Updates the module params
func (k msgServer) UpdateParams(goCtx context.Context, msg *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	if k.authority != msg.Authority {
		return nil, errorsmod.Wrapf(govtypes.ErrInvalidSigner, "invalid authority; expected %s, got %s", k.authority, msg.Authority)
	}

	err := k.Keeper.SetParams(ctx, msg.Params)
	if err != nil {
		return nil, err
	}

	return &types.MsgUpdateParamsResponse{}, nil
}

// Sets a whitelisted address pair
func (k msgServer) SetWhitelistedAddressPair(goCtx context.Context, msg *types.MsgSetWhitelistedAddressPair) (*types.MsgSetWhitelistedAddressPairResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := k.assertGovOrAdmin(ctx, msg.Authority); err != nil {
		return nil, err
	}

	k.Keeper.SetWhitelistedAddressPair(ctx,
		types.WhitelistedAddressPair{
			Sender:   msg.Sender,
			Receiver: msg.Receiver,
		},
	)

	return &types.MsgSetWhitelistedAddressPairResponse{}, nil
}

// Removes a whitelisted address pair
func (k msgServer) RemoveWhitelistedAddressPair(goCtx context.Context, msg *types.MsgRemoveWhitelistedAddressPair) (*types.MsgRemoveWhitelistedAddressPairResponse, error) {
	if err := msg.ValidateBasic(); err != nil {
		return nil, err
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	if err := k.assertGovOrAdmin(ctx, msg.Authority); err != nil {
		return nil, err
	}

	k.Keeper.RemoveWhitelistedAddressPair(ctx, msg.Sender, msg.Receiver)

	return &types.MsgRemoveWhitelistedAddressPairResponse{}, nil
}

func (k msgServer) assertGovOrAdmin(ctx sdk.Context, address string) error {
	if k.authority == address {
		return nil
	}
	admins := k.GetParams(ctx).Admins
	if !slices.Contains(admins, address) {
		return errortypes.ErrorInvalidSigner.Wrapf("invalid authority; expected gov or admin address, got %s", address)
	}
	return nil
}
