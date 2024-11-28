package validator

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
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

		if validator == (common.Address{}) {
			utils.Fatal(errors.New("private key or validator address is required"))
		}

		ec, err := cmdutils.NewEthClient(cmd)
		if err != nil {
			utils.Fatal(err)
		}

		doInfo(ec, validator)
	},
}

func doInfo(ec *ethclient.Client, validator common.Address) {
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
	nextEpoch := new(big.Int).Add(currentEpoch, common.Big1)

	currentInfo, err := stakemanager.GetValidatorInfo(callOpts, validator, currentEpoch)
	if err != nil {
		utils.Fatal(err)
	}

	nextInfo, err := stakemanager.GetValidatorInfo(callOpts, validator, nextEpoch)
	if err != nil {
		utils.Fatal(err)
	}

	commissions, err := stakemanager.GetCommissions(callOpts, validator, common.Big0)
	if err != nil {
		utils.Fatal(err)
	}

	slashes, err := stakemanager.GetBlockAndSlashes(callOpts, validator, currentEpoch)
	if err != nil {
		utils.Fatal(err)
	}

	status := "active"
	if !currentInfo.Active {
		status = "inactive"
	}

	joined := currentInfo.Operator != common.Address{}

	fmt.Printf("%s : %s\n", rightPad("Joined"), b2s(joined))
	fmt.Printf("%s : %s\n", rightPad("Status"), status)
	fmt.Printf("%s : %s\n", rightPad("Candidate"), b2s(currentInfo.Candidate))
	fmt.Printf("%s : %s\n", rightPad("Jailed"), b2s(currentInfo.Jailed))
	fmt.Printf("%s : %s\n", rightPad("Operator Address"), currentInfo.Operator.String())
	fmt.Printf("%s : %s\n", rightPad("BLS Public Key"), hexutil.Encode(currentInfo.BlsPublicKey))
	fmt.Printf("%s : %s\n", rightPad("Commissions"), utils.FormatWei(commissions))
	fmt.Printf("%s : %s\n", rightPad("Current Epoch Staking"), utils.FormatWei(currentInfo.Stakes))
	fmt.Printf("%s : %s\n", rightPad("Next Epoch Staking"), utils.FormatWei(nextInfo.Stakes))
	fmt.Printf("%s : %s\n", rightPad("Slash"), slashes.Slashes.String())
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
