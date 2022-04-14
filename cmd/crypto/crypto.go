package crypto

import "github.com/spf13/cobra"

const (
	cmdPrefix = "crypto:"
)

var (
	// command-line flags
	outputFlag = "out"
)

// AddCommand adds subcommands for crypto utility to RootCommand.
func AddCommand(rootCmd *cobra.Command) {
	// crypto:create-account
	rootCmd.AddCommand(createAccountCmd)
	createAccountCmd.Flags().String(outputFlag, "", "Output file for the created account. Default is stdout.")
}
