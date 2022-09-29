package staker

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/oasysgames/oasys-pos-cli/cmd/constants"
	cmdutils "github.com/oasysgames/oasys-pos-cli/cmd/utils"
	"github.com/oasysgames/oasys-pos-cli/contracts"
	"github.com/oasysgames/oasys-pos-cli/contracts/erc20"
	"github.com/oasysgames/oasys-pos-cli/eth"
	"github.com/oasysgames/oasys-pos-cli/utils"
)

var stakeCmd = &cobra.Command{
	Use:   cmdPrefix + "stake",
	Short: "Stake tokens to validator.",
	Run: func(cmd *cobra.Command, args []string) {
		wallet, err := cmdutils.NewEthWallet(cmd)
		if err != nil {
			utils.Fatal(err)
		}

		validator, err := parseValidatorFlag(cmd)
		if err != nil {
			utils.Fatal(err)
		}

		tokenType, amount, err := parseAmountFlag(cmd)
		if err != nil {
			utils.Fatal(err)
		}

		doStake(wallet, validator, tokenType, amount)
	},
}

func doStake(wallet *eth.Wallet, validator common.Address, tokenType uint8, amount int64) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(constants.RpcTimeout))
	defer cancel()

	stakemanager, err := contracts.NewStakeManager(wallet.Client)
	if err != nil {
		utils.Fatal(err)
	}

	// check validator joined.
	result, err := stakemanager.GetValidatorInfo(wallet.GetCallOpts(ctx), validator, common.Big0)
	if err != nil {
		utils.Fatal(err)
	}
	if result.Operator == (common.Address{}) {
		utils.Fatal(errors.New("validator is not join"))
	}

	txOpts, err := wallet.GetTransactOpts(ctx)
	if err != nil {
		utils.Fatal(err)
	}

	bamount := new(big.Int).Mul(utils.Ether, big.NewInt(amount))
	switch tokenType {
	case OASty:
		txOpts.Value = bamount // send oas

		// check balance.
		if balance, err := wallet.Client.BalanceAt(ctx, wallet.From, nil); err != nil {
			utils.Fatal(err)
		} else if balance.Cmp(bamount) == -1 {
			utils.Fatal(errors.New("insufficient funds"))
		}
	default:
		if err := approve(wallet, tokenType, bamount); err != nil {
			utils.Fatal(err)
		}
	}

	ctx, cancel = context.WithDeadline(context.Background(), time.Now().Add(constants.RpcTimeout))
	defer cancel()

	tx, err := stakemanager.Stake(txOpts, validator, tokenType, bamount)
	if err != nil {
		utils.Fatal(err)
	}

	fmt.Printf("staking (tx: %s)...", tx.Hash().String())

	receipt, err := wallet.WaitForTransactionReceipt(ctx, tx.Hash())
	if err != nil {
		utils.Fatal(err)
	}

	fmt.Printf(": success with %d gas\n", receipt.GasUsed)
}

func getERC20(wallet *eth.Wallet, tokenType uint8) (*erc20.Erc20, error) {
	var address common.Address
	switch tokenType {
	case WOASty:
		address = constants.WOASAddress
	case SOASty:
		address = constants.SOASAddress
	}
	return contracts.NewERC20(wallet.Client, address)
}

func approve(wallet *eth.Wallet, tokenType uint8, amount *big.Int) error {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(constants.RpcTimeout))
	defer cancel()

	callOpts := wallet.GetCallOpts(ctx)

	erc20, err := getERC20(wallet, tokenType)
	if err != nil {
		return err
	}

	allowance, err := erc20.Allowance(callOpts, wallet.From, constants.StakemanagerAddress)
	if err != nil {
		return err
	} else if allowance.Cmp(amount) >= 0 {
		return nil
	}

	allowance = new(big.Int).Sub(amount, allowance)
	if balance, err := erc20.BalanceOf(callOpts, wallet.From); err != nil {
		return err
	} else if balance.Cmp(allowance) == -1 {
		return errors.New("insufficient funds")
	}

	txOpts, err := wallet.GetTransactOpts(ctx)
	if err != nil {
		return err
	}

	tx, err := erc20.Approve(txOpts, constants.StakemanagerAddress, allowance)
	if err != nil {
		return fmt.Errorf("failed to approve token transfer: %w", err)
	}

	fmt.Printf("approving (tx: %s)...", tx.Hash().String())

	receipt, err := wallet.WaitForTransactionReceipt(ctx, tx.Hash())
	if err != nil {
		return fmt.Errorf("failed to approve token transfer: %w", err)
	}

	fmt.Printf(": success with %d gas\n", receipt.GasUsed)

	return nil
}
