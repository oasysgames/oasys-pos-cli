package staker

import (
	"errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/oasysgames/oasys-pos-cli/cmd/constants"
	cmdutils "github.com/oasysgames/oasys-pos-cli/cmd/utils"
)

const (
	// prefix of subcommands
	cmdPrefix = "staker:"

	// Token types
	OASty  = uint8(0)
	WOASty = uint8(1)
	SOASty = uint8(2)
)

var (
	tokenTypes = []uint8{OASty, WOASty, SOASty}
)

// AddCommand adds subcommands for staker to RootCommand.
func AddCommand(rootCmd *cobra.Command) {
	// Add common flags.
	for _, cmd := range []*cobra.Command{balanceCmd, stakesCmd, rewardsCmd} {
		cmd.Flags().String(constants.StakerFlag, "", "Address of the staker.")
	}
	for _, cmd := range []*cobra.Command{balanceCmd, stakesCmd, rewardsCmd, stakeCmd, unstakeCmd, claimRewardsCmd, claimUnstakesCmd} {
		cmd.Flags().String(constants.NetworkFlag, "", "Choice \"mainnet\" or \"testnet\" or \"custom\"")
		cmd.Flags().String(constants.RpcFlag, "", " JSON-RPC URL of the Hub-Layer(L1). ")
		cmd.Flags().Int64(constants.ChainIdFlag, 0, "Chain ID of the Verse-Layer(L2).")
	}
	for _, cmd := range []*cobra.Command{stakeCmd, unstakeCmd} {
		cmd.Flags().String(constants.ValidatorFlag, "", "Address of the validator.")
		cmd.MarkFlagRequired(constants.ValidatorFlag)

		cmd.Flags().Int64(constants.AmountFlag, 0, "Amount of OAS token to stake (unit: Ether).")
		cmd.Flags().Int64(constants.OASFlag, 0, "Amount of OAS token to stake (unit: Ether).")
		cmd.Flags().Int64(constants.WOASFlag, 0, "Amount of Wrapped OAS token to stake (unit: Ether).")
		cmd.Flags().Int64(constants.SOASFlag, 0, "Amount of Stakable OAS token to stake (unit: Ether).")
	}

	// staker:show-balance
	rootCmd.AddCommand(balanceCmd)

	// staker:show-stakes
	rootCmd.AddCommand(stakesCmd)
	stakesCmd.Flags().Bool(constants.NextEpochFlag, false, "Display the next epoch.")

	// staker:show-rewards
	rootCmd.AddCommand(rewardsCmd)

	// staker:stake
	rootCmd.AddCommand(stakeCmd)

	// staker:unstake
	rootCmd.AddCommand(unstakeCmd)

	// staker:claim-rewards
	rootCmd.AddCommand(claimRewardsCmd)
	claimRewardsCmd.Flags().String(constants.ValidatorFlag, "", "Address of the validator.")
	claimRewardsCmd.MarkFlagRequired(constants.ValidatorFlag)

	// staker:claim-unstakes
	rootCmd.AddCommand(claimUnstakesCmd)
}

func parseStakerFlag(cmd *cobra.Command) (common.Address, error) {
	var zero common.Address

	if flagv, err := cmd.Flags().GetString(constants.StakerFlag); err != nil {
		return zero, err
	} else if flagv != "" {
		return common.HexToAddress(flagv), nil
	} else if wallet, err := cmdutils.NewEthWallet(cmd); err == nil {
		return wallet.From, nil
	} else {
		return zero, errors.New("private key or staker address is required")
	}
}

func parseValidatorFlag(cmd *cobra.Command) (common.Address, error) {
	var zero common.Address

	if flagv, err := cmd.Flags().GetString(constants.ValidatorFlag); err != nil {
		return zero, err
	} else {
		return common.HexToAddress(flagv), nil
	}
}

func parseAmountFlag(cmd *cobra.Command) (uint8, int64, error) {
	var (
		tokenType uint8
		amount    int64
		err       error
	)
	if oas, _ := cmd.Flags().GetInt64(constants.OASFlag); oas > 0 {
		tokenType = OASty
		amount = oas
	} else if oas, _ := cmd.Flags().GetInt64(constants.AmountFlag); oas > 0 {
		tokenType = OASty
		amount = oas
	} else if woas, _ := cmd.Flags().GetInt64(constants.WOASFlag); woas > 0 {
		tokenType = WOASty
		amount = woas
	} else if soas, _ := cmd.Flags().GetInt64(constants.SOASFlag); soas > 0 {
		tokenType = SOASty
		amount = soas
	} else {
		err = errors.New("--oas or --woas or --soas required")
	}

	return tokenType, amount, err
}
