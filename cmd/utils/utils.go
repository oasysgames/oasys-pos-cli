package utils

import (
	"errors"
	"fmt"
	"math/big"
	"os"

	"github.com/oasysgames/oasys-pos-cli/cmd/constants"
	"github.com/oasysgames/oasys-pos-cli/eth"
	"github.com/spf13/cobra"
)

// Create a wallet from environment variables and command line flags.
func NewWallet(cmd *cobra.Command) (*eth.Wallet, error) {
	var (
		rpc     string
		chainId int64
	)

	network, err := cmd.Flags().GetString(constants.NetworkFlag)
	if err != nil {
		return nil, err
	}

	switch network {
	case "mainnet":
		rpc = constants.MainnetRPC
		chainId = constants.MainnetChainId
	case "testnet":
		rpc = constants.TestnetRPC
		chainId = constants.TestnetChainId
	default:
		rpc, err = cmd.Flags().GetString(constants.RpcFlag)
		if err != nil {
			return nil, err
		}

		chainId, err = cmd.Flags().GetInt64(constants.ChainIdFlag)
		if err != nil {
			return nil, err
		}

		if rpc == "" || chainId <= 0 {
			errmsg := ""
			errmsg += fmt.Sprintf("Usage: %s --%s mainnet|testnet\n", cmd.Use, constants.NetworkFlag)
			errmsg += "  or\n"
			errmsg += fmt.Sprintf("Usage: %s --%s http://127.0.0.1:845/ --%s 12345", cmd.Use, constants.RpcFlag, constants.ChainIdFlag)
			return nil, errors.New(errmsg)
		}
	}

	privateKey := os.Getenv(constants.PrivateKeyEnv)
	if privateKey == "" {
		return nil, fmt.Errorf("set the private key in the \"%s\" environment variable", constants.PrivateKeyEnv)
	}

	wallet, err := eth.NewWallet(rpc, big.NewInt(chainId), privateKey)
	if err != nil {
		return nil, err
	}

	return wallet, nil
}
