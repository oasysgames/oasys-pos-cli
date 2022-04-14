package validator

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oasysgames/oasys-pos-cli/eth"
	"github.com/spf13/cobra"
)

var claimCommissionsCmd = &cobra.Command{
	Use:   cmdPrefix + "claim-commissions",
	Short: "Withdraw validator commissions.",
	Run: func(cmd *cobra.Command, args []string) {
		validator, err := cmd.Flags().GetString(validatorFlag)
		if err != nil {
			fatal(err)
		}
		doClaimCommissions(getWallet(cmd), validator)
	},
}

func doClaimCommissions(wallet *eth.Wallet, validator string) {
	if validator == "" {
		validator = wallet.From.String()
	}

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(rpcTimeout))
	defer cancel()

	txOpts, err := wallet.GetTransactOpts(ctx)
	if err != nil {
		fatal(err)
	}

	stakermanager := getStakerManager(wallet.Client)
	tx, err := stakermanager.ClaimCommissions(txOpts, common.HexToAddress(validator), common.Big0)
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
