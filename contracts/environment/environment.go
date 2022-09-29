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
	CommissionRate     *big.Int
	ValidatorThreshold *big.Int
	JailThreshold      *big.Int
	JailPeriod         *big.Int
}

// EnvironmentMetaData contains all meta data concerning the Environment contract.
var EnvironmentMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyBlockProducer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyNotLastBlock\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PastEpoch\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"detail\",\"type\":\"string\"}],\"name\":\"ValidationError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"epoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_epoch\",\"type\":\"uint256\"}],\"name\":\"findValue\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailPeriod\",\"type\":\"uint256\"}],\"internalType\":\"structIEnvironment.EnvironmentValue\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailPeriod\",\"type\":\"uint256\"}],\"internalType\":\"structIEnvironment.EnvironmentValue\",\"name\":\"initialValue\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isFirstBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isLastBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextValue\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailPeriod\",\"type\":\"uint256\"}],\"internalType\":\"structIEnvironment.EnvironmentValue\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailPeriod\",\"type\":\"uint256\"}],\"internalType\":\"structIEnvironment.EnvironmentValue\",\"name\":\"newValue\",\"type\":\"tuple\"}],\"name\":\"updateValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"updates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"value\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailPeriod\",\"type\":\"uint256\"}],\"internalType\":\"structIEnvironment.EnvironmentValue\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"values\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailPeriod\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611807806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063900cf0cf11610071578063900cf0cf1461015c578063b4c2f7271461017a578063bd4eff19146101aa578063ccfb9935146101c6578063d4a53686146101e4578063fcbb371b14610202576100a9565b806308a54356146100ae578063158ef93e146100ca5780633fa4f245146100e85780634ba7ed39146101065780635e383d2114610124575b600080fd5b6100c860048036038101906100c391906110cf565b610232565b005b6100d2610318565b6040516100df9190611118565b60405180910390f35b6100f0610329565b6040516100fd91906111f9565b60405180910390f35b61010e6103cf565b60405161011b9190611118565b60405180910390f35b61013e60048036038101906101399190611215565b6103ef565b60405161015399989796959493929190611251565b60405180910390f35b61016461044d565b60405161017191906112de565b60405180910390f35b610194600480360381019061018f9190611215565b610496565b6040516101a191906112de565b60405180910390f35b6101c460048036038101906101bf91906110cf565b6104ba565b005b6101ce610618565b6040516101db91906111f9565b60405180910390f35b6101ec61070a565b6040516101f99190611118565b60405180910390f35b61021c60048036038101906102179190611215565b610736565b60405161022991906111f9565b60405180910390f35b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610297576040517f1cf4735900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60008054906101000a900460ff16156102dc576040517f0dc149f000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60016000806101000a81548160ff021916908315150217905550600081600001818152505060018160200181815250506103158161085c565b50565b60008054906101000a900460ff1681565b610331610ecd565b600061033b610a3e565b90506103504382610a7890919063ffffffff16565b6103615761035c610a8a565b610363565b805b60405180610120016040529081600082015481526020016001820154815260200160028201548152602001600382015481526020016004820154815260200160058201548152602001600682015481526020016007820154815260200160088201548152505091505090565b6000806103da610329565b60600151436103e99190611328565b14905090565b600281815481106103ff57600080fd5b90600052602060002090600902016000915090508060000154908060010154908060020154908060030154908060040154908060050154908060060154908060070154908060080154905089565b600080610458610a3e565b905061046d4382610a7890919063ffffffff16565b6104865761048161047c610a8a565b610afd565b610490565b61048f81610afd565b5b91505090565b600181815481106104a657600080fd5b906000526020600020016000915090505481565b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461051f576040517f1cf4735900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61052761070a565b1561055e576040517f1e59ccd900000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b61056661044d565b8160200151116105a2576040517f0eae4c9700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60006105ac610a3e565b90506105c14382610a7890919063ffffffff16565b156105e7576105d98282610b3490919063ffffffff16565b82600001818152505061060b565b610601826105f3610a8a565b610b3490919063ffffffff16565b8260000181815250505b6106148261085c565b5050565b610620610ecd565b600061062a610329565b90506000610636610a3e565b9050600082606001516001846020015161064e61044d565b6106589190611388565b61066291906113bc565b61066c9190611412565b836000015161067b91906113bc565b90506106908183610a7890919063ffffffff16565b61069a5782610702565b816040518061012001604052908160008201548152602001600182015481526020016002820154815260200160038201548152602001600482015481526020016005820154815260200160068201548152602001600782015481526020016008820154815250505b935050505090565b600080610715610329565b6060015160014361072691906113bc565b6107309190611328565b14905090565b61073e610ecd565b6108556002805480602002602001604051908101604052809291908181526020016000905b828210156107f0578382906000526020600020906009020160405180610120016040529081600082015481526020016001820154815260200160028201548152602001600382015481526020016004820154815260200160058201548152602001600682015481526020016007820154815260200160088201548152505081526020019060010190610763565b5050505083600180548060200260200160405190810160405280929190818152602001828054801561084157602002820191906000526020600020905b81548152602001906001019080831161082d575b5050505050610b709092919063ffffffff16565b9050919050565b61086581610c12565b6000600180549050905060008114806108ba57506108b943600260018461088c9190611388565b8154811061089d5761089c61146c565b5b9060005260206000209060090201610a7890919063ffffffff16565b5b1561097a576001826020015190806001815401808255809150506001900390600052602060002001600090919091909150556002829080600181540180825580915050600190039060005260206000209060090201600090919091909150600082015181600001556020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e0820151816007015561010082015181600801555050610a3a565b81602001516001808361098d9190611388565b8154811061099e5761099d61146c565b5b90600052602060002001819055508160026001836109bc9190611388565b815481106109cd576109cc61146c565b5b9060005260206000209060090201600082015181600001556020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e0820151816007015561010082015181600801559050505b5050565b600060026001600280549050610a549190611388565b81548110610a6557610a6461146c565b5b9060005260206000209060090201905090565b60008260000154821015905092915050565b60008060028054905090506001811415610ac9576002600081548110610ab357610ab261146c565b5b9060005260206000209060090201915050610afa565b60028082610ad79190611388565b81548110610ae857610ae761146c565b5b90600052602060002090600902019150505b90565b60008160030154826000015443610b149190611388565b610b1e919061149b565b8260010154610b2d91906113bc565b9050919050565b6000826003015483600101548360200151610b4f9190611388565b610b599190611412565b8360000154610b6891906113bc565b905092915050565b610b78610ecd565b6000845190508285600183610b8d9190611388565b81518110610b9e57610b9d61146c565b5b602002602001015111610bda5783600182610bb99190611388565b81518110610bca57610bc961146c565b5b6020026020010151915050610c0b565b6000610be98685600085610e0e565b9050848181518110610bfe57610bfd61146c565b5b6020026020010151925050505b9392505050565b600181604001511015610c5a576040517fa1b63ceb000000000000000000000000000000000000000000000000000000008152600401610c5190611529565b60405180910390fd5b600381606001511015610ca2576040517fa1b63ceb000000000000000000000000000000000000000000000000000000008152600401610c9990611595565b60405180910390fd5b606481608001511115610cea576040517fa1b63ceb000000000000000000000000000000000000000000000000000000008152600401610ce190611601565b60405180910390fd5b60648160a001511115610d32576040517fa1b63ceb000000000000000000000000000000000000000000000000000000008152600401610d299061166d565b60405180910390fd5b60018160c001511015610d7a576040517fa1b63ceb000000000000000000000000000000000000000000000000000000008152600401610d71906116d9565b60405180910390fd5b60018160e001511015610dc2576040517fa1b63ceb000000000000000000000000000000000000000000000000000000008152600401610db990611745565b60405180910390fd5b60018161010001511015610e0b576040517fa1b63ceb000000000000000000000000000000000000000000000000000000008152600401610e02906117b1565b60405180910390fd5b50565b600081831415610e2c57600182610e259190611388565b9050610ec5565b600060028385610e3c91906113bc565b610e46919061149b565b905084868281518110610e5c57610e5b61146c565b5b60200260200101511115610e7e57610e7686868684610e0e565b915050610ec5565b84868281518110610e9257610e9161146c565b5b60200260200101511015610ec057610eb88686600184610eb291906113bc565b86610e0e565b915050610ec5565b809150505b949350505050565b6040518061012001604052806000815260200160008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b6000604051905090565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610f7682610f2d565b810181811067ffffffffffffffff82111715610f9557610f94610f3e565b5b80604052505050565b6000610fa8610f19565b9050610fb48282610f6d565b919050565b6000819050919050565b610fcc81610fb9565b8114610fd757600080fd5b50565b600081359050610fe981610fc3565b92915050565b6000610120828403121561100657611005610f28565b5b611011610120610f9e565b9050600061102184828501610fda565b600083015250602061103584828501610fda565b602083015250604061104984828501610fda565b604083015250606061105d84828501610fda565b606083015250608061107184828501610fda565b60808301525060a061108584828501610fda565b60a08301525060c061109984828501610fda565b60c08301525060e06110ad84828501610fda565b60e0830152506101006110c284828501610fda565b6101008301525092915050565b600061012082840312156110e6576110e5610f23565b5b60006110f484828501610fef565b91505092915050565b60008115159050919050565b611112816110fd565b82525050565b600060208201905061112d6000830184611109565b92915050565b61113c81610fb9565b82525050565b610120820160008201516111596000850182611133565b50602082015161116c6020850182611133565b50604082015161117f6040850182611133565b5060608201516111926060850182611133565b5060808201516111a56080850182611133565b5060a08201516111b860a0850182611133565b5060c08201516111cb60c0850182611133565b5060e08201516111de60e0850182611133565b506101008201516111f3610100850182611133565b50505050565b60006101208201905061120f6000830184611142565b92915050565b60006020828403121561122b5761122a610f23565b5b600061123984828501610fda565b91505092915050565b61124b81610fb9565b82525050565b600061012082019050611267600083018c611242565b611274602083018b611242565b611281604083018a611242565b61128e6060830189611242565b61129b6080830188611242565b6112a860a0830187611242565b6112b560c0830186611242565b6112c260e0830185611242565b6112d0610100830184611242565b9a9950505050505050505050565b60006020820190506112f36000830184611242565b92915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061133382610fb9565b915061133e83610fb9565b92508261134e5761134d6112f9565b5b828206905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061139382610fb9565b915061139e83610fb9565b9250828210156113b1576113b0611359565b5b828203905092915050565b60006113c782610fb9565b91506113d283610fb9565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0382111561140757611406611359565b5b828201905092915050565b600061141d82610fb9565b915061142883610fb9565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561146157611460611359565b5b828202905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60006114a682610fb9565b91506114b183610fb9565b9250826114c1576114c06112f9565b5b828204905092915050565b600082825260208201905092915050565b7f626c6f636b506572696f6420697320746f6f20736d616c6c2e00000000000000600082015250565b60006115136019836114cc565b915061151e826114dd565b602082019050919050565b6000602082019050818103600083015261154281611506565b9050919050565b7f65706f6368506572696f6420697320746f6f20736d616c6c2e00000000000000600082015250565b600061157f6019836114cc565b915061158a82611549565b602082019050919050565b600060208201905081810360008301526115ae81611572565b9050919050565b7f7265776172645261746520697320746f6f206c617267652e0000000000000000600082015250565b60006115eb6018836114cc565b91506115f6826115b5565b602082019050919050565b6000602082019050818103600083015261161a816115de565b9050919050565b7f636f6d6d697373696f6e5261746520697320746f6f206c617267652e00000000600082015250565b6000611657601c836114cc565b915061166282611621565b602082019050919050565b600060208201905081810360008301526116868161164a565b9050919050565b7f76616c696461746f725468726573686f6c6420697320746f6f20736d616c6c2e600082015250565b60006116c36020836114cc565b91506116ce8261168d565b602082019050919050565b600060208201905081810360008301526116f2816116b6565b9050919050565b7f6a61696c5468726573686f6c6420697320746f6f20736d616c6c2e0000000000600082015250565b600061172f601b836114cc565b915061173a826116f9565b602082019050919050565b6000602082019050818103600083015261175e81611722565b9050919050565b7f6a61696c506572696f6420697320746f6f20736d616c6c2e0000000000000000600082015250565b600061179b6018836114cc565b91506117a682611765565b602082019050919050565b600060208201905081810360008301526117ca8161178e565b905091905056fea2646970667358221220ec81c7ab754b20653c3c9bd8c96df1dd00c5a8821ecd22f7e7710f4699cd35de64736f6c634300080c0033",
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

// FindValue is a free data retrieval call binding the contract method 0xfcbb371b.
//
// Solidity: function findValue(uint256 _epoch) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Environment *EnvironmentCaller) FindValue(opts *bind.CallOpts, _epoch *big.Int) (IEnvironmentEnvironmentValue, error) {
	var out []interface{}
	err := _Environment.contract.Call(opts, &out, "findValue", _epoch)

	if err != nil {
		return *new(IEnvironmentEnvironmentValue), err
	}

	out0 := *abi.ConvertType(out[0], new(IEnvironmentEnvironmentValue)).(*IEnvironmentEnvironmentValue)

	return out0, err

}

// FindValue is a free data retrieval call binding the contract method 0xfcbb371b.
//
// Solidity: function findValue(uint256 _epoch) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Environment *EnvironmentSession) FindValue(_epoch *big.Int) (IEnvironmentEnvironmentValue, error) {
	return _Environment.Contract.FindValue(&_Environment.CallOpts, _epoch)
}

// FindValue is a free data retrieval call binding the contract method 0xfcbb371b.
//
// Solidity: function findValue(uint256 _epoch) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Environment *EnvironmentCallerSession) FindValue(_epoch *big.Int) (IEnvironmentEnvironmentValue, error) {
	return _Environment.Contract.FindValue(&_Environment.CallOpts, _epoch)
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
// Solidity: function nextValue() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
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
// Solidity: function nextValue() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Environment *EnvironmentSession) NextValue() (IEnvironmentEnvironmentValue, error) {
	return _Environment.Contract.NextValue(&_Environment.CallOpts)
}

// NextValue is a free data retrieval call binding the contract method 0xccfb9935.
//
// Solidity: function nextValue() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
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
// Solidity: function value() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
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
// Solidity: function value() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Environment *EnvironmentSession) Value() (IEnvironmentEnvironmentValue, error) {
	return _Environment.Contract.Value(&_Environment.CallOpts)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Environment *EnvironmentCallerSession) Value() (IEnvironmentEnvironmentValue, error) {
	return _Environment.Contract.Value(&_Environment.CallOpts)
}

// Values is a free data retrieval call binding the contract method 0x5e383d21.
//
// Solidity: function values(uint256 ) view returns(uint256 startBlock, uint256 startEpoch, uint256 blockPeriod, uint256 epochPeriod, uint256 rewardRate, uint256 commissionRate, uint256 validatorThreshold, uint256 jailThreshold, uint256 jailPeriod)
func (_Environment *EnvironmentCaller) Values(opts *bind.CallOpts, arg0 *big.Int) (struct {
	StartBlock         *big.Int
	StartEpoch         *big.Int
	BlockPeriod        *big.Int
	EpochPeriod        *big.Int
	RewardRate         *big.Int
	CommissionRate     *big.Int
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
		CommissionRate     *big.Int
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
	outstruct.CommissionRate = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.ValidatorThreshold = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.JailThreshold = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.JailPeriod = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Values is a free data retrieval call binding the contract method 0x5e383d21.
//
// Solidity: function values(uint256 ) view returns(uint256 startBlock, uint256 startEpoch, uint256 blockPeriod, uint256 epochPeriod, uint256 rewardRate, uint256 commissionRate, uint256 validatorThreshold, uint256 jailThreshold, uint256 jailPeriod)
func (_Environment *EnvironmentSession) Values(arg0 *big.Int) (struct {
	StartBlock         *big.Int
	StartEpoch         *big.Int
	BlockPeriod        *big.Int
	EpochPeriod        *big.Int
	RewardRate         *big.Int
	CommissionRate     *big.Int
	ValidatorThreshold *big.Int
	JailThreshold      *big.Int
	JailPeriod         *big.Int
}, error) {
	return _Environment.Contract.Values(&_Environment.CallOpts, arg0)
}

// Values is a free data retrieval call binding the contract method 0x5e383d21.
//
// Solidity: function values(uint256 ) view returns(uint256 startBlock, uint256 startEpoch, uint256 blockPeriod, uint256 epochPeriod, uint256 rewardRate, uint256 commissionRate, uint256 validatorThreshold, uint256 jailThreshold, uint256 jailPeriod)
func (_Environment *EnvironmentCallerSession) Values(arg0 *big.Int) (struct {
	StartBlock         *big.Int
	StartEpoch         *big.Int
	BlockPeriod        *big.Int
	EpochPeriod        *big.Int
	RewardRate         *big.Int
	CommissionRate     *big.Int
	ValidatorThreshold *big.Int
	JailThreshold      *big.Int
	JailPeriod         *big.Int
}, error) {
	return _Environment.Contract.Values(&_Environment.CallOpts, arg0)
}

// Initialize is a paid mutator transaction binding the contract method 0x08a54356.
//
// Solidity: function initialize((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) initialValue) returns()
func (_Environment *EnvironmentTransactor) Initialize(opts *bind.TransactOpts, initialValue IEnvironmentEnvironmentValue) (*types.Transaction, error) {
	return _Environment.contract.Transact(opts, "initialize", initialValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x08a54356.
//
// Solidity: function initialize((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) initialValue) returns()
func (_Environment *EnvironmentSession) Initialize(initialValue IEnvironmentEnvironmentValue) (*types.Transaction, error) {
	return _Environment.Contract.Initialize(&_Environment.TransactOpts, initialValue)
}

// Initialize is a paid mutator transaction binding the contract method 0x08a54356.
//
// Solidity: function initialize((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) initialValue) returns()
func (_Environment *EnvironmentTransactorSession) Initialize(initialValue IEnvironmentEnvironmentValue) (*types.Transaction, error) {
	return _Environment.Contract.Initialize(&_Environment.TransactOpts, initialValue)
}

// UpdateValue is a paid mutator transaction binding the contract method 0xbd4eff19.
//
// Solidity: function updateValue((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) newValue) returns()
func (_Environment *EnvironmentTransactor) UpdateValue(opts *bind.TransactOpts, newValue IEnvironmentEnvironmentValue) (*types.Transaction, error) {
	return _Environment.contract.Transact(opts, "updateValue", newValue)
}

// UpdateValue is a paid mutator transaction binding the contract method 0xbd4eff19.
//
// Solidity: function updateValue((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) newValue) returns()
func (_Environment *EnvironmentSession) UpdateValue(newValue IEnvironmentEnvironmentValue) (*types.Transaction, error) {
	return _Environment.Contract.UpdateValue(&_Environment.TransactOpts, newValue)
}

// UpdateValue is a paid mutator transaction binding the contract method 0xbd4eff19.
//
// Solidity: function updateValue((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) newValue) returns()
func (_Environment *EnvironmentTransactorSession) UpdateValue(newValue IEnvironmentEnvironmentValue) (*types.Transaction, error) {
	return _Environment.Contract.UpdateValue(&_Environment.TransactOpts, newValue)
}
