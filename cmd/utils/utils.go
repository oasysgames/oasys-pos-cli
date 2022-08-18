package utils

import (
	"errors"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	"github.com/oasysgames/oasys-pos-cli/cmd/constants"
	"github.com/oasysgames/oasys-pos-cli/eth"
)

// Create a ethereum client from command line flags.
func NewEthClient(cmd *cobra.Command) (*ethclient.Client, error) {
	rpc, err := getRpc(cmd)
	if err != nil {
		return nil, err
	}

	if rpc == "" {
		msg := ""
		msg += fmt.Sprintf("Usage: %s --%s mainnet|testnet\n", cmd.Use, constants.NetworkFlag)
		msg += "  or\n"
		msg += fmt.Sprintf("Usage: %s --%s http://127.0.0.1:8545/", cmd.Use, constants.RpcFlag)
		return nil, errors.New(msg)
	}

	return ethclient.Dial(rpc)
}

// Create a wallet from environment variables and command line flags.
func NewEthWallet(cmd *cobra.Command) (*eth.Wallet, error) {
	rpc, err := getRpc(cmd)
	if err != nil {
		return nil, err
	}

	chainId, err := getChainId(cmd)
	if err != nil {
		return nil, err
	}

	if rpc == "" || chainId <= 0 {
		msg := ""
		msg += fmt.Sprintf("Usage: %s --%s mainnet|testnet\n", cmd.Use, constants.NetworkFlag)
		msg += "  or\n"
		msg += fmt.Sprintf("Usage: %s --%s http://127.0.0.1:8545/ --%s 12345", cmd.Use, constants.RpcFlag, constants.ChainIdFlag)
		return nil, errors.New(msg)
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

func getRpc(cmd *cobra.Command) (string, error) {
	network, err := cmd.Flags().GetString(constants.NetworkFlag)
	if err != nil {
		return "", err
	}

	switch network {
	case "mainnet":
		return constants.MainnetRPC, nil
	case "testnet":
		return constants.TestnetRPC, nil
	default:
		return cmd.Flags().GetString(constants.RpcFlag)
	}
}

func getChainId(cmd *cobra.Command) (int64, error) {
	network, err := cmd.Flags().GetString(constants.NetworkFlag)
	if err != nil {
		return 0, err
	}

	switch network {
	case "mainnet":
		return constants.MainnetChainId, nil
	case "testnet":
		return constants.TestnetChainId, nil
	default:
		return cmd.Flags().GetInt64(constants.ChainIdFlag)
	}
}
