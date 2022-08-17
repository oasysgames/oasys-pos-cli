package validator

import (
	"github.com/oasysgames/oasys-pos-cli/cmd/constants"
	"github.com/spf13/cobra"
)

const (
	// prefix of subcommands
	cmdPrefix = "validator:"
)

var (
	commands = []*cobra.Command{
		joinCmd,
		infoCmd,
		activateCmd,
		deactivateCmd,
		updateCommissionRateCmd,
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

	// validator:activate
	rootCmd.AddCommand(activateCmd)
	activateCmd.Flags().String(constants.ValidatorFlag, "", "Address of the validator. Default is transaction sender.")

	// validator:deactivate
	rootCmd.AddCommand(deactivateCmd)
	deactivateCmd.Flags().String(constants.ValidatorFlag, "", "Address of the validator. Default is transaction sender.")

	// validator:update-commission-rate
	rootCmd.AddCommand(updateCommissionRateCmd)
	updateCommissionRateCmd.Flags().Int64(constants.RateFlag, 0, "New commission rates(0% ~ 100%). Default is 0.")
	updateCommissionRateCmd.MarkFlagRequired(constants.RateFlag)

	// validator:claim-commissions
	rootCmd.AddCommand(claimCommissionsCmd)
	claimCommissionsCmd.Flags().String(constants.ValidatorFlag, "", "Address of the validator.")
}
