package validator

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/oasysgames/oasys-pos-cli/eth"
	"github.com/spf13/cobra"
)

var updateCommissionRateCmd = &cobra.Command{
	Use:   cmdPrefix + "update-commission-rate",
	Short: "Update validator commission rates.",
	Run: func(cmd *cobra.Command, args []string) {
		rate, err := cmd.Flags().GetInt64(rateFlag)
		if err != nil {
			fatal(err)
		} else if rate < 0 || rate > 100 {
			fatal(errors.New("commission rate must be 0% to 100%"))
		}
		doUpdateCommissionRate(getWallet(cmd), rate)
	},
}

func doUpdateCommissionRate(wallet *eth.Wallet, rate int64) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(rpcTimeout))
	defer cancel()

	txOpts, err := wallet.GetTransactOpts(ctx)
	if err != nil {
		fatal(err)
	}

	stakermanager := getStakerManager(wallet.Client)
	tx, err := stakermanager.UpdateCommissionRate(txOpts, big.NewInt(rate))
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
