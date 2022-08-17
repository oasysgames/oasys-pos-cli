package validator

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oasysgames/oasys-pos-cli/cmd/constants"
	"github.com/oasysgames/oasys-pos-cli/cmd/utils"
	"github.com/oasysgames/oasys-pos-cli/contracts"
	"github.com/oasysgames/oasys-pos-cli/eth"
	"github.com/oasysgames/oasys-pos-cli/util"
	"github.com/spf13/cobra"
)

var activateCmd = &cobra.Command{
	Use:   cmdPrefix + "activate",
	Short: "Change the validator status to active.",
	Run: func(cmd *cobra.Command, args []string) {
		validator, err := cmd.Flags().GetString(constants.ValidatorFlag)
		if err != nil {
			util.Fatal(err)
		}

		wallet, err := utils.NewWallet(cmd)
		if err != nil {
			util.Fatal(err)
		}

		doActivate(wallet, validator)
	},
}

func doActivate(wallet *eth.Wallet, validator string) {
	if validator == "" {
		validator = wallet.From.String()
	}

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(constants.RpcTimeout))
	defer cancel()

	environment, err := contracts.NewEnvironment(wallet.Client)
	if err != nil {
		util.Fatal(err)
	}

	stakermanager, err := contracts.NewStakeManager(wallet.Client)
	if err != nil {
		util.Fatal(err)
	}

	currentEpoch, err := environment.Epoch(wallet.GetCallOpts(ctx))
	if err != nil {
		util.Fatal(err)
	}
	nextEpoch := new(big.Int).Add(currentEpoch, common.Big1)

	to := common.HexToAddress(validator)
	result, err := stakermanager.GetValidatorInfo(wallet.GetCallOpts(ctx), to)
	if err != nil {
		util.Fatal(err)
	} else if result.Active {
		util.Fatal(errors.New("already active"))
	}

	txOpts, err := wallet.GetTransactOpts(ctx)
	if err != nil {
		util.Fatal(err)
	}
	tx, err := stakermanager.ActivateValidator(txOpts, to, []*big.Int{nextEpoch})
	if err != nil {
		util.Fatal(err)
	}

	fmt.Printf("sending (tx: %s)...", tx.Hash().String())

	receipt, err := wallet.WaitForTransactionReceipt(ctx, tx.Hash())
	if err != nil {
		util.Fatal(err)
	}

	fmt.Printf(": success with %d gas\n", receipt.GasUsed)
}
