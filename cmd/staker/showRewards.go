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

var rewardsCmd = &cobra.Command{
	Use:   cmdPrefix + "show-rewards",
	Short: "Show balance of staking rewards.",
	Run: func(cmd *cobra.Command, args []string) {
		staker, err := parseStakerFlag(cmd)
		if err != nil {
			utils.Fatal(err)
		}

		ec, err := cmdutils.NewEthClient(cmd)
		if err != nil {
			utils.Fatal(err)
		}

		doRewards(ec, staker)
	},
}

func doRewards(ec *ethclient.Client, staker common.Address) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(constants.RpcTimeout))
	defer cancel()

	callOpts := &bind.CallOpts{Context: ctx}

	stakemanager, err := contracts.NewStakeManager(ec)
	if err != nil {
		utils.Fatal(err)
	}

	validators, _, _, _, err := cmdutils.GetValidators(ctx, stakemanager, common.Big0)
	if err != nil {
		utils.Fatal(err)
	}

	total := big.NewInt(0)
	rewards := map[common.Address]*big.Int{}
	for _, validator := range validators {
		reward, err := stakemanager.GetRewards(callOpts, staker, validator, common.Big0)
		if err != nil {
			utils.Fatal(err)
		}

		total.Add(total, reward)
		rewards[validator] = reward
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoFormatHeaders(false)
	table.SetHeader([]string{"Validator", "Reward"})
	for _, v := range validators {
		table.Append([]string{v.String(), utils.FormatWei(rewards[v])})
	}
	table.SetFooterAlignment(tablewriter.ALIGN_LEFT)
	table.SetFooter([]string{"Total", utils.FormatWei(total)})
	table.Render()
}
