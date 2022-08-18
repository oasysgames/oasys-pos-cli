package validator

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cobra"

	"github.com/oasysgames/oasys-pos-cli/cmd/constants"
	cmdutils "github.com/oasysgames/oasys-pos-cli/cmd/utils"
	"github.com/oasysgames/oasys-pos-cli/contracts"
	"github.com/oasysgames/oasys-pos-cli/utils"
)

var infoCmd = &cobra.Command{
	Use:   cmdPrefix + "info",
	Short: "Show validator information.",
	Run: func(cmd *cobra.Command, args []string) {
		validator, err := cmd.Flags().GetString(constants.ValidatorFlag)
		if err != nil {
			utils.Fatal(err)
		}

		var to common.Address
		if validator != "" {
			to = common.HexToAddress(validator)
		} else {
			wallet, err := cmdutils.NewEthWallet(cmd)
			if err == nil {
				to = wallet.From
			}
		}

		if to == (common.Address{}) {
			utils.Fatal(errors.New("private key or validator address is required"))
		}

		ec, err := cmdutils.NewEthClient(cmd)
		if err != nil {
			utils.Fatal(err)
		}

		doInfo(ec, to)
	},
}

func doInfo(ec *ethclient.Client, to common.Address) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(constants.RpcTimeout))
	defer cancel()

	callOpts := &bind.CallOpts{Context: ctx}

	environment, err := contracts.NewEnvironment(ec)
	if err != nil {
		utils.Fatal(err)
	}

	stakermanager, err := contracts.NewStakeManager(ec)
	if err != nil {
		utils.Fatal(err)
	}

	balance, err := ec.BalanceAt(ctx, to, nil)
	if err != nil {
		utils.Fatal(err)
	}

	currentEpoch, err := environment.Epoch(callOpts)
	if err != nil {
		utils.Fatal(err)
	}
	nextEpoch := new(big.Int).Add(currentEpoch, common.Big1)

	currentInfo, err := stakermanager.GetValidatorInfo(callOpts, to)
	if err != nil {
		utils.Fatal(err)
	}

	nextInfo, err := stakermanager.GetValidatorInfo0(callOpts, to, nextEpoch)
	if err != nil {
		utils.Fatal(err)
	}

	commissions, err := stakermanager.GetCommissions(callOpts, to, common.Big0)
	if err != nil {
		utils.Fatal(err)
	}

	status := "active"
	if !currentInfo.Active {
		status = "inactive"
	}

	joined := currentInfo.Operator != common.Address{}

	fmt.Printf("%s : %s\n", rightPad("Balance"), utils.FormatWei(balance))
	fmt.Printf("%s : %s\n", rightPad("Joined"), b2s(joined))
	fmt.Printf("%s : %s\n", rightPad("Status"), status)
	fmt.Printf("%s : %s\n", rightPad("Jailed"), b2s(currentInfo.Jailed))
	fmt.Printf("%s : %s\n", rightPad("Operator Address"), currentInfo.Operator.String())
	fmt.Printf("%s : %s\n", rightPad("Commissions"), utils.FormatWei(commissions))
	fmt.Printf("%s : %s\n", rightPad("Current Epoch Staking"), utils.FormatWei(currentInfo.Stakes))
	fmt.Printf("%s : %s %%\n", rightPad("Current Epoch Commission Rate"), currentInfo.CommissionRate.String())
	fmt.Printf("%s : %s\n", rightPad("Next Epoch Staking"), utils.FormatWei(nextInfo.Stakes))
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
