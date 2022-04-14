package validator

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/oasysgames/oasys-pos-cli/eth"
	"github.com/spf13/cobra"
)

var joinCmd = &cobra.Command{
	Use:   cmdPrefix + "join",
	Short: "Join as a validator in the proof-of-stake.",
	Run: func(cmd *cobra.Command, args []string) {
		operator, err := cmd.Flags().GetString(operatorFlag)
		if err != nil {
			fatal(err)
		}
		doJoin(getWallet(cmd), operator)
	},
}

func doJoin(wallet *eth.Wallet, operator string) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(rpcTimeout))
	defer cancel()

	txOpts, err := wallet.GetTransactOpts(ctx)
	if err != nil {
		fatal(err)
	}

	stakermanager := getStakerManager(wallet.Client)
	tx, err := stakermanager.JoinValidator(txOpts, common.HexToAddress(operator))
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
