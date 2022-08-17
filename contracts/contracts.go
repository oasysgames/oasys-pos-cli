package contracts

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/oasysgames/oasys-pos-cli/cmd/constants"
	"github.com/oasysgames/oasys-pos-cli/contracts/environment"
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
