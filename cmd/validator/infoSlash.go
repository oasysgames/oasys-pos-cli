package validator

import (
	"context"
	"fmt"
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

const MAX_BACk_EPOCH = 1024

var infoSlashCmd = &cobra.Command{
	Use:   cmdPrefix + "info-slash",
	Short: "List of the number of slashings occurred in each recent epoch",
	Run: func(cmd *cobra.Command, args []string) {
		var validator common.Address
		if s, err := cmd.Flags().GetString(constants.ValidatorFlag); err != nil {
			utils.Fatal(err)
		} else if s != "" {
			validator = common.HexToAddress(s)
		} else {
			if wallet, err := cmdutils.NewEthWallet(cmd); err == nil {
				validator = wallet.From
			}
		}

		ec, err := cmdutils.NewEthClient(cmd)
		if err != nil {
			utils.Fatal(err)
		}

		backEpoch, err := cmd.Flags().GetUint16(constants.BackEpochFlag)
		if err != nil {
			utils.Fatal(err)
		}

		if backEpoch > MAX_BACk_EPOCH {
			utils.Fatal(fmt.Errorf("back-epoch must be less than or equal to %d", MAX_BACk_EPOCH))
		}

		doInfoSlash(ec, validator, backEpoch)
	},
}

func doInfoSlash(ec *ethclient.Client, validator common.Address, backEpoch uint16) {
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

	currentEpoch, err := environment.Epoch(callOpts)
	if err != nil {
		utils.Fatal(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetAutoFormatHeaders(false)
	table.SetHeader([]string{
		"ID",
		"Epoch",
		"Slash",
	})

	epoch := big.NewInt(0)
	for i := int(backEpoch) - 1; 0 <= i; i-- {
		epoch.Sub(currentEpoch, big.NewInt(int64(i)))
		result, err := stakemanager.GetBlockAndSlashes(callOpts, validator, epoch)
		if err != nil {
			utils.Fatal(err)
		}

		table.Append([]string{
			fmt.Sprintf("%d", int(backEpoch)-i),
			fmt.Sprintf("%d", epoch),
			fmt.Sprintf("%s", result.Slashes.String()),
		})
	}

	table.Render()
}
