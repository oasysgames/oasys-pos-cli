package validator

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/oasysgames/oasys-pos-cli/cmd/constants"
	cmdutils "github.com/oasysgames/oasys-pos-cli/cmd/utils"
	"github.com/oasysgames/oasys-pos-cli/contracts"
	"github.com/oasysgames/oasys-pos-cli/eth"
	"github.com/oasysgames/oasys-pos-cli/utils"
)

var updateOperatorCmd = &cobra.Command{
	Use:   cmdPrefix + "update-operator",
	Short: "Update the block signing address.",
	Run: func(cmd *cobra.Command, args []string) {
		operator, err := cmd.Flags().GetString(constants.OperatorFlag)
		if err != nil {
			utils.Fatal(err)
		}

		wallet, err := cmdutils.NewEthWallet(cmd)
		if err != nil {
			utils.Fatal(err)
		}

		doUpdateOperator(wallet, operator)
	},
}

func doUpdateOperator(wallet *eth.Wallet, operator string) {
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

	result, err := stakemanager.GetValidatorInfo(wallet.GetCallOpts(ctx), wallet.From)
	if err != nil {
		utils.Fatal(err)
	}

	operatorAddr := common.HexToAddress(operator)
	if operatorAddr == result.Operator {
		utils.Fatal(errors.New("already updated"))
	}

	tx, err := stakemanager.UpdateOperator(txOpts, operatorAddr)
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
