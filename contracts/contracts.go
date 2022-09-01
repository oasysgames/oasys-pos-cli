package contracts

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/oasysgames/oasys-pos-cli/cmd/constants"
	"github.com/oasysgames/oasys-pos-cli/contracts/environment"
	"github.com/oasysgames/oasys-pos-cli/contracts/erc20"
	"github.com/oasysgames/oasys-pos-cli/contracts/stakemanager"
)

// Returns Environment contract.
func NewEnvironment(ec *ethclient.Client) (*environment.Environment, error) {
	instance, err := environment.NewEnvironment(constants.EnvironmentAddress, ec)
	if err != nil {
		return nil, err
	}
	return instance, nil
}

// Returns StakeManager contract.
func NewStakeManager(ec *ethclient.Client) (*stakemanager.Stakemanager, error) {
	instance, err := stakemanager.NewStakemanager(constants.StakemanagerAddress, ec)
	if err != nil {
		return nil, err
	}
	return instance, nil
}

// Returns ERC20 contract.
func NewERC20(ec *ethclient.Client, address common.Address) (*erc20.Erc20, error) {
	instance, err := erc20.NewErc20(address, ec)
	if err != nil {
		return nil, err
	}
	return instance, nil
}
