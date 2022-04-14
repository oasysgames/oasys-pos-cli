package validator

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oasysgames/oasys-pos-cli/eth"
	"github.com/spf13/cobra"
)

var activateCmd = &cobra.Command{
	Use:   cmdPrefix + "activate",
	Short: "Change the validator status to active.",
	Run: func(cmd *cobra.Command, args []string) {
		validator, err := cmd.Flags().GetString(validatorFlag)
		if err != nil {
			fatal(err)
		}
		doActivate(getWallet(cmd), validator)
	},
}

func doActivate(wallet *eth.Wallet, validator string) {
	if validator == "" {
		validator = wallet.From.String()
	}

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(rpcTimeout))
	defer cancel()

	to := common.HexToAddress(validator)
	stakermanager := getStakerManager(wallet.Client)

	result, err := stakermanager.GetValidatorInfo(wallet.GetCallOpts(ctx), to)
	if err != nil {
		fatal(err)
	} else if result.Active {
		fatal(errors.New("already active"))
	}

	txOpts, err := wallet.GetTransactOpts(ctx)
	if err != nil {
		fatal(err)
	}
	tx, err := stakermanager.ActivateValidator(txOpts, to)
	if err != nil {
		fatal(err)
	}

	fmt.Printf("sending (tx: %s)...", tx.Hash().String())

	receipt, err := wallet.WaitForTransactionReceipt(ctx, tx.Hash())
	if err != nil {
		fatal(err)
	}

	fmt.Printf(": success with %d gas\n", receipt.GasUsed)
}
