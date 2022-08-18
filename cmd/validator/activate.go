package validator

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

var activateCmd = &cobra.Command{
	Use:   cmdPrefix + "activate",
	Short: "Change the validator status to active.",
	Run: func(cmd *cobra.Command, args []string) {
		validator, err := cmd.Flags().GetString(constants.ValidatorFlag)
		if err != nil {
			utils.Fatal(err)
		}

		wallet, err := cmdutils.NewEthWallet(cmd)
		if err != nil {
			utils.Fatal(err)
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
		utils.Fatal(err)
	}

	stakermanager, err := contracts.NewStakeManager(wallet.Client)
	if err != nil {
		utils.Fatal(err)
	}

	currentEpoch, err := environment.Epoch(wallet.GetCallOpts(ctx))
	if err != nil {
		utils.Fatal(err)
	}
	nextEpoch := new(big.Int).Add(currentEpoch, common.Big1)

	to := common.HexToAddress(validator)
	result, err := stakermanager.GetValidatorInfo(wallet.GetCallOpts(ctx), to)
	if err != nil {
		utils.Fatal(err)
	} else if result.Active {
		utils.Fatal(errors.New("already active"))
	}

	txOpts, err := wallet.GetTransactOpts(ctx)
	if err != nil {
		utils.Fatal(err)
	}
	tx, err := stakermanager.ActivateValidator(txOpts, to, []*big.Int{nextEpoch})
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
