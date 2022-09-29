package validator

import (
	"github.com/spf13/cobra"

	"github.com/oasysgames/oasys-pos-cli/cmd/constants"
)

const (
	// prefix of subcommands
	cmdPrefix = "validator:"
)

var (
	commands = []*cobra.Command{
		joinCmd,
		infoCmd,
		infoAllCmd,
		activateCmd,
		deactivateCmd,
		updateOperatorCmd,
		claimCommissionsCmd,
	}
)

// AddCommand adds subcommands for validator to RootCommand.
func AddCommand(rootCmd *cobra.Command) {
	// Add common flags.
	for _, cmd := range commands {
		cmd.Flags().String(constants.NetworkFlag, "", "Choice \"mainnet\" or \"testnet\" or \"custom\"")
		cmd.Flags().String(constants.RpcFlag, "", " JSON-RPC URL of the Hub-Layer(L1). ")
		cmd.Flags().Int64(constants.ChainIdFlag, 0, "Chain ID of the Verse-Layer(L2).")
	}

	// validator:join
	rootCmd.AddCommand(joinCmd)
	joinCmd.Flags().String(constants.OperatorFlag, "", "Address used for block signing.")
	joinCmd.MarkFlagRequired(constants.OperatorFlag)

	// validator:info
	rootCmd.AddCommand(infoCmd)
	infoCmd.Flags().String(constants.ValidatorFlag, "", "Address of the validator owner. Default is transaction sender.")

	// validator:info-all
	rootCmd.AddCommand(infoAllCmd)
	infoAllCmd.Flags().Bool(constants.NextEpochFlag, false, "Display the next epoch.")

	// validator:activate
	rootCmd.AddCommand(activateCmd)
	activateCmd.Flags().String(constants.ValidatorFlag, "", "Address of the validator. Default is transaction sender.")

	// validator:deactivate
	rootCmd.AddCommand(deactivateCmd)
	deactivateCmd.Flags().String(constants.ValidatorFlag, "", "Address of the validator. Default is transaction sender.")

	// validator:update-operator
	rootCmd.AddCommand(updateOperatorCmd)
	updateOperatorCmd.Flags().String(constants.OperatorFlag, "", "New address used for block signing.")
	updateOperatorCmd.MarkFlagRequired(constants.OperatorFlag)

	// validator:claim-commissions
	rootCmd.AddCommand(claimCommissionsCmd)
	claimCommissionsCmd.Flags().String(constants.ValidatorFlag, "", "Address of the validator.")
}
