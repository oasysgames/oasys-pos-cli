package validator

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/oasysgames/oasys-pos-cli/cmd/constants"
	"github.com/oasysgames/oasys-pos-cli/cmd/utils"
	"github.com/oasysgames/oasys-pos-cli/contracts"
	"github.com/oasysgames/oasys-pos-cli/eth"
	"github.com/oasysgames/oasys-pos-cli/util"
	"github.com/spf13/cobra"
)

var updateCommissionRateCmd = &cobra.Command{
	Use:   cmdPrefix + "update-commission-rate",
	Short: "Update validator commission rates.",
	Run: func(cmd *cobra.Command, args []string) {
		rate, err := cmd.Flags().GetInt64(constants.RateFlag)
		if err != nil {
			util.Fatal(err)
		} else if rate < 0 || rate > 100 {
			util.Fatal(errors.New("commission rate must be 0% to 100%"))
		}

		wallet, err := utils.NewWallet(cmd)
		if err != nil {
			util.Fatal(err)
		}

		doUpdateCommissionRate(wallet, rate)
	},
}

func doUpdateCommissionRate(wallet *eth.Wallet, rate int64) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(constants.RpcTimeout))
	defer cancel()

	txOpts, err := wallet.GetTransactOpts(ctx)
	if err != nil {
		util.Fatal(err)
	}

	stakermanager, err := contracts.NewStakeManager(wallet.Client)
	if err != nil {
		util.Fatal(err)
	}

	tx, err := stakermanager.UpdateCommissionRate(txOpts, big.NewInt(rate))
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
