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

var unstakeCmd = &cobra.Command{
	Use:   cmdPrefix + "unstake",
	Short: "Unstake tokens from validator.",
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

		doUnstake(wallet, validator, tokenType, amount)
	},
}

func doUnstake(wallet *eth.Wallet, validator common.Address, tokenType uint8, amount int64) {
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

	result, err := stakemanager.GetValidatorInfo(wallet.GetCallOpts(ctx), validator, common.Big0)
	if err != nil {
		utils.Fatal(err)
	}
	if result.Operator == (common.Address{}) {
		utils.Fatal(errors.New("validator is not join"))
	}

	bamount := new(big.Int).Mul(utils.Ether, big.NewInt(amount))
	tx, err := stakemanager.Unstake(txOpts, validator, tokenType, bamount)
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
