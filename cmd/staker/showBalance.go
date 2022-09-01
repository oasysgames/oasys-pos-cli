package staker

import (
	"context"
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

var balanceCmd = &cobra.Command{
	Use:   cmdPrefix + "show-balance",
	Short: "Show token balance.",
	Run: func(cmd *cobra.Command, args []string) {
		staker, err := parseStakerFlag(cmd)
		if err != nil {
			utils.Fatal(err)
		}

		ec, err := cmdutils.NewEthClient(cmd)
		if err != nil {
			utils.Fatal(err)
		}

		doBalance(ec, staker)
	},
}

func doBalance(ec *ethclient.Client, staker common.Address) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(constants.RpcTimeout))
	defer cancel()

	callOpts := &bind.CallOpts{Context: ctx}

	woas, err := contracts.NewERC20(ec, constants.WOASAddress)
	if err != nil {
		utils.Fatal(err)
	}

	soas, err := contracts.NewERC20(ec, constants.SOASAddress)
	if err != nil {
		utils.Fatal(err)
	}

	oasBalance, err := ec.BalanceAt(ctx, staker, nil)
	if err != nil {
		utils.Fatal(err)
	}

	woasBalance, err := woas.BalanceOf(callOpts, staker)
	if err != nil {
		utils.Fatal(err)
	}

	soasBalance, err := soas.BalanceOf(callOpts, staker)
	if err != nil {
		utils.Fatal(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoFormatHeaders(false)
	table.SetHeader([]string{"OAS", "WOAS", "SOAS"})
	table.Append([]string{
		utils.FormatWei(oasBalance),
		utils.FormatWei(woasBalance),
		utils.FormatWei(soasBalance),
	})
	table.Render()
}
