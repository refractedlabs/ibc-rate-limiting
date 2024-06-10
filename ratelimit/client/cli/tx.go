package cli

import (
	"fmt"
	"github.com/Stride-Labs/ibc-rate-limiting/ratelimit/types"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdAddRateLimit())
	cmd.AddCommand(CmdUpdateRateLimit())
	cmd.AddCommand(CmdRemoveRateLimit())
	cmd.AddCommand(CmdResetRateLimit())
	cmd.AddCommand(CmdSetWhitelistedAddressPair())
	cmd.AddCommand(CmdRemoveWhitelistedAddressPair())
	// this line is used by starport scaffolding # 1

	return cmd
}

func CmdAddRateLimit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add-rate-limit [denom] [channel-id] [max-percent-send] [max-percent-recv] [duration-hours]",
		Short: "Broadcast message add-rate-limit",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			maxPercentSend, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return fmt.Errorf("invalid int %s", args[2])
			}

			maxPercentRecv, ok := sdk.NewIntFromString(args[3])
			if !ok {
				return fmt.Errorf("invalid int %s", args[3])
			}

			durationHours, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}

			msg := &types.MsgAddRateLimit{
				Authority:      clientCtx.GetFromAddress().String(),
				Denom:          args[0],
				ChannelId:      args[1],
				MaxPercentSend: maxPercentSend,
				MaxPercentRecv: maxPercentRecv,
				DurationHours:  durationHours,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdUpdateRateLimit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-rate-limit [denom] [channel-id] [max-percent-send] [max-percent-recv] [duration-hours]",
		Short: "Broadcast message update-rate-limit",
		Args:  cobra.ExactArgs(5),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			maxPercentSend, ok := sdk.NewIntFromString(args[2])
			if !ok {
				return fmt.Errorf("invalid int %s", args[2])
			}

			maxPercentRecv, ok := sdk.NewIntFromString(args[3])
			if !ok {
				return fmt.Errorf("invalid int %s", args[3])
			}

			durationHours, err := cast.ToUint64E(args[4])
			if err != nil {
				return err
			}

			msg := &types.MsgUpdateRateLimit{
				Authority:      clientCtx.GetFromAddress().String(),
				Denom:          args[0],
				ChannelId:      args[1],
				MaxPercentSend: maxPercentSend,
				MaxPercentRecv: maxPercentRecv,
				DurationHours:  durationHours,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdRemoveRateLimit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-rate-limit [denom] [channel-id]",
		Short: "Broadcast message remove-rate-limit",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &types.MsgRemoveRateLimit{
				Authority: clientCtx.GetFromAddress().String(),
				Denom:     args[0],
				ChannelId: args[1],
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdResetRateLimit() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "reset-rate-limit [denom] [channel-id]",
		Short: "Broadcast message reset-rate-limit",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &types.MsgResetRateLimit{
				Authority: clientCtx.GetFromAddress().String(),
				Denom:     args[0],
				ChannelId: args[1],
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdSetWhitelistedAddressPair() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set-whitelisted-address-pair [sender] [receiver]",
		Short: "Broadcast message add-rate-limit",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &types.MsgSetWhitelistedAddressPair{
				Authority: clientCtx.GetFromAddress().String(),
				Sender:    args[0],
				Receiver:  args[1],
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func CmdRemoveWhitelistedAddressPair() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-whitelisted-address-pair [sender] [receiver]",
		Short: "Broadcast message remove-whitelisted-address-pair",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) (err error) {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := &types.MsgRemoveWhitelistedAddressPair{
				Authority: clientCtx.GetFromAddress().String(),
				Sender:    args[0],
				Receiver:  args[1],
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
