package staker

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

var stakesCmd = &cobra.Command{
	Use:   cmdPrefix + "show-stakes",
	Short: "Show stake information.",
	Run: func(cmd *cobra.Command, args []string) {
		staker, err := parseStakerFlag(cmd)
		if err != nil {
			utils.Fatal(err)
		}

		ec, err := cmdutils.NewEthClient(cmd)
		if err != nil {
			utils.Fatal(err)
		}

		showNextEpoch, err := cmd.Flags().GetBool(constants.NextEpochFlag)
		if err != nil {
			utils.Fatal(err)
		}

		doStakes(ec, staker, showNextEpoch)
	},
}

func doStakes(ec *ethclient.Client, staker common.Address, showNextEpoch bool) {
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

	epoch, err := environment.Epoch(callOpts)
	if err != nil {
		utils.Fatal(err)
	}
	if showNextEpoch {
		epoch.Add(epoch, common.Big1)
	}

	validators := []common.Address{}
	stakes := map[common.Address]map[uint8]*big.Int{}
	total := map[uint8]*big.Int{}
	for i, tt := range tokenTypes {
		if total[tt] == nil {
			total[tt] = new(big.Int)
		}

		result, err := stakermanager.GetStakerStakes(callOpts, staker, tt, epoch)
		if err != nil {
			utils.Fatal(err)
		}
		if i == 0 {
			validators = append(validators, result.Validators...)
		}

		for i, validator := range result.Validators {
			if stakes[validator] == nil {
				stakes[validator] = map[uint8]*big.Int{}
			}
			stakes[validator][tt] = result.Stakes[i]
			total[tt].Add(total[tt], result.Stakes[i])
		}
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoFormatHeaders(false)
	table.SetHeader([]string{"Validator", "OAS", "WOAS", "SOAS"})

	i := 0
	for {
		var (
			fc      func([]string)
			values  map[uint8]*big.Int
			columns []string
		)

		if i < len(validators) {
			fc = table.Append
			values = stakes[validators[i]]
			columns = append(columns, validators[i].String())
		} else if i == len(validators) {
			fc = table.SetFooter
			values = total
			columns = append(columns, "Total")
		} else {
			break
		}

		for _, tt := range tokenTypes {
			columns = append(columns, utils.FormatWei(values[tt]))
		}
		fc(columns)
		i++
	}
	table.SetFooterAlignment(tablewriter.ALIGN_LEFT)
	table.Render()
}
