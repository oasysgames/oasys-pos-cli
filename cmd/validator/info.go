package validator

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/oasysgames/oasys-pos-cli/cmd/constants"
	"github.com/oasysgames/oasys-pos-cli/cmd/utils"
	"github.com/oasysgames/oasys-pos-cli/contracts"
	"github.com/oasysgames/oasys-pos-cli/eth"
	"github.com/oasysgames/oasys-pos-cli/util"
)

var infoCmd = &cobra.Command{
	Use:   cmdPrefix + "info",
	Short: "Show validator information.",
	Run: func(cmd *cobra.Command, args []string) {
		validator, err := cmd.Flags().GetString(constants.ValidatorFlag)
		if err != nil {
			util.Fatal(err)
		}

		wallet, err := utils.NewWallet(cmd)
		if err != nil {
			util.Fatal(err)
		}

		doInfo(wallet, validator)
	},
}

func doInfo(wallet *eth.Wallet, validator string) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(constants.RpcTimeout))
	defer cancel()

	var to common.Address
	if validator != "" {
		to = common.HexToAddress(validator)
	} else {
		to = wallet.From
	}

	environment, err := contracts.NewEnvironment(wallet.Client)
	if err != nil {
		util.Fatal(err)
	}

	stakermanager, err := contracts.NewStakeManager(wallet.Client)
	if err != nil {
		util.Fatal(err)
	}

	balance, err := wallet.Client.BalanceAt(ctx, to, nil)
	if err != nil {
		util.Fatal(err)
	}

	callOpts := wallet.GetCallOpts(ctx)
	currentEpoch, err := environment.Epoch(callOpts)
	if err != nil {
		util.Fatal(err)
	}
	nextEpoch := new(big.Int).Add(currentEpoch, common.Big1)

	currentInfo, err := stakermanager.GetValidatorInfo(callOpts, to)
	if err != nil {
		util.Fatal(err)
	}

	nextInfo, err := stakermanager.GetValidatorInfo0(callOpts, to, nextEpoch)
	if err != nil {
		util.Fatal(err)
	}

	commissions, err := stakermanager.GetCommissions(callOpts, to, common.Big0)
	if err != nil {
		util.Fatal(err)
	}

	status := "active"
	if !currentInfo.Active {
		status = "inactive"
	}

	joined := currentInfo.Operator != common.Address{}

	fmt.Printf("%s : %s\n", rightPad("Balance"), util.FormatWei(balance))
	fmt.Printf("%s : %s\n", rightPad("Joined"), b2s(joined))
	fmt.Printf("%s : %s\n", rightPad("Status"), status)
	fmt.Printf("%s : %s\n", rightPad("Jailed"), b2s(currentInfo.Jailed))
	fmt.Printf("%s : %s\n", rightPad("Operator Address"), currentInfo.Operator.String())
	fmt.Printf("%s : %s\n", rightPad("Commissions"), util.FormatWei(commissions))
	fmt.Printf("%s : %s\n", rightPad("Current Epoch Staking"), util.FormatWei(currentInfo.Stakes))
	fmt.Printf("%s : %s %%\n", rightPad("Current Epoch Commission Rate"), currentInfo.CommissionRate.String())
	fmt.Printf("%s : %s\n", rightPad("Next Epoch Staking"), util.FormatWei(nextInfo.Stakes))
	fmt.Printf("%s : %s %%\n", rightPad("Next Epoch Commission Rate"), nextInfo.CommissionRate.String())
}

func rightPad(s string) string {
	return fmt.Sprintf("%-29s", s)
}

func b2s(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}
