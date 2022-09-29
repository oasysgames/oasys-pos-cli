package utils

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/oasysgames/oasys-pos-cli/contracts/stakemanager"
)

func GetValidators(
	ctx context.Context,
	stakemanager *stakemanager.Stakemanager,
	epoch *big.Int,
) (
	owners []common.Address,
	operators []common.Address,
	stakes []*big.Int,
	candidates []bool,
	err error,
) {
	var (
		callOpts = &bind.CallOpts{Context: ctx}
		cursor   = big.NewInt(0)
		howMany  = big.NewInt(200)
	)
	for {
		result, err := stakemanager.GetValidators(callOpts, epoch, cursor, howMany)
		if err != nil {
			return nil, nil, nil, nil, err
		}
		if len(result.Owners) == 0 {
			break
		}

		cursor = result.NewCursor
		owners = append(owners, result.Owners...)
		operators = append(operators, result.Operators...)
		stakes = append(stakes, result.Stakes...)
		candidates = append(candidates, result.Candidates...)
	}

	return
}
