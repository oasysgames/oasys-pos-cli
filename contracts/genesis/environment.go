// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package genesis

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

// EnvironmentEnvironmentValue is an auto generated low-level Go binding around an user-defined struct.
type EnvironmentEnvironmentValue struct {
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
	ABI: "[{\"inputs\":[],\"name\":\"epoch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"epochAndValues\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"epochs\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailPeriod\",\"type\":\"uint256\"}],\"internalType\":\"structEnvironment.EnvironmentValue[]\",\"name\":\"_values\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailPeriod\",\"type\":\"uint256\"}],\"internalType\":\"structEnvironment.EnvironmentValue\",\"name\":\"initialValue\",\"type\":\"tuple\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isFirstBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isLastBlock\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextValue\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailPeriod\",\"type\":\"uint256\"}],\"internalType\":\"structEnvironment.EnvironmentValue\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailPeriod\",\"type\":\"uint256\"}],\"internalType\":\"structEnvironment.EnvironmentValue\",\"name\":\"newValue\",\"type\":\"tuple\"}],\"name\":\"updateValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"updates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"value\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailPeriod\",\"type\":\"uint256\"}],\"internalType\":\"structEnvironment.EnvironmentValue\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"values\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"startBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"epochPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"validatorThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailPeriod\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611988806100206000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063900cf0cf11610071578063900cf0cf1461015b5780639cf370d714610179578063b4c2f72714610198578063ccfb9935146101c8578063d4a53686146101e6578063f38e08e114610204576100a9565b806311bff261146100ae578063158ef93e146100ca5780633fa4f245146100e85780634ba7ed39146101065780635e383d2114610124575b600080fd5b6100c860048036038101906100c39190610ed2565b610220565b005b6100d2610399565b6040516100df9190610f1b565b60405180910390f35b6100f06103aa565b6040516100fd9190610fe7565b60405180910390f35b61010e610446565b60405161011b9190610f1b565b60405180910390f35b61013e60048036038101906101399190611003565b610466565b60405161015298979695949392919061103f565b60405180910390f35b6101636104be565b60405161017091906110bd565b60405180910390f35b610181610507565b60405161018f9291906112d9565b60405180910390f35b6101b260048036038101906101ad9190611003565b610612565b6040516101bf91906110bd565b60405180910390f35b6101d0610636565b6040516101dd9190610fe7565b60405180910390f35b6101ee6106de565b6040516101fb9190610f1b565b60405180910390f35b61021e60048036038101906102199190610ed2565b61070a565b005b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461028e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102859061136d565b60405180910390fd5b6102966106de565b156102d6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102cd906113ff565b60405180910390fd5b6102de6104be565b816020015111610323576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161031a9061146b565b60405180910390fd5b600061032d610802565b9050610342438261087690919063ffffffff16565b156103685761035a828261088890919063ffffffff16565b82600001818152505061038c565b610382826103746108c4565b61088890919063ffffffff16565b8260000181815250505b61039582610937565b5050565b60008054906101000a900460ff1681565b6103b2610ced565b60006103bc610802565b90506103d1438261087690919063ffffffff16565b6103e2576103dd6108c4565b6103e4565b805b604051806101000160405290816000820154815260200160018201548152602001600282015481526020016003820154815260200160048201548152602001600582015481526020016006820154815260200160078201548152505091505090565b6000806104516103aa565b606001514361046091906114ba565b14905090565b6002818154811061047657600080fd5b90600052602060002090600802016000915090508060000154908060010154908060020154908060030154908060040154908060050154908060060154908060070154905088565b6000806104c9610802565b90506104de438261087690919063ffffffff16565b6104f7576104f26104ed6108c4565b610b03565b610501565b61050081610b03565b5b91505090565b606080600160028180548060200260200160405190810160405280929190818152602001828054801561055957602002820191906000526020600020905b815481526020019060010190808311610545575b5050505050915080805480602002602001604051908101604052809291908181526020016000905b828210156106045783829060005260206000209060080201604051806101000160405290816000820154815260200160018201548152602001600282015481526020016003820154815260200160048201548152602001600582015481526020016006820154815260200160078201548152505081526020019060010190610581565b505050509050915091509091565b6001818154811061062257600080fd5b906000526020600020016000915090505481565b61063e610ced565b6000610648610802565b905061066960014361065a919061151a565b8261087690919063ffffffff16565b61067a576106756108c4565b61067c565b805b604051806101000160405290816000820154815260200160018201548152602001600282015481526020016003820154815260200160048201548152602001600582015481526020016006820154815260200160078201548152505091505090565b6000806106e96103aa565b606001516001436106fa919061151a565b61070491906114ba565b14905090565b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610778576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161076f9061136d565b60405180910390fd5b60008054906101000a900460ff16156107c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107bd906115bc565b60405180910390fd5b60016000806101000a81548160ff021916908315150217905550600081600001818152505060018160200181815250506107ff81610937565b50565b6000806002805490509050600181141561084157600260008154811061082b5761082a6115dc565b5b9060005260206000209060080201915050610873565b6002600182610850919061160b565b81548110610861576108606115dc565b5b90600052602060002090600802019150505b90565b60008260000154821015905092915050565b60008260030154836001015483602001516108a3919061160b565b6108ad919061163f565b83600001546108bc919061151a565b905092915050565b600080600280549050905060018114156109035760026000815481106108ed576108ec6115dc565b5b9060005260206000209060080201915050610934565b60028082610911919061160b565b81548110610922576109216115dc565b5b90600052602060002090600802019150505b90565b61094081610b3a565b6000600180549050905060008114806109955750610994436002600184610967919061160b565b81548110610978576109776115dc565b5b906000526020600020906008020161087690919063ffffffff16565b5b15610a4a576001826020015190806001815401808255809150506001900390600052602060002001600090919091909150556002829080600181540180825580915050600190039060005260206000209060080201600090919091909150600082015181600001556020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e082015181600701555050610aff565b816020015160018083610a5d919061160b565b81548110610a6e57610a6d6115dc565b5b9060005260206000200181905550816002600183610a8c919061160b565b81548110610a9d57610a9c6115dc565b5b9060005260206000209060080201600082015181600001556020820151816001015560408201518160020155606082015181600301556080820151816004015560a0820151816005015560c0820151816006015560e082015181600701559050505b5050565b60008160030154826000015443610b1a919061160b565b610b249190611699565b8260010154610b33919061151a565b9050919050565b600181604001511015610b82576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b7990611716565b60405180910390fd5b600381606001511015610bca576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bc190611782565b60405180910390fd5b606481608001511115610c12576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c09906117ee565b60405180910390fd5b60018160a001511015610c5a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c519061185a565b60405180910390fd5b60018160c001511015610ca2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c99906118c6565b60405180910390fd5b60018160e001511015610cea576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ce190611932565b60405180910390fd5b50565b60405180610100016040528060008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b6000604051905090565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610d8f82610d46565b810181811067ffffffffffffffff82111715610dae57610dad610d57565b5b80604052505050565b6000610dc1610d32565b9050610dcd8282610d86565b919050565b6000819050919050565b610de581610dd2565b8114610df057600080fd5b50565b600081359050610e0281610ddc565b92915050565b60006101008284031215610e1f57610e1e610d41565b5b610e2a610100610db7565b90506000610e3a84828501610df3565b6000830152506020610e4e84828501610df3565b6020830152506040610e6284828501610df3565b6040830152506060610e7684828501610df3565b6060830152506080610e8a84828501610df3565b60808301525060a0610e9e84828501610df3565b60a08301525060c0610eb284828501610df3565b60c08301525060e0610ec684828501610df3565b60e08301525092915050565b60006101008284031215610ee957610ee8610d3c565b5b6000610ef784828501610e08565b91505092915050565b60008115159050919050565b610f1581610f00565b82525050565b6000602082019050610f306000830184610f0c565b92915050565b610f3f81610dd2565b82525050565b61010082016000820151610f5c6000850182610f36565b506020820151610f6f6020850182610f36565b506040820151610f826040850182610f36565b506060820151610f956060850182610f36565b506080820151610fa86080850182610f36565b5060a0820151610fbb60a0850182610f36565b5060c0820151610fce60c0850182610f36565b5060e0820151610fe160e0850182610f36565b50505050565b600061010082019050610ffd6000830184610f45565b92915050565b60006020828403121561101957611018610d3c565b5b600061102784828501610df3565b91505092915050565b61103981610dd2565b82525050565b600061010082019050611055600083018b611030565b611062602083018a611030565b61106f6040830189611030565b61107c6060830188611030565b6110896080830187611030565b61109660a0830186611030565b6110a360c0830185611030565b6110b060e0830184611030565b9998505050505050505050565b60006020820190506110d26000830184611030565b92915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b60006111108383610f36565b60208301905092915050565b6000602082019050919050565b6000611134826110d8565b61113e81856110e3565b9350611149836110f4565b8060005b8381101561117a5781516111618882611104565b975061116c8361111c565b92505060018101905061114d565b5085935050505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b610100820160008201516111ca6000850182610f36565b5060208201516111dd6020850182610f36565b5060408201516111f06040850182610f36565b5060608201516112036060850182610f36565b5060808201516112166080850182610f36565b5060a082015161122960a0850182610f36565b5060c082015161123c60c0850182610f36565b5060e082015161124f60e0850182610f36565b50505050565b600061126183836111b3565b6101008301905092915050565b6000602082019050919050565b600061128682611187565b6112908185611192565b935061129b836111a3565b8060005b838110156112cc5781516112b38882611255565b97506112be8361126e565b92505060018101905061129f565b5085935050505092915050565b600060408201905081810360008301526112f38185611129565b90508181036020830152611307818461127b565b90509392505050565b600082825260208201905092915050565b7f73656e646572206d75737420626520626c6f636b2070726f64756365722e0000600082015250565b6000611357601e83611310565b915061136282611321565b602082019050919050565b600060208201905081810360008301526113868161134a565b9050919050565b7f6e6f742065786563757461626c6520696e20746865206c61737420626c6f636b60008201527f206f662065706f63682e00000000000000000000000000000000000000000000602082015250565b60006113e9602a83611310565b91506113f48261138d565b604082019050919050565b60006020820190508181036000830152611418816113dc565b9050919050565b7f737461727445706f6368206d757374206265206675747572652e000000000000600082015250565b6000611455601a83611310565b91506114608261141f565b602082019050919050565b6000602082019050818103600083015261148481611448565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60006114c582610dd2565b91506114d083610dd2565b9250826114e0576114df61148b565b5b828206905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600061152582610dd2565b915061153083610dd2565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611565576115646114eb565b5b828201905092915050565b7f616c726561647920696e697469616c697a65642e000000000000000000000000600082015250565b60006115a6601483611310565b91506115b182611570565b602082019050919050565b600060208201905081810360008301526115d581611599565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b600061161682610dd2565b915061162183610dd2565b925082821015611634576116336114eb565b5b828203905092915050565b600061164a82610dd2565b915061165583610dd2565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561168e5761168d6114eb565b5b828202905092915050565b60006116a482610dd2565b91506116af83610dd2565b9250826116bf576116be61148b565b5b828204905092915050565b7f626c6f636b506572696f6420697320746f6f20736d616c6c2e00000000000000600082015250565b6000611700601983611310565b915061170b826116ca565b602082019050919050565b6000602082019050818103600083015261172f816116f3565b9050919050565b7f65706f6368506572696f6420697320746f6f20736d616c6c2e00000000000000600082015250565b600061176c601983611310565b915061177782611736565b602082019050919050565b6000602082019050818103600083015261179b8161175f565b9050919050565b7f7265776172645261746520697320746f6f206c617267652e0000000000000000600082015250565b60006117d8601883611310565b91506117e3826117a2565b602082019050919050565b60006020820190508181036000830152611807816117cb565b9050919050565b7f76616c696461746f725468726573686f6c6420697320746f6f20736d616c6c2e600082015250565b6000611844602083611310565b915061184f8261180e565b602082019050919050565b6000602082019050818103600083015261187381611837565b9050919050565b7f6a61696c5468726573686f6c6420697320746f6f20736d616c6c2e0000000000600082015250565b60006118b0601b83611310565b91506118bb8261187a565b602082019050919050565b600060208201905081810360008301526118df816118a3565b9050919050565b7f6a61696c506572696f6420697320746f6f20736d616c6c2e0000000000000000600082015250565b600061191c601883611310565b9150611927826118e6565b602082019050919050565b6000602082019050818103600083015261194b8161190f565b905091905056fea26469706673582212200d7deabcaf0fc37872d77e575ea67bf006271225731019be9957a0a4b3004f6d64736f6c63430008090033",
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
	Values []EnvironmentEnvironmentValue
}, error) {
	var out []interface{}
	err := _Environment.contract.Call(opts, &out, "epochAndValues")

	outstruct := new(struct {
		Epochs []*big.Int
		Values []EnvironmentEnvironmentValue
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Epochs = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.Values = *abi.ConvertType(out[1], new([]EnvironmentEnvironmentValue)).(*[]EnvironmentEnvironmentValue)

	return *outstruct, err

}

// EpochAndValues is a free data retrieval call binding the contract method 0x9cf370d7.
//
// Solidity: function epochAndValues() view returns(uint256[] epochs, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] _values)
func (_Environment *EnvironmentSession) EpochAndValues() (struct {
	Epochs []*big.Int
	Values []EnvironmentEnvironmentValue
}, error) {
	return _Environment.Contract.EpochAndValues(&_Environment.CallOpts)
}

// EpochAndValues is a free data retrieval call binding the contract method 0x9cf370d7.
//
// Solidity: function epochAndValues() view returns(uint256[] epochs, (uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256)[] _values)
func (_Environment *EnvironmentCallerSession) EpochAndValues() (struct {
	Epochs []*big.Int
	Values []EnvironmentEnvironmentValue
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
func (_Environment *EnvironmentCaller) NextValue(opts *bind.CallOpts) (EnvironmentEnvironmentValue, error) {
	var out []interface{}
	err := _Environment.contract.Call(opts, &out, "nextValue")

	if err != nil {
		return *new(EnvironmentEnvironmentValue), err
	}

	out0 := *abi.ConvertType(out[0], new(EnvironmentEnvironmentValue)).(*EnvironmentEnvironmentValue)

	return out0, err

}

// NextValue is a free data retrieval call binding the contract method 0xccfb9935.
//
// Solidity: function nextValue() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Environment *EnvironmentSession) NextValue() (EnvironmentEnvironmentValue, error) {
	return _Environment.Contract.NextValue(&_Environment.CallOpts)
}

// NextValue is a free data retrieval call binding the contract method 0xccfb9935.
//
// Solidity: function nextValue() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Environment *EnvironmentCallerSession) NextValue() (EnvironmentEnvironmentValue, error) {
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
func (_Environment *EnvironmentCaller) Value(opts *bind.CallOpts) (EnvironmentEnvironmentValue, error) {
	var out []interface{}
	err := _Environment.contract.Call(opts, &out, "value")

	if err != nil {
		return *new(EnvironmentEnvironmentValue), err
	}

	out0 := *abi.ConvertType(out[0], new(EnvironmentEnvironmentValue)).(*EnvironmentEnvironmentValue)

	return out0, err

}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Environment *EnvironmentSession) Value() (EnvironmentEnvironmentValue, error) {
	return _Environment.Contract.Value(&_Environment.CallOpts)
}

// Value is a free data retrieval call binding the contract method 0x3fa4f245.
//
// Solidity: function value() view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))
func (_Environment *EnvironmentCallerSession) Value() (EnvironmentEnvironmentValue, error) {
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
func (_Environment *EnvironmentTransactor) Initialize(opts *bind.TransactOpts, initialValue EnvironmentEnvironmentValue) (*types.Transaction, error) {
	return _Environment.contract.Transact(opts, "initialize", initialValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xf38e08e1.
//
// Solidity: function initialize((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) initialValue) returns()
func (_Environment *EnvironmentSession) Initialize(initialValue EnvironmentEnvironmentValue) (*types.Transaction, error) {
	return _Environment.Contract.Initialize(&_Environment.TransactOpts, initialValue)
}

// Initialize is a paid mutator transaction binding the contract method 0xf38e08e1.
//
// Solidity: function initialize((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) initialValue) returns()
func (_Environment *EnvironmentTransactorSession) Initialize(initialValue EnvironmentEnvironmentValue) (*types.Transaction, error) {
	return _Environment.Contract.Initialize(&_Environment.TransactOpts, initialValue)
}

// UpdateValue is a paid mutator transaction binding the contract method 0x11bff261.
//
// Solidity: function updateValue((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) newValue) returns()
func (_Environment *EnvironmentTransactor) UpdateValue(opts *bind.TransactOpts, newValue EnvironmentEnvironmentValue) (*types.Transaction, error) {
	return _Environment.contract.Transact(opts, "updateValue", newValue)
}

// UpdateValue is a paid mutator transaction binding the contract method 0x11bff261.
//
// Solidity: function updateValue((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) newValue) returns()
func (_Environment *EnvironmentSession) UpdateValue(newValue EnvironmentEnvironmentValue) (*types.Transaction, error) {
	return _Environment.Contract.UpdateValue(&_Environment.TransactOpts, newValue)
}

// UpdateValue is a paid mutator transaction binding the contract method 0x11bff261.
//
// Solidity: function updateValue((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) newValue) returns()
func (_Environment *EnvironmentTransactorSession) UpdateValue(newValue EnvironmentEnvironmentValue) (*types.Transaction, error) {
	return _Environment.Contract.UpdateValue(&_Environment.TransactOpts, newValue)
}
