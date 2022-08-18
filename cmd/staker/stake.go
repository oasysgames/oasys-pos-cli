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
	"github.com/oasysgames/oasys-pos-cli/eth"
	"github.com/oasysgames/oasys-pos-cli/utils"
)

var stakeCmd = &cobra.Command{
	Use:   cmdPrefix + "stake",
	Short: "Stake tokens to validator.",
	Run: func(cmd *cobra.Command, args []string) {
		validator, err := cmd.Flags().GetString(constants.ValidatorFlag)
		if err != nil {
			utils.Fatal(err)
		}

		amount, err := cmd.Flags().GetString(constants.AmountFlag)
		if err != nil {
			utils.Fatal(err)
		}

		wallet, err := cmdutils.NewEthWallet(cmd)
		if err != nil {
			utils.Fatal(err)
		}

		doStake(wallet, validator, amount)
	},
}

func doStake(wallet *eth.Wallet, validator string, amount string) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(constants.RpcTimeout))
	defer cancel()

	txOpts, err := wallet.GetTransactOpts(ctx)
	if err != nil {
		utils.Fatal(err)
	}

	stakemanager, err := contracts.NewStakeManager(wallet.Client)
	if err != nil {
		utils.Fatal(err)
	}

	to := common.HexToAddress(validator)
	result, err := stakemanager.GetValidatorInfo(wallet.GetCallOpts(ctx), to)
	if err != nil {
		utils.Fatal(err)
	}

	if result.Operator == (common.Address{}) {
		utils.Fatal(errors.New("validator is not join"))
	}

	token := uint8(0)
	bamount, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		utils.Fatal(err)
	}

	txOpts.Value = new(big.Int).Mul(bamount, utils.Ether)

	balance, err := wallet.Client.BalanceAt(ctx, to, nil)
	if err != nil {
		utils.Fatal(err)
	}

	if balance.Cmp(txOpts.Value) == -1 {
		utils.Fatal(errors.New("insufficient funds"))
	}

	tx, err := stakemanager.Stake(txOpts, to, token, txOpts.Value)
	if err != nil {
		utils.Fatal(err)
	}

	fmt.Printf("sending (tx: %s)...", tx.Hash().String())

	receipt, err := wallet.WaitForTransactionReceipt(ctx, tx.Hash())
	if err != nil {
		utils.Fatal(err)
	}

	fmt.Printf(": success with %d gas\n", receipt.GasUsed)
}
