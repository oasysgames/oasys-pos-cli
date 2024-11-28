package validator

import (
	"context"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

	owners, operators, blsPublicKeys, stakes, candidates, err := cmdutils.GetValidators(ctx, stakemanager, epoch)
	if err != nil {
		utils.Fatal(err)
	}

	actives := make([]bool, len(owners))
	jails := make([]bool, len(owners))
	for i, owner := range owners {
		info, err := stakemanager.GetValidatorInfo(callOpts, owner, epoch)
		if err != nil {
			utils.Fatal(err)
		}
		actives[i] = info.Active
		jails[i] = info.Jailed
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoFormatHeaders(false)
	table.SetHeader([]string{
		"Owner",
		"Operator",
		"BLS Public Key",
		"Total Stake",
		"Status",
		"Candidate",
		"Jailed",
	})
	for i, owner := range owners {
		status := "active"
		if !actives[i] {
			status = "inactive"
		}

		table.Append([]string{
			owner.String(),
			operators[i].String(),
			hexutil.Encode(blsPublicKeys[i]),
			utils.FormatWei(stakes[i]),
			status,
			b2s(candidates[i]),
			b2s(jails[i]),
		})
	}
	table.Render()
}
