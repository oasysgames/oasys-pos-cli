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

// StakemanagerMetaData contains all meta data concerning the Stakemanager contract.
var StakemanagerMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Staked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Unstaked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorActivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"ValidatorCommissionRateUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorDeactivated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"ValidatorJailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"ValidatorSlashed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"activateValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allowlist\",\"outputs\":[{\"internalType\":\"contractIAllowlist\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epochs\",\"type\":\"uint256\"}],\"name\":\"claimCommissions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epochs\",\"type\":\"uint256\"}],\"name\":\"claimRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimUnstakes\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"currentValidators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"deactivateValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"environment\",\"outputs\":[{\"internalType\":\"contractEnvironment\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getBlockAndSlashes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"blocks\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"slashes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epochs\",\"type\":\"uint256\"}],\"name\":\"getCommissions\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"commissions\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"owners\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"operators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"stakes\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epochs\",\"type\":\"uint256\"}],\"name\":\"getRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"}],\"name\":\"getStakerInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"stakes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unstakes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"staker\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"}],\"name\":\"getStakerStakes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_validators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"stakes\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"stakeRequests\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"unstakeRequests\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"page\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"perPage\",\"type\":\"uint256\"}],\"name\":\"getStakers\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"epochs\",\"type\":\"uint256\"}],\"name\":\"getTotalRewards\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rewards\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"getValidatorInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"stakes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"commissionRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"jailEpoch\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"epoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"page\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"perPage\",\"type\":\"uint256\"}],\"name\":\"getValidatorStakes\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"_stakers\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"stakes\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidators\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractEnvironment\",\"name\":\"_environment\",\"type\":\"address\"},{\"internalType\":\"contractIAllowlist\",\"name\":\"_allowlist\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"initialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"joinValidator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"operatorToOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"slash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"}],\"name\":\"stake\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"stakerSigners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"stakers\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"signer\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"validator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"unstake\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newRate\",\"type\":\"uint256\"}],\"name\":\"updateCommissionRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"updateOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"operators\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"counts\",\"type\":\"uint256[]\"}],\"name\":\"updateValidatorBlocks\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"updateValidators\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validatorOwners\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"validatorUpdates\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"validators\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"jailEpoch\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastClaimCommission\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b5061874f80620000226000396000f3fe6080604052600436106102035760003560e01c80638a11d7c911610118578063ba50b879116100a0578063cbc0fac61161006f578063cbc0fac6146107ab578063dbd61d87146107d4578063f4fa1dde14610811578063f65a5ed21461084e578063fa52c7d81461088b57610203565b8063ba50b87914610703578063c2a672e01461072c578063c5f9dff014610755578063c96be4cb1461078257610203565b80639a99b4f0116100e75780639a99b4f014610620578063ac7475ed14610649578063ad71bd3614610672578063b46e5520146106af578063b7ab4db5146106d857610203565b80638a11d7c9146105255780638bca036b14610566578063900eb5a8146105a65780639168ae72146105e357610203565b806346dfce7b1161019b5780636b2b33691161016a5780636b2b33691461042d578063733bdef01461045657806374e2b63c146104945780637b520aa8146104bf57806388d0eb7d146104fc57610203565b806346dfce7b1461034c578063485cc9551461038a578063552164ee146103b35780635efc766e146103f057610203565b806326476204116101d757806326476204146102d75780632b47da52146102f35780632d497ba21461031e5780632ffb688d1461033557610203565b8062fa3d5014610208578063158ef93e14610231578063195afea11461025c5780632222636714610299575b600080fd5b34801561021457600080fd5b5061022f600480360381019061022a9190616e72565b6108cc565b005b34801561023d57600080fd5b50610246610a66565b6040516102539190616eba565b60405180910390f35b34801561026857600080fd5b50610283600480360381019061027e9190616f33565b610a77565b6040516102909190616f82565b60405180910390f35b3480156102a557600080fd5b506102c060048036038101906102bb9190616f33565b610af9565b6040516102ce929190616f9d565b60405180910390f35b6102f160048036038101906102ec9190616fc6565b610c0a565b005b3480156102ff57600080fd5b50610308611024565b6040516103159190617052565b60405180910390f35b34801561032a57600080fd5b5061033361104a565b005b34801561034157600080fd5b5061034a61186e565b005b34801561035857600080fd5b50610373600480360381019061036e919061706d565b6119b5565b604051610381929190617250565b60405180910390f35b34801561039657600080fd5b506103b160048036038101906103ac9190617303565b611d2e565b005b3480156103bf57600080fd5b506103da60048036038101906103d59190616e72565b611e8a565b6040516103e79190616f82565b60405180910390f35b3480156103fc57600080fd5b5061041760048036038101906104129190616e72565b612104565b6040516104249190617352565b60405180910390f35b34801561043957600080fd5b50610454600480360381019061044f9190616fc6565b612143565b005b34801561046257600080fd5b5061047d60048036038101906104789190616fc6565b612362565b60405161048b929190616f9d565b60405180910390f35b3480156104a057600080fd5b506104a961249c565b6040516104b6919061738e565b60405180910390f35b3480156104cb57600080fd5b506104e660048036038101906104e19190616fc6565b6124c2565b6040516104f39190617352565b60405180910390f35b34801561050857600080fd5b50610523600480360381019061051e91906175c5565b6124f5565b005b34801561053157600080fd5b5061054c60048036038101906105479190616fc6565b612909565b60405161055d95949392919061763d565b60405180910390f35b34801561057257600080fd5b5061058d60048036038101906105889190616f33565b612b07565b60405161059d9493929190617690565b60405180910390f35b3480156105b257600080fd5b506105cd60048036038101906105c89190616e72565b612cad565b6040516105da9190617352565b60405180910390f35b3480156105ef57600080fd5b5061060a60048036038101906106059190616fc6565b612cec565b6040516106179190617352565b60405180910390f35b34801561062c57600080fd5b5061064760048036038101906106429190616f33565b612d2a565b005b34801561065557600080fd5b50610670600480360381019061066b9190616fc6565b612f89565b005b34801561067e57600080fd5b50610699600480360381019061069491906176f1565b61312f565b6040516106a69190617731565b60405180910390f35b3480156106bb57600080fd5b506106d660048036038101906106d19190616fc6565b6132a6565b005b3480156106e457600080fd5b506106ed613571565b6040516106fa9190617731565b60405180910390f35b34801561070f57600080fd5b5061072a60048036038101906107259190616fc6565b6135ff565b005b34801561073857600080fd5b50610753600480360381019061074e9190616f33565b6138ca565b005b34801561076157600080fd5b5061076a613cb3565b60405161077993929190617753565b60405180910390f35b34801561078e57600080fd5b506107a960048036038101906107a49190616fc6565b61412a565b005b3480156107b757600080fd5b506107d260048036038101906107cd9190616f33565b6144ef565b005b3480156107e057600080fd5b506107fb60048036038101906107f6919061779f565b6147a6565b6040516108089190616f82565b60405180910390f35b34801561081d57600080fd5b5061083860048036038101906108339190616e72565b614869565b6040516108459190617352565b60405180910390f35b34801561085a57600080fd5b5061087560048036038101906108709190616e72565b61489c565b6040516108829190617352565b60405180910390f35b34801561089757600080fd5b506108b260048036038101906108ad9190616fc6565b6148db565b6040516108c39594939291906177f2565b60405180910390f35b33600073ffffffffffffffffffffffffffffffffffffffff16600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561099f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610996906178a2565b60405180910390fd5b610a14600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1683600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002061495e9092919063ffffffff16565b3373ffffffffffffffffffffffffffffffffffffffff167f658f52b34d965c76f54298d43b21ae83217d77033214bdafed998bf6df6b9b0583604051610a5a9190616f82565b60405180910390a25050565b60008054906101000a900460ff1681565b6000610aee600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1683600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020614a4d9092919063ffffffff16565b508091505092915050565b600080610bfb60008411610bac57600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b815260040160206040518083038186803b158015610b6f57600080fd5b505afa158015610b83573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ba791906178d7565b610bae565b835b600360008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020614c4390919063ffffffff16565b80925081935050509250929050565b80600073ffffffffffffffffffffffffffffffffffffffff16600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610cdd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cd4906178a2565b60405180910390fd5b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d4a536866040518163ffffffff1660e01b815260040160206040518083038186803b158015610d4557600080fd5b505afa158015610d59573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d7d9190617930565b15610dbd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610db4906179cf565b60405180910390fd5b60003411610e00576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610df790617a3b565b60405180910390fd5b6000600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050600073ffffffffffffffffffffffffffffffffffffffff168160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415610f4357338160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506007339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b610fba600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000203484614c7d909392919063ffffffff16565b8273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f5dac0c1b1112564a045ba943c9d50270893e8e826c49be8e7073adc713ab7bd7346040516110179190616f82565b60405180910390a3505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d4a536866040518163ffffffff1660e01b815260040160206040518083038186803b1580156110b257600080fd5b505afa1580156110c6573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110ea9190617930565b611129576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161112090617acd565b60405180910390fd5b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611197576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161118e90617b39565b60405180910390fd5b60008060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b815260040160206040518083038186803b15801561120057600080fd5b505afa158015611214573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061123891906178d7565b9050600073ffffffffffffffffffffffffffffffffffffffff166008600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146112dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016112d390617ba5565b60405180910390fd5b60008060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633fa4f2456040518163ffffffff1660e01b81526004016101006040518083038186803b15801561134657600080fd5b505afa15801561135a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061137e9190617c94565b905060005b6002805490508110156115355760006003600060056000600286815481106113ae576113ad617cc2565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050838160020154148061149a57508260c0015181600c01600086815260200190815260200160002054105b156114a55750611522565b8381600201819055508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167feb7d7a49847ec491969db21a0e31b234565a9923145a2d1b56a75c9e95825802856040516115189190616f82565b60405180910390a2505b808061152d90617d20565b915050611383565b50600060048054905067ffffffffffffffff811115611557576115566173bf565b5b6040519080825280602002602001820160405280156115855781602001602082028036833780820191505090505b5090506000805b60048054905081101561172057600060036000600484815481106115b3576115b2617cc2565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905061164f600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682614e2690919063ffffffff16565b1561170c5780600a016001876116659190617d69565b90806001815401808255809150506001900390600052602060002001600090919091909150558060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168484815181106116c3576116c2617cc2565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050828061170890617d20565b9350505b50808061171890617d20565b91505061158c565b508067ffffffffffffffff81111561173b5761173a6173bf565b5b6040519080825280602002602001820160405280156117695781602001602082028036833780820191505090505b506002908051906020019061177f929190616cce565b5060005b818110156118155782818151811061179e5761179d617cc2565b5b6020026020010151600282815481106117ba576117b9617cc2565b5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808061180d90617d20565b915050611783565b50336008600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b600073ffffffffffffffffffffffffffffffffffffffff16600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415611940576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161193790617e0b565b60405180910390fd5b6119b3600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020614fbb90919063ffffffff16565b565b60608060008511611a6557600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b815260040160206040518083038186803b158015611a2857600080fd5b505afa158015611a3c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a6091906178d7565b611a67565b845b945060008411611a78576001611a7a565b835b935060008311611a8b576032611a8d565b825b92506000600360008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090506000816008018054905090506000858787611aed9190617e2b565b611af79190617e85565b90508567ffffffffffffffff811115611b1357611b126173bf565b5b604051908082528060200260200182016040528015611b415781602001602082028036833780820191505090505b5094508567ffffffffffffffff811115611b5e57611b5d6173bf565b5b604051908082528060200260200182016040528015611b8c5781602001602082028036833780820191505090505b50935060005b8787611b9e9190617e2b565b821015611d215782821415611bb257611d21565b600060066000866008018581548110611bce57611bcd617cc2565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16878381518110611c6d57611c6c617cc2565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050611ce08560000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168b836151c89092919063ffffffff16565b868381518110611cf357611cf2617cc2565b5b6020026020010181815250508180611d0a90617d20565b925050508180611d1990617d20565b925050611b92565b5050505094509492505050565b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611d9c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d9390617b39565b60405180910390fd5b60008054906101000a900460ff1615611dea576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611de190617f05565b60405180910390fd5b60016000806101000a81548160ff02191690831515021790555081600060016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b600080600183600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b815260040160206040518083038186803b158015611ef857600080fd5b505afa158015611f0c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f3091906178d7565b611f3a9190617e85565b611f449190617e85565b9050600080600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639cf370d76040518163ffffffff1660e01b815260040160006040518083038186803b158015611fb157600080fd5b505afa158015611fc5573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190611fee9190618081565b9150915060005b858110156120fb5760018461200a9190617d69565b935060006120238386866152689092919063ffffffff16565b905060005b6004805490508110156120e6576120c68287600360006004868154811061205257612051617cc2565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206153059092919063ffffffff16565b876120d19190617d69565b965080806120de90617d20565b915050612028565b505080806120f390617d20565b915050611ff5565b50505050919050565b6004818154811061211457600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663322433e3336040518263ffffffff1660e01b815260040161219e9190617352565b60206040518083038186803b1580156121b657600080fd5b505afa1580156121ca573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906121ee9190617930565b61222d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161222490618145565b60405180910390fd5b61227e81600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002061543a90919063ffffffff16565b6004339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555033600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000806000600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020905061245d6004600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b815260040160206040518083038186803b15801561241557600080fd5b505afa158015612429573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061244d91906178d7565b8361553b9092919063ffffffff16565b612492600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16836155bf90919063ffffffff16565b9250925050915091565b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60056020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634ba7ed396040518163ffffffff1660e01b815260040160206040518083038186803b15801561255d57600080fd5b505afa158015612571573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125959190617930565b6125d4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016125cb906181d7565b60405180910390fd5b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614612642576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161263990617b39565b60405180910390fd5b60008060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b815260040160206040518083038186803b1580156126ab57600080fd5b505afa1580156126bf573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906126e391906178d7565b9050600080600090505b845181101561281c5783818151811061270957612708617cc2565b5b6020026020010151600360006005600089868151811061272c5761272b617cc2565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600b016000858152602001908152602001600020819055508381815181106127f4576127f3617cc2565b5b6020026020010151826128079190617d69565b9150808061281490617d20565b9150506126ed565b50600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633fa4f2456040518163ffffffff1660e01b81526004016101006040518083038186803b15801561288657600080fd5b505afa15801561289a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906128be9190617c94565b606001518114612903576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016128fa90618243565b60405180910390fd5b50505050565b600080600080600080600360008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002090508060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168160010160149054906101000a900460ff16612a3c600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b815260040160206040518083038186803b1580156129f557600080fd5b505afa158015612a09573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612a2d91906178d7565b8461574890919063ffffffff16565b612aee600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b815260040160206040518083038186803b158015612aa757600080fd5b505afa158015612abb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612adf91906178d7565b8561576d90919063ffffffff16565b8460020154955095509550955095505091939590929450565b6060806060806004805480602002602001604051908101604052809291908181526020018280548015612b8f57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311612b45575b50505050509350612c98600460008711612c4857600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b815260040160206040518083038186803b158015612c0b57600080fd5b505afa158015612c1f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612c4391906178d7565b612c4a565b865b600660008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206157929092919063ffffffff16565b80935081945082955050505092959194509250565b60028181548110612cbd57600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60066020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905081565b81600073ffffffffffffffffffffffffffffffffffffffff16600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415612dfd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612df4906178a2565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415612ecf576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612ec690617e0b565b60405180910390fd5b612f84600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002084600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206159e3909392919063ffffffff16565b505050565b33600073ffffffffffffffffffffffffffffffffffffffff16600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561305c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613053906178a2565b60405180910390fd5b6130ad82600360003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020615add90919063ffffffff16565b33600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b606060008311613140576001613142565b825b925060008211613153576032613155565b815b91506000600780549050905060008385856131709190617e2b565b61317a9190617e85565b905060008467ffffffffffffffff811115613198576131976173bf565b5b6040519080825280602002602001820160405280156131c65781602001602082028036833780820191505090505b50905060005b86866131d89190617e2b565b83101561329957838314156131ec57613299565b60078381548110613200576131ff617cc2565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682828151811061323e5761323d617cc2565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff1681525050808061328390617d20565b915050828061329190617d20565b9350506131cc565b8194505050505092915050565b80600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806133d05750600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b61340f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613406906182af565b60405180910390fd5b81600073ffffffffffffffffffffffffffffffffffffffff16600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614156134e2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016134d9906178a2565b60405180910390fd5b613529600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020615c27565b8273ffffffffffffffffffffffffffffffffffffffff167fb2b8c5f713f27009f7a78ec2167999c472e5ee8bc591102da17bc4a6551236c860405160405180910390a2505050565b606060048054806020026020016040519081016040528092919081815260200182805480156135f557602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116135ab575b5050505050905090565b80600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806137295750600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b613768576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161375f906182af565b60405180910390fd5b81600073ffffffffffffffffffffffffffffffffffffffff16600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561383b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613832906178a2565b60405180910390fd5b613882600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020615c47565b8273ffffffffffffffffffffffffffffffffffffffff167f23d934bfe7f1275bc6fd70432159c9cc1c0075d069f89da6a40f43bfe7a94ed360405160405180910390a2505050565b81600073ffffffffffffffffffffffffffffffffffffffff16600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561399d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613994906178a2565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415613a6f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613a6690617e0b565b60405180910390fd5b600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d4a536866040518163ffffffff1660e01b815260040160206040518083038186803b158015613ad757600080fd5b505afa158015613aeb573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613b0f9190617930565b15613b4f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613b46906179cf565b60405180910390fd5b60008211613b92576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613b8990617a3b565b60405180910390fd5b613c47600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002084600660003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020615c67909392919063ffffffff16565b91508273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167fd8654fcc8cf5b36d30b3f5e4688fc78118e6d68de60b9994e09902268b57c3e384604051613ca69190616f82565b60405180910390a3505050565b606080606060028054905067ffffffffffffffff811115613cd757613cd66173bf565b5b604051908082528060200260200182016040528015613d055781602001602082028036833780820191505090505b50925060028054905067ffffffffffffffff811115613d2757613d266173bf565b5b604051908082528060200260200182016040528015613d555781602001602082028036833780820191505090505b50915060028054905067ffffffffffffffff811115613d7757613d766173bf565b5b604051908082528060200260200182016040528015613da55781602001602082028036833780820191505090505b50905060008060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b815260040160206040518083038186803b158015613e1157600080fd5b505afa158015613e25573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613e4991906178d7565b9050600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663d4a536866040518163ffffffff1660e01b815260040160206040518083038186803b158015613eb357600080fd5b505afa158015613ec7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613eeb9190617930565b15613eff578080613efb90617d20565b9150505b60005b600280549050811015614123576005600060028381548110613f2757613f26617cc2565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16858281518110613fc057613fbf617cc2565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506002818154811061400e5761400d617cc2565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684828151811061404c5761404b617cc2565b5b602002602001019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250506140f182600360008885815181106140a1576140a0617cc2565b5b602002602001015173ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002061574890919063ffffffff16565b83828151811061410457614103617cc2565b5b602002602001018181525050808061411b90617d20565b915050613f02565b5050909192565b600560008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff16600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561425c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614253906178a2565b60405180910390fd5b4173ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146142ca576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016142c190617b39565b60405180910390fd5b6001600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b815260040160206040518083038186803b15801561433457600080fd5b505afa158015614348573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061436c91906178d7565b116143ac576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016143a390618341565b60405180910390fd5b600060036000600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000209050614483600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1682615f4b90919063ffffffff16565b8060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f1647efd0ce9727dc31dc201c9d8d35ac687f7370adcacbd454afc6485ddabfda60405160405180910390a2505050565b81600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806146195750600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b614658576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161464f906182af565b60405180910390fd5b82600073ffffffffffffffffffffffffffffffffffffffff16600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141561472b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614722906178a2565b60405180910390fd5b6147a0600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684600360008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020615ff99092919063ffffffff16565b50505050565b600061485d600060019054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600360008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002084600660008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020616090909392919063ffffffff16565b50809150509392505050565b60086020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600781815481106148ac57600080fd5b906000526020600020016000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60036020528060005260406000206000915090508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160149054906101000a900460ff16908060020154908060070154905085565b60648111156149a2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614999906183ad565b60405180910390fd5b614a488360040160018473ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b815260040160206040518083038186803b1580156149f157600080fd5b505afa158015614a05573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614a2991906178d7565b614a339190617d69565b8386600301616328909392919063ffffffff16565b505050565b60008084600701549050600060018573ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b815260040160206040518083038186803b158015614aa157600080fd5b505afa158015614ab5573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614ad991906178d7565b614ae39190617e85565b90506000841480614afe5750808285614afc9190617d69565b115b15614b12578181614b0f9190617e85565b93505b6000808673ffffffffffffffffffffffffffffffffffffffff16639cf370d76040518163ffffffff1660e01b815260040160006040518083038186803b158015614b5b57600080fd5b505afa158015614b6f573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190614b989190618081565b9150915060005b86811015614c3757600185614bb49190617d69565b94506000614bd78a614bd18589886152689092919063ffffffff16565b88615305565b90506000811415614be85750614c24565b6000614bf48b8861576d565b90506000811415614c06575050614c24565b614c1482826064601961636b565b88614c1f9190617d69565b975050505b8080614c2f90617d20565b915050614b9f565b50505050935093915050565b60008083600b0160008481526020019081526020016000205484600c01600085815260200190815260200160002054915091509250929050565b614de58460020160008460000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060018573ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b815260040160206040518083038186803b158015614d2d57600080fd5b505afa158015614d41573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614d6591906178d7565b614d6f9190617d69565b838760010160008760000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206163a4909392919063ffffffff16565b614e20838560000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1683856163f9909392919063ffffffff16565b50505050565b60008260010160149054906101000a900460ff16614e475760009050614fb5565b60008273ffffffffffffffffffffffffffffffffffffffff16633fa4f2456040518163ffffffff1660e01b81526004016101006040518083038186803b158015614e9057600080fd5b505afa158015614ea4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614ec89190617c94565b905060008373ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b815260040160206040518083038186803b158015614f1257600080fd5b505afa158015614f26573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190614f4a91906178d7565b905060008560020154118015614f7257508160e00151856002015482614f709190617e85565b105b15614f8257600092505050614fb5565b8160a00151614f9d86600184614f989190617d69565b615748565b1015614fae57600092505050614fb5565b6001925050505b92915050565b6000614fc783836155bf565b90506000811415614fd857506151c4565b6000836003018054905090508273ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b815260040160206040518083038186803b15801561502a57600080fd5b505afa15801561503e573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061506291906178d7565b846003016001836150739190617e85565b8154811061508457615083617cc2565b5b9060005260206000200154116150b9578360030160006150a49190616d58565b8360040160006150b49190616d58565b615156565b6040518060200160405280856003016001846150d59190617e85565b815481106150e6576150e5617cc2565b5b906000526020600020015481525084600301906001615106929190616d79565b506040518060200160405280856004016001846151239190617e85565b8154811061513457615133617cc2565b5b906000526020600020015481525084600401906001615154929190616d79565b505b8360000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f193505050501580156151c0573d6000803e3d6000fd5b5050505b5050565b600061525f8460020160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020838660010160008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206165b89092919063ffffffff16565b90509392505050565b615270616dc6565b8184600186516152809190617e85565b8151811061529157615290617cc2565b5b6020026020010151116152cd5782600185516152ad9190617e85565b815181106152be576152bd617cc2565b5b602002602001015190506152fe565b60006152dd85846000885161669e565b90508381815181106152f2576152f1617cc2565b5b60200260200101519150505b9392505050565b6000806153128584615748565b90506000811415615327576000915050615433565b600085600b016000858152602001908152602001600020549050600086600c016000868152602001908152602001600020549050600082148061536a5750818110155b1561537b5760009350505050615433565b60006019600a61538b9190618500565b61539c88608001516064601961675d565b856153a79190617e2b565b6153b1919061857a565b905060008114156153c9576000945050505050615433565b6153ec876060015188604001516153e09190617e2b565b6301e13380601961675d565b816153f79190617e2b565b90506019600a6154079190618500565b81615412919061857a565b905061542c8183856154249190617e85565b85601961636b565b9450505050505b9392505050565b600073ffffffffffffffffffffffffffffffffffffffff168260000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146154cd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016154c4906185f7565b60405180910390fd5b338260000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060018260010160146101000a81548160ff0219169083151502179055506155378282615add565b5050565b600080600090505b83805490508110156155b7576155978585838154811061556657615565617cc2565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16856151c8565b826155a29190617d69565b915080806155af90617d20565b915050615543565b509392505050565b6000808360030180549050905060008114156155df576000915050615742565b60008373ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b815260040160206040518083038186803b15801561562757600080fd5b505afa15801561563b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061565f91906178d7565b905060006001836156709190617e85565b90506000811180156156a157508186600301828154811061569457615693617cc2565b5b9060005260206000200154115b156156b55780806156b190618617565b9150505b818660030182815481106156cc576156cb617cc2565b5b906000526020600020015411156156e95760009350505050615742565b600080600090505b8281116157395787600401818154811061570e5761570d617cc2565b5b9060005260206000200154826157249190617d69565b9150808061573190617d20565b9150506156f1565b50809450505050505b92915050565b60006157658360060183856005016165b89092919063ffffffff16565b905092915050565b600061578a8360040183856003016165b89092919063ffffffff16565b905092915050565b6060806060848054905067ffffffffffffffff8111156157b5576157b46173bf565b5b6040519080825280602002602001820160405280156157e35781602001602082028036833780820191505090505b509250848054905067ffffffffffffffff811115615804576158036173bf565b5b6040519080825280602002602001820160405280156158325781602001602082028036833780820191505090505b509150848054905067ffffffffffffffff811115615853576158526173bf565b5b6040519080825280602002602001820160405280156158815781602001602082028036833780820191505090505b50905060005b85805490508110156159d95760006158dd888884815481106158ac576158ab617cc2565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16886151c8565b90506000615935898985815481106158f8576158f7617cc2565b5b9060005260206000200160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660018a6159309190617d69565b6151c8565b90508186848151811061594b5761594a617cc2565b5b6020026020010181815250508181111561598f57818161596b9190617e85565b85848151811061597e5761597d617cc2565b5b6020026020010181815250506159c4565b818110156159c35780826159a39190617e85565b8484815181106159b6576159b5617cc2565b5b6020026020010181815250505b5b505080806159d190617d20565b915050615887565b5093509350939050565b6000806159f286868686616090565b91509150808660050160008660000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506000821115615ad5578560000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f19350505050158015615ad3573d6000803e3d6000fd5b505b505050505050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415615b4d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615b449061868d565b60405180910390fd5b8160000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415615be0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615bd7906186f9565b60405180910390fd5b808260010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b60018160010160146101000a81548160ff02191690831515021790555050565b60008160010160146101000a81548160ff02191690831515021790555050565b6000808473ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b815260040160206040518083038186803b158015615cb057600080fd5b505afa158015615cc4573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615ce891906178d7565b90506000615d1b878660000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16846151c8565b90506000615d5a888760000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600186615d559190617d69565b6151c8565b9050615e478860020160008860000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600185615dd19190617d69565b878b60010160008b60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206167b4909392919063ffffffff16565b94506000851415615e5e5760009350505050615f43565b615e7387868861685f9092919063ffffffff16565b6000859050600083831115615eb0578383615e8e9190617e85565b9050808710615e9d5780615e9f565b865b90508082615ead9190617e85565b91505b6000821115615ec557615ec48a8a8461690b565b5b6000811115615f3a578960000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc829081150290604051600060405180830381858888f19350505050158015615f38573d6000803e3d6000fd5b505b86955050505050505b949350505050565b600182600c0160008373ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b815260040160206040518083038186803b158015615f9957600080fd5b505afa158015615fad573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190615fd191906178d7565b81526020019081526020016000206000828254615fee9190617d69565b925050819055505050565b600080616007858585614a4d565b915091508085600701819055506000821115616089578460000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc839081150290604051600060405180830381858888f19350505050158015616087573d6000803e3d6000fd5b505b5050505050565b6000808560050160008560000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020549050600060018673ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b815260040160206040518083038186803b15801561614557600080fd5b505afa158015616159573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061617d91906178d7565b6161879190617e85565b905060008414806161a257508082856161a09190617d69565b115b156161b65781816161b39190617e85565b93505b6000808773ffffffffffffffffffffffffffffffffffffffff16639cf370d76040518163ffffffff1660e01b815260040160006040518083038186803b1580156161ff57600080fd5b505afa158015616213573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061623c9190618081565b9150915060005b8681101561631b576001856162589190617d69565b9450600061628b8b8a60000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16886151c8565b9050600081141561629c5750616308565b60006162c76162b68589886152689092919063ffffffff16565b888c616a8d9092919063ffffffff16565b905060008114156162d9575050616308565b6162f881836162f18a8e61574890919063ffffffff16565b601961636b565b886163039190617d69565b975050505b808061631390617d20565b915050616243565b5050505094509492505050565b616333848484616af6565b8083600186805490506163469190617e85565b8154811061635757616356617cc2565b5b906000526020600020018190555050505050565b600081600a61637a9190618500565b61638585858561675d565b866163909190617e2b565b61639a919061857a565b9050949350505050565b6163af848484616af6565b8083600186805490506163c29190617e85565b815481106163d3576163d2617cc2565b5b9060005260206000200160008282546163ec9190617d69565b9250508190555050505050565b8360090160008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1661650c5760018460090160008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555083600801829080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b6165b28460060160018573ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b815260040160206040518083038186803b15801561655b57600080fd5b505afa15801561656f573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061659391906178d7565b61659d9190617d69565b83876005016163a4909392919063ffffffff16565b50505050565b600080848054905014806165e9575081846000815481106165dc576165db617cc2565b5b9060005260206000200154115b156165f75760009050616697565b81846001868054905061660a9190617e85565b8154811061661b5761661a617cc2565b5b90600052602060002001541161666057826001858054905061663d9190617e85565b8154811061664e5761664d617cc2565b5b90600052602060002001549050616697565b6000616673858460008880549050616c09565b905083818154811061668857616687617cc2565b5b90600052602060002001549150505b9392505050565b6000818314156166bc576001826166b59190617e85565b9050616755565b6000600283856166cc9190617d69565b6166d6919061857a565b9050848682815181106166ec576166eb617cc2565b5b6020026020010151111561670e576167068686868461669e565b915050616755565b8486828151811061672257616721617cc2565b5b602002602001015110156167505761674886866001846167429190617d69565b8661669e565b915050616755565b809150505b949350505050565b60008060018361676d9190617d69565b600a6167799190618500565b856167849190617e2b565b9050600a60058583616796919061857a565b6167a09190617d69565b6167aa919061857a565b9150509392505050565b60006167c1858585616af6565b600084600187805490506167d59190617e85565b815481106167e6576167e5617cc2565b5b90600052602060002001549050808311156168015780616803565b825b92506000831115616853578285600188805490506168219190617e85565b8154811061683257616831617cc2565b5b90600052602060002001600082825461684b9190617e85565b925050819055505b82915050949350505050565b6169058360060160018473ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b815260040160206040518083038186803b1580156168ae57600080fd5b505afa1580156168c2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906168e691906178d7565b6168f09190617d69565b83866005016167b4909392919063ffffffff16565b50505050565b600060018373ffffffffffffffffffffffffffffffffffffffff1663900cf0cf6040518163ffffffff1660e01b815260040160206040518083038186803b15801561695557600080fd5b505afa158015616969573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061698d91906178d7565b6169979190617d69565b905060008460030180549050905060008114806169e0575081856003016001836169c19190617e85565b815481106169d2576169d1617cc2565b5b906000526020600020015414155b15616a425784600301829080600181540180825580915050600190039060005260206000200160009091909190915055846004018390806001815401808255809150506001900390600052602060002001600090919091909150555050616a88565b8285600401600183616a549190617e85565b81548110616a6557616a64617cc2565b5b906000526020600020016000828254616a7e9190617d69565b9250508190555050505b505050565b600080616a9b858585615305565b90506000811415616ab0576000915050616aef565b6000616abc868561576d565b90506000811415616ad1578192505050616aef565b616adf82826064601961636b565b82616aea9190617e85565b925050505b9392505050565b6000838054905090506000811415616b5257838290806001815401808255809150506001900390600052602060002001600090919091909150558260018160018154018082558091505003906000526020600020505050616c04565b600084600183616b629190617e85565b81548110616b7357616b72617cc2565b5b90600052602060002001549050828114616c0157848390806001815401808255809150506001900390600052602060002001600090919091909150558384600184616bbe9190617e85565b81548110616bcf57616bce617cc2565b5b906000526020600020015490806001815401808255809150506001900390600052602060002001600090919091909150555b50505b505050565b600081831415616c2757600182616c209190617e85565b9050616cc6565b600060028385616c379190617d69565b616c41919061857a565b905084868281548110616c5757616c56617cc2565b5b90600052602060002001541115616c7c57616c7486868684616c09565b915050616cc6565b84868281548110616c9057616c8f617cc2565b5b90600052602060002001541015616cc157616cb98686600184616cb39190617d69565b86616c09565b915050616cc6565b809150505b949350505050565b828054828255906000526020600020908101928215616d47579160200282015b82811115616d465782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190616cee565b5b509050616d549190616e0b565b5090565b5080546000825590600052602060002090810190616d769190616e0b565b50565b828054828255906000526020600020908101928215616db5579160200282015b82811115616db4578251825591602001919060010190616d99565b5b509050616dc29190616e0b565b5090565b60405180610100016040528060008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b5b80821115616e24576000816000905550600101616e0c565b5090565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b616e4f81616e3c565b8114616e5a57600080fd5b50565b600081359050616e6c81616e46565b92915050565b600060208284031215616e8857616e87616e32565b5b6000616e9684828501616e5d565b91505092915050565b60008115159050919050565b616eb481616e9f565b82525050565b6000602082019050616ecf6000830184616eab565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000616f0082616ed5565b9050919050565b616f1081616ef5565b8114616f1b57600080fd5b50565b600081359050616f2d81616f07565b92915050565b60008060408385031215616f4a57616f49616e32565b5b6000616f5885828601616f1e565b9250506020616f6985828601616e5d565b9150509250929050565b616f7c81616e3c565b82525050565b6000602082019050616f976000830184616f73565b92915050565b6000604082019050616fb26000830185616f73565b616fbf6020830184616f73565b9392505050565b600060208284031215616fdc57616fdb616e32565b5b6000616fea84828501616f1e565b91505092915050565b6000819050919050565b600061701861701361700e84616ed5565b616ff3565b616ed5565b9050919050565b600061702a82616ffd565b9050919050565b600061703c8261701f565b9050919050565b61704c81617031565b82525050565b60006020820190506170676000830184617043565b92915050565b6000806000806080858703121561708757617086616e32565b5b600061709587828801616f1e565b94505060206170a687828801616e5d565b93505060406170b787828801616e5d565b92505060606170c887828801616e5d565b91505092959194509250565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61710981616ef5565b82525050565b600061711b8383617100565b60208301905092915050565b6000602082019050919050565b600061713f826170d4565b61714981856170df565b9350617154836170f0565b8060005b8381101561718557815161716c888261710f565b975061717783617127565b925050600181019050617158565b5085935050505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6171c781616e3c565b82525050565b60006171d983836171be565b60208301905092915050565b6000602082019050919050565b60006171fd82617192565b617207818561719d565b9350617212836171ae565b8060005b8381101561724357815161722a88826171cd565b9750617235836171e5565b925050600181019050617216565b5085935050505092915050565b6000604082019050818103600083015261726a8185617134565b9050818103602083015261727e81846171f2565b90509392505050565b600061729282616ef5565b9050919050565b6172a281617287565b81146172ad57600080fd5b50565b6000813590506172bf81617299565b92915050565b60006172d082616ef5565b9050919050565b6172e0816172c5565b81146172eb57600080fd5b50565b6000813590506172fd816172d7565b92915050565b6000806040838503121561731a57617319616e32565b5b6000617328858286016172b0565b9250506020617339858286016172ee565b9150509250929050565b61734c81616ef5565b82525050565b60006020820190506173676000830184617343565b92915050565b60006173788261701f565b9050919050565b6173888161736d565b82525050565b60006020820190506173a3600083018461737f565b92915050565b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6173f7826173ae565b810181811067ffffffffffffffff82111715617416576174156173bf565b5b80604052505050565b6000617429616e28565b905061743582826173ee565b919050565b600067ffffffffffffffff821115617455576174546173bf565b5b602082029050602081019050919050565b600080fd5b600061747e6174798461743a565b61741f565b905080838252602082019050602084028301858111156174a1576174a0617466565b5b835b818110156174ca57806174b68882616f1e565b8452602084019350506020810190506174a3565b5050509392505050565b600082601f8301126174e9576174e86173a9565b5b81356174f984826020860161746b565b91505092915050565b600067ffffffffffffffff82111561751d5761751c6173bf565b5b602082029050602081019050919050565b600061754161753c84617502565b61741f565b9050808382526020820190506020840283018581111561756457617563617466565b5b835b8181101561758d57806175798882616e5d565b845260208401935050602081019050617566565b5050509392505050565b600082601f8301126175ac576175ab6173a9565b5b81356175bc84826020860161752e565b91505092915050565b600080604083850312156175dc576175db616e32565b5b600083013567ffffffffffffffff8111156175fa576175f9616e37565b5b617606858286016174d4565b925050602083013567ffffffffffffffff81111561762757617626616e37565b5b61763385828601617597565b9150509250929050565b600060a0820190506176526000830188617343565b61765f6020830187616eab565b61766c6040830186616f73565b6176796060830185616f73565b6176866080830184616f73565b9695505050505050565b600060808201905081810360008301526176aa8187617134565b905081810360208301526176be81866171f2565b905081810360408301526176d281856171f2565b905081810360608301526176e681846171f2565b905095945050505050565b6000806040838503121561770857617707616e32565b5b600061771685828601616e5d565b925050602061772785828601616e5d565b9150509250929050565b6000602082019050818103600083015261774b8184617134565b905092915050565b6000606082019050818103600083015261776d8186617134565b905081810360208301526177818185617134565b9050818103604083015261779581846171f2565b9050949350505050565b6000806000606084860312156177b8576177b7616e32565b5b60006177c686828701616f1e565b93505060206177d786828701616f1e565b92505060406177e886828701616e5d565b9150509250925092565b600060a0820190506178076000830188617343565b6178146020830187617343565b6178216040830186616eab565b61782e6060830185616f73565b61783b6080830184616f73565b9695505050505050565b600082825260208201905092915050565b7f76616c696461746f7220646f6573206e6f742065786973742e00000000000000600082015250565b600061788c601983617845565b915061789782617856565b602082019050919050565b600060208201905081810360008301526178bb8161787f565b9050919050565b6000815190506178d181616e46565b92915050565b6000602082840312156178ed576178ec616e32565b5b60006178fb848285016178c2565b91505092915050565b61790d81616e9f565b811461791857600080fd5b50565b60008151905061792a81617904565b92915050565b60006020828403121561794657617945616e32565b5b60006179548482850161791b565b91505092915050565b7f6e6f742065786563757461626c6520696e20746865206c61737420626c6f636b60008201527f206f662065706f63682e00000000000000000000000000000000000000000000602082015250565b60006179b9602a83617845565b91506179c48261795d565b604082019050919050565b600060208201905081810360008301526179e8816179ac565b9050919050565b7f616d6f756e74206973207a65726f2e0000000000000000000000000000000000600082015250565b6000617a25600f83617845565b9150617a30826179ef565b602082019050919050565b60006020820190508181036000830152617a5481617a18565b9050919050565b7f6f6e6c792065786563757461626c6520696e20746865206c61737420626c6f6360008201527f6b206f662065706f63682e000000000000000000000000000000000000000000602082015250565b6000617ab7602b83617845565b9150617ac282617a5b565b604082019050919050565b60006020820190508181036000830152617ae681617aaa565b9050919050565b7f73656e646572206d75737420626520626c6f636b2070726f64756365722e0000600082015250565b6000617b23601e83617845565b9150617b2e82617aed565b602082019050919050565b60006020820190508181036000830152617b5281617b16565b9050919050565b7f616c726561647920757064617465642e00000000000000000000000000000000600082015250565b6000617b8f601083617845565b9150617b9a82617b59565b602082019050919050565b60006020820190508181036000830152617bbe81617b82565b9050919050565b600080fd5b60006101008284031215617be157617be0617bc5565b5b617bec61010061741f565b90506000617bfc848285016178c2565b6000830152506020617c10848285016178c2565b6020830152506040617c24848285016178c2565b6040830152506060617c38848285016178c2565b6060830152506080617c4c848285016178c2565b60808301525060a0617c60848285016178c2565b60a08301525060c0617c74848285016178c2565b60c08301525060e0617c88848285016178c2565b60e08301525092915050565b60006101008284031215617cab57617caa616e32565b5b6000617cb984828501617bca565b91505092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000617d2b82616e3c565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415617d5e57617d5d617cf1565b5b600182019050919050565b6000617d7482616e3c565b9150617d7f83616e3c565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115617db457617db3617cf1565b5b828201905092915050565b7f7374616b657220646f6573206e6f742065786973742e00000000000000000000600082015250565b6000617df5601683617845565b9150617e0082617dbf565b602082019050919050565b60006020820190508181036000830152617e2481617de8565b9050919050565b6000617e3682616e3c565b9150617e4183616e3c565b9250817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615617e7a57617e79617cf1565b5b828202905092915050565b6000617e9082616e3c565b9150617e9b83616e3c565b925082821015617eae57617ead617cf1565b5b828203905092915050565b7f616c726561647920696e697469616c697a65642e000000000000000000000000600082015250565b6000617eef601483617845565b9150617efa82617eb9565b602082019050919050565b60006020820190508181036000830152617f1e81617ee2565b9050919050565b6000617f38617f3384617502565b61741f565b90508083825260208201905060208402830185811115617f5b57617f5a617466565b5b835b81811015617f845780617f7088826178c2565b845260208401935050602081019050617f5d565b5050509392505050565b600082601f830112617fa357617fa26173a9565b5b8151617fb3848260208601617f25565b91505092915050565b600067ffffffffffffffff821115617fd757617fd66173bf565b5b602082029050602081019050919050565b6000617ffb617ff684617fbc565b61741f565b905080838252602082019050610100840283018581111561801f5761801e617466565b5b835b8181101561804957806180348882617bca565b84526020840193505061010081019050618021565b5050509392505050565b600082601f830112618068576180676173a9565b5b8151618078848260208601617fe8565b91505092915050565b6000806040838503121561809857618097616e32565b5b600083015167ffffffffffffffff8111156180b6576180b5616e37565b5b6180c285828601617f8e565b925050602083015167ffffffffffffffff8111156180e3576180e2616e37565b5b6180ef85828601618053565b9150509250929050565b7f6e6f7420616c6c6f7765642e0000000000000000000000000000000000000000600082015250565b600061812f600c83617845565b915061813a826180f9565b602082019050919050565b6000602082019050818103600083015261815e81618122565b9050919050565b7f6f6e6c792065786563757461626c6520696e2074686520666972737420626c6f60008201527f636b206f662065706f63682e0000000000000000000000000000000000000000602082015250565b60006181c1602c83617845565b91506181cc82618165565b604082019050919050565b600060208201905081810360008301526181f0816181b4565b9050919050565b7f626c6f636b20636f756e74206973206d69736d617463682e0000000000000000600082015250565b600061822d601883617845565b9150618238826181f7565b602082019050919050565b6000602082019050818103600083015261825c81618220565b9050919050565b7f796f7520617265206e6f74206f776e6572206f72206f70657261746f722e0000600082015250565b6000618299601e83617845565b91506182a482618263565b602082019050919050565b600060208201905081810360008301526182c88161828c565b9050919050565b7f6e6f742065786563757461626c6520696e207468652066697273742065706f6360008201527f682e000000000000000000000000000000000000000000000000000000000000602082015250565b600061832b602283617845565b9150618336826182cf565b604082019050919050565b6000602082019050818103600083015261835a8161831e565b9050919050565b7f6d757374206265206c657373207468616e203130302e00000000000000000000600082015250565b6000618397601683617845565b91506183a282618361565b602082019050919050565b600060208201905081810360008301526183c68161838a565b9050919050565b60008160011c9050919050565b6000808291508390505b600185111561842457808604811115618400576183ff617cf1565b5b600185161561840f5780820291505b808102905061841d856183cd565b94506183e4565b94509492505050565b60008261843d57600190506184f9565b8161844b57600090506184f9565b8160018114618461576002811461846b5761849a565b60019150506184f9565b60ff84111561847d5761847c617cf1565b5b8360020a91508482111561849457618493617cf1565b5b506184f9565b5060208310610133831016604e8410600b84101617156184cf5782820a9050838111156184ca576184c9617cf1565b5b6184f9565b6184dc84848460016183da565b925090508184048111156184f3576184f2617cf1565b5b81810290505b9392505050565b600061850b82616e3c565b915061851683616e3c565b92506185437fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff848461842d565b905092915050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600061858582616e3c565b915061859083616e3c565b9250826185a05761859f61854b565b5b828204905092915050565b7f616c7265616479206a6f696e65642e0000000000000000000000000000000000600082015250565b60006185e1600f83617845565b91506185ec826185ab565b602082019050919050565b60006020820190508181036000830152618610816185d4565b9050919050565b600061862282616e3c565b9150600082141561863657618635617cf1565b5b600182039050919050565b7f6f70657261746f72206973207a65726f20616464726573732e00000000000000600082015250565b6000618677601983617845565b915061868282618641565b602082019050919050565b600060208201905081810360008301526186a68161866a565b9050919050565b7f6f70657261746f722069732073616d65206173206f776e65722e000000000000600082015250565b60006186e3601a83617845565b91506186ee826186ad565b602082019050919050565b60006020820190508181036000830152618712816186d6565b905091905056fea2646970667358221220519296a4dab2c35f750e0e5fd953e7524cea32f614776b21afb8c49f76bbb18364736f6c63430008090033",
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

// CurrentValidators is a free data retrieval call binding the contract method 0x900eb5a8.
//
// Solidity: function currentValidators(uint256 ) view returns(address)
func (_Stakemanager *StakemanagerCaller) CurrentValidators(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "currentValidators", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CurrentValidators is a free data retrieval call binding the contract method 0x900eb5a8.
//
// Solidity: function currentValidators(uint256 ) view returns(address)
func (_Stakemanager *StakemanagerSession) CurrentValidators(arg0 *big.Int) (common.Address, error) {
	return _Stakemanager.Contract.CurrentValidators(&_Stakemanager.CallOpts, arg0)
}

// CurrentValidators is a free data retrieval call binding the contract method 0x900eb5a8.
//
// Solidity: function currentValidators(uint256 ) view returns(address)
func (_Stakemanager *StakemanagerCallerSession) CurrentValidators(arg0 *big.Int) (common.Address, error) {
	return _Stakemanager.Contract.CurrentValidators(&_Stakemanager.CallOpts, arg0)
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

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(address staker) view returns(uint256 stakes, uint256 unstakes)
func (_Stakemanager *StakemanagerCaller) GetStakerInfo(opts *bind.CallOpts, staker common.Address) (struct {
	Stakes   *big.Int
	Unstakes *big.Int
}, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getStakerInfo", staker)

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

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(address staker) view returns(uint256 stakes, uint256 unstakes)
func (_Stakemanager *StakemanagerSession) GetStakerInfo(staker common.Address) (struct {
	Stakes   *big.Int
	Unstakes *big.Int
}, error) {
	return _Stakemanager.Contract.GetStakerInfo(&_Stakemanager.CallOpts, staker)
}

// GetStakerInfo is a free data retrieval call binding the contract method 0x733bdef0.
//
// Solidity: function getStakerInfo(address staker) view returns(uint256 stakes, uint256 unstakes)
func (_Stakemanager *StakemanagerCallerSession) GetStakerInfo(staker common.Address) (struct {
	Stakes   *big.Int
	Unstakes *big.Int
}, error) {
	return _Stakemanager.Contract.GetStakerInfo(&_Stakemanager.CallOpts, staker)
}

// GetStakerStakes is a free data retrieval call binding the contract method 0x8bca036b.
//
// Solidity: function getStakerStakes(address staker, uint256 epoch) view returns(address[] _validators, uint256[] stakes, uint256[] stakeRequests, uint256[] unstakeRequests)
func (_Stakemanager *StakemanagerCaller) GetStakerStakes(opts *bind.CallOpts, staker common.Address, epoch *big.Int) (struct {
	Validators      []common.Address
	Stakes          []*big.Int
	StakeRequests   []*big.Int
	UnstakeRequests []*big.Int
}, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getStakerStakes", staker, epoch)

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

// GetStakerStakes is a free data retrieval call binding the contract method 0x8bca036b.
//
// Solidity: function getStakerStakes(address staker, uint256 epoch) view returns(address[] _validators, uint256[] stakes, uint256[] stakeRequests, uint256[] unstakeRequests)
func (_Stakemanager *StakemanagerSession) GetStakerStakes(staker common.Address, epoch *big.Int) (struct {
	Validators      []common.Address
	Stakes          []*big.Int
	StakeRequests   []*big.Int
	UnstakeRequests []*big.Int
}, error) {
	return _Stakemanager.Contract.GetStakerStakes(&_Stakemanager.CallOpts, staker, epoch)
}

// GetStakerStakes is a free data retrieval call binding the contract method 0x8bca036b.
//
// Solidity: function getStakerStakes(address staker, uint256 epoch) view returns(address[] _validators, uint256[] stakes, uint256[] stakeRequests, uint256[] unstakeRequests)
func (_Stakemanager *StakemanagerCallerSession) GetStakerStakes(staker common.Address, epoch *big.Int) (struct {
	Validators      []common.Address
	Stakes          []*big.Int
	StakeRequests   []*big.Int
	UnstakeRequests []*big.Int
}, error) {
	return _Stakemanager.Contract.GetStakerStakes(&_Stakemanager.CallOpts, staker, epoch)
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
// Solidity: function getValidatorInfo(address validator) view returns(address operator, bool active, uint256 stakes, uint256 commissionRate, uint256 jailEpoch)
func (_Stakemanager *StakemanagerCaller) GetValidatorInfo(opts *bind.CallOpts, validator common.Address) (struct {
	Operator       common.Address
	Active         bool
	Stakes         *big.Int
	CommissionRate *big.Int
	JailEpoch      *big.Int
}, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "getValidatorInfo", validator)

	outstruct := new(struct {
		Operator       common.Address
		Active         bool
		Stakes         *big.Int
		CommissionRate *big.Int
		JailEpoch      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Operator = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Active = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Stakes = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CommissionRate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.JailEpoch = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetValidatorInfo is a free data retrieval call binding the contract method 0x8a11d7c9.
//
// Solidity: function getValidatorInfo(address validator) view returns(address operator, bool active, uint256 stakes, uint256 commissionRate, uint256 jailEpoch)
func (_Stakemanager *StakemanagerSession) GetValidatorInfo(validator common.Address) (struct {
	Operator       common.Address
	Active         bool
	Stakes         *big.Int
	CommissionRate *big.Int
	JailEpoch      *big.Int
}, error) {
	return _Stakemanager.Contract.GetValidatorInfo(&_Stakemanager.CallOpts, validator)
}

// GetValidatorInfo is a free data retrieval call binding the contract method 0x8a11d7c9.
//
// Solidity: function getValidatorInfo(address validator) view returns(address operator, bool active, uint256 stakes, uint256 commissionRate, uint256 jailEpoch)
func (_Stakemanager *StakemanagerCallerSession) GetValidatorInfo(validator common.Address) (struct {
	Operator       common.Address
	Active         bool
	Stakes         *big.Int
	CommissionRate *big.Int
	JailEpoch      *big.Int
}, error) {
	return _Stakemanager.Contract.GetValidatorInfo(&_Stakemanager.CallOpts, validator)
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

// ValidatorUpdates is a free data retrieval call binding the contract method 0xf4fa1dde.
//
// Solidity: function validatorUpdates(uint256 ) view returns(address)
func (_Stakemanager *StakemanagerCaller) ValidatorUpdates(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "validatorUpdates", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ValidatorUpdates is a free data retrieval call binding the contract method 0xf4fa1dde.
//
// Solidity: function validatorUpdates(uint256 ) view returns(address)
func (_Stakemanager *StakemanagerSession) ValidatorUpdates(arg0 *big.Int) (common.Address, error) {
	return _Stakemanager.Contract.ValidatorUpdates(&_Stakemanager.CallOpts, arg0)
}

// ValidatorUpdates is a free data retrieval call binding the contract method 0xf4fa1dde.
//
// Solidity: function validatorUpdates(uint256 ) view returns(address)
func (_Stakemanager *StakemanagerCallerSession) ValidatorUpdates(arg0 *big.Int) (common.Address, error) {
	return _Stakemanager.Contract.ValidatorUpdates(&_Stakemanager.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(address owner, address operator, bool active, uint256 jailEpoch, uint256 lastClaimCommission)
func (_Stakemanager *StakemanagerCaller) Validators(opts *bind.CallOpts, arg0 common.Address) (struct {
	Owner               common.Address
	Operator            common.Address
	Active              bool
	JailEpoch           *big.Int
	LastClaimCommission *big.Int
}, error) {
	var out []interface{}
	err := _Stakemanager.contract.Call(opts, &out, "validators", arg0)

	outstruct := new(struct {
		Owner               common.Address
		Operator            common.Address
		Active              bool
		JailEpoch           *big.Int
		LastClaimCommission *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Operator = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Active = *abi.ConvertType(out[2], new(bool)).(*bool)
	outstruct.JailEpoch = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LastClaimCommission = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(address owner, address operator, bool active, uint256 jailEpoch, uint256 lastClaimCommission)
func (_Stakemanager *StakemanagerSession) Validators(arg0 common.Address) (struct {
	Owner               common.Address
	Operator            common.Address
	Active              bool
	JailEpoch           *big.Int
	LastClaimCommission *big.Int
}, error) {
	return _Stakemanager.Contract.Validators(&_Stakemanager.CallOpts, arg0)
}

// Validators is a free data retrieval call binding the contract method 0xfa52c7d8.
//
// Solidity: function validators(address ) view returns(address owner, address operator, bool active, uint256 jailEpoch, uint256 lastClaimCommission)
func (_Stakemanager *StakemanagerCallerSession) Validators(arg0 common.Address) (struct {
	Owner               common.Address
	Operator            common.Address
	Active              bool
	JailEpoch           *big.Int
	LastClaimCommission *big.Int
}, error) {
	return _Stakemanager.Contract.Validators(&_Stakemanager.CallOpts, arg0)
}

// ActivateValidator is a paid mutator transaction binding the contract method 0xb46e5520.
//
// Solidity: function activateValidator(address validator) returns()
func (_Stakemanager *StakemanagerTransactor) ActivateValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "activateValidator", validator)
}

// ActivateValidator is a paid mutator transaction binding the contract method 0xb46e5520.
//
// Solidity: function activateValidator(address validator) returns()
func (_Stakemanager *StakemanagerSession) ActivateValidator(validator common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.ActivateValidator(&_Stakemanager.TransactOpts, validator)
}

// ActivateValidator is a paid mutator transaction binding the contract method 0xb46e5520.
//
// Solidity: function activateValidator(address validator) returns()
func (_Stakemanager *StakemanagerTransactorSession) ActivateValidator(validator common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.ActivateValidator(&_Stakemanager.TransactOpts, validator)
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

// DeactivateValidator is a paid mutator transaction binding the contract method 0xba50b879.
//
// Solidity: function deactivateValidator(address validator) returns()
func (_Stakemanager *StakemanagerTransactor) DeactivateValidator(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "deactivateValidator", validator)
}

// DeactivateValidator is a paid mutator transaction binding the contract method 0xba50b879.
//
// Solidity: function deactivateValidator(address validator) returns()
func (_Stakemanager *StakemanagerSession) DeactivateValidator(validator common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.DeactivateValidator(&_Stakemanager.TransactOpts, validator)
}

// DeactivateValidator is a paid mutator transaction binding the contract method 0xba50b879.
//
// Solidity: function deactivateValidator(address validator) returns()
func (_Stakemanager *StakemanagerTransactorSession) DeactivateValidator(validator common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.DeactivateValidator(&_Stakemanager.TransactOpts, validator)
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

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address operator) returns()
func (_Stakemanager *StakemanagerTransactor) Slash(opts *bind.TransactOpts, operator common.Address) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "slash", operator)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address operator) returns()
func (_Stakemanager *StakemanagerSession) Slash(operator common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.Slash(&_Stakemanager.TransactOpts, operator)
}

// Slash is a paid mutator transaction binding the contract method 0xc96be4cb.
//
// Solidity: function slash(address operator) returns()
func (_Stakemanager *StakemanagerTransactorSession) Slash(operator common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.Slash(&_Stakemanager.TransactOpts, operator)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address validator) payable returns()
func (_Stakemanager *StakemanagerTransactor) Stake(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "stake", validator)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address validator) payable returns()
func (_Stakemanager *StakemanagerSession) Stake(validator common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.Stake(&_Stakemanager.TransactOpts, validator)
}

// Stake is a paid mutator transaction binding the contract method 0x26476204.
//
// Solidity: function stake(address validator) payable returns()
func (_Stakemanager *StakemanagerTransactorSession) Stake(validator common.Address) (*types.Transaction, error) {
	return _Stakemanager.Contract.Stake(&_Stakemanager.TransactOpts, validator)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address validator, uint256 amount) returns()
func (_Stakemanager *StakemanagerTransactor) Unstake(opts *bind.TransactOpts, validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "unstake", validator, amount)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address validator, uint256 amount) returns()
func (_Stakemanager *StakemanagerSession) Unstake(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.Unstake(&_Stakemanager.TransactOpts, validator, amount)
}

// Unstake is a paid mutator transaction binding the contract method 0xc2a672e0.
//
// Solidity: function unstake(address validator, uint256 amount) returns()
func (_Stakemanager *StakemanagerTransactorSession) Unstake(validator common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.Unstake(&_Stakemanager.TransactOpts, validator, amount)
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

// UpdateValidatorBlocks is a paid mutator transaction binding the contract method 0x88d0eb7d.
//
// Solidity: function updateValidatorBlocks(address[] operators, uint256[] counts) returns()
func (_Stakemanager *StakemanagerTransactor) UpdateValidatorBlocks(opts *bind.TransactOpts, operators []common.Address, counts []*big.Int) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "updateValidatorBlocks", operators, counts)
}

// UpdateValidatorBlocks is a paid mutator transaction binding the contract method 0x88d0eb7d.
//
// Solidity: function updateValidatorBlocks(address[] operators, uint256[] counts) returns()
func (_Stakemanager *StakemanagerSession) UpdateValidatorBlocks(operators []common.Address, counts []*big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.UpdateValidatorBlocks(&_Stakemanager.TransactOpts, operators, counts)
}

// UpdateValidatorBlocks is a paid mutator transaction binding the contract method 0x88d0eb7d.
//
// Solidity: function updateValidatorBlocks(address[] operators, uint256[] counts) returns()
func (_Stakemanager *StakemanagerTransactorSession) UpdateValidatorBlocks(operators []common.Address, counts []*big.Int) (*types.Transaction, error) {
	return _Stakemanager.Contract.UpdateValidatorBlocks(&_Stakemanager.TransactOpts, operators, counts)
}

// UpdateValidators is a paid mutator transaction binding the contract method 0x2d497ba2.
//
// Solidity: function updateValidators() returns()
func (_Stakemanager *StakemanagerTransactor) UpdateValidators(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakemanager.contract.Transact(opts, "updateValidators")
}

// UpdateValidators is a paid mutator transaction binding the contract method 0x2d497ba2.
//
// Solidity: function updateValidators() returns()
func (_Stakemanager *StakemanagerSession) UpdateValidators() (*types.Transaction, error) {
	return _Stakemanager.Contract.UpdateValidators(&_Stakemanager.TransactOpts)
}

// UpdateValidators is a paid mutator transaction binding the contract method 0x2d497ba2.
//
// Solidity: function updateValidators() returns()
func (_Stakemanager *StakemanagerTransactorSession) UpdateValidators() (*types.Transaction, error) {
	return _Stakemanager.Contract.UpdateValidators(&_Stakemanager.TransactOpts)
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
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStaked is a free log retrieval operation binding the contract event 0x5dac0c1b1112564a045ba943c9d50270893e8e826c49be8e7073adc713ab7bd7.
//
// Solidity: event Staked(address indexed staker, address indexed validator, uint256 amount)
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

// WatchStaked is a free log subscription operation binding the contract event 0x5dac0c1b1112564a045ba943c9d50270893e8e826c49be8e7073adc713ab7bd7.
//
// Solidity: event Staked(address indexed staker, address indexed validator, uint256 amount)
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

// ParseStaked is a log parse operation binding the contract event 0x5dac0c1b1112564a045ba943c9d50270893e8e826c49be8e7073adc713ab7bd7.
//
// Solidity: event Staked(address indexed staker, address indexed validator, uint256 amount)
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
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUnstaked is a free log retrieval operation binding the contract event 0xd8654fcc8cf5b36d30b3f5e4688fc78118e6d68de60b9994e09902268b57c3e3.
//
// Solidity: event Unstaked(address indexed staker, address indexed validator, uint256 amount)
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

// WatchUnstaked is a free log subscription operation binding the contract event 0xd8654fcc8cf5b36d30b3f5e4688fc78118e6d68de60b9994e09902268b57c3e3.
//
// Solidity: event Unstaked(address indexed staker, address indexed validator, uint256 amount)
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

// ParseUnstaked is a log parse operation binding the contract event 0xd8654fcc8cf5b36d30b3f5e4688fc78118e6d68de60b9994e09902268b57c3e3.
//
// Solidity: event Unstaked(address indexed staker, address indexed validator, uint256 amount)
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
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorActivated is a free log retrieval operation binding the contract event 0xb2b8c5f713f27009f7a78ec2167999c472e5ee8bc591102da17bc4a6551236c8.
//
// Solidity: event ValidatorActivated(address indexed validator)
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

// WatchValidatorActivated is a free log subscription operation binding the contract event 0xb2b8c5f713f27009f7a78ec2167999c472e5ee8bc591102da17bc4a6551236c8.
//
// Solidity: event ValidatorActivated(address indexed validator)
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

// ParseValidatorActivated is a log parse operation binding the contract event 0xb2b8c5f713f27009f7a78ec2167999c472e5ee8bc591102da17bc4a6551236c8.
//
// Solidity: event ValidatorActivated(address indexed validator)
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
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorDeactivated is a free log retrieval operation binding the contract event 0x23d934bfe7f1275bc6fd70432159c9cc1c0075d069f89da6a40f43bfe7a94ed3.
//
// Solidity: event ValidatorDeactivated(address indexed validator)
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

// WatchValidatorDeactivated is a free log subscription operation binding the contract event 0x23d934bfe7f1275bc6fd70432159c9cc1c0075d069f89da6a40f43bfe7a94ed3.
//
// Solidity: event ValidatorDeactivated(address indexed validator)
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

// ParseValidatorDeactivated is a log parse operation binding the contract event 0x23d934bfe7f1275bc6fd70432159c9cc1c0075d069f89da6a40f43bfe7a94ed3.
//
// Solidity: event ValidatorDeactivated(address indexed validator)
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
	Epoch     *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorJailed is a free log retrieval operation binding the contract event 0xeb7d7a49847ec491969db21a0e31b234565a9923145a2d1b56a75c9e95825802.
//
// Solidity: event ValidatorJailed(address indexed validator, uint256 epoch)
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
// Solidity: event ValidatorJailed(address indexed validator, uint256 epoch)
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
// Solidity: event ValidatorJailed(address indexed validator, uint256 epoch)
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
