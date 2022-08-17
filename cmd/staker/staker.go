package staker

import (
	"github.com/oasysgames/oasys-pos-cli/cmd/constants"
	"github.com/spf13/cobra"
)

const (
	// prefix of subcommands
	cmdPrefix = "staker:"
)

var (
	commands = []*cobra.Command{
		stakeCmd,
		unstakeCmd,
	}
)

// AddCommand adds subcommands for staker to RootCommand.
func AddCommand(rootCmd *cobra.Command) {
	// Add common flags.
	for _, cmd := range commands {
		cmd.Flags().String(constants.NetworkFlag, "", "Choice \"mainnet\" or \"testnet\" or \"custom\"")
		cmd.Flags().String(constants.RpcFlag, "", " JSON-RPC URL of the Hub-Layer(L1). ")
		cmd.Flags().Int64(constants.ChainIdFlag, 0, "Chain ID of the Verse-Layer(L2).")
		cmd.Flags().String(constants.ValidatorFlag, "", "Address of the validator owner.")
		cmd.MarkFlagRequired(constants.ValidatorFlag)
		cmd.Flags().String(constants.AmountFlag, "", "Amount of tokens to stake (unit: Ether).")
		cmd.MarkFlagRequired(constants.AmountFlag)
	}

	// staker:stake
	rootCmd.AddCommand(stakeCmd)

	// staker:unstake
	rootCmd.AddCommand(unstakeCmd)
}
