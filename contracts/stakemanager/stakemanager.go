// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stakemanager

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StakemanagerMetaData contains all meta data concerning the Stakemanager contract.
var StakemanagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyJoined\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountMismatched\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotZeroAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyBlockProducer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotLastBlock\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OverRate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameAsOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakerDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"enumToken.Type\",\"name\":\"token\",\"type\":\"uint8\"}],\"name\":\"TransferFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedSender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedValidator\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnknownToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorDoesNotExist\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumToken.Type\",\"name\":\"token\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumToken.Type\",\"name\":\"token\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"epochs\",\"type\":\"uint256[]\"}],\"name\":\"ValidatorActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"ValidatorCommissionRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"epochs\",\"type\":\"uint256[]\"}],\"name\":\"ValidatorDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"until\",\"type\":\"uint256\"}],\"name\":\"ValidatorJailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorSlashed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"epochs\",\"type\":\"uint256[]\"}],\"name\":\"activateValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowlist\",\"outputs\":[{\"internalType\":\"contractIAllowlist\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epochs\",\"type\":\"uint256\"}],\"name\":\"claimCommissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epochs\",\"type\":\"uint256\"}],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimUnstakes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"epochs\",\"type\":\"uint256[]\"}],\"name\":\"deactivateValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"environment\",\"outputs\":[{\"internalType\":\"contractIEnvironment\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getBlockAndSlashes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epochs\",\"type\":\"uint256\"}],\"name\":\"getCommissions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"commissions\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"operators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"stakes\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"operators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"stakes\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epochs\",\"type\":\"uint256\"}],\"name\":\"getRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"enumToken.Type\",\"name\":\"token\",\"type\":\"uint8\"}],\"name\":\"getStakerInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stakes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"enumToken.Type\",\"name\":\"token\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getStakerStakes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"stakes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"stakeRequests\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"unstakeRequests\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"page\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"perPage\",\"type\":\"uint256\"}],\"name\":\"getStakers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochs\",\"type\":\"uint256\"}],\"name\":\"getTotalRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"getValidatorInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"jailed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"stakes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getValidatorInfo\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"jailed\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"stakes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"page\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"perPage\",\"type\":\"uint256\"}],\"name\":\"getValidatorStakes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_stakers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"stakes\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIEnvironment\",\"name\":\"_environment\",\"type\":\"address\"},{\"internalType\":\"contractIAllowlist\",\"name\":\"_allowlist\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"joinValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operatorToOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"blocks\",\"type\":\"uint256\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"enumToken.Type\",\"name\":\"token\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakerSigners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"enumToken.Type\",\"name\":\"token\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newRate\",\"type\":\"uint256\"}],\"name\":\"updateCommissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"updateOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validatorOwners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"lastClaimCommission\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50617ebe80620000216000396000f3fe6080604052600436106101ed5760003560e01c806374e2b63c1161010d578063b7ab4db5116100a0578063dbd61d871161006f578063dbd61d8714610747578063e1aca34114610784578063e4a8a343146107ad578063f65a5ed2146107eb578063fa52c7d814610828576101ed565b8063b7ab4db514610686578063c5f9dff0146106b1578063cbc0fac6146106de578063d1f18ee114610707576101ed565b80639168ae72116100dc5780639168ae72146105ba5780639a99b4f0146105f7578063ac7475ed14610620578063ad71bd3614610649576101ed565b806374e2b63c146104f55780637b520aa8146105205780637befa74f1461055d5780638a11d7c914610579576101ed565b806340cddab311610185578063533a339c11610154578063533a339c14610412578063552164ee146104525780635efc766e1461048f5780636b2b3369146104cc576101ed565b806340cddab314610355578063428e85621461038257806346dfce7b146103ab578063485cc955146103e9576101ed565b8063195afea1116101c1578063195afea11461029857806322226367146102d55780632b47da52146103135780632ffb688d1461033e576101ed565b8062fa3d50146101f257806302fb4d851461021b578063158ef93e146102445780631903cf161461026f575b600080fd5b3480156101fe57600080fd5b5061021960048036038101906102149190616c68565b610867565b005b34801561022757600080fd5b50610242600480360381019061023d9190616cf3565b6109f8565b005b34801561025057600080fd5b50610259610e4e565b6040516102669190616d4e565b60405180910390f35b34801561027b57600080fd5b5061029660048036038101906102919190616ec2565b610e5f565b005b3480156102a457600080fd5b506102bf60048036038101906102ba9190616cf3565b611254565b6040516102cc9190616f2d565b60405180910390f35b3480156102e157600080fd5b506102fc60048036038101906102f79190616cf3565b6112d6565b60405161030a929190616f48565b60405180910390f35b34801561031f57600080fd5b506103286113d8565b6040516103359190616fd0565b60405180910390f35b34801561034a57600080fd5b506103536113fe565b005b34801561036157600080fd5b5061036a61153c565b60405161037993929190617167565b60405180910390f35b34801561038e57600080fd5b506103a960048036038101906103a49190616ec2565b611689565b005b3480156103b757600080fd5b506103d260048036038101906103cd91906171b3565b611a7e565b6040516103e092919061721a565b60405180910390f35b3480156103f557600080fd5b50610410600480360381019061040b91906172cd565b611e72565b005b34801561041e57600080fd5b5061043960048036038101906104349190617332565b611fbc565b6040516104499493929190617385565b60405180910390f35b34801561045e57600080fd5b5061047960048036038101906104749190616c68565b612155565b6040516104869190616f2d565b60405180910390f35b34801561049b57600080fd5b506104b660048036038101906104b19190616c68565b6123b7565b6040516104c391906173f5565b60405180910390f35b3480156104d857600080fd5b506104f360048036038101906104ee9190617410565b6123f6565b005b34801561050157600080fd5b5061050a6125fd565b604051610517919061745e565b60405180910390f35b34801561052c57600080fd5b5061054760048036038101906105429190617410565b612623565b60405161055491906173f5565b60405180910390f35b61057760048036038101906105729190617332565b612656565b005b34801561058557600080fd5b506105a0600480360381019061059b9190617410565b612a58565b6040516105b1959493929190617479565b60405180910390f35b3480156105c657600080fd5b506105e160048036038101906105dc9190617410565b612bbd565b6040516105ee91906173f5565b60405180910390f35b34801561060357600080fd5b5061061e60048036038101906106199190616cf3565b612bfb565b005b34801561062c57600080fd5b5061064760048036038101906106429190617410565b612e48565b005b34801561065557600080fd5b50610670600480360381019061066b91906174cc565b612fe5565b60405161067d919061750c565b60405180910390f35b34801561069257600080fd5b5061069b613157565b6040516106a8919061750c565b60405180910390f35b3480156106bd57600080fd5b506106c66131e5565b6040516106d593929190617167565b60405180910390f35b3480156106ea57600080fd5b5061070560048036038101906107009190616cf3565b613326565b005b34801561071357600080fd5b5061072e60048036038101906107299190616cf3565b613597565b60405161073e949392919061752e565b60405180910390f35b34801561075357600080fd5b5061076e60048036038101906107699190617573565b61363f565b60405161077b9190616f2d565b60405180910390f35b34801561079057600080fd5b506107ab60048036038101906107a69190617332565b613702565b005b3480156107b957600080fd5b506107d460048036038101906107cf91906175c6565b613abe565b6040516107e2929190616f48565b60405180910390f35b3480156107f757600080fd5b50610812600480360381019061080d9190616c68565b613bef565b60405161081f91906173f5565b60405180910390f35b34801561083457600080fd5b5061084f600480360381019061084a9190617410565b613c2e565b60405161085e93929190617606565b60405180910390f35b33600073ffffffffffffffffffffffffffffffffffffffff16600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610931576040517fe51315d200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6109a6600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1683600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020613c989092919063ffffffff16565b3373ffffffffffffffffffffffffffffffffffffffff167f658f52b34d965c76f54298d43b21ae83217d77033214bdafed998bf6df6b9b05836040516109ec9190616f2d565b60405180910390a25050565b600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff16600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610b21576040517fe51315d200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610b86576040517f1cf4735900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060026000600460008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000610d62600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633fa4f2456040518163ffffffff1660e01b815260040161010060405180830381865afa158015610c9b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610cbf9190617721565b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610d2c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d50919061774f565b8685613d6f909392919063ffffffff16565b90508160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f1647efd0ce9727dc31dc201c9d8d35ac687f7370adcacbd454afc6485ddabfda60405160405180910390a26000811115610e47578160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167feb7d7a49847ec491969db21a0e31b234565a9923145a2d1b56a75c9e9582580282604051610e3e9190616f2d565b60405180910390a25b5050505050565b60008054906101000a900460ff1681565b816000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614158015610f5357508060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b15610f8a576040517f0809490800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b83600073ffffffffffffffffffffffffffffffffffffffff16600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415611054576040517fe51315d200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d4a536866040518163ffffffff1660e01b8152600401602060405180830381865afa1580156110c1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110e591906177a8565b1561111c576040517f1e59ccd900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6111ff600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561118c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111b0919061774f565b85600260008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020613e949092919063ffffffff16565b8473ffffffffffffffffffffffffffffffffffffffff167fc11dfc9c24621433bb10587dc4bbae26a33a4aff53914e0d4c9fddf224a8cb7d8560405161124591906177d5565b60405180910390a25050505050565b60006112cb600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1683600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020613ea69092919063ffffffff16565b508091505092915050565b6000806113c96000841161137a57600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611351573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611375919061774f565b61137c565b835b600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002061407e90919063ffffffff16565b80925081935050509250929050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600073ffffffffffffffffffffffffffffffffffffffff16600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156114c7576040517fcf83d93d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61153a600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206140b890919063ffffffff16565b565b6060806060611678600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ccfb99356040518163ffffffff1660e01b815260040161010060405180830381865afa1580156115b2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906115d69190617721565b6001600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611645573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611669919061774f565b6116739190617826565b6140e0565b809350819450829550505050909192565b816000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415801561177d57508060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b156117b4576040517f0809490800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b83600073ffffffffffffffffffffffffffffffffffffffff16600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561187e576040517fe51315d200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d4a536866040518163ffffffff1660e01b8152600401602060405180830381865afa1580156118eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061190f91906177a8565b15611946576040517f1e59ccd900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b611a29600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156119b6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906119da919061774f565b85600260008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002061455c9092919063ffffffff16565b8473ffffffffffffffffffffffffffffffffffffffff167f0ad9bf1b8c026a174a2f30954417002a6ea00c9b08c1b8846c7951c687be809585604051611a6f91906177d5565b60405180910390a25050505050565b60608060008511611b1f57600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611af6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b1a919061774f565b611b21565b845b945060008411611b32576001611b34565b835b935060008311611b45576032611b47565b825b92506000600260008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060008160090180549050905060008686611ba6919061787c565b905060008682611bb691906178d6565b90508667ffffffffffffffff811115611bd257611bd1616d7f565b5b604051908082528060200260200182016040528015611c005781602001602082028036833780820191505090505b5095508667ffffffffffffffff811115611c1d57611c1c616d7f565b5b604051908082528060200260200182016040528015611c4b5781602001602082028036833780820191505090505b50945060005b82821015611e645783821415611c6657611e64565b600060056000876009018581548110611c8257611c8161790a565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16888381518110611d2157611d2061790a565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050611d978660000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660028d8461456e909392919063ffffffff16565b611dd38760000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660018e8561456e909392919063ffffffff16565b611e0f8860000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660008f8661456e909392919063ffffffff16565b611e199190617826565b611e239190617826565b878381518110611e3657611e3561790a565b5b6020026020010181815250508180611e4d90617939565b925050508180611e5c90617939565b925050611c51565b505050505094509492505050565b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611ed7576040517f1cf4735900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008054906101000a900460ff1615611f1c576040517f0dc149f000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60016000806101000a81548160ff02191690831515021790555081600060016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b606080606080600380548060200260200160405190810160405280929190818152602001828054801561204457602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611ffa575b50505050509350612140600387600088116120ef57600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156120c6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906120ea919061774f565b6120f1565b875b600560008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020614679909392919063ffffffff16565b80935081945082955050505093509350935093565b600080600183600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156121c8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121ec919061774f565b6121f691906178d6565b61220091906178d6565b9050600080600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639cf370d76040518163ffffffff1660e01b8152600401600060405180830381865afa158015612272573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061229b9190617ade565b915091506000600380549050905060005b868110156123ad576001856122c19190617826565b945060006122da8487876148d39092919063ffffffff16565b905060005b8381101561239857612378828860026000600386815481106123045761230361790a565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206149759092919063ffffffff16565b886123839190617826565b9750808061239090617939565b9150506122df565b505080806123a590617939565b9150506122ac565b5050505050919050565b600381815481106123c757600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663322433e3336040518263ffffffff1660e01b815260040161245191906173f5565b602060405180830381865afa15801561246e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061249291906177a8565b6124c8576040517f8460af8a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61251981600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020614abf90919063ffffffff16565b6003339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555033600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60046020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b82600073ffffffffffffffffffffffffffffffffffffffff16600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415612720576040517fe51315d200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d4a536866040518163ffffffff1660e01b8152600401602060405180830381865afa15801561278d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906127b191906177a8565b156127e8576040517f1e59ccd900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000821415612823576040517ff792180a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61282e833384614b99565b6000600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600073ffffffffffffffffffffffffffffffffffffffff168160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561297157338160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506006339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b6129ea600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020868685614d1390949392919063ffffffff16565b8473ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f8fc656e319452025372383dc27d933046d412b8253de50a10739eeaa59862be68686604051612a49929190617bcd565b60405180910390a35050505050565b600080600080600080600260008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905060008060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612b10573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612b34919061774f565b90508160010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16612b6e8284614f1890919063ffffffff16565b15612b828385614f4590919063ffffffff16565b612b958486614f7290919063ffffffff16565b612ba88587614f9790919063ffffffff16565b96509650965096509650505091939590929450565b60056020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081565b81600073ffffffffffffffffffffffffffffffffffffffff16600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415612cc5576040517fe51315d200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415612d8e576040517fcf83d93d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612e43600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002084600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020614fbc909392919063ffffffff16565b505050565b33600073ffffffffffffffffffffffffffffffffffffffff16600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415612f12576040517fe51315d200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b612f6382600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002061505790919063ffffffff16565b33600460008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b606060008311612ff6576001612ff8565b825b92506000821161300957603261300b565b815b91506000600680549050905060008484613025919061787c565b90506000848261303591906178d6565b905060008567ffffffffffffffff81111561305357613052616d7f565b5b6040519080825280602002602001820160405280156130815781602001602082028036833780820191505090505b50905060005b83831015613149578483141561309c57613149565b600683815481106130b0576130af61790a565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168282815181106130ee576130ed61790a565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050808061313390617939565b915050828061314190617939565b935050613087565b819550505050505092915050565b606060038054806020026020016040519081016040528092919081815260200182805480156131db57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311613191575b5050505050905090565b6060806060613315600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633fa4f2456040518163ffffffff1660e01b815260040161010060405180830381865afa15801561325b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061327f9190617721565b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156132ec573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613310919061774f565b6140e0565b809350819450829550505050909192565b816000600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415801561341a57508060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614155b15613451576040517f0809490800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b83600073ffffffffffffffffffffffffffffffffffffffff16600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561351b576040517fe51315d200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613590600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1685600260008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002061518f9092919063ffffffff16565b5050505050565b6000806000806000600260008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506135f38682614f1890919063ffffffff16565b156136078783614f4590919063ffffffff16565b61361a8884614f7290919063ffffffff16565b61362d8985614f9790919063ffffffff16565b94509450945094505092959194509250565b60006136f6600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002084600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206151eb909392919063ffffffff16565b50809150509392505050565b82600073ffffffffffffffffffffffffffffffffffffffff16600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156137cc576040517fe51315d200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415613895576040517fcf83d93d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d4a536866040518163ffffffff1660e01b8152600401602060405180830381865afa158015613902573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061392691906177a8565b1561395d576040517f1e59ccd900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000821415613998576040517ff792180a00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613a4f600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208585600560003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206154dd90949392919063ffffffff16565b91508373ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167ff2812c3df2511a467cbe777b1ee98b1ddb9918bb0f09568a269d2fb58233cb528585604051613ab0929190617bcd565b60405180910390a350505050565b6000806000600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050613bac600385600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613b77573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613b9b919061774f565b846157c0909392919063ffffffff16565b613be3600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1686846158489092919063ffffffff16565b92509250509250929050565b60068181548110613bff57600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60026020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060080154905083565b6064811115613cd3576040517f755451f400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b613d6a8360050160018473ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613d27573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613d4b919061774f565b613d559190617826565b8386600401615a97909392919063ffffffff16565b505050565b60008085600b016000858152602001908152602001600020541415613da9578185600b016000858152602001908152602001600020819055505b6000600186600c01600086815260200190815260200160002054613dcd9190617826565b90508086600c016000868152602001908152602001600020819055508460c001518110158015613e295750856003016000600186613e0b9190617826565b815260200190815260200160002060009054906101000a900460ff16155b15613e8b578460e0015184613e3e9190617826565b91505b81841015613e8a578380613e5490617939565b945050600186600301600086815260200190815260200160002060006101000a81548160ff021916908315150217905550613e41565b5b50949350505050565b613ea18383836000615ada565b505050565b60008084600801549050600060018573ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613eff573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613f23919061774f565b613f2d91906178d6565b90506000841480613f485750808285613f469190617826565b115b15613f5c578181613f5991906178d6565b93505b6000808673ffffffffffffffffffffffffffffffffffffffff16639cf370d76040518163ffffffff1660e01b8152600401600060405180830381865afa158015613faa573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190613fd39190617ade565b9150915060005b8681101561407257600185613fef9190617826565b945060006140128a61400c8589886148d39092919063ffffffff16565b88614975565b90506000811415614023575061405f565b600061402f8b88614f97565b9050600081141561404157505061405f565b61404f828260646019615b8e565b8861405a9190617826565b975050505b808061406a90617939565b915050613fda565b50505050935093915050565b60008083600b0160008481526020019081526020016000205484600c01600085815260200190815260200160002054915091509250929050565b6140c482826001615bc7565b6140d082826002615bc7565b6140dc82826000615bc7565b5050565b6060806060600060038054905067ffffffffffffffff81111561410657614105616d7f565b5b6040519080825280602002602001820160405280156141345781602001602082028036833780820191505090505b5090506000805b6003805490508110156142e857600060026000600384815481106141625761416161790a565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905080600201600089815260200190815260200160002060009054906101000a900460ff16156141f757506142d5565b61420a8882614f4590919063ffffffff16565b1561421557506142d5565b8860a0015161422d8983614f7290919063ffffffff16565b101561423957506142d5565b6003828154811061424d5761424c61790a565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684848151811061428b5761428a61790a565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505082806142d090617939565b935050505b80806142e090617939565b91505061413b565b508067ffffffffffffffff81111561430357614302616d7f565b5b6040519080825280602002602001820160405280156143315781602001602082028036833780820191505090505b5094508067ffffffffffffffff81111561434e5761434d616d7f565b5b60405190808252806020026020018201604052801561437c5781602001602082028036833780820191505090505b5093508067ffffffffffffffff81111561439957614398616d7f565b5b6040519080825280602002602001820160405280156143c75781602001602082028036833780820191505090505b50925060005b81811015614552578281815181106143e8576143e761790a565b5b60200260200101518682815181106144035761440261790a565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506000600260008584815181106144565761445561790a565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168683815181106144d2576144d161790a565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff168152505061451f8882614f7290919063ffffffff16565b8583815181106145325761453161790a565b5b60200260200101818152505050808061454a90617939565b9150506143cd565b5050509250925092565b6145698383836001615ada565b505050565b600061466f85600201600085600281111561458c5761458b617b56565b5b600281111561459e5761459d617b56565b5b815260200190815260200160002060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208387600101600087600281111561460357614602617b56565b5b600281111561461557614614617b56565b5b815260200190815260200160002060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020615f0f9092919063ffffffff16565b9050949350505050565b6060806060858054905067ffffffffffffffff81111561469c5761469b616d7f565b5b6040519080825280602002602001820160405280156146ca5781602001602082028036833780820191505090505b509250858054905067ffffffffffffffff8111156146eb576146ea616d7f565b5b6040519080825280602002602001820160405280156147195781602001602082028036833780820191505090505b509150858054905067ffffffffffffffff81111561473a57614739616d7f565b5b6040519080825280602002602001820160405280156147685781602001602082028036833780820191505090505b50905060008680549050905060005b818110156148c75760006147ca8a8a84815481106147985761479761790a565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168a8a61456e565b905060006148238b8b85815481106147e5576147e461790a565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168b60018c61481e9190617826565b61456e565b9050818784815181106148395761483861790a565b5b6020026020010181815250508181111561487d57818161485991906178d6565b86848151811061486c5761486b61790a565b5b6020026020010181815250506148b2565b818110156148b157808261489191906178d6565b8584815181106148a4576148a361790a565b5b6020026020010181815250505b5b505080806148bf90617939565b915050614777565b50509450945094915050565b6148db616b4e565b60008451905082856001836148f091906178d6565b815181106149015761490061790a565b5b60200260200101511161493d578360018261491c91906178d6565b8151811061492d5761492c61790a565b5b602002602001015191505061496e565b600061494c8685600085615ff1565b90508481815181106149615761496061790a565b5b6020026020010151925050505b9392505050565b60006149818483614f18565b8061499257506149918483614f45565b5b156149a05760009050614ab8565b60006149ac8584614f72565b905060008114156149c1576000915050614ab8565b60006019600a6149d19190617d29565b6149e28660800151606460196160b0565b836149ed919061787c565b6149f79190617da3565b90506000811415614a0d57600092505050614ab8565b614a3085606001518660400151614a24919061787c565b6301e1338060196160b0565b81614a3b919061787c565b90506019600a614a4b9190617d29565b81614a569190617da3565b9050600086600c0160008681526020019081526020016000205490506000811115614ab157600087600b016000878152602001908152602001600020549050614aad838383614aa591906178d6565b836019615b8e565b9250505b8193505050505b9392505050565b600073ffffffffffffffffffffffffffffffffffffffff168260000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614614b48576040517e3b268200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b338260000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550614b958282615057565b5050565b60006002811115614bad57614bac617b56565b5b836002811115614bc057614bbf617b56565b5b1415614c0457803414614bff576040517f1fcb60ca00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b614d0e565b60003414614c3e576040517fa745ac8500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000614c4984616107565b73ffffffffffffffffffffffffffffffffffffffff166323b872dd8430856040518463ffffffff1660e01b8152600401614c8593929190617606565b6020604051808303816000875af1158015614ca4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614cc891906177a8565b905080614d0c57836040517f1b6a68ea000000000000000000000000000000000000000000000000000000008152600401614d039190617dd4565b60405180910390fd5b505b505050565b614ed6856002016000846002811115614d2f57614d2e617b56565b5b6002811115614d4157614d40617b56565b5b815260200190815260200160002060008560000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060018673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015614dfd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614e21919061774f565b614e2b9190617826565b83886001016000876002811115614e4557614e44617b56565b5b6002811115614e5757614e56617b56565b5b815260200190815260200160002060008860000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206161d1909392919063ffffffff16565b614f11848660000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168386616226909392919063ffffffff16565b5050505050565b600082600201600083815260200190815260200160002060009054906101000a900460ff16905092915050565b600082600301600083815260200190815260200160002060009054906101000a900460ff16905092915050565b6000614f8f836007018385600601615f0f9092919063ffffffff16565b905092915050565b6000614fb4836005018385600401615f0f9092919063ffffffff16565b905092915050565b600080614fcb868686866151eb565b91509150808660050160008660000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600082111561504f5761504e600033846163d6565b5b505050505050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156150be576040517f7138356f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415615148576040517fe037058f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b808260010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b60008061519d858585613ea6565b9150915080856008018190555060008211156151e4576151e360008660000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846163d6565b5b5050505050565b6000808560050160008560000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050600060018673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156152a5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906152c9919061774f565b6152d391906178d6565b905060008414806152ee57508082856152ec9190617826565b115b156153025781816152ff91906178d6565b93505b6000808773ffffffffffffffffffffffffffffffffffffffff16639cf370d76040518163ffffffff1660e01b8152600401600060405180830381865afa158015615350573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906153799190617ade565b9150915060005b868110156154d0576001856153959190617826565b945060006153ca8b8a60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660028961456e565b6153fb8c8b60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660018a61456e565b61542c8d8c60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660008b61456e565b6154369190617826565b6154409190617826565b9050600081141561545157506154bd565b600061547c61546b8589886148d39092919063ffffffff16565b888c6165469092919063ffffffff16565b9050600081141561548e5750506154bd565b6154ad81836154a68a8e614f7290919063ffffffff16565b6019615b8e565b886154b89190617826565b975050505b80806154c890617939565b915050615380565b5050505094509492505050565b6000808573ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561552b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061554f919061774f565b90506000615583888760000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16878561456e565b905060006155c3898860000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16886001876155be9190617826565b61456e565b905061571a8960020160008860028111156155e1576155e0617b56565b5b60028111156155f3576155f2617b56565b5b815260200190815260200160002060008960000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060018561566f9190617826565b878c60010160008b600281111561568957615688617b56565b5b600281111561569b5761569a617b56565b5b815260200190815260200160002060008c60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206165af909392919063ffffffff16565b9450600085141561573157600093505050506157b7565b61574688868961665c9092919063ffffffff16565b600085905060008383111561578357838361576191906178d6565b90508087106157705780615772565b865b9050808261578091906178d6565b91505b6000821115615799576157988b8b8a856166f9565b5b60008111156157ae576157ad8833836163d6565b5b86955050505050505b95945050505050565b6000808480549050905060005b8181101561583e5761581e878783815481106157ec576157eb61790a565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16878761456e565b836158299190617826565b9250808061583690617939565b9150506157cd565b5050949350505050565b60008084600301600084600281111561586457615863617b56565b5b600281111561587657615875617b56565b5b8152602001908152602001600020805490509050600081141561589d576000915050615a90565b60008473ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156158ea573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061590e919061774f565b9050600060018361591f91906178d6565b905060008111801561598557508187600301600087600281111561594657615945617b56565b5b600281111561595857615957617b56565b5b815260200190815260200160002082815481106159785761597761790a565b5b9060005260206000200154115b1561599957808061599590617def565b9150505b818760030160008760028111156159b3576159b2617b56565b5b60028111156159c5576159c4617b56565b5b815260200190815260200160002082815481106159e5576159e461790a565b5b90600052602060002001541115615a025760009350505050615a90565b600080600090505b828111615a8757886004016000886002811115615a2a57615a29617b56565b5b6002811115615a3c57615a3b617b56565b5b81526020019081526020016000208181548110615a5c57615a5b61790a565b5b906000526020600020015482615a729190617826565b91508080615a7f90617939565b915050615a0a565b50809450505050505b9392505050565b615aa2848484616976565b808360018680549050615ab591906178d6565b81548110615ac657615ac561790a565b5b906000526020600020018190555050505050565b60008251905060005b81811015615b86576000848281518110615b0057615aff61790a565b5b602002602001015190508581118015615b3f575083151587600201600083815260200190815260200160002060009054906101000a900460ff16151514155b15615b72578387600201600083815260200190815260200160002060006101000a81548160ff0219169083151502179055505b508080615b7e90617939565b915050615ae3565b505050505050565b600081600a615b9d9190617d29565b615ba88585856160b0565b86615bb3919061787c565b615bbd9190617da3565b9050949350505050565b6000615bd4848484615848565b90506000811415615be55750615f0a565b6000846003016000846002811115615c0057615bff617b56565b5b6002811115615c1257615c11617b56565b5b81526020019081526020016000208054905090508373ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015615c71573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615c95919061774f565b856003016000856002811115615cae57615cad617b56565b5b6002811115615cc057615cbf617b56565b5b8152602001908152602001600020600183615cdb91906178d6565b81548110615cec57615ceb61790a565b5b906000526020600020015411615d8b57846003016000846002811115615d1557615d14617b56565b5b6002811115615d2757615d26617b56565b5b81526020019081526020016000206000615d419190616b93565b846004016000846002811115615d5a57615d59617b56565b5b6002811115615d6c57615d6b617b56565b5b81526020019081526020016000206000615d869190616b93565b615efc565b6040518060200160405280866003016000866002811115615daf57615dae617b56565b5b6002811115615dc157615dc0617b56565b5b8152602001908152602001600020600184615ddc91906178d6565b81548110615ded57615dec61790a565b5b9060005260206000200154815250856003016000856002811115615e1457615e13617b56565b5b6002811115615e2657615e25617b56565b5b8152602001908152602001600020906001615e42929190616bb4565b506040518060200160405280866004016000866002811115615e6757615e66617b56565b5b6002811115615e7957615e78617b56565b5b8152602001908152602001600020600184615e9491906178d6565b81548110615ea557615ea461790a565b5b9060005260206000200154815250856004016000856002811115615ecc57615ecb617b56565b5b6002811115615ede57615edd617b56565b5b8152602001908152602001600020906001615efa929190616bb4565b505b615f078333846163d6565b50505b505050565b600080848054905090506000811480615f4557508285600081548110615f3857615f3761790a565b5b9060005260206000200154115b15615f54576000915050615fea565b8285600183615f6391906178d6565b81548110615f7457615f7361790a565b5b906000526020600020015411615fb65783600182615f9291906178d6565b81548110615fa357615fa261790a565b5b9060005260206000200154915050615fea565b6000615fc58685600085616a89565b9050848181548110615fda57615fd961790a565b5b9060005260206000200154925050505b9392505050565b60008183141561600f5760018261600891906178d6565b90506160a8565b60006002838561601f9190617826565b6160299190617da3565b90508486828151811061603f5761603e61790a565b5b602002602001015111156160615761605986868684615ff1565b9150506160a8565b848682815181106160755761607461790a565b5b602002602001015110156160a35761609b86866001846160959190617826565b86615ff1565b9150506160a8565b809150505b949350505050565b6000806001836160c09190617826565b600a6160cc9190617d29565b856160d7919061787c565b9050600a600585836160e99190617da3565b6160f39190617826565b6160fd9190617da3565b9150509392505050565b60006001600281111561611d5761611c617b56565b5b8260028111156161305761612f617b56565b5b14156161525773520000000000000000000000000000000000000190506161cc565b60028081111561616557616164617b56565b5b82600281111561617857616177617b56565b5b141561619a5773520000000000000000000000000000000000000290506161cc565b6040517f8698bf3700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b919050565b6161dc848484616976565b8083600186805490506161ef91906178d6565b81548110616200576161ff61790a565b5b9060005260206000200160008282546162199190617826565b9250508190555050505050565b83600a0160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1661633957600184600a0160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555083600901829080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b6163d08460070160018573ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561638d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906163b1919061774f565b6163bb9190617826565b83876006016161d1909392919063ffffffff16565b50505050565b60008060028111156163eb576163ea617b56565b5b8460028111156163fe576163fd617b56565b5b1415616475578273ffffffffffffffffffffffffffffffffffffffff168260405161642890617e4a565b60006040518083038185875af1925050503d8060008114616465576040519150601f19603f3d011682016040523d82523d6000602084013e61646a565b606091505b5050809150506164fe565b61647e84616107565b73ffffffffffffffffffffffffffffffffffffffff1663a9059cbb84846040518363ffffffff1660e01b81526004016164b8929190617e5f565b6020604051808303816000875af11580156164d7573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906164fb91906177a8565b90505b8061654057836040517f1b6a68ea0000000000000000000000000000000000000000000000000000000081526004016165379190617dd4565b60405180910390fd5b50505050565b600080616554858585614975565b905060008114156165695760009150506165a8565b60006165758685614f97565b9050600081141561658a5781925050506165a8565b616598828260646019615b8e565b826165a391906178d6565b925050505b9392505050565b60006165bc858585616976565b6000858054905090506000856001836165d591906178d6565b815481106165e6576165e561790a565b5b90600052602060002001549050808411156166015780616603565b835b9350600084111561664f57838660018461661d91906178d6565b8154811061662e5761662d61790a565b5b90600052602060002001600082825461664791906178d6565b925050819055505b8392505050949350505050565b6166f38360070160018473ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156166b0573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906166d4919061774f565b6166de9190617826565b83866006016165af909392919063ffffffff16565b50505050565b600060018473ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b8152600401602060405180830381865afa158015616748573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061676c919061774f565b6167769190617826565b9050600085600301600085600281111561679357616792617b56565b5b60028111156167a5576167a4617b56565b5b815260200190815260200160002080549050905060008114806168295750818660030160008660028111156167dd576167dc617b56565b5b60028111156167ef576167ee617b56565b5b815260200190815260200160002060018361680a91906178d6565b8154811061681b5761681a61790a565b5b906000526020600020015414155b156168f55785600301600085600281111561684757616846617b56565b5b600281111561685957616858617b56565b5b81526020019081526020016000208290806001815401808255809150506001900390600052602060002001600090919091909150558560040160008560028111156168a7576168a6617b56565b5b60028111156168b9576168b8617b56565b5b81526020019081526020016000208390806001815401808255809150506001900390600052602060002001600090919091909150555050616970565b8286600401600086600281111561690f5761690e617b56565b5b600281111561692157616920617b56565b5b815260200190815260200160002060018361693c91906178d6565b8154811061694d5761694c61790a565b5b9060005260206000200160008282546169669190617826565b9250508190555050505b50505050565b60008380549050905060008114156169d257838290806001815401808255809150506001900390600052602060002001600090919091909150558260018160018154018082558091505003906000526020600020505050616a84565b6000846001836169e291906178d6565b815481106169f3576169f261790a565b5b90600052602060002001549050828114616a8157848390806001815401808255809150506001900390600052602060002001600090919091909150558384600184616a3e91906178d6565b81548110616a4f57616a4e61790a565b5b906000526020600020015490806001815401808255809150506001900390600052602060002001600090919091909150555b50505b505050565b600081831415616aa757600182616aa091906178d6565b9050616b46565b600060028385616ab79190617826565b616ac19190617da3565b905084868281548110616ad757616ad661790a565b5b90600052602060002001541115616afc57616af486868684616a89565b915050616b46565b84868281548110616b1057616b0f61790a565b5b90600052602060002001541015616b4157616b398686600184616b339190617826565b86616a89565b915050616b46565b809150505b949350505050565b60405180610100016040528060008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b5080546000825590600052602060002090810190616bb19190616c01565b50565b828054828255906000526020600020908101928215616bf0579160200282015b82811115616bef578251825591602001919060010190616bd4565b5b509050616bfd9190616c01565b5090565b5b80821115616c1a576000816000905550600101616c02565b5090565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b616c4581616c32565b8114616c5057600080fd5b50565b600081359050616c6281616c3c565b92915050565b600060208284031215616c7e57616c7d616c28565b5b6000616c8c84828501616c53565b91505092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000616cc082616c95565b9050919050565b616cd081616cb5565b8114616cdb57600080fd5b50565b600081359050616ced81616cc7565b92915050565b60008060408385031215616d0a57616d09616c28565b5b6000616d1885828601616cde565b9250506020616d2985828601616c53565b9150509250929050565b60008115159050919050565b616d4881616d33565b82525050565b6000602082019050616d636000830184616d3f565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b616db782616d6e565b810181811067ffffffffffffffff82111715616dd657616dd5616d7f565b5b80604052505050565b6000616de9616c1e565b9050616df58282616dae565b919050565b600067ffffffffffffffff821115616e1557616e14616d7f565b5b602082029050602081019050919050565b600080fd5b6000616e3e616e3984616dfa565b616ddf565b90508083825260208201905060208402830185811115616e6157616e60616e26565b5b835b81811015616e8a5780616e768882616c53565b845260208401935050602081019050616e63565b5050509392505050565b600082601f830112616ea957616ea8616d69565b5b8135616eb9848260208601616e2b565b91505092915050565b60008060408385031215616ed957616ed8616c28565b5b6000616ee785828601616cde565b925050602083013567ffffffffffffffff811115616f0857616f07616c2d565b5b616f1485828601616e94565b9150509250929050565b616f2781616c32565b82525050565b6000602082019050616f426000830184616f1e565b92915050565b6000604082019050616f5d6000830185616f1e565b616f6a6020830184616f1e565b9392505050565b6000819050919050565b6000616f96616f91616f8c84616c95565b616f71565b616c95565b9050919050565b6000616fa882616f7b565b9050919050565b6000616fba82616f9d565b9050919050565b616fca81616faf565b82525050565b6000602082019050616fe56000830184616fc1565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61702081616cb5565b82525050565b60006170328383617017565b60208301905092915050565b6000602082019050919050565b600061705682616feb565b6170608185616ff6565b935061706b83617007565b8060005b8381101561709c5781516170838882617026565b975061708e8361703e565b92505060018101905061706f565b5085935050505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6170de81616c32565b82525050565b60006170f083836170d5565b60208301905092915050565b6000602082019050919050565b6000617114826170a9565b61711e81856170b4565b9350617129836170c5565b8060005b8381101561715a57815161714188826170e4565b975061714c836170fc565b92505060018101905061712d565b5085935050505092915050565b60006060820190508181036000830152617181818661704b565b90508181036020830152617195818561704b565b905081810360408301526171a98184617109565b9050949350505050565b600080600080608085870312156171cd576171cc616c28565b5b60006171db87828801616cde565b94505060206171ec87828801616c53565b93505060406171fd87828801616c53565b925050606061720e87828801616c53565b91505092959194509250565b60006040820190508181036000830152617234818561704b565b905081810360208301526172488184617109565b90509392505050565b600061725c82616cb5565b9050919050565b61726c81617251565b811461727757600080fd5b50565b60008135905061728981617263565b92915050565b600061729a82616cb5565b9050919050565b6172aa8161728f565b81146172b557600080fd5b50565b6000813590506172c7816172a1565b92915050565b600080604083850312156172e4576172e3616c28565b5b60006172f28582860161727a565b9250506020617303858286016172b8565b9150509250929050565b6003811061731a57600080fd5b50565b60008135905061732c8161730d565b92915050565b60008060006060848603121561734b5761734a616c28565b5b600061735986828701616cde565b935050602061736a8682870161731d565b925050604061737b86828701616c53565b9150509250925092565b6000608082019050818103600083015261739f818761704b565b905081810360208301526173b38186617109565b905081810360408301526173c78185617109565b905081810360608301526173db8184617109565b905095945050505050565b6173ef81616cb5565b82525050565b600060208201905061740a60008301846173e6565b92915050565b60006020828403121561742657617425616c28565b5b600061743484828501616cde565b91505092915050565b600061744882616f9d565b9050919050565b6174588161743d565b82525050565b6000602082019050617473600083018461744f565b92915050565b600060a08201905061748e60008301886173e6565b61749b6020830187616d3f565b6174a86040830186616d3f565b6174b56060830185616f1e565b6174c26080830184616f1e565b9695505050505050565b600080604083850312156174e3576174e2616c28565b5b60006174f185828601616c53565b925050602061750285828601616c53565b9150509250929050565b60006020820190508181036000830152617526818461704b565b905092915050565b60006080820190506175436000830187616d3f565b6175506020830186616d3f565b61755d6040830185616f1e565b61756a6060830184616f1e565b95945050505050565b60008060006060848603121561758c5761758b616c28565b5b600061759a86828701616cde565b93505060206175ab86828701616cde565b92505060406175bc86828701616c53565b9150509250925092565b600080604083850312156175dd576175dc616c28565b5b60006175eb85828601616cde565b92505060206175fc8582860161731d565b9150509250929050565b600060608201905061761b60008301866173e6565b61762860208301856173e6565b6176356040830184616f1e565b949350505050565b600080fd5b60008151905061765181616c3c565b92915050565b6000610100828403121561766e5761766d61763d565b5b617679610100616ddf565b9050600061768984828501617642565b600083015250602061769d84828501617642565b60208301525060406176b184828501617642565b60408301525060606176c584828501617642565b60608301525060806176d984828501617642565b60808301525060a06176ed84828501617642565b60a08301525060c061770184828501617642565b60c08301525060e061771584828501617642565b60e08301525092915050565b6000610100828403121561773857617737616c28565b5b600061774684828501617657565b91505092915050565b60006020828403121561776557617764616c28565b5b600061777384828501617642565b91505092915050565b61778581616d33565b811461779057600080fd5b50565b6000815190506177a28161777c565b92915050565b6000602082840312156177be576177bd616c28565b5b60006177cc84828501617793565b91505092915050565b600060208201905081810360008301526177ef8184617109565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061783182616c32565b915061783c83616c32565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115617871576178706177f7565b5b828201905092915050565b600061788782616c32565b915061789283616c32565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156178cb576178ca6177f7565b5b828202905092915050565b60006178e182616c32565b91506178ec83616c32565b9250828210156178ff576178fe6177f7565b5b828203905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600061794482616c32565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415617977576179766177f7565b5b600182019050919050565b600061799561799084616dfa565b616ddf565b905080838252602082019050602084028301858111156179b8576179b7616e26565b5b835b818110156179e157806179cd8882617642565b8452602084019350506020810190506179ba565b5050509392505050565b600082601f830112617a00576179ff616d69565b5b8151617a10848260208601617982565b91505092915050565b600067ffffffffffffffff821115617a3457617a33616d7f565b5b602082029050602081019050919050565b6000617a58617a5384617a19565b616ddf565b9050808382526020820190506101008402830185811115617a7c57617a7b616e26565b5b835b81811015617aa65780617a918882617657565b84526020840193505061010081019050617a7e565b5050509392505050565b600082601f830112617ac557617ac4616d69565b5b8151617ad5848260208601617a45565b91505092915050565b60008060408385031215617af557617af4616c28565b5b600083015167ffffffffffffffff811115617b1357617b12616c2d565b5b617b1f858286016179eb565b925050602083015167ffffffffffffffff811115617b4057617b3f616c2d565b5b617b4c85828601617ab0565b9150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60038110617b9657617b95617b56565b5b50565b6000819050617ba782617b85565b919050565b6000617bb782617b99565b9050919050565b617bc781617bac565b82525050565b6000604082019050617be26000830185617bbe565b617bef6020830184616f1e565b9392505050565b60008160011c9050919050565b6000808291508390505b6001851115617c4d57808604811115617c2957617c286177f7565b5b6001851615617c385780820291505b8081029050617c4685617bf6565b9450617c0d565b94509492505050565b600082617c665760019050617d22565b81617c745760009050617d22565b8160018114617c8a5760028114617c9457617cc3565b6001915050617d22565b60ff841115617ca657617ca56177f7565b5b8360020a915084821115617cbd57617cbc6177f7565b5b50617d22565b5060208310610133831016604e8410600b8410161715617cf85782820a905083811115617cf357617cf26177f7565b5b617d22565b617d058484846001617c03565b92509050818404811115617d1c57617d1b6177f7565b5b81810290505b9392505050565b6000617d3482616c32565b9150617d3f83616c32565b9250617d6c7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8484617c56565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000617dae82616c32565b9150617db983616c32565b925082617dc957617dc8617d74565b5b828204905092915050565b6000602082019050617de96000830184617bbe565b92915050565b6000617dfa82616c32565b91506000821415617e0e57617e0d6177f7565b5b600182039050919050565b600081905092915050565b50565b6000617e34600083617e19565b9150617e3f82617e24565b600082019050919050565b6000617e5582617e27565b9150819050919050565b6000604082019050617e7460008301856173e6565b617e816020830184616f1e565b939250505056fea26469706673582212209c9417a214756c78a4e6e81e9b9b11abe93a4bd887bbbb08707f86d8a832a71a64736f6c634300080c0033",
}

// StakemanagerABI is the input ABI used to generate the binding from.
// Deprecated: Use StakemanagerMetaData.ABI instead.
var StakemanagerABI = StakemanagerMetaData.ABI

// StakemanagerBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakemanagerMetaData.Bin instead.
var StakemanagerBin = StakemanagerMetaData.Bin

// DeployStakemanager deploys a new Ethereum contract, binding an instance of Stakemanager to it.
func DeployStakemanager(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Stakemanager, error) {
	parsed, err := StakemanagerMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakemanagerBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Stakemanager{StakemanagerCaller: StakemanagerCaller{contract: contract}, StakemanagerTransactor: StakemanagerTransactor{contract: contract}, StakemanagerFilterer: StakemanagerFilterer{contract: contract}}, nil
}

// Stakemanager is an auto generated Go binding around an Ethereum contract.
type Stakemanager struct {
	StakemanagerCaller     // Read-only binding to the contract
	StakemanagerTransactor // Write-only binding to the contract
	StakemanagerFilterer   // Log filterer for contract events
}

// StakemanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakemanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakemanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakemanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakemanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakemanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakemanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakemanagerSession struct {
	Contract     *Stakemanager     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakemanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakemanagerCallerSession struct {
	Contract *StakemanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// StakemanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakemanagerTransactorSession struct {
	Contract     *StakemanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// StakemanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakemanagerRaw struct {
	Contract *Stakemanager // Generic contract binding to access the raw methods on
}

// StakemanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakemanagerCallerRaw struct {
	Contract *StakemanagerCaller // Generic read-only contract binding to access the raw methods on
}

// StakemanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakemanagerTransactorRaw struct {
	Contract *StakemanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakemanager creates a new instance of Stakemanager, bound to a specific deployed contract.
func NewStakemanager(address common.Address, backend bind.ContractBackend) (*Stakemanager, error) {
	contract, err := bindStakemanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stakemanager{StakemanagerCaller: StakemanagerCaller{contract: contract}, StakemanagerTransactor: StakemanagerTransactor{contract: contract}, StakemanagerFilterer: StakemanagerFilterer{contract: contract}}, nil
}

// NewStakemanagerCaller creates a new read-only instance of Stakemanager, bound to a specific deployed contract.
func NewStakemanagerCaller(address common.Address, caller bind.ContractCaller) (*StakemanagerCaller, error) {
	contract, err := bindStakemanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakemanagerCaller{contract: contract}, nil
}

// NewStakemanagerTransactor creates a new write-only instance of Stakemanager, bound to a specific deployed contract.
func NewStakemanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*StakemanagerTransactor, error) {
	contract, err := bindStakemanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakemanagerTransactor{contract: contract}, nil
}

// NewStakemanagerFilterer creates a new log filterer instance of Stakemanager, bound to a specific deployed contract.
func NewStakemanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*StakemanagerFilterer, error) {
	contract, err := bindStakemanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakemanagerFilterer{contract: contract}, nil
}

// bindStakemanager binds a generic wrapper to an already deployed contract.
func bindStakemanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakemanagerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stakemanager *StakemanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stakemanager.Contract.StakemanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stakemanager *StakemanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakemanager.Contract.StakemanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stakemanager *StakemanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stakemanager.Contract.StakemanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stakemanager *StakemanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stakemanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stakemanager *StakemanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakemanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stakemanager *StakemanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stakemanager.Contract.contract.Transact(opts, method, params...)
}

// Allowlist is a free data retrieval call binding the contract method 0x2b47da52.
//
// Solidity: function allowlist() view returns(address)
func (_Stakemanager *StakemanagerCaller) Allowlist(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "allowlist")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Allowlist is a free data retrieval call binding the contract method 0x2b47da52.
//
// Solidity: function allowlist() view returns(address)
func (_Stakemanager *StakemanagerSession) Allowlist() (common.Address, error) {
	return _Stakemanager.Contract.Allowlist(&_Stakemanager.CallOpts)
}

// Allowlist is a free data retrieval call binding the contract method 0x2b47da52.
//
// Solidity: function allowlist() view returns(address)
func (_Stakemanager *StakemanagerCallerSession) Allowlist() (common.Address, error) {
	return _Stakemanager.Contract.Allowlist(&_Stakemanager.CallOpts)
}

// Environment is a free data retrieval call binding the contract method 0x74e2b63c.
//
// Solidity: function environment() view returns(address)
func (_Stakemanager *StakemanagerCaller) Environment(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "environment")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Environment is a free data retrieval call binding the contract method 0x74e2b63c.
//
// Solidity: function environment() view returns(address)
func (_Stakemanager *StakemanagerSession) Environment() (common.Address, error) {
	return _Stakemanager.Contract.Environment(&_Stakemanager.CallOpts)
}

// Environment is a free data retrieval call binding the contract method 0x74e2b63c.
//
// Solidity: function environment() view returns(address)
func (_Stakemanager *StakemanagerCallerSession) Environment() (common.Address, error) {
	return _Stakemanager.Contract.Environment(&_Stakemanager.CallOpts)
}

// GetBlockAndSlashes is a free data retrieval call binding the contract method 0x22226367.
//
// Solidity: function getBlockAndSlashes(address validator, uint256 epoch) view returns(uint256 blocks, uint256 slashes)
func (_Stakemanager *StakemanagerCaller) GetBlockAndSlashes(opts *bind.CallOpts, validator common.Address, epoch *big.Int) (struct {
	Blocks  *big.Int
	Slashes *big.Int
}, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getBlockAndSlashes", validator, epoch)

	outstruct := new(struct {
		Blocks  *big.Int
		Slashes *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Blocks = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Slashes = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBlockAndSlashes is a free data retrieval call binding the contract method 0x22226367.
//
// Solidity: function getBlockAndSlashes(address validator, uint256 epoch) view returns(uint256 blocks, uint256 slashes)
func (_Stakemanager *StakemanagerSession) GetBlockAndSlashes(validator common.Address, epoch *big.Int) (struct {
	Blocks  *big.Int
	Slashes *big.Int
}, error) {
	return _Stakemanager.Contract.GetBlockAndSlashes(&_Stakemanager.CallOpts, validator, epoch)
}

// GetBlockAndSlashes is a free data retrieval call binding the contract method 0x22226367.
//
// Solidity: function getBlockAndSlashes(address validator, uint256 epoch) view returns(uint256 blocks, uint256 slashes)
func (_Stakemanager *StakemanagerCallerSession) GetBlockAndSlashes(validator common.Address, epoch *big.Int) (struct {
	Blocks  *big.Int
	Slashes *big.Int
}, error) {
	return _Stakemanager.Contract.GetBlockAndSlashes(&_Stakemanager.CallOpts, validator, epoch)
}

// GetCommissions is a free data retrieval call binding the contract method 0x195afea1.
//
// Solidity: function getCommissions(address validator, uint256 epochs) view returns(uint256 commissions)
func (_Stakemanager *StakemanagerCaller) GetCommissions(opts *bind.CallOpts, validator common.Address, epochs *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getCommissions", validator, epochs)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCommissions is a free data retrieval call binding the contract method 0x195afea1.
//
// Solidity: function getCommissions(address validator, uint256 epochs) view returns(uint256 commissions)
func (_Stakemanager *StakemanagerSession) GetCommissions(validator common.Address, epochs *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.GetCommissions(&_Stakemanager.CallOpts, validator, epochs)
}

// GetCommissions is a free data retrieval call binding the contract method 0x195afea1.
//
// Solidity: function getCommissions(address validator, uint256 epochs) view returns(uint256 commissions)
func (_Stakemanager *StakemanagerCallerSession) GetCommissions(validator common.Address, epochs *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.GetCommissions(&_Stakemanager.CallOpts, validator, epochs)
}

// GetCurrentValidators is a free data retrieval call binding the contract method 0xc5f9dff0.
//
// Solidity: function getCurrentValidators() view returns(address[] owners, address[] operators, uint256[] stakes)
func (_Stakemanager *StakemanagerCaller) GetCurrentValidators(opts *bind.CallOpts) (struct {
	Owners    []common.Address
	Operators []common.Address
	Stakes    []*big.Int
}, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getCurrentValidators")

	outstruct := new(struct {
		Owners    []common.Address
		Operators []common.Address
		Stakes    []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owners = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Operators = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Stakes = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetCurrentValidators is a free data retrieval call binding the contract method 0xc5f9dff0.
//
// Solidity: function getCurrentValidators() view returns(address[] owners, address[] operators, uint256[] stakes)
func (_Stakemanager *StakemanagerSession) GetCurrentValidators() (struct {
	Owners    []common.Address
	Operators []common.Address
	Stakes    []*big.Int
}, error) {
	return _Stakemanager.Contract.GetCurrentValidators(&_Stakemanager.CallOpts)
}

// GetCurrentValidators is a free data retrieval call binding the contract method 0xc5f9dff0.
//
// Solidity: function getCurrentValidators() view returns(address[] owners, address[] operators, uint256[] stakes)
func (_Stakemanager *StakemanagerCallerSession) GetCurrentValidators() (struct {
	Owners    []common.Address
	Operators []common.Address
	Stakes    []*big.Int
}, error) {
	return _Stakemanager.Contract.GetCurrentValidators(&_Stakemanager.CallOpts)
}

// GetNextValidators is a free data retrieval call binding the contract method 0x40cddab3.
//
// Solidity: function getNextValidators() view returns(address[] owners, address[] operators, uint256[] stakes)
func (_Stakemanager *StakemanagerCaller) GetNextValidators(opts *bind.CallOpts) (struct {
	Owners    []common.Address
	Operators []common.Address
	Stakes    []*big.Int
}, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getNextValidators")

	outstruct := new(struct {
		Owners    []common.Address
		Operators []common.Address
		Stakes    []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owners = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Operators = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)
	outstruct.Stakes = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetNextValidators is a free data retrieval call binding the contract method 0x40cddab3.
//
// Solidity: function getNextValidators() view returns(address[] owners, address[] operators, uint256[] stakes)
func (_Stakemanager *StakemanagerSession) GetNextValidators() (struct {
	Owners    []common.Address
	Operators []common.Address
	Stakes    []*big.Int
}, error) {
	return _Stakemanager.Contract.GetNextValidators(&_Stakemanager.CallOpts)
}

// GetNextValidators is a free data retrieval call binding the contract method 0x40cddab3.
//
// Solidity: function getNextValidators() view returns(address[] owners, address[] operators, uint256[] stakes)
func (_Stakemanager *StakemanagerCallerSession) GetNextValidators() (struct {
	Owners    []common.Address
	Operators []common.Address
	Stakes    []*big.Int
}, error) {
	return _Stakemanager.Contract.GetNextValidators(&_Stakemanager.CallOpts)
}

// GetRewards is a free data retrieval call binding the contract method 0xdbd61d87.
//
// Solidity: function getRewards(address staker, address validator, uint256 epochs) view returns(uint256 rewards)
func (_Stakemanager *StakemanagerCaller) GetRewards(opts *bind.CallOpts, staker common.Address, validator common.Address, epochs *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getRewards", staker, validator, epochs)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRewards is a free data retrieval call binding the contract method 0xdbd61d87.
//
// Solidity: function getRewards(address staker, address validator, uint256 epochs) view returns(uint256 rewards)
func (_Stakemanager *StakemanagerSession) GetRewards(staker common.Address, validator common.Address, epochs *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.GetRewards(&_Stakemanager.CallOpts, staker, validator, epochs)
}

// GetRewards is a free data retrieval call binding the contract method 0xdbd61d87.
//
// Solidity: function getRewards(address staker, address validator, uint256 epochs) view returns(uint256 rewards)
func (_Stakemanager *StakemanagerCallerSession) GetRewards(staker common.Address, validator common.Address, epochs *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.GetRewards(&_Stakemanager.CallOpts, staker, validator, epochs)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0xe4a8a343.
//
// Solidity: function getStakerInfo(address staker, uint8 token) view returns(uint256 stakes, uint256 unstakes)
func (_Stakemanager *StakemanagerCaller) GetStakerInfo(opts *bind.CallOpts, staker common.Address, token uint8) (struct {
	Stakes   *big.Int
	Unstakes *big.Int
}, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getStakerInfo", staker, token)

	outstruct := new(struct {
		Stakes   *big.Int
		Unstakes *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Stakes = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Unstakes = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetStakerInfo is a free data retrieval call binding the contract method 0xe4a8a343.
//
// Solidity: function getStakerInfo(address staker, uint8 token) view returns(uint256 stakes, uint256 unstakes)
func (_Stakemanager *StakemanagerSession) GetStakerInfo(staker common.Address, token uint8) (struct {
	Stakes   *big.Int
	Unstakes *big.Int
}, error) {
	return _Stakemanager.Contract.GetStakerInfo(&_Stakemanager.CallOpts, staker, token)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0xe4a8a343.
//
// Solidity: function getStakerInfo(address staker, uint8 token) view returns(uint256 stakes, uint256 unstakes)
func (_Stakemanager *StakemanagerCallerSession) GetStakerInfo(staker common.Address, token uint8) (struct {
	Stakes   *big.Int
	Unstakes *big.Int
}, error) {
	return _Stakemanager.Contract.GetStakerInfo(&_Stakemanager.CallOpts, staker, token)
}

// GetStakerStakes is a free data retrieval call binding the contract method 0x533a339c.
//
// Solidity: function getStakerStakes(address staker, uint8 token, uint256 epoch) view returns(address[] _validators, uint256[] stakes, uint256[] stakeRequests, uint256[] unstakeRequests)
func (_Stakemanager *StakemanagerCaller) GetStakerStakes(opts *bind.CallOpts, staker common.Address, token uint8, epoch *big.Int) (struct {
	Validators      []common.Address
	Stakes          []*big.Int
	StakeRequests   []*big.Int
	UnstakeRequests []*big.Int
}, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getStakerStakes", staker, token, epoch)

	outstruct := new(struct {
		Validators      []common.Address
		Stakes          []*big.Int
		StakeRequests   []*big.Int
		UnstakeRequests []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Validators = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Stakes = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.StakeRequests = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	outstruct.UnstakeRequests = *abi.ConvertType(out[3], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetStakerStakes is a free data retrieval call binding the contract method 0x533a339c.
//
// Solidity: function getStakerStakes(address staker, uint8 token, uint256 epoch) view returns(address[] _validators, uint256[] stakes, uint256[] stakeRequests, uint256[] unstakeRequests)
func (_Stakemanager *StakemanagerSession) GetStakerStakes(staker common.Address, token uint8, epoch *big.Int) (struct {
	Validators      []common.Address
	Stakes          []*big.Int
	StakeRequests   []*big.Int
	UnstakeRequests []*big.Int
}, error) {
	return _Stakemanager.Contract.GetStakerStakes(&_Stakemanager.CallOpts, staker, token, epoch)
}

// GetStakerStakes is a free data retrieval call binding the contract method 0x533a339c.
//
// Solidity: function getStakerStakes(address staker, uint8 token, uint256 epoch) view returns(address[] _validators, uint256[] stakes, uint256[] stakeRequests, uint256[] unstakeRequests)
func (_Stakemanager *StakemanagerCallerSession) GetStakerStakes(staker common.Address, token uint8, epoch *big.Int) (struct {
	Validators      []common.Address
	Stakes          []*big.Int
	StakeRequests   []*big.Int
	UnstakeRequests []*big.Int
}, error) {
	return _Stakemanager.Contract.GetStakerStakes(&_Stakemanager.CallOpts, staker, token, epoch)
}

// GetStakers is a free data retrieval call binding the contract method 0xad71bd36.
//
// Solidity: function getStakers(uint256 page, uint256 perPage) view returns(address[])
func (_Stakemanager *StakemanagerCaller) GetStakers(opts *bind.CallOpts, page *big.Int, perPage *big.Int) ([]common.Address, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getStakers", page, perPage)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetStakers is a free data retrieval call binding the contract method 0xad71bd36.
//
// Solidity: function getStakers(uint256 page, uint256 perPage) view returns(address[])
func (_Stakemanager *StakemanagerSession) GetStakers(page *big.Int, perPage *big.Int) ([]common.Address, error) {
	return _Stakemanager.Contract.GetStakers(&_Stakemanager.CallOpts, page, perPage)
}

// GetStakers is a free data retrieval call binding the contract method 0xad71bd36.
//
// Solidity: function getStakers(uint256 page, uint256 perPage) view returns(address[])
func (_Stakemanager *StakemanagerCallerSession) GetStakers(page *big.Int, perPage *big.Int) ([]common.Address, error) {
	return _Stakemanager.Contract.GetStakers(&_Stakemanager.CallOpts, page, perPage)
}

// GetTotalRewards is a free data retrieval call binding the contract method 0x552164ee.
//
// Solidity: function getTotalRewards(uint256 epochs) view returns(uint256 rewards)
func (_Stakemanager *StakemanagerCaller) GetTotalRewards(opts *bind.CallOpts, epochs *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getTotalRewards", epochs)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalRewards is a free data retrieval call binding the contract method 0x552164ee.
//
// Solidity: function getTotalRewards(uint256 epochs) view returns(uint256 rewards)
func (_Stakemanager *StakemanagerSession) GetTotalRewards(epochs *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.GetTotalRewards(&_Stakemanager.CallOpts, epochs)
}

// GetTotalRewards is a free data retrieval call binding the contract method 0x552164ee.
//
// Solidity: function getTotalRewards(uint256 epochs) view returns(uint256 rewards)
func (_Stakemanager *StakemanagerCallerSession) GetTotalRewards(epochs *big.Int) (*big.Int, error) {
	return _Stakemanager.Contract.GetTotalRewards(&_Stakemanager.CallOpts, epochs)
}

// GetValidatorInfo is a free data retrieval call binding the contract method 0x8a11d7c9.
//
// Solidity: function getValidatorInfo(address validator) view returns(address operator, bool active, bool jailed, uint256 stakes, uint256 commissionRate)
func (_Stakemanager *StakemanagerCaller) GetValidatorInfo(opts *bind.CallOpts, validator common.Address) (struct {
	Operator       common.Address
	Active         bool
	Jailed         bool
	Stakes         *big.Int
	CommissionRate *big.Int
}, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getValidatorInfo", validator)

	outstruct := new(struct {
		Operator       common.Address
		Active         bool
		Jailed         bool
		Stakes         *big.Int
		CommissionRate *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Operator = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Active = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Jailed = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.Stakes = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.CommissionRate = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetValidatorInfo is a free data retrieval call binding the contract method 0x8a11d7c9.
//
// Solidity: function getValidatorInfo(address validator) view returns(address operator, bool active, bool jailed, uint256 stakes, uint256 commissionRate)
func (_Stakemanager *StakemanagerSession) GetValidatorInfo(validator common.Address) (struct {
	Operator       common.Address
	Active         bool
	Jailed         bool
	Stakes         *big.Int
	CommissionRate *big.Int
}, error) {
	return _Stakemanager.Contract.GetValidatorInfo(&_Stakemanager.CallOpts, validator)
}

// GetValidatorInfo is a free data retrieval call binding the contract method 0x8a11d7c9.
//
// Solidity: function getValidatorInfo(address validator) view returns(address operator, bool active, bool jailed, uint256 stakes, uint256 commissionRate)
func (_Stakemanager *StakemanagerCallerSession) GetValidatorInfo(validator common.Address) (struct {
	Operator       common.Address
	Active         bool
	Jailed         bool
	Stakes         *big.Int
	CommissionRate *big.Int
}, error) {
	return _Stakemanager.Contract.GetValidatorInfo(&_Stakemanager.CallOpts, validator)
}

// GetValidatorInfo0 is a free data retrieval call binding the contract method 0xd1f18ee1.
//
// Solidity: function getValidatorInfo(address validator, uint256 epoch) view returns(bool active, bool jailed, uint256 stakes, uint256 commissionRate)
func (_Stakemanager *StakemanagerCaller) GetValidatorInfo0(opts *bind.CallOpts, validator common.Address, epoch *big.Int) (struct {
	Active         bool
	Jailed         bool
	Stakes         *big.Int
	CommissionRate *big.Int
}, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getValidatorInfo0", validator, epoch)

	outstruct := new(struct {
		Active         bool
		Jailed         bool
		Stakes         *big.Int
		CommissionRate *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Active = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Jailed = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Stakes = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CommissionRate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetValidatorInfo0 is a free data retrieval call binding the contract method 0xd1f18ee1.
//
// Solidity: function getValidatorInfo(address validator, uint256 epoch) view returns(bool active, bool jailed, uint256 stakes, uint256 commissionRate)
func (_Stakemanager *StakemanagerSession) GetValidatorInfo0(validator common.Address, epoch *big.Int) (struct {
	Active         bool
	Jailed         bool
	Stakes         *big.Int
	CommissionRate *big.Int
}, error) {
	return _Stakemanager.Contract.GetValidatorInfo0(&_Stakemanager.CallOpts, validator, epoch)
}

// GetValidatorInfo0 is a free data retrieval call binding the contract method 0xd1f18ee1.
//
// Solidity: function getValidatorInfo(address validator, uint256 epoch) view returns(bool active, bool jailed, uint256 stakes, uint256 commissionRate)
func (_Stakemanager *StakemanagerCallerSession) GetValidatorInfo0(validator common.Address, epoch *big.Int) (struct {
	Active         bool
	Jailed         bool
	Stakes         *big.Int
	CommissionRate *big.Int
}, error) {
	return _Stakemanager.Contract.GetValidatorInfo0(&_Stakemanager.CallOpts, validator, epoch)
}

// GetValidatorStakes is a free data retrieval call binding the contract method 0x46dfce7b.
//
// Solidity: function getValidatorStakes(address validator, uint256 epoch, uint256 page, uint256 perPage) view returns(address[] _stakers, uint256[] stakes)
func (_Stakemanager *StakemanagerCaller) GetValidatorStakes(opts *bind.CallOpts, validator common.Address, epoch *big.Int, page *big.Int, perPage *big.Int) (struct {
	Stakers []common.Address
	Stakes  []*big.Int
}, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getValidatorStakes", validator, epoch, page, perPage)

	outstruct := new(struct {
		Stakers []common.Address
		Stakes  []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Stakers = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Stakes = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetValidatorStakes is a free data retrieval call binding the contract method 0x46dfce7b.
//
// Solidity: function getValidatorStakes(address validator, uint256 epoch, uint256 page, uint256 perPage) view returns(address[] _stakers, uint256[] stakes)
func (_Stakemanager *StakemanagerSession) GetValidatorStakes(validator common.Address, epoch *big.Int, page *big.Int, perPage *big.Int) (struct {
	Stakers []common.Address
	Stakes  []*big.Int
}, error) {
	return _Stakemanager.Contract.GetValidatorStakes(&_Stakemanager.CallOpts, validator, epoch, page, perPage)
}

// GetValidatorStakes is a free data retrieval call binding the contract method 0x46dfce7b.
//
// Solidity: function getValidatorStakes(address validator, uint256 epoch, uint256 page, uint256 perPage) view returns(address[] _stakers, uint256[] stakes)
func (_Stakemanager *StakemanagerCallerSession) GetValidatorStakes(validator common.Address, epoch *big.Int, page *big.Int, perPage *big.Int) (struct {
	Stakers []common.Address
	Stakes  []*big.Int
}, error) {
	return _Stakemanager.Contract.GetValidatorStakes(&_Stakemanager.CallOpts, validator, epoch, page, perPage)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_Stakemanager *StakemanagerCaller) GetValidators(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getValidators")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_Stakemanager *StakemanagerSession) GetValidators() ([]common.Address, error) {
	return _Stakemanager.Contract.GetValidators(&_Stakemanager.CallOpts)
}

// GetValidators is a free data retrieval call binding the contract method 0xb7ab4db5.
//
// Solidity: function getValidators() view returns(address[])
func (_Stakemanager *StakemanagerCallerSession) GetValidators() ([]common.Address, error) {
	return _Stakemanager.Contract.GetValidators(&_Stakemanager.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Stakemanager *StakemanagerCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Stakemanager *StakemanagerSession) Initialized() (bool, error) {
	return _Stakemanager.Contract.Initialized(&_Stakemanager.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Stakemanager *StakemanagerCallerSession) Initialized() (bool, error) {
	return _Stakemanager.Contract.Initialized(&_Stakemanager.CallOpts)
}

// OperatorToOwner is a free data retrieval call binding the contract method 0x7b520aa8.
//
// Solidity: function operatorToOwner(address ) view returns(address)
func (_Stakemanager *StakemanagerCaller) OperatorToOwner(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "operatorToOwner", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OperatorToOwner is a free data retrieval call binding the contract method 0x7b520aa8.
//
// Solidity: function operatorToOwner(address ) view returns(address)
func (_Stakemanager *StakemanagerSession) OperatorToOwner(arg0 common.Address) (common.Address, error) {
	return _Stakemanager.Contract.OperatorToOwner(&_Stakemanager.CallOpts, arg0)
}

// OperatorToOwner is a free data retrieval call binding the contract method 0x7b520aa8.
//
// Solidity: function operatorToOwner(address ) view returns(address)
func (_Stakemanager *StakemanagerCallerSession) OperatorToOwner(arg0 common.Address) (common.Address, error) {
	return _Stakemanager.Contract.OperatorToOwner(&_Stakemanager.CallOpts, arg0)
}

// StakerSigners is a free data retrieval call binding the contract method 0xf65a5ed2.
//
// Solidity: function stakerSigners(uint256 ) view returns(address)
func (_Stakemanager *StakemanagerCaller) StakerSigners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "stakerSigners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// StakerSigners is a free data retrieval call binding the contract method 0xf65a5ed2.
//
// Solidity: function stakerSigners(uint256 ) view returns(address)
func (_Stakemanager *StakemanagerSession) StakerSigners(arg0 *big.Int) (common.Address, error) {
	return _Stakemanager.Contract.StakerSigners(&_Stakemanager.CallOpts, arg0)
}

// StakerSigners is a free data retrieval call binding the contract method 0xf65a5ed2.
//
// Solidity: function stakerSigners(uint256 ) view returns(address)
func (_Stakemanager *StakemanagerCallerSession) StakerSigners(arg0 *big.Int) (common.Address, error) {
	return _Stakemanager.Contract.StakerSigners(&_Stakemanager.CallOpts, arg0)
}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address ) view returns(address signer)
func (_Stakemanager *StakemanagerCaller) Stakers(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "stakers", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address ) view returns(address signer)
func (_Stakemanager *StakemanagerSession) Stakers(arg0 common.Address) (common.Address, error) {
	return _Stakemanager.Contract.Stakers(&_Stakemanager.CallOpts, arg0)
}

// Stakers is a free data retrieval call binding the contract method 0x9168ae72.
//
// Solidity: function stakers(address ) view returns(address signer)
func (_Stakemanager *StakemanagerCallerSession) Stakers(arg0 common.Address) (common.Address, error) {
	return _Stakemanager.Contract.Stakers(&_Stakemanager.CallOpts, arg0)
}

// ValidatorOwners is a free data retrieval call binding the contract method 0x5efc766e.
//
// Solidity: function validatorOwners(uint256 ) view returns(address)
func (_Stakemanager *StakemanagerCaller) ValidatorOwners(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "validatorOwners", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorOwners is a free data retrieval call binding the contract method 0x5efc766e.
//
// Solidity: function validatorOwners(uint256 ) view returns(address)
func (_Stakemanager *StakemanagerSession) ValidatorOwners(arg0 *big.Int) (common.Address, error) {
	return _Stakemanager.Contract.ValidatorOwners(&_Stakemanager.CallOpts, arg0)
}

// ValidatorOwners is a free data retrieval call binding the contract method 0x5efc766e.
//
// Solidity: function validatorOwners(uint256 ) view returns(address)
func (_Stakemanager *StakemanagerCallerSession) ValidatorOwners(arg0 *big.Int) (common.Address, error) {
	return _Stakemanager.Contract.ValidatorOwners(&_Stakemanager.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(address owner, address operator, uint256 lastClaimCommission)
func (_Stakemanager *StakemanagerCaller) Validators(opts *bind.CallOpts, arg0 common.Address) (struct {
	Owner               common.Address
	Operator            common.Address
	LastClaimCommission *big.Int
}, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "validators", arg0)

	outstruct := new(struct {
		Owner               common.Address
		Operator            common.Address
		LastClaimCommission *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Operator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.LastClaimCommission = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(address owner, address operator, uint256 lastClaimCommission)
func (_Stakemanager *StakemanagerSession) Validators(arg0 common.Address) (struct {
	Owner               common.Address
	Operator            common.Address
	LastClaimCommission *big.Int
}, error) {
	return _Stakemanager.Contract.Validators(&_Stakemanager.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(address owner, address operator, uint256 lastClaimCommission)
func (_Stakemanager *StakemanagerCallerSession) Validators(arg0 common.Address) (struct {
	Owner               common.Address
	Operator            common.Address
	LastClaimCommission *big.Int
}, error) {
	return _Stakemanager.Contract.Validators(&_Stakemanager.CallOpts, arg0)
}

// ActivateValidator is a paid mutator transaction binding the contract method 0x1903cf16.
//
// Solidity: function activateValidator(address validator, uint256[] epochs) returns()
func (_Stakemanager *StakemanagerTransactor) ActivateValidator(opts *bind.TransactOpts, validator common.Address, epochs []*big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "activateValidator", validator, epochs)
}

// ActivateValidator is a paid mutator transaction binding the contract method 0x1903cf16.
//
// Solidity: function activateValidator(address validator, uint256[] epochs) returns()
func (_Stakemanager *StakemanagerSession) ActivateValidator(validator common.Address, epochs []*big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.ActivateValidator(&_Stakemanager.TransactOpts, validator, epochs)
}

// ActivateValidator is a paid mutator transaction binding the contract method 0x1903cf16.
//
// Solidity: function activateValidator(address validator, uint256[] epochs) returns()
func (_Stakemanager *StakemanagerTransactorSession) ActivateValidator(validator common.Address, epochs []*big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.ActivateValidator(&_Stakemanager.TransactOpts, validator, epochs)
}

// ClaimCommissions is a paid mutator transaction binding the contract method 0xcbc0fac6.
//
// Solidity: function claimCommissions(address validator, uint256 epochs) returns()
func (_Stakemanager *StakemanagerTransactor) ClaimCommissions(opts *bind.TransactOpts, validator common.Address, epochs *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "claimCommissions", validator, epochs)
}

// ClaimCommissions is a paid mutator transaction binding the contract method 0xcbc0fac6.
//
// Solidity: function claimCommissions(address validator, uint256 epochs) returns()
func (_Stakemanager *StakemanagerSession) ClaimCommissions(validator common.Address, epochs *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.ClaimCommissions(&_Stakemanager.TransactOpts, validator, epochs)
}

// ClaimCommissions is a paid mutator transaction binding the contract method 0xcbc0fac6.
//
// Solidity: function claimCommissions(address validator, uint256 epochs) returns()
func (_Stakemanager *StakemanagerTransactorSession) ClaimCommissions(validator common.Address, epochs *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.ClaimCommissions(&_Stakemanager.TransactOpts, validator, epochs)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x9a99b4f0.
//
// Solidity: function claimRewards(address validator, uint256 epochs) returns()
func (_Stakemanager *StakemanagerTransactor) ClaimRewards(opts *bind.TransactOpts, validator common.Address, epochs *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "claimRewards", validator, epochs)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x9a99b4f0.
//
// Solidity: function claimRewards(address validator, uint256 epochs) returns()
func (_Stakemanager *StakemanagerSession) ClaimRewards(validator common.Address, epochs *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.ClaimRewards(&_Stakemanager.TransactOpts, validator, epochs)
}

// ClaimRewards is a paid mutator transaction binding the contract method 0x9a99b4f0.
//
// Solidity: function claimRewards(address validator, uint256 epochs) returns()
func (_Stakemanager *StakemanagerTransactorSession) ClaimRewards(validator common.Address, epochs *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.ClaimRewards(&_Stakemanager.TransactOpts, validator, epochs)
}

// ClaimUnstakes is a paid mutator transaction binding the contract method 0x2ffb688d.
//
// Solidity: function claimUnstakes() returns()
func (_Stakemanager *StakemanagerTransactor) ClaimUnstakes(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "claimUnstakes")
}

// ClaimUnstakes is a paid mutator transaction binding the contract method 0x2ffb688d.
//
// Solidity: function claimUnstakes() returns()
func (_Stakemanager *StakemanagerSession) ClaimUnstakes() (*types.Transaction, error) {
	return _Stakemanager.Contract.ClaimUnstakes(&_Stakemanager.TransactOpts)
}

// ClaimUnstakes is a paid mutator transaction binding the contract method 0x2ffb688d.
//
// Solidity: function claimUnstakes() returns()
func (_Stakemanager *StakemanagerTransactorSession) ClaimUnstakes() (*types.Transaction, error) {
	return _Stakemanager.Contract.ClaimUnstakes(&_Stakemanager.TransactOpts)
}

// DeactivateValidator is a paid mutator transaction binding the contract method 0x428e8562.
//
// Solidity: function deactivateValidator(address validator, uint256[] epochs) returns()
func (_Stakemanager *StakemanagerTransactor) DeactivateValidator(opts *bind.TransactOpts, validator common.Address, epochs []*big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "deactivateValidator", validator, epochs)
}

// DeactivateValidator is a paid mutator transaction binding the contract method 0x428e8562.
//
// Solidity: function deactivateValidator(address validator, uint256[] epochs) returns()
func (_Stakemanager *StakemanagerSession) DeactivateValidator(validator common.Address, epochs []*big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.DeactivateValidator(&_Stakemanager.TransactOpts, validator, epochs)
}

// DeactivateValidator is a paid mutator transaction binding the contract method 0x428e8562.
//
// Solidity: function deactivateValidator(address validator, uint256[] epochs) returns()
func (_Stakemanager *StakemanagerTransactorSession) DeactivateValidator(validator common.Address, epochs []*big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.DeactivateValidator(&_Stakemanager.TransactOpts, validator, epochs)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _environment, address _allowlist) returns()
func (_Stakemanager *StakemanagerTransactor) Initialize(opts *bind.TransactOpts, _environment common.Address, _allowlist common.Address) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "initialize", _environment, _allowlist)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _environment, address _allowlist) returns()
func (_Stakemanager *StakemanagerSession) Initialize(_environment common.Address, _allowlist common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.Initialize(&_Stakemanager.TransactOpts, _environment, _allowlist)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address _environment, address _allowlist) returns()
func (_Stakemanager *StakemanagerTransactorSession) Initialize(_environment common.Address, _allowlist common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.Initialize(&_Stakemanager.TransactOpts, _environment, _allowlist)
}

// JoinValidator is a paid mutator transaction binding the contract method 0x6b2b3369.
//
// Solidity: function joinValidator(address operator) returns()
func (_Stakemanager *StakemanagerTransactor) JoinValidator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "joinValidator", operator)
}

// JoinValidator is a paid mutator transaction binding the contract method 0x6b2b3369.
//
// Solidity: function joinValidator(address operator) returns()
func (_Stakemanager *StakemanagerSession) JoinValidator(operator common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.JoinValidator(&_Stakemanager.TransactOpts, operator)
}

// JoinValidator is a paid mutator transaction binding the contract method 0x6b2b3369.
//
// Solidity: function joinValidator(address operator) returns()
func (_Stakemanager *StakemanagerTransactorSession) JoinValidator(operator common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.JoinValidator(&_Stakemanager.TransactOpts, operator)
}

// Slash is a paid mutator transaction binding the contract method 0x02fb4d85.
//
// Solidity: function slash(address operator, uint256 blocks) returns()
func (_Stakemanager *StakemanagerTransactor) Slash(opts *bind.TransactOpts, operator common.Address, blocks *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "slash", operator, blocks)
}

// Slash is a paid mutator transaction binding the contract method 0x02fb4d85.
//
// Solidity: function slash(address operator, uint256 blocks) returns()
func (_Stakemanager *StakemanagerSession) Slash(operator common.Address, blocks *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.Slash(&_Stakemanager.TransactOpts, operator, blocks)
}

// Slash is a paid mutator transaction binding the contract method 0x02fb4d85.
//
// Solidity: function slash(address operator, uint256 blocks) returns()
func (_Stakemanager *StakemanagerTransactorSession) Slash(operator common.Address, blocks *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.Slash(&_Stakemanager.TransactOpts, operator, blocks)
}

// Stake is a paid mutator transaction binding the contract method 0x7befa74f.
//
// Solidity: function stake(address validator, uint8 token, uint256 amount) payable returns()
func (_Stakemanager *StakemanagerTransactor) Stake(opts *bind.TransactOpts, validator common.Address, token uint8, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "stake", validator, token, amount)
}

// Stake is a paid mutator transaction binding the contract method 0x7befa74f.
//
// Solidity: function stake(address validator, uint8 token, uint256 amount) payable returns()
func (_Stakemanager *StakemanagerSession) Stake(validator common.Address, token uint8, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.Stake(&_Stakemanager.TransactOpts, validator, token, amount)
}

// Stake is a paid mutator transaction binding the contract method 0x7befa74f.
//
// Solidity: function stake(address validator, uint8 token, uint256 amount) payable returns()
func (_Stakemanager *StakemanagerTransactorSession) Stake(validator common.Address, token uint8, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.Stake(&_Stakemanager.TransactOpts, validator, token, amount)
}

// Unstake is a paid mutator transaction binding the contract method 0xe1aca341.
//
// Solidity: function unstake(address validator, uint8 token, uint256 amount) returns()
func (_Stakemanager *StakemanagerTransactor) Unstake(opts *bind.TransactOpts, validator common.Address, token uint8, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "unstake", validator, token, amount)
}

// Unstake is a paid mutator transaction binding the contract method 0xe1aca341.
//
// Solidity: function unstake(address validator, uint8 token, uint256 amount) returns()
func (_Stakemanager *StakemanagerSession) Unstake(validator common.Address, token uint8, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.Unstake(&_Stakemanager.TransactOpts, validator, token, amount)
}

// Unstake is a paid mutator transaction binding the contract method 0xe1aca341.
//
// Solidity: function unstake(address validator, uint8 token, uint256 amount) returns()
func (_Stakemanager *StakemanagerTransactorSession) Unstake(validator common.Address, token uint8, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.Unstake(&_Stakemanager.TransactOpts, validator, token, amount)
}

// UpdateCommissionRate is a paid mutator transaction binding the contract method 0x00fa3d50.
//
// Solidity: function updateCommissionRate(uint256 newRate) returns()
func (_Stakemanager *StakemanagerTransactor) UpdateCommissionRate(opts *bind.TransactOpts, newRate *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "updateCommissionRate", newRate)
}

// UpdateCommissionRate is a paid mutator transaction binding the contract method 0x00fa3d50.
//
// Solidity: function updateCommissionRate(uint256 newRate) returns()
func (_Stakemanager *StakemanagerSession) UpdateCommissionRate(newRate *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.UpdateCommissionRate(&_Stakemanager.TransactOpts, newRate)
}

// UpdateCommissionRate is a paid mutator transaction binding the contract method 0x00fa3d50.
//
// Solidity: function updateCommissionRate(uint256 newRate) returns()
func (_Stakemanager *StakemanagerTransactorSession) UpdateCommissionRate(newRate *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.UpdateCommissionRate(&_Stakemanager.TransactOpts, newRate)
}

// UpdateOperator is a paid mutator transaction binding the contract method 0xac7475ed.
//
// Solidity: function updateOperator(address operator) returns()
func (_Stakemanager *StakemanagerTransactor) UpdateOperator(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "updateOperator", operator)
}

// UpdateOperator is a paid mutator transaction binding the contract method 0xac7475ed.
//
// Solidity: function updateOperator(address operator) returns()
func (_Stakemanager *StakemanagerSession) UpdateOperator(operator common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.UpdateOperator(&_Stakemanager.TransactOpts, operator)
}

// UpdateOperator is a paid mutator transaction binding the contract method 0xac7475ed.
//
// Solidity: function updateOperator(address operator) returns()
func (_Stakemanager *StakemanagerTransactorSession) UpdateOperator(operator common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.UpdateOperator(&_Stakemanager.TransactOpts, operator)
}

// StakemanagerStakedIterator is returned from FilterStaked and is used to iterate over the raw logs and unpacked data for Staked events raised by the Stakemanager contract.
type StakemanagerStakedIterator struct {
	Event *StakemanagerStaked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakemanagerStakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerStaked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakemanagerStaked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakemanagerStakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerStakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerStaked represents a Staked event raised by the Stakemanager contract.
type StakemanagerStaked struct {
	Staker    common.Address
	Validator common.Address
	Token     uint8
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x8fc656e319452025372383dc27d933046d412b8253de50a10739eeaa59862be6.
//
// Solidity: event Staked(address indexed staker, address indexed validator, uint8 token, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) FilterStaked(opts *bind.FilterOpts, staker []common.Address, validator []common.Address) (*StakemanagerStakedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "Staked", stakerRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerStakedIterator{contract: _Stakemanager.contract, event: "Staked", logs: logs, sub: sub}, nil
}

// WatchStaked is a free log subscription operation binding the contract event 0x8fc656e319452025372383dc27d933046d412b8253de50a10739eeaa59862be6.
//
// Solidity: event Staked(address indexed staker, address indexed validator, uint8 token, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) WatchStaked(opts *bind.WatchOpts, sink chan<- *StakemanagerStaked, staker []common.Address, validator []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "Staked", stakerRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerStaked)
				if err := _Stakemanager.contract.UnpackLog(event, "Staked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseStaked is a log parse operation binding the contract event 0x8fc656e319452025372383dc27d933046d412b8253de50a10739eeaa59862be6.
//
// Solidity: event Staked(address indexed staker, address indexed validator, uint8 token, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) ParseStaked(log types.Log) (*StakemanagerStaked, error) {
	event := new(StakemanagerStaked)
	if err := _Stakemanager.contract.UnpackLog(event, "Staked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerUnstakedIterator is returned from FilterUnstaked and is used to iterate over the raw logs and unpacked data for Unstaked events raised by the Stakemanager contract.
type StakemanagerUnstakedIterator struct {
	Event *StakemanagerUnstaked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakemanagerUnstakedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerUnstaked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakemanagerUnstaked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakemanagerUnstakedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerUnstakedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerUnstaked represents a Unstaked event raised by the Stakemanager contract.
type StakemanagerUnstaked struct {
	Staker    common.Address
	Validator common.Address
	Token     uint8
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0xf2812c3df2511a467cbe777b1ee98b1ddb9918bb0f09568a269d2fb58233cb52.
//
// Solidity: event Unstaked(address indexed staker, address indexed validator, uint8 token, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) FilterUnstaked(opts *bind.FilterOpts, staker []common.Address, validator []common.Address) (*StakemanagerUnstakedIterator, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "Unstaked", stakerRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerUnstakedIterator{contract: _Stakemanager.contract, event: "Unstaked", logs: logs, sub: sub}, nil
}

// WatchUnstaked is a free log subscription operation binding the contract event 0xf2812c3df2511a467cbe777b1ee98b1ddb9918bb0f09568a269d2fb58233cb52.
//
// Solidity: event Unstaked(address indexed staker, address indexed validator, uint8 token, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) WatchUnstaked(opts *bind.WatchOpts, sink chan<- *StakemanagerUnstaked, staker []common.Address, validator []common.Address) (event.Subscription, error) {

	var stakerRule []interface{}
	for _, stakerItem := range staker {
		stakerRule = append(stakerRule, stakerItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "Unstaked", stakerRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerUnstaked)
				if err := _Stakemanager.contract.UnpackLog(event, "Unstaked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnstaked is a log parse operation binding the contract event 0xf2812c3df2511a467cbe777b1ee98b1ddb9918bb0f09568a269d2fb58233cb52.
//
// Solidity: event Unstaked(address indexed staker, address indexed validator, uint8 token, uint256 amount)
func (_Stakemanager *StakemanagerFilterer) ParseUnstaked(log types.Log) (*StakemanagerUnstaked, error) {
	event := new(StakemanagerUnstaked)
	if err := _Stakemanager.contract.UnpackLog(event, "Unstaked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerValidatorActivatedIterator is returned from FilterValidatorActivated and is used to iterate over the raw logs and unpacked data for ValidatorActivated events raised by the Stakemanager contract.
type StakemanagerValidatorActivatedIterator struct {
	Event *StakemanagerValidatorActivated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakemanagerValidatorActivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerValidatorActivated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakemanagerValidatorActivated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakemanagerValidatorActivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerValidatorActivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerValidatorActivated represents a ValidatorActivated event raised by the Stakemanager contract.
type StakemanagerValidatorActivated struct {
	Validator common.Address
	Epochs    []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorActivated is a free log retrieval operation binding the contract event 0xc11dfc9c24621433bb10587dc4bbae26a33a4aff53914e0d4c9fddf224a8cb7d.
//
// Solidity: event ValidatorActivated(address indexed validator, uint256[] epochs)
func (_Stakemanager *StakemanagerFilterer) FilterValidatorActivated(opts *bind.FilterOpts, validator []common.Address) (*StakemanagerValidatorActivatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "ValidatorActivated", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerValidatorActivatedIterator{contract: _Stakemanager.contract, event: "ValidatorActivated", logs: logs, sub: sub}, nil
}

// WatchValidatorActivated is a free log subscription operation binding the contract event 0xc11dfc9c24621433bb10587dc4bbae26a33a4aff53914e0d4c9fddf224a8cb7d.
//
// Solidity: event ValidatorActivated(address indexed validator, uint256[] epochs)
func (_Stakemanager *StakemanagerFilterer) WatchValidatorActivated(opts *bind.WatchOpts, sink chan<- *StakemanagerValidatorActivated, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "ValidatorActivated", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerValidatorActivated)
				if err := _Stakemanager.contract.UnpackLog(event, "ValidatorActivated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorActivated is a log parse operation binding the contract event 0xc11dfc9c24621433bb10587dc4bbae26a33a4aff53914e0d4c9fddf224a8cb7d.
//
// Solidity: event ValidatorActivated(address indexed validator, uint256[] epochs)
func (_Stakemanager *StakemanagerFilterer) ParseValidatorActivated(log types.Log) (*StakemanagerValidatorActivated, error) {
	event := new(StakemanagerValidatorActivated)
	if err := _Stakemanager.contract.UnpackLog(event, "ValidatorActivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerValidatorCommissionRateUpdatedIterator is returned from FilterValidatorCommissionRateUpdated and is used to iterate over the raw logs and unpacked data for ValidatorCommissionRateUpdated events raised by the Stakemanager contract.
type StakemanagerValidatorCommissionRateUpdatedIterator struct {
	Event *StakemanagerValidatorCommissionRateUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakemanagerValidatorCommissionRateUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerValidatorCommissionRateUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakemanagerValidatorCommissionRateUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakemanagerValidatorCommissionRateUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerValidatorCommissionRateUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerValidatorCommissionRateUpdated represents a ValidatorCommissionRateUpdated event raised by the Stakemanager contract.
type StakemanagerValidatorCommissionRateUpdated struct {
	Validator common.Address
	Rate      *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorCommissionRateUpdated is a free log retrieval operation binding the contract event 0x658f52b34d965c76f54298d43b21ae83217d77033214bdafed998bf6df6b9b05.
//
// Solidity: event ValidatorCommissionRateUpdated(address indexed validator, uint256 rate)
func (_Stakemanager *StakemanagerFilterer) FilterValidatorCommissionRateUpdated(opts *bind.FilterOpts, validator []common.Address) (*StakemanagerValidatorCommissionRateUpdatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "ValidatorCommissionRateUpdated", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerValidatorCommissionRateUpdatedIterator{contract: _Stakemanager.contract, event: "ValidatorCommissionRateUpdated", logs: logs, sub: sub}, nil
}

// WatchValidatorCommissionRateUpdated is a free log subscription operation binding the contract event 0x658f52b34d965c76f54298d43b21ae83217d77033214bdafed998bf6df6b9b05.
//
// Solidity: event ValidatorCommissionRateUpdated(address indexed validator, uint256 rate)
func (_Stakemanager *StakemanagerFilterer) WatchValidatorCommissionRateUpdated(opts *bind.WatchOpts, sink chan<- *StakemanagerValidatorCommissionRateUpdated, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "ValidatorCommissionRateUpdated", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerValidatorCommissionRateUpdated)
				if err := _Stakemanager.contract.UnpackLog(event, "ValidatorCommissionRateUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorCommissionRateUpdated is a log parse operation binding the contract event 0x658f52b34d965c76f54298d43b21ae83217d77033214bdafed998bf6df6b9b05.
//
// Solidity: event ValidatorCommissionRateUpdated(address indexed validator, uint256 rate)
func (_Stakemanager *StakemanagerFilterer) ParseValidatorCommissionRateUpdated(log types.Log) (*StakemanagerValidatorCommissionRateUpdated, error) {
	event := new(StakemanagerValidatorCommissionRateUpdated)
	if err := _Stakemanager.contract.UnpackLog(event, "ValidatorCommissionRateUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerValidatorDeactivatedIterator is returned from FilterValidatorDeactivated and is used to iterate over the raw logs and unpacked data for ValidatorDeactivated events raised by the Stakemanager contract.
type StakemanagerValidatorDeactivatedIterator struct {
	Event *StakemanagerValidatorDeactivated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakemanagerValidatorDeactivatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerValidatorDeactivated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakemanagerValidatorDeactivated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakemanagerValidatorDeactivatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerValidatorDeactivatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerValidatorDeactivated represents a ValidatorDeactivated event raised by the Stakemanager contract.
type StakemanagerValidatorDeactivated struct {
	Validator common.Address
	Epochs    []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorDeactivated is a free log retrieval operation binding the contract event 0x0ad9bf1b8c026a174a2f30954417002a6ea00c9b08c1b8846c7951c687be8095.
//
// Solidity: event ValidatorDeactivated(address indexed validator, uint256[] epochs)
func (_Stakemanager *StakemanagerFilterer) FilterValidatorDeactivated(opts *bind.FilterOpts, validator []common.Address) (*StakemanagerValidatorDeactivatedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "ValidatorDeactivated", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerValidatorDeactivatedIterator{contract: _Stakemanager.contract, event: "ValidatorDeactivated", logs: logs, sub: sub}, nil
}

// WatchValidatorDeactivated is a free log subscription operation binding the contract event 0x0ad9bf1b8c026a174a2f30954417002a6ea00c9b08c1b8846c7951c687be8095.
//
// Solidity: event ValidatorDeactivated(address indexed validator, uint256[] epochs)
func (_Stakemanager *StakemanagerFilterer) WatchValidatorDeactivated(opts *bind.WatchOpts, sink chan<- *StakemanagerValidatorDeactivated, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "ValidatorDeactivated", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerValidatorDeactivated)
				if err := _Stakemanager.contract.UnpackLog(event, "ValidatorDeactivated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorDeactivated is a log parse operation binding the contract event 0x0ad9bf1b8c026a174a2f30954417002a6ea00c9b08c1b8846c7951c687be8095.
//
// Solidity: event ValidatorDeactivated(address indexed validator, uint256[] epochs)
func (_Stakemanager *StakemanagerFilterer) ParseValidatorDeactivated(log types.Log) (*StakemanagerValidatorDeactivated, error) {
	event := new(StakemanagerValidatorDeactivated)
	if err := _Stakemanager.contract.UnpackLog(event, "ValidatorDeactivated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerValidatorJailedIterator is returned from FilterValidatorJailed and is used to iterate over the raw logs and unpacked data for ValidatorJailed events raised by the Stakemanager contract.
type StakemanagerValidatorJailedIterator struct {
	Event *StakemanagerValidatorJailed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakemanagerValidatorJailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerValidatorJailed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakemanagerValidatorJailed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakemanagerValidatorJailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerValidatorJailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerValidatorJailed represents a ValidatorJailed event raised by the Stakemanager contract.
type StakemanagerValidatorJailed struct {
	Validator common.Address
	Until     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorJailed is a free log retrieval operation binding the contract event 0xeb7d7a49847ec491969db21a0e31b234565a9923145a2d1b56a75c9e95825802.
//
// Solidity: event ValidatorJailed(address indexed validator, uint256 until)
func (_Stakemanager *StakemanagerFilterer) FilterValidatorJailed(opts *bind.FilterOpts, validator []common.Address) (*StakemanagerValidatorJailedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "ValidatorJailed", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerValidatorJailedIterator{contract: _Stakemanager.contract, event: "ValidatorJailed", logs: logs, sub: sub}, nil
}

// WatchValidatorJailed is a free log subscription operation binding the contract event 0xeb7d7a49847ec491969db21a0e31b234565a9923145a2d1b56a75c9e95825802.
//
// Solidity: event ValidatorJailed(address indexed validator, uint256 until)
func (_Stakemanager *StakemanagerFilterer) WatchValidatorJailed(opts *bind.WatchOpts, sink chan<- *StakemanagerValidatorJailed, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "ValidatorJailed", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerValidatorJailed)
				if err := _Stakemanager.contract.UnpackLog(event, "ValidatorJailed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorJailed is a log parse operation binding the contract event 0xeb7d7a49847ec491969db21a0e31b234565a9923145a2d1b56a75c9e95825802.
//
// Solidity: event ValidatorJailed(address indexed validator, uint256 until)
func (_Stakemanager *StakemanagerFilterer) ParseValidatorJailed(log types.Log) (*StakemanagerValidatorJailed, error) {
	event := new(StakemanagerValidatorJailed)
	if err := _Stakemanager.contract.UnpackLog(event, "ValidatorJailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakemanagerValidatorSlashedIterator is returned from FilterValidatorSlashed and is used to iterate over the raw logs and unpacked data for ValidatorSlashed events raised by the Stakemanager contract.
type StakemanagerValidatorSlashedIterator struct {
	Event *StakemanagerValidatorSlashed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakemanagerValidatorSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakemanagerValidatorSlashed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakemanagerValidatorSlashed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakemanagerValidatorSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakemanagerValidatorSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakemanagerValidatorSlashed represents a ValidatorSlashed event raised by the Stakemanager contract.
type StakemanagerValidatorSlashed struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorSlashed is a free log retrieval operation binding the contract event 0x1647efd0ce9727dc31dc201c9d8d35ac687f7370adcacbd454afc6485ddabfda.
//
// Solidity: event ValidatorSlashed(address indexed validator)
func (_Stakemanager *StakemanagerFilterer) FilterValidatorSlashed(opts *bind.FilterOpts, validator []common.Address) (*StakemanagerValidatorSlashedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.FilterLogs(opts, "ValidatorSlashed", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakemanagerValidatorSlashedIterator{contract: _Stakemanager.contract, event: "ValidatorSlashed", logs: logs, sub: sub}, nil
}

// WatchValidatorSlashed is a free log subscription operation binding the contract event 0x1647efd0ce9727dc31dc201c9d8d35ac687f7370adcacbd454afc6485ddabfda.
//
// Solidity: event ValidatorSlashed(address indexed validator)
func (_Stakemanager *StakemanagerFilterer) WatchValidatorSlashed(opts *bind.WatchOpts, sink chan<- *StakemanagerValidatorSlashed, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Stakemanager.contract.WatchLogs(opts, "ValidatorSlashed", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakemanagerValidatorSlashed)
				if err := _Stakemanager.contract.UnpackLog(event, "ValidatorSlashed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseValidatorSlashed is a log parse operation binding the contract event 0x1647efd0ce9727dc31dc201c9d8d35ac687f7370adcacbd454afc6485ddabfda.
//
// Solidity: event ValidatorSlashed(address indexed validator)
func (_Stakemanager *StakemanagerFilterer) ParseValidatorSlashed(log types.Log) (*StakemanagerValidatorSlashed, error) {
	event := new(StakemanagerValidatorSlashed)
	if err := _Stakemanager.contract.UnpackLog(event, "ValidatorSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
