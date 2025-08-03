// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Version

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
	_ = abi.ConvertType
)

// VersionMetaData contains all meta data concerning the Version contract.
var VersionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"version_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"getVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"version_\",\"type\":\"string\"}],\"name\":\"setVersion\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b50604051610b9a380380610b9a83398181016040528101906100319190610193565b805f908161003f91906103ea565b50506104b9565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b5f601f19601f8301169050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b6100a58261005f565b810181811067ffffffffffffffff821117156100c4576100c361006f565b5b80604052505050565b5f6100d6610046565b90506100e2828261009c565b919050565b5f67ffffffffffffffff8211156101015761010061006f565b5b61010a8261005f565b9050602081019050919050565b8281835e5f83830152505050565b5f610137610132846100e7565b6100cd565b9050828152602081018484840111156101535761015261005b565b5b61015e848285610117565b509392505050565b5f82601f83011261017a57610179610057565b5b815161018a848260208601610125565b91505092915050565b5f602082840312156101a8576101a761004f565b5b5f82015167ffffffffffffffff8111156101c5576101c4610053565b5b6101d184828501610166565b91505092915050565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061022857607f821691505b60208210810361023b5761023a6101e4565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f6008830261029d7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610262565b6102a78683610262565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6102eb6102e66102e1846102bf565b6102c8565b6102bf565b9050919050565b5f819050919050565b610304836102d1565b610318610310826102f2565b84845461026e565b825550505050565b5f5f905090565b61032f610320565b61033a8184846102fb565b505050565b5b8181101561035d576103525f82610327565b600181019050610340565b5050565b601f8211156103a25761037381610241565b61037c84610253565b8101602085101561038b578190505b61039f61039785610253565b83018261033f565b50505b505050565b5f82821c905092915050565b5f6103c25f19846008026103a7565b1980831691505092915050565b5f6103da83836103b3565b9150826002028217905092915050565b6103f3826101da565b67ffffffffffffffff81111561040c5761040b61006f565b5b6104168254610211565b610421828285610361565b5f60209050601f831160018114610452575f8415610440578287015190505b61044a85826103cf565b8655506104b1565b601f19841661046086610241565b5f5b8281101561048757848901518255600182019150602085019450602081019050610462565b868310156104a457848901516104a0601f8916826103b3565b8355505b6001600288020188555050505b505050505050565b6106d4806104c65f395ff3fe608060405234801561000f575f5ffd5b5060043610610034575f3560e01c80630d8e6e2c14610038578063788bc78c14610056575b5f5ffd5b610040610086565b60405161004d9190610225565b60405180910390f35b610070600480360381019061006b9190610382565b610115565b60405161007d9190610225565b60405180910390f35b60605f8054610094906103f6565b80601f01602080910402602001604051908101604052809291908181526020018280546100c0906103f6565b801561010b5780601f106100e25761010080835404028352916020019161010b565b820191905f5260205f20905b8154815290600101906020018083116100ee57829003601f168201915b5050505050905090565b6060815f908161012591906105cf565b505f8054610132906103f6565b80601f016020809104026020016040519081016040528092919081815260200182805461015e906103f6565b80156101a95780601f10610180576101008083540402835291602001916101a9565b820191905f5260205f20905b81548152906001019060200180831161018c57829003601f168201915b50505050509050919050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f6101f7826101b5565b61020181856101bf565b93506102118185602086016101cf565b61021a816101dd565b840191505092915050565b5f6020820190508181035f83015261023d81846101ed565b905092915050565b5f604051905090565b5f5ffd5b5f5ffd5b5f5ffd5b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b610294826101dd565b810181811067ffffffffffffffff821117156102b3576102b261025e565b5b80604052505050565b5f6102c5610245565b90506102d1828261028b565b919050565b5f67ffffffffffffffff8211156102f0576102ef61025e565b5b6102f9826101dd565b9050602081019050919050565b828183375f83830152505050565b5f610326610321846102d6565b6102bc565b9050828152602081018484840111156103425761034161025a565b5b61034d848285610306565b509392505050565b5f82601f83011261036957610368610256565b5b8135610379848260208601610314565b91505092915050565b5f602082840312156103975761039661024e565b5b5f82013567ffffffffffffffff8111156103b4576103b3610252565b5b6103c084828501610355565b91505092915050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061040d57607f821691505b6020821081036104205761041f6103c9565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026104827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610447565b61048c8683610447565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f6104d06104cb6104c6846104a4565b6104ad565b6104a4565b9050919050565b5f819050919050565b6104e9836104b6565b6104fd6104f5826104d7565b848454610453565b825550505050565b5f5f905090565b610514610505565b61051f8184846104e0565b505050565b5b81811015610542576105375f8261050c565b600181019050610525565b5050565b601f8211156105875761055881610426565b61056184610438565b81016020851015610570578190505b61058461057c85610438565b830182610524565b50505b505050565b5f82821c905092915050565b5f6105a75f198460080261058c565b1980831691505092915050565b5f6105bf8383610598565b9150826002028217905092915050565b6105d8826101b5565b67ffffffffffffffff8111156105f1576105f061025e565b5b6105fb82546103f6565b610606828285610546565b5f60209050601f831160018114610637575f8415610625578287015190505b61062f85826105b4565b865550610696565b601f19841661064586610426565b5f5b8281101561066c57848901518255600182019150602085019450602081019050610647565b868310156106895784890151610685601f891682610598565b8355505b6001600288020188555050505b50505050505056fea2646970667358221220c03a179eba7dd11d1636661af92de78b00246815e4f2cfa0c6a83d786625949064736f6c634300081e0033",
}

// VersionABI is the input ABI used to generate the binding from.
// Deprecated: Use VersionMetaData.ABI instead.
var VersionABI = VersionMetaData.ABI

// VersionBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VersionMetaData.Bin instead.
var VersionBin = VersionMetaData.Bin

// DeployVersion deploys a new Ethereum contract, binding an instance of Version to it.
func DeployVersion(auth *bind.TransactOpts, backend bind.ContractBackend, version_ string) (common.Address, *types.Transaction, *Version, error) {
	parsed, err := VersionMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VersionBin), backend, version_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Version{VersionCaller: VersionCaller{contract: contract}, VersionTransactor: VersionTransactor{contract: contract}, VersionFilterer: VersionFilterer{contract: contract}}, nil
}

// Version is an auto generated Go binding around an Ethereum contract.
type Version struct {
	VersionCaller     // Read-only binding to the contract
	VersionTransactor // Write-only binding to the contract
	VersionFilterer   // Log filterer for contract events
}

// VersionCaller is an auto generated read-only Go binding around an Ethereum contract.
type VersionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VersionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VersionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VersionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VersionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VersionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VersionSession struct {
	Contract     *Version          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VersionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VersionCallerSession struct {
	Contract *VersionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// VersionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VersionTransactorSession struct {
	Contract     *VersionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// VersionRaw is an auto generated low-level Go binding around an Ethereum contract.
type VersionRaw struct {
	Contract *Version // Generic contract binding to access the raw methods on
}

// VersionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VersionCallerRaw struct {
	Contract *VersionCaller // Generic read-only contract binding to access the raw methods on
}

// VersionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VersionTransactorRaw struct {
	Contract *VersionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVersion creates a new instance of Version, bound to a specific deployed contract.
func NewVersion(address common.Address, backend bind.ContractBackend) (*Version, error) {
	contract, err := bindVersion(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Version{VersionCaller: VersionCaller{contract: contract}, VersionTransactor: VersionTransactor{contract: contract}, VersionFilterer: VersionFilterer{contract: contract}}, nil
}

// NewVersionCaller creates a new read-only instance of Version, bound to a specific deployed contract.
func NewVersionCaller(address common.Address, caller bind.ContractCaller) (*VersionCaller, error) {
	contract, err := bindVersion(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VersionCaller{contract: contract}, nil
}

// NewVersionTransactor creates a new write-only instance of Version, bound to a specific deployed contract.
func NewVersionTransactor(address common.Address, transactor bind.ContractTransactor) (*VersionTransactor, error) {
	contract, err := bindVersion(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VersionTransactor{contract: contract}, nil
}

// NewVersionFilterer creates a new log filterer instance of Version, bound to a specific deployed contract.
func NewVersionFilterer(address common.Address, filterer bind.ContractFilterer) (*VersionFilterer, error) {
	contract, err := bindVersion(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VersionFilterer{contract: contract}, nil
}

// bindVersion binds a generic wrapper to an already deployed contract.
func bindVersion(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VersionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Version *VersionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Version.Contract.VersionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Version *VersionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Version.Contract.VersionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Version *VersionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Version.Contract.VersionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Version *VersionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Version.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Version *VersionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Version.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Version *VersionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Version.Contract.contract.Transact(opts, method, params...)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns(string)
func (_Version *VersionCaller) GetVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Version.contract.Call(opts, &out, "getVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns(string)
func (_Version *VersionSession) GetVersion() (string, error) {
	return _Version.Contract.GetVersion(&_Version.CallOpts)
}

// GetVersion is a free data retrieval call binding the contract method 0x0d8e6e2c.
//
// Solidity: function getVersion() view returns(string)
func (_Version *VersionCallerSession) GetVersion() (string, error) {
	return _Version.Contract.GetVersion(&_Version.CallOpts)
}

// SetVersion is a paid mutator transaction binding the contract method 0x788bc78c.
//
// Solidity: function setVersion(string version_) returns(string)
func (_Version *VersionTransactor) SetVersion(opts *bind.TransactOpts, version_ string) (*types.Transaction, error) {
	return _Version.contract.Transact(opts, "setVersion", version_)
}

// SetVersion is a paid mutator transaction binding the contract method 0x788bc78c.
//
// Solidity: function setVersion(string version_) returns(string)
func (_Version *VersionSession) SetVersion(version_ string) (*types.Transaction, error) {
	return _Version.Contract.SetVersion(&_Version.TransactOpts, version_)
}

// SetVersion is a paid mutator transaction binding the contract method 0x788bc78c.
//
// Solidity: function setVersion(string version_) returns(string)
func (_Version *VersionTransactorSession) SetVersion(version_ string) (*types.Transaction, error) {
	return _Version.Contract.SetVersion(&_Version.TransactOpts, version_)
}
