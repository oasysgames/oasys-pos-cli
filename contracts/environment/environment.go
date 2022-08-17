// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package environment

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

// IEnvironmentEnvironmentValue is an auto generated low-level Go binding around an user-defined struct.
type IEnvironmentEnvironmentValue struct {
	StartBlock         *big.Int
	StartEpoch         *big.Int
	BlockPeriod        *big.Int
	EpochPeriod        *big.Int
	RewardRate         *big.Int
	ValidatorThreshold *big.Int
	JailThreshold      *big.Int
	JailPeriod         *big.Int
}

// EnvironmentMetaData contains all meta data concerning the Environment contract.
var EnvironmentMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyBlockProducer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotLastBlock\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PastEpoch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"detail\",\"type\":\"string\"}],\"name\":\"ValidationError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"epoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochAndValues\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"epochs\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailPeriod\",\"type\":\"uint256\"}],\"internalType\":\"structIEnvironment.EnvironmentValue[]\",\"name\":\"_values\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailPeriod\",\"type\":\"uint256\"}],\"internalType\":\"structIEnvironment.EnvironmentValue\",\"name\":\"initialValue\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isFirstBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isLastBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextValue\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailPeriod\",\"type\":\"uint256\"}],\"internalType\":\"structIEnvironment.EnvironmentValue\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailPeriod\",\"type\":\"uint256\"}],\"internalType\":\"structIEnvironment.EnvironmentValue\",\"name\":\"newValue\",\"type\":\"tuple\"}],\"name\":\"updateValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"updates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"value\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailPeriod\",\"type\":\"uint256\"}],\"internalType\":\"structIEnvironment.EnvironmentValue\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"values\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailPeriod\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061178b806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063900cf0cf11610071578063900cf0cf1461015b5780639cf370d714610179578063b4c2f72714610198578063ccfb9935146101c8578063d4a53686146101e6578063f38e08e114610204576100a9565b806311bff261146100ae578063158ef93e146100ca5780633fa4f245146100e85780634ba7ed39146101065780635e383d2114610124575b600080fd5b6100c860048036038101906100c39190610eab565b610220565b005b6100d261037e565b6040516100df9190610ef4565b60405180910390f35b6100f061038f565b6040516100fd9190610fc0565b60405180910390f35b61010e61042b565b60405161011b9190610ef4565b60405180910390f35b61013e60048036038101906101399190610fdc565b61044b565b604051610152989796959493929190611018565b60405180910390f35b6101636104a3565b6040516101709190611096565b60405180910390f35b6101816104ec565b60405161018f9291906112b2565b60405180910390f35b6101b260048036038101906101ad9190610fdc565b6105f7565b6040516101bf9190611096565b60405180910390f35b6101d061061b565b6040516101dd9190610fc0565b60405180910390f35b6101ee610703565b6040516101fb9190610ef4565b60405180910390f35b61021e60048036038101906102199190610eab565b61072f565b005b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610285576040517f1cf4735900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61028d610703565b156102c4576040517f1e59ccd900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6102cc6104a3565b816020015111610308576040517f0eae4c9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000610312610815565b9050610327438261084f90919063ffffffff16565b1561034d5761033f828261086190919063ffffffff16565b826000018181525050610371565b6103678261035961089d565b61086190919063ffffffff16565b8260000181815250505b61037a82610910565b5050565b60008054906101000a900460ff1681565b610397610cc6565b60006103a1610815565b90506103b6438261084f90919063ffffffff16565b6103c7576103c261089d565b6103c9565b805b604051806101000160405290816000820154815260200160018201548152602001600282015481526020016003820154815260200160048201548152602001600582015481526020016006820154815260200160078201548152505091505090565b60008061043661038f565b60600151436104459190611318565b14905090565b6002818154811061045b57600080fd5b90600052602060002090600802016000915090508060000154908060010154908060020154908060030154908060040154908060050154908060060154908060070154905088565b6000806104ae610815565b90506104c3438261084f90919063ffffffff16565b6104dc576104d76104d261089d565b610adc565b6104e6565b6104e581610adc565b5b91505090565b606080600160028180548060200260200160405190810160405280929190818152602001828054801561053e57602002820191906000526020600020905b81548152602001906001019080831161052a575b5050505050915080805480602002602001604051908101604052809291908181526020016000905b828210156105e95783829060005260206000209060080201604051806101000160405290816000820154815260200160018201548152602001600282015481526020016003820154815260200160048201548152602001600582015481526020016006820154815260200160078201548152505081526020019060010190610566565b505050509050915091509091565b6001818154811061060757600080fd5b906000526020600020016000915090505481565b610623610cc6565b600061062d61038f565b90506000610639610815565b905060008260600151600184602001516106516104a3565b61065b9190611378565b61066591906113ac565b61066f9190611402565b836000015161067e91906113ac565b9050610693818361084f90919063ffffffff16565b61069d57826106fb565b8160405180610100016040529081600082015481526020016001820154815260200160028201548152602001600382015481526020016004820154815260200160058201548152602001600682015481526020016007820154815250505b935050505090565b60008061070e61038f565b6060015160014361071f91906113ac565b6107299190611318565b14905090565b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610794576040517f1cf4735900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008054906101000a900460ff16156107d9576040517f0dc149f000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60016000806101000a81548160ff0219169083151502179055506000816000018181525050600181602001818152505061081281610910565b50565b60006002600160028054905061082b9190611378565b8154811061083c5761083b61145c565b5b9060005260206000209060080201905090565b60008260000154821015905092915050565b600082600301548360010154836020015161087c9190611378565b6108869190611402565b836000015461089591906113ac565b905092915050565b600080600280549050905060018114156108dc5760026000815481106108c6576108c561145c565b5b906000526020600020906008020191505061090d565b600280826108ea9190611378565b815481106108fb576108fa61145c565b5b90600052602060002090600802019150505b90565b61091981610b13565b60006001805490509050600081148061096e575061096d4360026001846109409190611378565b815481106109515761095061145c565b5b906000526020600020906008020161084f90919063ffffffff16565b5b15610a23576001826020015190806001815401808255809150506001900390600052602060002001600090919091909150556002829080600181540180825580915050600190039060005260206000209060080201600090919091909150600082015181600001556020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e082015181600701555050610ad8565b816020015160018083610a369190611378565b81548110610a4757610a4661145c565b5b9060005260206000200181905550816002600183610a659190611378565b81548110610a7657610a7561145c565b5b9060005260206000209060080201600082015181600001556020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e082015181600701559050505b5050565b60008160030154826000015443610af39190611378565b610afd919061148b565b8260010154610b0c91906113ac565b9050919050565b600181604001511015610b5b576040517fa1b63ceb000000000000000000000000000000000000000000000000000000008152600401610b5290611519565b60405180910390fd5b600381606001511015610ba3576040517fa1b63ceb000000000000000000000000000000000000000000000000000000008152600401610b9a90611585565b60405180910390fd5b606481608001511115610beb576040517fa1b63ceb000000000000000000000000000000000000000000000000000000008152600401610be2906115f1565b60405180910390fd5b60018160a001511015610c33576040517fa1b63ceb000000000000000000000000000000000000000000000000000000008152600401610c2a9061165d565b60405180910390fd5b60018160c001511015610c7b576040517fa1b63ceb000000000000000000000000000000000000000000000000000000008152600401610c72906116c9565b60405180910390fd5b60018160e001511015610cc3576040517fa1b63ceb000000000000000000000000000000000000000000000000000000008152600401610cba90611735565b60405180910390fd5b50565b60405180610100016040528060008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b6000604051905090565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610d6882610d1f565b810181811067ffffffffffffffff82111715610d8757610d86610d30565b5b80604052505050565b6000610d9a610d0b565b9050610da68282610d5f565b919050565b6000819050919050565b610dbe81610dab565b8114610dc957600080fd5b50565b600081359050610ddb81610db5565b92915050565b60006101008284031215610df857610df7610d1a565b5b610e03610100610d90565b90506000610e1384828501610dcc565b6000830152506020610e2784828501610dcc565b6020830152506040610e3b84828501610dcc565b6040830152506060610e4f84828501610dcc565b6060830152506080610e6384828501610dcc565b60808301525060a0610e7784828501610dcc565b60a08301525060c0610e8b84828501610dcc565b60c08301525060e0610e9f84828501610dcc565b60e08301525092915050565b60006101008284031215610ec257610ec1610d15565b5b6000610ed084828501610de1565b91505092915050565b60008115159050919050565b610eee81610ed9565b82525050565b6000602082019050610f096000830184610ee5565b92915050565b610f1881610dab565b82525050565b61010082016000820151610f356000850182610f0f565b506020820151610f486020850182610f0f565b506040820151610f5b6040850182610f0f565b506060820151610f6e6060850182610f0f565b506080820151610f816080850182610f0f565b5060a0820151610f9460a0850182610f0f565b5060c0820151610fa760c0850182610f0f565b5060e0820151610fba60e0850182610f0f565b50505050565b600061010082019050610fd66000830184610f1e565b92915050565b600060208284031215610ff257610ff1610d15565b5b600061100084828501610dcc565b91505092915050565b61101281610dab565b82525050565b60006101008201905061102e600083018b611009565b61103b602083018a611009565b6110486040830189611009565b6110556060830188611009565b6110626080830187611009565b61106f60a0830186611009565b61107c60c0830185611009565b61108960e0830184611009565b9998505050505050505050565b60006020820190506110ab6000830184611009565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b60006110e98383610f0f565b60208301905092915050565b6000602082019050919050565b600061110d826110b1565b61111781856110bc565b9350611122836110cd565b8060005b8381101561115357815161113a88826110dd565b9750611145836110f5565b925050600181019050611126565b5085935050505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610100820160008201516111a36000850182610f0f565b5060208201516111b66020850182610f0f565b5060408201516111c96040850182610f0f565b5060608201516111dc6060850182610f0f565b5060808201516111ef6080850182610f0f565b5060a082015161120260a0850182610f0f565b5060c082015161121560c0850182610f0f565b5060e082015161122860e0850182610f0f565b50505050565b600061123a838361118c565b6101008301905092915050565b6000602082019050919050565b600061125f82611160565b611269818561116b565b93506112748361117c565b8060005b838110156112a557815161128c888261122e565b975061129783611247565b925050600181019050611278565b5085935050505092915050565b600060408201905081810360008301526112cc8185611102565b905081810360208301526112e08184611254565b90509392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061132382610dab565b915061132e83610dab565b92508261133e5761133d6112e9565b5b828206905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061138382610dab565b915061138e83610dab565b9250828210156113a1576113a0611349565b5b828203905092915050565b60006113b782610dab565b91506113c283610dab565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156113f7576113f6611349565b5b828201905092915050565b600061140d82610dab565b915061141883610dab565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561145157611450611349565b5b828202905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600061149682610dab565b91506114a183610dab565b9250826114b1576114b06112e9565b5b828204905092915050565b600082825260208201905092915050565b7f626c6f636b506572696f6420697320746f6f20736d616c6c2e00000000000000600082015250565b60006115036019836114bc565b915061150e826114cd565b602082019050919050565b60006020820190508181036000830152611532816114f6565b9050919050565b7f65706f6368506572696f6420697320746f6f20736d616c6c2e00000000000000600082015250565b600061156f6019836114bc565b915061157a82611539565b602082019050919050565b6000602082019050818103600083015261159e81611562565b9050919050565b7f7265776172645261746520697320746f6f206c617267652e0000000000000000600082015250565b60006115db6018836114bc565b91506115e6826115a5565b602082019050919050565b6000602082019050818103600083015261160a816115ce565b9050919050565b7f76616c696461746f725468726573686f6c6420697320746f6f20736d616c6c2e600082015250565b60006116476020836114bc565b915061165282611611565b602082019050919050565b600060208201905081810360008301526116768161163a565b9050919050565b7f6a61696c5468726573686f6c6420697320746f6f20736d616c6c2e0000000000600082015250565b60006116b3601b836114bc565b91506116be8261167d565b602082019050919050565b600060208201905081810360008301526116e2816116a6565b9050919050565b7f6a61696c506572696f6420697320746f6f20736d616c6c2e0000000000000000600082015250565b600061171f6018836114bc565b915061172a826116e9565b602082019050919050565b6000602082019050818103600083015261174e81611712565b905091905056fea2646970667358221220579b4e13d7bca19dd54eb3f45052d9622ced5b6b12c0fa682401bc63c4f9355064736f6c634300080c0033",
}

// EnvironmentABI is the input ABI used to generate the binding from.
// Deprecated: Use EnvironmentMetaData.ABI instead.
var EnvironmentABI = EnvironmentMetaData.ABI

// EnvironmentBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EnvironmentMetaData.Bin instead.
var EnvironmentBin = EnvironmentMetaData.Bin

// DeployEnvironment deploys a new Ethereum contract, binding an instance of Environment to it.
func DeployEnvironment(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Environment, error) {
	parsed, err := EnvironmentMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EnvironmentBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Environment{EnvironmentCaller: EnvironmentCaller{contract: contract}, EnvironmentTransactor: EnvironmentTransactor{contract: contract}, EnvironmentFilterer: EnvironmentFilterer{contract: contract}}, nil
}

// Environment is an auto generated Go binding around an Ethereum contract.
type Environment struct {
	EnvironmentCaller     // Read-only binding to the contract
	EnvironmentTransactor // Write-only binding to the contract
	EnvironmentFilterer   // Log filterer for contract events
}

// EnvironmentCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnvironmentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnvironmentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnvironmentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnvironmentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnvironmentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnvironmentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnvironmentSession struct {
	Contract     *Environment      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EnvironmentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnvironmentCallerSession struct {
	Contract *EnvironmentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// EnvironmentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnvironmentTransactorSession struct {
	Contract     *EnvironmentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// EnvironmentRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnvironmentRaw struct {
	Contract *Environment // Generic contract binding to access the raw methods on
}

// EnvironmentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnvironmentCallerRaw struct {
	Contract *EnvironmentCaller // Generic read-only contract binding to access the raw methods on
}

// EnvironmentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnvironmentTransactorRaw struct {
	Contract *EnvironmentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnvironment creates a new instance of Environment, bound to a specific deployed contract.
func NewEnvironment(address common.Address, backend bind.ContractBackend) (*Environment, error) {
	contract, err := bindEnvironment(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Environment{EnvironmentCaller: EnvironmentCaller{contract: contract}, EnvironmentTransactor: EnvironmentTransactor{contract: contract}, EnvironmentFilterer: EnvironmentFilterer{contract: contract}}, nil
}

// NewEnvironmentCaller creates a new read-only instance of Environment, bound to a specific deployed contract.
func NewEnvironmentCaller(address common.Address, caller bind.ContractCaller) (*EnvironmentCaller, error) {
	contract, err := bindEnvironment(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnvironmentCaller{contract: contract}, nil
}

// NewEnvironmentTransactor creates a new write-only instance of Environment, bound to a specific deployed contract.
func NewEnvironmentTransactor(address common.Address, transactor bind.ContractTransactor) (*EnvironmentTransactor, error) {
	contract, err := bindEnvironment(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnvironmentTransactor{contract: contract}, nil
}

// NewEnvironmentFilterer creates a new log filterer instance of Environment, bound to a specific deployed contract.
func NewEnvironmentFilterer(address common.Address, filterer bind.ContractFilterer) (*EnvironmentFilterer, error) {
	contract, err := bindEnvironment(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnvironmentFilterer{contract: contract}, nil
}

// bindEnvironment binds a generic wrapper to an already deployed contract.
func bindEnvironment(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EnvironmentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Environment *EnvironmentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Environment.Contract.EnvironmentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Environment *EnvironmentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Environment.Contract.EnvironmentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Environment *EnvironmentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Environment.Contract.EnvironmentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Environment *EnvironmentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Environment.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Environment *EnvironmentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Environment.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Environment *EnvironmentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Environment.Contract.contract.Transact(opts, method, params...)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Environment *EnvironmentCaller) Epoch(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Environment.contract.Call(opts, &out, "epoch")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Environment *EnvironmentSession) Epoch() (*big.Int, error) {
	return _Environment.Contract.Epoch(&_Environment.CallOpts)
}

// Epoch is a free data retrieval call binding the contract method 0x900cf0cf.
//
// Solidity: function epoch() view returns(uint256)
func (_Environment *EnvironmentCallerSession) Epoch() (*big.Int, error) {
	return _Environment.Contract.Epoch(&_Environment.CallOpts)
}

// EpochAndValues is a free data retrieval call binding the contract method 0x9cf370d7.
//
// Solidity: function epochAndValues() view returns(uint256[] epochs, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] _values)
func (_Environment *EnvironmentCaller) EpochAndValues(opts *bind.CallOpts) (struct {
	Epochs []*big.Int
	Values []IEnvironmentEnvironmentValue
}, error) {
	var out []interface{}
	err := _Environment.contract.Call(opts, &out, "epochAndValues")

	outstruct := new(struct {
		Epochs []*big.Int
		Values []IEnvironmentEnvironmentValue
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Epochs = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Values = *abi.ConvertType(out[1], new([]IEnvironmentEnvironmentValue)).(*[]IEnvironmentEnvironmentValue)

	return *outstruct, err

}

// EpochAndValues is a free data retrieval call binding the contract method 0x9cf370d7.
//
// Solidity: function epochAndValues() view returns(uint256[] epochs, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] _values)
func (_Environment *EnvironmentSession) EpochAndValues() (struct {
	Epochs []*big.Int
	Values []IEnvironmentEnvironmentValue
}, error) {
	return _Environment.Contract.EpochAndValues(&_Environment.CallOpts)
}

// EpochAndValues is a free data retrieval call binding the contract method 0x9cf370d7.
//
// Solidity: function epochAndValues() view returns(uint256[] epochs, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] _values)
func (_Environment *EnvironmentCallerSession) EpochAndValues() (struct {
	Epochs []*big.Int
	Values []IEnvironmentEnvironmentValue
}, error) {
	return _Environment.Contract.EpochAndValues(&_Environment.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Environment *EnvironmentCaller) Initialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Environment.contract.Call(opts, &out, "initialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Environment *EnvironmentSession) Initialized() (bool, error) {
	return _Environment.Contract.Initialized(&_Environment.CallOpts)
}

// Initialized is a free data retrieval call binding the contract method 0x158ef93e.
//
// Solidity: function initialized() view returns(bool)
func (_Environment *EnvironmentCallerSession) Initialized() (bool, error) {
	return _Environment.Contract.Initialized(&_Environment.CallOpts)
}

// IsFirstBlock is a free data retrieval call binding the contract method 0x4ba7ed39.
//
// Solidity: function isFirstBlock() view returns(bool)
func (_Environment *EnvironmentCaller) IsFirstBlock(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Environment.contract.Call(opts, &out, "isFirstBlock")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFirstBlock is a free data retrieval call binding the contract method 0x4ba7ed39.
//
// Solidity: function isFirstBlock() view returns(bool)
func (_Environment *EnvironmentSession) IsFirstBlock() (bool, error) {
	return _Environment.Contract.IsFirstBlock(&_Environment.CallOpts)
}

// IsFirstBlock is a free data retrieval call binding the contract method 0x4ba7ed39.
//
// Solidity: function isFirstBlock() view returns(bool)
func (_Environment *EnvironmentCallerSession) IsFirstBlock() (bool, error) {
	return _Environment.Contract.IsFirstBlock(&_Environment.CallOpts)
}

// IsLastBlock is a free data retrieval call binding the contract method 0xd4a53686.
//
// Solidity: function isLastBlock() view returns(bool)
func (_Environment *EnvironmentCaller) IsLastBlock(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Environment.contract.Call(opts, &out, "isLastBlock")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsLastBlock is a free data retrieval call binding the contract method 0xd4a53686.
//
// Solidity: function isLastBlock() view returns(bool)
func (_Environment *EnvironmentSession) IsLastBlock() (bool, error) {
	return _Environment.Contract.IsLastBlock(&_Environment.CallOpts)
}

// IsLastBlock is a free data retrieval call binding the contract method 0xd4a53686.
//
// Solidity: function isLastBlock() view returns(bool)
func (_Environment *EnvironmentCallerSession) IsLastBlock() (bool, error) {
	return _Environment.Contract.IsLastBlock(&_Environment.CallOpts)
}

// NextValue is a free data retrieval call binding the contract method 0xccfb9935.
//
// Solidity: function nextValue() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Environment *EnvironmentCaller) NextValue(opts *bind.CallOpts) (IEnvironmentEnvironmentValue, error) {
	var out []interface{}
	err := _Environment.contract.Call(opts, &out, "nextValue")

	if err != nil {
		return *new(IEnvironmentEnvironmentValue), err
	}

	out0 := *abi.ConvertType(out[0], new(IEnvironmentEnvironmentValue)).(*IEnvironmentEnvironmentValue)

	return out0, err

}

// NextValue is a free data retrieval call binding the contract method 0xccfb9935.
//
// Solidity: function nextValue() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Environment *EnvironmentSession) NextValue() (IEnvironmentEnvironmentValue, error) {
	return _Environment.Contract.NextValue(&_Environment.CallOpts)
}

// NextValue is a free data retrieval call binding the contract method 0xccfb9935.
//
// Solidity: function nextValue() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Environment *EnvironmentCallerSession) NextValue() (IEnvironmentEnvironmentValue, error) {
	return _Environment.Contract.NextValue(&_Environment.CallOpts)
}

// Updates is a free data retrieval call binding the contract method 0xb4c2f727.
//
// Solidity: function updates(uint256 ) view returns(uint256)
func (_Environment *EnvironmentCaller) Updates(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Environment.contract.Call(opts, &out, "updates", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Updates is a free data retrieval call binding the contract method 0xb4c2f727.
//
// Solidity: function updates(uint256 ) view returns(uint256)
func (_Environment *EnvironmentSession) Updates(arg0 *big.Int) (*big.Int, error) {
	return _Environment.Contract.Updates(&_Environment.CallOpts, arg0)
}

// Updates is a free data retrieval call binding the contract method 0xb4c2f727.
//
// Solidity: function updates(uint256 ) view returns(uint256)
func (_Environment *EnvironmentCallerSession) Updates(arg0 *big.Int) (*big.Int, error) {
	return _Environment.Contract.Updates(&_Environment.CallOpts, arg0)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Environment *EnvironmentCaller) Value(opts *bind.CallOpts) (IEnvironmentEnvironmentValue, error) {
	var out []interface{}
	err := _Environment.contract.Call(opts, &out, "value")

	if err != nil {
		return *new(IEnvironmentEnvironmentValue), err
	}

	out0 := *abi.ConvertType(out[0], new(IEnvironmentEnvironmentValue)).(*IEnvironmentEnvironmentValue)

	return out0, err

}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Environment *EnvironmentSession) Value() (IEnvironmentEnvironmentValue, error) {
	return _Environment.Contract.Value(&_Environment.CallOpts)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Environment *EnvironmentCallerSession) Value() (IEnvironmentEnvironmentValue, error) {
	return _Environment.Contract.Value(&_Environment.CallOpts)
}

// Values is a free data retrieval call binding the contract method 0x5e383d21.
//
// Solidity: function values(uint256 ) view returns(uint256 startBlock, uint256 startEpoch, uint256 blockPeriod, uint256 epochPeriod, uint256 rewardRate, uint256 validatorThreshold, uint256 jailThreshold, uint256 jailPeriod)
func (_Environment *EnvironmentCaller) Values(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StartBlock         *big.Int
	StartEpoch         *big.Int
	BlockPeriod        *big.Int
	EpochPeriod        *big.Int
	RewardRate         *big.Int
	ValidatorThreshold *big.Int
	JailThreshold      *big.Int
	JailPeriod         *big.Int
}, error) {
	var out []interface{}
	err := _Environment.contract.Call(opts, &out, "values", arg0)

	outstruct := new(struct {
		StartBlock         *big.Int
		StartEpoch         *big.Int
		BlockPeriod        *big.Int
		EpochPeriod        *big.Int
		RewardRate         *big.Int
		ValidatorThreshold *big.Int
		JailThreshold      *big.Int
		JailPeriod         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.StartBlock = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.StartEpoch = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BlockPeriod = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EpochPeriod = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.RewardRate = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.ValidatorThreshold = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.JailThreshold = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.JailPeriod = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Values is a free data retrieval call binding the contract method 0x5e383d21.
//
// Solidity: function values(uint256 ) view returns(uint256 startBlock, uint256 startEpoch, uint256 blockPeriod, uint256 epochPeriod, uint256 rewardRate, uint256 validatorThreshold, uint256 jailThreshold, uint256 jailPeriod)
func (_Environment *EnvironmentSession) Values(arg0 *big.Int) (struct {
	StartBlock         *big.Int
	StartEpoch         *big.Int
	BlockPeriod        *big.Int
	EpochPeriod        *big.Int
	RewardRate         *big.Int
	ValidatorThreshold *big.Int
	JailThreshold      *big.Int
	JailPeriod         *big.Int
}, error) {
	return _Environment.Contract.Values(&_Environment.CallOpts, arg0)
}

// Values is a free data retrieval call binding the contract method 0x5e383d21.
//
// Solidity: function values(uint256 ) view returns(uint256 startBlock, uint256 startEpoch, uint256 blockPeriod, uint256 epochPeriod, uint256 rewardRate, uint256 validatorThreshold, uint256 jailThreshold, uint256 jailPeriod)
func (_Environment *EnvironmentCallerSession) Values(arg0 *big.Int) (struct {
	StartBlock         *big.Int
	StartEpoch         *big.Int
	BlockPeriod        *big.Int
	EpochPeriod        *big.Int
	RewardRate         *big.Int
	ValidatorThreshold *big.Int
	JailThreshold      *big.Int
	JailPeriod         *big.Int
}, error) {
	return _Environment.Contract.Values(&_Environment.CallOpts, arg0)
}

// Initialize is a paid mutator transaction binding the contract method 0xf38e08e1.
//
// Solidity: function initialize((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) initialValue) returns()
func (_Environment *EnvironmentTransactor) Initialize(opts *bind.TransactOpts, initialValue IEnvironmentEnvironmentValue) (*types.Transaction, error) {
	return _Environment.contract.Transact(opts, "initialize", initialValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xf38e08e1.
//
// Solidity: function initialize((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) initialValue) returns()
func (_Environment *EnvironmentSession) Initialize(initialValue IEnvironmentEnvironmentValue) (*types.Transaction, error) {
	return _Environment.Contract.Initialize(&_Environment.TransactOpts, initialValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xf38e08e1.
//
// Solidity: function initialize((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) initialValue) returns()
func (_Environment *EnvironmentTransactorSession) Initialize(initialValue IEnvironmentEnvironmentValue) (*types.Transaction, error) {
	return _Environment.Contract.Initialize(&_Environment.TransactOpts, initialValue)
}

// UpdateValue is a paid mutator transaction binding the contract method 0x11bff261.
//
// Solidity: function updateValue((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) newValue) returns()
func (_Environment *EnvironmentTransactor) UpdateValue(opts *bind.TransactOpts, newValue IEnvironmentEnvironmentValue) (*types.Transaction, error) {
	return _Environment.contract.Transact(opts, "updateValue", newValue)
}

// UpdateValue is a paid mutator transaction binding the contract method 0x11bff261.
//
// Solidity: function updateValue((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) newValue) returns()
func (_Environment *EnvironmentSession) UpdateValue(newValue IEnvironmentEnvironmentValue) (*types.Transaction, error) {
	return _Environment.Contract.UpdateValue(&_Environment.TransactOpts, newValue)
}

// UpdateValue is a paid mutator transaction binding the contract method 0x11bff261.
//
// Solidity: function updateValue((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) newValue) returns()
func (_Environment *EnvironmentTransactorSession) UpdateValue(newValue IEnvironmentEnvironmentValue) (*types.Transaction, error) {
	return _Environment.Contract.UpdateValue(&_Environment.TransactOpts, newValue)
}
