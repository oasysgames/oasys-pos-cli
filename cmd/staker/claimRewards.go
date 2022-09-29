package staker

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/spf13/cobra"

	"github.com/oasysgames/oasys-pos-cli/cmd/constants"
	cmdutils "github.com/oasysgames/oasys-pos-cli/cmd/utils"
	"github.com/oasysgames/oasys-pos-cli/contracts"
	"github.com/oasysgames/oasys-pos-cli/eth"
	"github.com/oasysgames/oasys-pos-cli/utils"
)

var claimRewardsCmd = &cobra.Command{
	Use:   cmdPrefix + "claim-rewards",
	Short: "Withdraw staking rewards.",
	Run: func(cmd *cobra.Command, args []string) {
		wallet, err := cmdutils.NewEthWallet(cmd)
		if err != nil {
			utils.Fatal(err)
		}

		validator, err := parseValidatorFlag(cmd)
		if err != nil {
			utils.Fatal(err)
		}

		doClaimRewards(wallet, validator)
	},
}

func doClaimRewards(wallet *eth.Wallet, validator common.Address) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(constants.RpcTimeout))
	defer cancel()

	txOpts, err := wallet.GetTransactOpts(ctx)
	if err != nil {
		utils.Fatal(err)
	}

	stakemanager, err := contracts.NewStakeManager(wallet.Client)
	if err != nil {
		utils.Fatal(err)
	}

	tx, err := stakemanager.ClaimRewards(txOpts, wallet.From, validator, common.Big0)
	if err != nil {
		utils.Fatal(err)
	}

	fmt.Printf("sending (tx: %s)...", tx.Hash().String())

	receipt, err := wallet.WaitForTransactionReceipt(ctx, tx.Hash())
	if err != nil {
		utils.Fatal(err)
	}

	fmt.Printf(": success with %d gas\n", receipt.GasUsed)
}
