package validator

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/oasysgames/oasys-pos-cli/eth"
)

var infoCmd = &cobra.Command{
	Use:   cmdPrefix + "info",
	Short: "Show validator information.",
	Run: func(cmd *cobra.Command, args []string) {
		validator, err := cmd.Flags().GetString(validatorFlag)
		if err != nil {
			fatal(err)
		}
		doInfo(getWallet(cmd), validator)
	},
}

func doInfo(wallet *eth.Wallet, validator string) {
	var to common.Address
	if validator != "" {
		to = common.HexToAddress(validator)
	} else {
		to = wallet.From
	}

	showInfo(wallet, to)
	showNextEpochStaking(wallet, to)
}

func showInfo(wallet *eth.Wallet, to common.Address) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(rpcTimeout))
	defer cancel()

	balance, err := wallet.Client.BalanceAt(ctx, to, nil)
	if err != nil {
		fatal(err)
	}

	stakermanager := getStakerManager(wallet.Client)
	result, err := stakermanager.GetValidatorInfo(wallet.GetCallOpts(ctx), to)
	if err != nil {
		fatal(err)
	}

	commissions, err := stakermanager.GetCommissions(wallet.GetCallOpts(ctx), to, common.Big0)
	if err != nil {
		fatal(err)
	}

	status := "active"
	if !result.Active {
		status = "inactive"
	}

	fmt.Printf("%s : %s\n", rightPad("Balance"), formatWei(balance))
	fmt.Printf("%s : %s\n", rightPad("Status"), status)
	fmt.Printf("%s : %s\n", rightPad("Operator Address"), result.Operator.String())
	fmt.Printf("%s : %s %%\n", rightPad("Commission Rate"), result.CommissionRate.String())
	fmt.Printf("%s : %s\n", rightPad("Commissions"), formatWei(commissions))
	fmt.Printf("%s : %s\n", rightPad("Jailed Epoch"), result.JailEpoch.String())
	fmt.Printf("%s : %s\n", rightPad("Current Epoch Staking"), formatWei(result.Stakes))
}

func showNextEpochStaking(wallet *eth.Wallet, to common.Address) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(rpcTimeout))
	defer cancel()

	callOpts := wallet.GetCallOpts(ctx)

	stakermanager := getStakerManager(wallet.Client)
	environment := getEnvironment(wallet.Client)

	currentEpoch, err := environment.Epoch(callOpts)
	if err != nil {
		fatal(err)
	}

	result, err := stakermanager.GetValidatorStakes(callOpts, to,
		new(big.Int).Add(currentEpoch, common.Big1), common.Big0, common.Big0)
	if err != nil {
		fatal(err)
	}

	total := new(big.Int)
	for i, stake := range result.Stakes {
		if result.Stakers[i] == (common.Address{}) {
			break
		}
		total = new(big.Int).Add(total, stake)
	}

	fmt.Printf("%s : %s\n", rightPad("Next Epoch Staking"), formatWei(total))
}

func rightPad(s string) string {
	return fmt.Sprintf("%-21s", s)
}
