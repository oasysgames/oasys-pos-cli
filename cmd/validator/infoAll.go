package validator

import (
	"context"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"

	"github.com/oasysgames/oasys-pos-cli/cmd/constants"
	cmdutils "github.com/oasysgames/oasys-pos-cli/cmd/utils"
	"github.com/oasysgames/oasys-pos-cli/contracts"
	"github.com/oasysgames/oasys-pos-cli/utils"
)

var infoAllCmd = &cobra.Command{
	Use:   cmdPrefix + "info-all",
	Short: "Show all validator information.",
	Run: func(cmd *cobra.Command, args []string) {
		ec, err := cmdutils.NewEthClient(cmd)
		if err != nil {
			utils.Fatal(err)
		}

		showNextEpoch, err := cmd.Flags().GetBool(constants.NextEpochFlag)
		if err != nil {
			utils.Fatal(err)
		}

		doInfoAll(ec, showNextEpoch)
	},
}

func doInfoAll(ec *ethclient.Client, showNextEpoch bool) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(constants.RpcTimeout))
	defer cancel()

	callOpts := &bind.CallOpts{Context: ctx}

	environment, err := contracts.NewEnvironment(ec)
	if err != nil {
		utils.Fatal(err)
	}

	stakemanager, err := contracts.NewStakeManager(ec)
	if err != nil {
		utils.Fatal(err)
	}

	epoch, err := environment.Epoch(callOpts)
	if err != nil {
		utils.Fatal(err)
	}
	if showNextEpoch {
		epoch.Add(epoch, common.Big1)
	}

	validators, _, _, _, err := cmdutils.GetValidators(ctx, stakemanager, epoch)
	if err != nil {
		utils.Fatal(err)
	}

	type info struct {
		operator common.Address
		active   bool
		jailed   bool
		stakes   *big.Int
	}

	infos := map[common.Address]*info{}
	for _, validator := range validators {
		result1, err := stakemanager.GetValidatorInfo(callOpts, validator, epoch)
		if err != nil {
			utils.Fatal(err)
		}
		infos[validator] = &info{
			active: result1.Active,
			jailed: result1.Jailed,
			stakes: result1.Stakes,
		}

		result2, err := stakemanager.Validators(callOpts, validator)
		if err != nil {
			utils.Fatal(err)
		}
		infos[validator].operator = result2.Operator
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoFormatHeaders(false)
	table.SetHeader([]string{
		"Owner",
		"Operator",
		"Total Stake",
		"Status",
		"Jailed",
	})
	for _, owner := range validators {
		info := infos[owner]

		status := "active"
		if !info.active {
			status = "inactive"
		}

		table.Append([]string{
			owner.String(),
			info.operator.String(),
			utils.FormatWei(info.stakes),
			status,
			b2s(info.jailed),
		})
	}
	table.Render()
}
