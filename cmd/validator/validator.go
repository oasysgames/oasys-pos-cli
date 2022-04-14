package validator

import (
	"errors"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	genesis "github.com/oasysgames/oasys-pos-cli/contracts/genesis"
	"github.com/oasysgames/oasys-pos-cli/eth"
	"github.com/oasysgames/oasys-pos-cli/ethutil"
	"github.com/spf13/cobra"
)

const (
	// prefix of subcommands
	cmdPrefix = "validator:"
	// environment variables
	privateKeyEnv = "PRIVATE_KEY"
	// command-line flags
	networkFlag   = "network"
	rpcFlag       = "rpc"
	chainIdFlag   = "chain-id"
	validatorFlag = "validator"
	operatorFlag  = "operator"
	rateFlag      = "rate"
	// Hub-Layer(L1)
	mainnetRPC     = "https://rpc.mainnet.oasys.games/"
	testnetRPC     = "https://rpc.testnet.oasys.games/"
	mainnetChainId = 248
	testnetChainId = 9372
)

var (
	cmdName string

	// genesis contract addresses
	environmentAddress  = common.HexToAddress("0x0000000000000000000000000000000000001000")
	stakemanagerAddress = common.HexToAddress("0x0000000000000000000000000000000000001001")

	rpcTimeout = 30 * time.Second

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
	cmdName = rootCmd.Use

	// Add common flags.
	for _, cmd := range commands {
		cmd.Flags().String(networkFlag, "", "Choice \"mainnet\" or \"testnet\" or \"custom\"")
		cmd.Flags().String(rpcFlag, "", " JSON-RPC URL of the Hub-Layer(L1). ")
		cmd.Flags().Int64(chainIdFlag, 0, "Chain ID of the Verse-Layer(L2).")
	}

	// validator:join
	rootCmd.AddCommand(joinCmd)
	joinCmd.Flags().String(operatorFlag, "", "Address used for block signing.")
	joinCmd.MarkFlagRequired(operatorFlag)

	// validator:info
	rootCmd.AddCommand(infoCmd)
	infoCmd.Flags().String(validatorFlag, "", "Address of the validator owner. Default is transaction sender.")

	// validator:activate
	rootCmd.AddCommand(activateCmd)
	activateCmd.Flags().String(validatorFlag, "", "Address of the validator. Default is transaction sender.")

	// validator:deactivate
	rootCmd.AddCommand(deactivateCmd)
	deactivateCmd.Flags().String(validatorFlag, "", "Address of the validator. Default is transaction sender.")

	// validator:update-commission-rate
	rootCmd.AddCommand(updateCommissionRateCmd)
	updateCommissionRateCmd.Flags().Int64(rateFlag, 0, "New commission rates(0% ~ 100%). Default is 0.")
	updateCommissionRateCmd.MarkFlagRequired(rateFlag)

	// validator:claim-commissions
	rootCmd.AddCommand(claimCommissionsCmd)
	claimCommissionsCmd.Flags().String(validatorFlag, "", "Address of the validator.")
}

// Returns Environment contract.
func getEnvironment(ec *ethclient.Client) *genesis.Environment {
	instance, err := genesis.NewEnvironment(environmentAddress, ec)
	if err != nil {
		fatal(err)
	}
	return instance
}

// Returns StakerManager contract.
func getStakerManager(ec *ethclient.Client) *genesis.Stakemanager {
	instance, err := genesis.NewStakemanager(stakemanagerAddress, ec)
	if err != nil {
		fatal(err)
	}
	return instance
}

// Create a wallet from environment variables and command line flags.
func getWallet(cmd *cobra.Command) *eth.Wallet {
	var (
		rpc     string
		chainId int64
		err     error
	)

	network, err := cmd.Flags().GetString(networkFlag)
	if err != nil {
		fatal(err)
	}

	switch network {
	case "mainnet":
		rpc = mainnetRPC
		chainId = mainnetChainId
	case "testnet":
		rpc = testnetRPC
		chainId = testnetChainId
	default:
		rpc, err = cmd.Flags().GetString(rpcFlag)
		if err != nil {
			fatal(err)
		}
		chainId, err = cmd.Flags().GetInt64(chainIdFlag)
		if err != nil {
			fatal(err)
		}
		if rpc == "" || chainId <= 0 {
			errmsg := ""
			errmsg += fmt.Sprintf("Usage: %s --%s mainnet|testnet\n", cmdName, networkFlag)
			errmsg += "  or\n"
			errmsg += fmt.Sprintf("Usage: %s --%s http://127.0.0.1:845/ --%s 12345", cmdName, rpcFlag, chainIdFlag)
			fatal(errors.New(errmsg))
		}
	}

	privateKey := os.Getenv(privateKeyEnv)
	if privateKey == "" {
		fatal(fmt.Errorf("set the private key in the \"%s\" environment variable", privateKeyEnv))
	}

	wallet, err := eth.NewWallet(rpc, big.NewInt(chainId), privateKey)
	if err != nil {
		fatal(err)
	}

	return wallet
}

// Adjust units and separate by 3 digits.
func formatWei(amount *big.Int) string {
	if amount.Cmp(ethutil.Ether) > 0 {
		return ethutil.ToDigit(ethutil.ToEther(amount)) + " OAS"
	} else if amount.Cmp(ethutil.GWei) > 0 {
		return ethutil.ToDigit(ethutil.ToGWei(amount)) + " GWei"
	}
	return ethutil.ToDigit(amount) + " Wei"
}

func fatal(err error) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
