// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aavev3

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

// AaveV3AddressesProviderMetaData contains all meta data concerning the AaveV3AddressesProvider contract.
var AaveV3AddressesProviderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"marketId\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"ACLAdminUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"ACLManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"AddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proxyAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldImplementationAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newImplementationAddress\",\"type\":\"address\"}],\"name\":\"AddressSetAsProxy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"oldMarketId\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"string\",\"name\":\"newMarketId\",\"type\":\"string\"}],\"name\":\"MarketIdSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PoolConfiguratorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PoolDataProviderUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PoolUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PriceOracleSentinelUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PriceOracleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"proxyAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementationAddress\",\"type\":\"address\"}],\"name\":\"ProxyCreated\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"getACLAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getACLManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMarketId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolConfigurator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolDataProvider\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriceOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriceOracleSentinel\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAclAdmin\",\"type\":\"address\"}],\"name\":\"setACLAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAclManager\",\"type\":\"address\"}],\"name\":\"setACLManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"newImplementationAddress\",\"type\":\"address\"}],\"name\":\"setAddressAsProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newMarketId\",\"type\":\"string\"}],\"name\":\"setMarketId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPoolConfiguratorImpl\",\"type\":\"address\"}],\"name\":\"setPoolConfiguratorImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newDataProvider\",\"type\":\"address\"}],\"name\":\"setPoolDataProvider\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPoolImpl\",\"type\":\"address\"}],\"name\":\"setPoolImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPriceOracle\",\"type\":\"address\"}],\"name\":\"setPriceOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newPriceOracleSentinel\",\"type\":\"address\"}],\"name\":\"setPriceOracleSentinel\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// AaveV3AddressesProviderABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveV3AddressesProviderMetaData.ABI instead.
var AaveV3AddressesProviderABI = AaveV3AddressesProviderMetaData.ABI

// AaveV3AddressesProvider is an auto generated Go binding around an Ethereum contract.
type AaveV3AddressesProvider struct {
	AaveV3AddressesProviderCaller     // Read-only binding to the contract
	AaveV3AddressesProviderTransactor // Write-only binding to the contract
	AaveV3AddressesProviderFilterer   // Log filterer for contract events
}

// AaveV3AddressesProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type AaveV3AddressesProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveV3AddressesProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveV3AddressesProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveV3AddressesProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveV3AddressesProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveV3AddressesProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveV3AddressesProviderSession struct {
	Contract     *AaveV3AddressesProvider // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AaveV3AddressesProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveV3AddressesProviderCallerSession struct {
	Contract *AaveV3AddressesProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// AaveV3AddressesProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveV3AddressesProviderTransactorSession struct {
	Contract     *AaveV3AddressesProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// AaveV3AddressesProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type AaveV3AddressesProviderRaw struct {
	Contract *AaveV3AddressesProvider // Generic contract binding to access the raw methods on
}

// AaveV3AddressesProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveV3AddressesProviderCallerRaw struct {
	Contract *AaveV3AddressesProviderCaller // Generic read-only contract binding to access the raw methods on
}

// AaveV3AddressesProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveV3AddressesProviderTransactorRaw struct {
	Contract *AaveV3AddressesProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveV3AddressesProvider creates a new instance of AaveV3AddressesProvider, bound to a specific deployed contract.
func NewAaveV3AddressesProvider(address common.Address, backend bind.ContractBackend) (*AaveV3AddressesProvider, error) {
	contract, err := bindAaveV3AddressesProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveV3AddressesProvider{AaveV3AddressesProviderCaller: AaveV3AddressesProviderCaller{contract: contract}, AaveV3AddressesProviderTransactor: AaveV3AddressesProviderTransactor{contract: contract}, AaveV3AddressesProviderFilterer: AaveV3AddressesProviderFilterer{contract: contract}}, nil
}

// NewAaveV3AddressesProviderCaller creates a new read-only instance of AaveV3AddressesProvider, bound to a specific deployed contract.
func NewAaveV3AddressesProviderCaller(address common.Address, caller bind.ContractCaller) (*AaveV3AddressesProviderCaller, error) {
	contract, err := bindAaveV3AddressesProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveV3AddressesProviderCaller{contract: contract}, nil
}

// NewAaveV3AddressesProviderTransactor creates a new write-only instance of AaveV3AddressesProvider, bound to a specific deployed contract.
func NewAaveV3AddressesProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*AaveV3AddressesProviderTransactor, error) {
	contract, err := bindAaveV3AddressesProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveV3AddressesProviderTransactor{contract: contract}, nil
}

// NewAaveV3AddressesProviderFilterer creates a new log filterer instance of AaveV3AddressesProvider, bound to a specific deployed contract.
func NewAaveV3AddressesProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*AaveV3AddressesProviderFilterer, error) {
	contract, err := bindAaveV3AddressesProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveV3AddressesProviderFilterer{contract: contract}, nil
}

// bindAaveV3AddressesProvider binds a generic wrapper to an already deployed contract.
func bindAaveV3AddressesProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AaveV3AddressesProviderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveV3AddressesProvider *AaveV3AddressesProviderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveV3AddressesProvider.Contract.AaveV3AddressesProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveV3AddressesProvider *AaveV3AddressesProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.Contract.AaveV3AddressesProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveV3AddressesProvider *AaveV3AddressesProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.Contract.AaveV3AddressesProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveV3AddressesProvider *AaveV3AddressesProviderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveV3AddressesProvider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveV3AddressesProvider *AaveV3AddressesProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveV3AddressesProvider *AaveV3AddressesProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.Contract.contract.Transact(opts, method, params...)
}

// GetACLAdmin is a free data retrieval call binding the contract method 0x0e67178c.
//
// Solidity: function getACLAdmin() view returns(address)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderCaller) GetACLAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveV3AddressesProvider.contract.Call(opts, &out, "getACLAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetACLAdmin is a free data retrieval call binding the contract method 0x0e67178c.
//
// Solidity: function getACLAdmin() view returns(address)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderSession) GetACLAdmin() (common.Address, error) {
	return _AaveV3AddressesProvider.Contract.GetACLAdmin(&_AaveV3AddressesProvider.CallOpts)
}

// GetACLAdmin is a free data retrieval call binding the contract method 0x0e67178c.
//
// Solidity: function getACLAdmin() view returns(address)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderCallerSession) GetACLAdmin() (common.Address, error) {
	return _AaveV3AddressesProvider.Contract.GetACLAdmin(&_AaveV3AddressesProvider.CallOpts)
}

// GetACLManager is a free data retrieval call binding the contract method 0x707cd716.
//
// Solidity: function getACLManager() view returns(address)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderCaller) GetACLManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveV3AddressesProvider.contract.Call(opts, &out, "getACLManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetACLManager is a free data retrieval call binding the contract method 0x707cd716.
//
// Solidity: function getACLManager() view returns(address)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderSession) GetACLManager() (common.Address, error) {
	return _AaveV3AddressesProvider.Contract.GetACLManager(&_AaveV3AddressesProvider.CallOpts)
}

// GetACLManager is a free data retrieval call binding the contract method 0x707cd716.
//
// Solidity: function getACLManager() view returns(address)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderCallerSession) GetACLManager() (common.Address, error) {
	return _AaveV3AddressesProvider.Contract.GetACLManager(&_AaveV3AddressesProvider.CallOpts)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) view returns(address)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderCaller) GetAddress(opts *bind.CallOpts, id [32]byte) (common.Address, error) {
	var out []interface{}
	err := _AaveV3AddressesProvider.contract.Call(opts, &out, "getAddress", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) view returns(address)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderSession) GetAddress(id [32]byte) (common.Address, error) {
	return _AaveV3AddressesProvider.Contract.GetAddress(&_AaveV3AddressesProvider.CallOpts, id)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) view returns(address)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderCallerSession) GetAddress(id [32]byte) (common.Address, error) {
	return _AaveV3AddressesProvider.Contract.GetAddress(&_AaveV3AddressesProvider.CallOpts, id)
}

// GetMarketId is a free data retrieval call binding the contract method 0x568ef470.
//
// Solidity: function getMarketId() view returns(string)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderCaller) GetMarketId(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AaveV3AddressesProvider.contract.Call(opts, &out, "getMarketId")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMarketId is a free data retrieval call binding the contract method 0x568ef470.
//
// Solidity: function getMarketId() view returns(string)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderSession) GetMarketId() (string, error) {
	return _AaveV3AddressesProvider.Contract.GetMarketId(&_AaveV3AddressesProvider.CallOpts)
}

// GetMarketId is a free data retrieval call binding the contract method 0x568ef470.
//
// Solidity: function getMarketId() view returns(string)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderCallerSession) GetMarketId() (string, error) {
	return _AaveV3AddressesProvider.Contract.GetMarketId(&_AaveV3AddressesProvider.CallOpts)
}

// GetPool is a free data retrieval call binding the contract method 0x026b1d5f.
//
// Solidity: function getPool() view returns(address)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderCaller) GetPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveV3AddressesProvider.contract.Call(opts, &out, "getPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPool is a free data retrieval call binding the contract method 0x026b1d5f.
//
// Solidity: function getPool() view returns(address)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderSession) GetPool() (common.Address, error) {
	return _AaveV3AddressesProvider.Contract.GetPool(&_AaveV3AddressesProvider.CallOpts)
}

// GetPool is a free data retrieval call binding the contract method 0x026b1d5f.
//
// Solidity: function getPool() view returns(address)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderCallerSession) GetPool() (common.Address, error) {
	return _AaveV3AddressesProvider.Contract.GetPool(&_AaveV3AddressesProvider.CallOpts)
}

// GetPoolConfigurator is a free data retrieval call binding the contract method 0x631adfca.
//
// Solidity: function getPoolConfigurator() view returns(address)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderCaller) GetPoolConfigurator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveV3AddressesProvider.contract.Call(opts, &out, "getPoolConfigurator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolConfigurator is a free data retrieval call binding the contract method 0x631adfca.
//
// Solidity: function getPoolConfigurator() view returns(address)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderSession) GetPoolConfigurator() (common.Address, error) {
	return _AaveV3AddressesProvider.Contract.GetPoolConfigurator(&_AaveV3AddressesProvider.CallOpts)
}

// GetPoolConfigurator is a free data retrieval call binding the contract method 0x631adfca.
//
// Solidity: function getPoolConfigurator() view returns(address)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderCallerSession) GetPoolConfigurator() (common.Address, error) {
	return _AaveV3AddressesProvider.Contract.GetPoolConfigurator(&_AaveV3AddressesProvider.CallOpts)
}

// GetPoolDataProvider is a free data retrieval call binding the contract method 0xe860accb.
//
// Solidity: function getPoolDataProvider() view returns(address)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderCaller) GetPoolDataProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveV3AddressesProvider.contract.Call(opts, &out, "getPoolDataProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolDataProvider is a free data retrieval call binding the contract method 0xe860accb.
//
// Solidity: function getPoolDataProvider() view returns(address)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderSession) GetPoolDataProvider() (common.Address, error) {
	return _AaveV3AddressesProvider.Contract.GetPoolDataProvider(&_AaveV3AddressesProvider.CallOpts)
}

// GetPoolDataProvider is a free data retrieval call binding the contract method 0xe860accb.
//
// Solidity: function getPoolDataProvider() view returns(address)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderCallerSession) GetPoolDataProvider() (common.Address, error) {
	return _AaveV3AddressesProvider.Contract.GetPoolDataProvider(&_AaveV3AddressesProvider.CallOpts)
}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderCaller) GetPriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveV3AddressesProvider.contract.Call(opts, &out, "getPriceOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderSession) GetPriceOracle() (common.Address, error) {
	return _AaveV3AddressesProvider.Contract.GetPriceOracle(&_AaveV3AddressesProvider.CallOpts)
}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderCallerSession) GetPriceOracle() (common.Address, error) {
	return _AaveV3AddressesProvider.Contract.GetPriceOracle(&_AaveV3AddressesProvider.CallOpts)
}

// GetPriceOracleSentinel is a free data retrieval call binding the contract method 0x5eb88d3d.
//
// Solidity: function getPriceOracleSentinel() view returns(address)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderCaller) GetPriceOracleSentinel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveV3AddressesProvider.contract.Call(opts, &out, "getPriceOracleSentinel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPriceOracleSentinel is a free data retrieval call binding the contract method 0x5eb88d3d.
//
// Solidity: function getPriceOracleSentinel() view returns(address)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderSession) GetPriceOracleSentinel() (common.Address, error) {
	return _AaveV3AddressesProvider.Contract.GetPriceOracleSentinel(&_AaveV3AddressesProvider.CallOpts)
}

// GetPriceOracleSentinel is a free data retrieval call binding the contract method 0x5eb88d3d.
//
// Solidity: function getPriceOracleSentinel() view returns(address)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderCallerSession) GetPriceOracleSentinel() (common.Address, error) {
	return _AaveV3AddressesProvider.Contract.GetPriceOracleSentinel(&_AaveV3AddressesProvider.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveV3AddressesProvider.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderSession) Owner() (common.Address, error) {
	return _AaveV3AddressesProvider.Contract.Owner(&_AaveV3AddressesProvider.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderCallerSession) Owner() (common.Address, error) {
	return _AaveV3AddressesProvider.Contract.Owner(&_AaveV3AddressesProvider.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderSession) RenounceOwnership() (*types.Transaction, error) {
	return _AaveV3AddressesProvider.Contract.RenounceOwnership(&_AaveV3AddressesProvider.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AaveV3AddressesProvider.Contract.RenounceOwnership(&_AaveV3AddressesProvider.TransactOpts)
}

// SetACLAdmin is a paid mutator transaction binding the contract method 0x76d84ffc.
//
// Solidity: function setACLAdmin(address newAclAdmin) returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderTransactor) SetACLAdmin(opts *bind.TransactOpts, newAclAdmin common.Address) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.contract.Transact(opts, "setACLAdmin", newAclAdmin)
}

// SetACLAdmin is a paid mutator transaction binding the contract method 0x76d84ffc.
//
// Solidity: function setACLAdmin(address newAclAdmin) returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderSession) SetACLAdmin(newAclAdmin common.Address) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.Contract.SetACLAdmin(&_AaveV3AddressesProvider.TransactOpts, newAclAdmin)
}

// SetACLAdmin is a paid mutator transaction binding the contract method 0x76d84ffc.
//
// Solidity: function setACLAdmin(address newAclAdmin) returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderTransactorSession) SetACLAdmin(newAclAdmin common.Address) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.Contract.SetACLAdmin(&_AaveV3AddressesProvider.TransactOpts, newAclAdmin)
}

// SetACLManager is a paid mutator transaction binding the contract method 0xed301ca9.
//
// Solidity: function setACLManager(address newAclManager) returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderTransactor) SetACLManager(opts *bind.TransactOpts, newAclManager common.Address) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.contract.Transact(opts, "setACLManager", newAclManager)
}

// SetACLManager is a paid mutator transaction binding the contract method 0xed301ca9.
//
// Solidity: function setACLManager(address newAclManager) returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderSession) SetACLManager(newAclManager common.Address) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.Contract.SetACLManager(&_AaveV3AddressesProvider.TransactOpts, newAclManager)
}

// SetACLManager is a paid mutator transaction binding the contract method 0xed301ca9.
//
// Solidity: function setACLManager(address newAclManager) returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderTransactorSession) SetACLManager(newAclManager common.Address) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.Contract.SetACLManager(&_AaveV3AddressesProvider.TransactOpts, newAclManager)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 id, address newAddress) returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderTransactor) SetAddress(opts *bind.TransactOpts, id [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.contract.Transact(opts, "setAddress", id, newAddress)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 id, address newAddress) returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderSession) SetAddress(id [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.Contract.SetAddress(&_AaveV3AddressesProvider.TransactOpts, id, newAddress)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 id, address newAddress) returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderTransactorSession) SetAddress(id [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.Contract.SetAddress(&_AaveV3AddressesProvider.TransactOpts, id, newAddress)
}

// SetAddressAsProxy is a paid mutator transaction binding the contract method 0x5dcc528c.
//
// Solidity: function setAddressAsProxy(bytes32 id, address newImplementationAddress) returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderTransactor) SetAddressAsProxy(opts *bind.TransactOpts, id [32]byte, newImplementationAddress common.Address) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.contract.Transact(opts, "setAddressAsProxy", id, newImplementationAddress)
}

// SetAddressAsProxy is a paid mutator transaction binding the contract method 0x5dcc528c.
//
// Solidity: function setAddressAsProxy(bytes32 id, address newImplementationAddress) returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderSession) SetAddressAsProxy(id [32]byte, newImplementationAddress common.Address) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.Contract.SetAddressAsProxy(&_AaveV3AddressesProvider.TransactOpts, id, newImplementationAddress)
}

// SetAddressAsProxy is a paid mutator transaction binding the contract method 0x5dcc528c.
//
// Solidity: function setAddressAsProxy(bytes32 id, address newImplementationAddress) returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderTransactorSession) SetAddressAsProxy(id [32]byte, newImplementationAddress common.Address) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.Contract.SetAddressAsProxy(&_AaveV3AddressesProvider.TransactOpts, id, newImplementationAddress)
}

// SetMarketId is a paid mutator transaction binding the contract method 0xf67b1847.
//
// Solidity: function setMarketId(string newMarketId) returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderTransactor) SetMarketId(opts *bind.TransactOpts, newMarketId string) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.contract.Transact(opts, "setMarketId", newMarketId)
}

// SetMarketId is a paid mutator transaction binding the contract method 0xf67b1847.
//
// Solidity: function setMarketId(string newMarketId) returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderSession) SetMarketId(newMarketId string) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.Contract.SetMarketId(&_AaveV3AddressesProvider.TransactOpts, newMarketId)
}

// SetMarketId is a paid mutator transaction binding the contract method 0xf67b1847.
//
// Solidity: function setMarketId(string newMarketId) returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderTransactorSession) SetMarketId(newMarketId string) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.Contract.SetMarketId(&_AaveV3AddressesProvider.TransactOpts, newMarketId)
}

// SetPoolConfiguratorImpl is a paid mutator transaction binding the contract method 0xe4ca28b7.
//
// Solidity: function setPoolConfiguratorImpl(address newPoolConfiguratorImpl) returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderTransactor) SetPoolConfiguratorImpl(opts *bind.TransactOpts, newPoolConfiguratorImpl common.Address) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.contract.Transact(opts, "setPoolConfiguratorImpl", newPoolConfiguratorImpl)
}

// SetPoolConfiguratorImpl is a paid mutator transaction binding the contract method 0xe4ca28b7.
//
// Solidity: function setPoolConfiguratorImpl(address newPoolConfiguratorImpl) returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderSession) SetPoolConfiguratorImpl(newPoolConfiguratorImpl common.Address) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.Contract.SetPoolConfiguratorImpl(&_AaveV3AddressesProvider.TransactOpts, newPoolConfiguratorImpl)
}

// SetPoolConfiguratorImpl is a paid mutator transaction binding the contract method 0xe4ca28b7.
//
// Solidity: function setPoolConfiguratorImpl(address newPoolConfiguratorImpl) returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderTransactorSession) SetPoolConfiguratorImpl(newPoolConfiguratorImpl common.Address) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.Contract.SetPoolConfiguratorImpl(&_AaveV3AddressesProvider.TransactOpts, newPoolConfiguratorImpl)
}

// SetPoolDataProvider is a paid mutator transaction binding the contract method 0xe44e9ed1.
//
// Solidity: function setPoolDataProvider(address newDataProvider) returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderTransactor) SetPoolDataProvider(opts *bind.TransactOpts, newDataProvider common.Address) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.contract.Transact(opts, "setPoolDataProvider", newDataProvider)
}

// SetPoolDataProvider is a paid mutator transaction binding the contract method 0xe44e9ed1.
//
// Solidity: function setPoolDataProvider(address newDataProvider) returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderSession) SetPoolDataProvider(newDataProvider common.Address) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.Contract.SetPoolDataProvider(&_AaveV3AddressesProvider.TransactOpts, newDataProvider)
}

// SetPoolDataProvider is a paid mutator transaction binding the contract method 0xe44e9ed1.
//
// Solidity: function setPoolDataProvider(address newDataProvider) returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderTransactorSession) SetPoolDataProvider(newDataProvider common.Address) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.Contract.SetPoolDataProvider(&_AaveV3AddressesProvider.TransactOpts, newDataProvider)
}

// SetPoolImpl is a paid mutator transaction binding the contract method 0xa1564406.
//
// Solidity: function setPoolImpl(address newPoolImpl) returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderTransactor) SetPoolImpl(opts *bind.TransactOpts, newPoolImpl common.Address) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.contract.Transact(opts, "setPoolImpl", newPoolImpl)
}

// SetPoolImpl is a paid mutator transaction binding the contract method 0xa1564406.
//
// Solidity: function setPoolImpl(address newPoolImpl) returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderSession) SetPoolImpl(newPoolImpl common.Address) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.Contract.SetPoolImpl(&_AaveV3AddressesProvider.TransactOpts, newPoolImpl)
}

// SetPoolImpl is a paid mutator transaction binding the contract method 0xa1564406.
//
// Solidity: function setPoolImpl(address newPoolImpl) returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderTransactorSession) SetPoolImpl(newPoolImpl common.Address) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.Contract.SetPoolImpl(&_AaveV3AddressesProvider.TransactOpts, newPoolImpl)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address newPriceOracle) returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderTransactor) SetPriceOracle(opts *bind.TransactOpts, newPriceOracle common.Address) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.contract.Transact(opts, "setPriceOracle", newPriceOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address newPriceOracle) returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderSession) SetPriceOracle(newPriceOracle common.Address) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.Contract.SetPriceOracle(&_AaveV3AddressesProvider.TransactOpts, newPriceOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address newPriceOracle) returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderTransactorSession) SetPriceOracle(newPriceOracle common.Address) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.Contract.SetPriceOracle(&_AaveV3AddressesProvider.TransactOpts, newPriceOracle)
}

// SetPriceOracleSentinel is a paid mutator transaction binding the contract method 0x74944cec.
//
// Solidity: function setPriceOracleSentinel(address newPriceOracleSentinel) returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderTransactor) SetPriceOracleSentinel(opts *bind.TransactOpts, newPriceOracleSentinel common.Address) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.contract.Transact(opts, "setPriceOracleSentinel", newPriceOracleSentinel)
}

// SetPriceOracleSentinel is a paid mutator transaction binding the contract method 0x74944cec.
//
// Solidity: function setPriceOracleSentinel(address newPriceOracleSentinel) returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderSession) SetPriceOracleSentinel(newPriceOracleSentinel common.Address) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.Contract.SetPriceOracleSentinel(&_AaveV3AddressesProvider.TransactOpts, newPriceOracleSentinel)
}

// SetPriceOracleSentinel is a paid mutator transaction binding the contract method 0x74944cec.
//
// Solidity: function setPriceOracleSentinel(address newPriceOracleSentinel) returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderTransactorSession) SetPriceOracleSentinel(newPriceOracleSentinel common.Address) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.Contract.SetPriceOracleSentinel(&_AaveV3AddressesProvider.TransactOpts, newPriceOracleSentinel)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.Contract.TransferOwnership(&_AaveV3AddressesProvider.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AaveV3AddressesProvider *AaveV3AddressesProviderTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AaveV3AddressesProvider.Contract.TransferOwnership(&_AaveV3AddressesProvider.TransactOpts, newOwner)
}

// AaveV3AddressesProviderACLAdminUpdatedIterator is returned from FilterACLAdminUpdated and is used to iterate over the raw logs and unpacked data for ACLAdminUpdated events raised by the AaveV3AddressesProvider contract.
type AaveV3AddressesProviderACLAdminUpdatedIterator struct {
	Event *AaveV3AddressesProviderACLAdminUpdated // Event containing the contract specifics and raw log

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
func (it *AaveV3AddressesProviderACLAdminUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3AddressesProviderACLAdminUpdated)
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
		it.Event = new(AaveV3AddressesProviderACLAdminUpdated)
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
func (it *AaveV3AddressesProviderACLAdminUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3AddressesProviderACLAdminUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3AddressesProviderACLAdminUpdated represents a ACLAdminUpdated event raised by the AaveV3AddressesProvider contract.
type AaveV3AddressesProviderACLAdminUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterACLAdminUpdated is a free log retrieval operation binding the contract event 0xe9cf53972264dc95304fd424458745019ddfca0e37ae8f703d74772c41ad115b.
//
// Solidity: event ACLAdminUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) FilterACLAdminUpdated(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*AaveV3AddressesProviderACLAdminUpdatedIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveV3AddressesProvider.contract.FilterLogs(opts, "ACLAdminUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3AddressesProviderACLAdminUpdatedIterator{contract: _AaveV3AddressesProvider.contract, event: "ACLAdminUpdated", logs: logs, sub: sub}, nil
}

// WatchACLAdminUpdated is a free log subscription operation binding the contract event 0xe9cf53972264dc95304fd424458745019ddfca0e37ae8f703d74772c41ad115b.
//
// Solidity: event ACLAdminUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) WatchACLAdminUpdated(opts *bind.WatchOpts, sink chan<- *AaveV3AddressesProviderACLAdminUpdated, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveV3AddressesProvider.contract.WatchLogs(opts, "ACLAdminUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3AddressesProviderACLAdminUpdated)
				if err := _AaveV3AddressesProvider.contract.UnpackLog(event, "ACLAdminUpdated", log); err != nil {
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

// ParseACLAdminUpdated is a log parse operation binding the contract event 0xe9cf53972264dc95304fd424458745019ddfca0e37ae8f703d74772c41ad115b.
//
// Solidity: event ACLAdminUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) ParseACLAdminUpdated(log types.Log) (*AaveV3AddressesProviderACLAdminUpdated, error) {
	event := new(AaveV3AddressesProviderACLAdminUpdated)
	if err := _AaveV3AddressesProvider.contract.UnpackLog(event, "ACLAdminUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3AddressesProviderACLManagerUpdatedIterator is returned from FilterACLManagerUpdated and is used to iterate over the raw logs and unpacked data for ACLManagerUpdated events raised by the AaveV3AddressesProvider contract.
type AaveV3AddressesProviderACLManagerUpdatedIterator struct {
	Event *AaveV3AddressesProviderACLManagerUpdated // Event containing the contract specifics and raw log

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
func (it *AaveV3AddressesProviderACLManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3AddressesProviderACLManagerUpdated)
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
		it.Event = new(AaveV3AddressesProviderACLManagerUpdated)
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
func (it *AaveV3AddressesProviderACLManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3AddressesProviderACLManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3AddressesProviderACLManagerUpdated represents a ACLManagerUpdated event raised by the AaveV3AddressesProvider contract.
type AaveV3AddressesProviderACLManagerUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterACLManagerUpdated is a free log retrieval operation binding the contract event 0xb30efa04327bb8a537d61cc1e5c48095345ad18ef7cc04e6bacf7dfb6caaf507.
//
// Solidity: event ACLManagerUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) FilterACLManagerUpdated(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*AaveV3AddressesProviderACLManagerUpdatedIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveV3AddressesProvider.contract.FilterLogs(opts, "ACLManagerUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3AddressesProviderACLManagerUpdatedIterator{contract: _AaveV3AddressesProvider.contract, event: "ACLManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchACLManagerUpdated is a free log subscription operation binding the contract event 0xb30efa04327bb8a537d61cc1e5c48095345ad18ef7cc04e6bacf7dfb6caaf507.
//
// Solidity: event ACLManagerUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) WatchACLManagerUpdated(opts *bind.WatchOpts, sink chan<- *AaveV3AddressesProviderACLManagerUpdated, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveV3AddressesProvider.contract.WatchLogs(opts, "ACLManagerUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3AddressesProviderACLManagerUpdated)
				if err := _AaveV3AddressesProvider.contract.UnpackLog(event, "ACLManagerUpdated", log); err != nil {
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

// ParseACLManagerUpdated is a log parse operation binding the contract event 0xb30efa04327bb8a537d61cc1e5c48095345ad18ef7cc04e6bacf7dfb6caaf507.
//
// Solidity: event ACLManagerUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) ParseACLManagerUpdated(log types.Log) (*AaveV3AddressesProviderACLManagerUpdated, error) {
	event := new(AaveV3AddressesProviderACLManagerUpdated)
	if err := _AaveV3AddressesProvider.contract.UnpackLog(event, "ACLManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3AddressesProviderAddressSetIterator is returned from FilterAddressSet and is used to iterate over the raw logs and unpacked data for AddressSet events raised by the AaveV3AddressesProvider contract.
type AaveV3AddressesProviderAddressSetIterator struct {
	Event *AaveV3AddressesProviderAddressSet // Event containing the contract specifics and raw log

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
func (it *AaveV3AddressesProviderAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3AddressesProviderAddressSet)
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
		it.Event = new(AaveV3AddressesProviderAddressSet)
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
func (it *AaveV3AddressesProviderAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3AddressesProviderAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3AddressesProviderAddressSet represents a AddressSet event raised by the AaveV3AddressesProvider contract.
type AaveV3AddressesProviderAddressSet struct {
	Id         [32]byte
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddressSet is a free log retrieval operation binding the contract event 0x9ef0e8c8e52743bb38b83b17d9429141d494b8041ca6d616a6c77cebae9cd8b7.
//
// Solidity: event AddressSet(bytes32 indexed id, address indexed oldAddress, address indexed newAddress)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) FilterAddressSet(opts *bind.FilterOpts, id [][32]byte, oldAddress []common.Address, newAddress []common.Address) (*AaveV3AddressesProviderAddressSetIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveV3AddressesProvider.contract.FilterLogs(opts, "AddressSet", idRule, oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3AddressesProviderAddressSetIterator{contract: _AaveV3AddressesProvider.contract, event: "AddressSet", logs: logs, sub: sub}, nil
}

// WatchAddressSet is a free log subscription operation binding the contract event 0x9ef0e8c8e52743bb38b83b17d9429141d494b8041ca6d616a6c77cebae9cd8b7.
//
// Solidity: event AddressSet(bytes32 indexed id, address indexed oldAddress, address indexed newAddress)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) WatchAddressSet(opts *bind.WatchOpts, sink chan<- *AaveV3AddressesProviderAddressSet, id [][32]byte, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveV3AddressesProvider.contract.WatchLogs(opts, "AddressSet", idRule, oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3AddressesProviderAddressSet)
				if err := _AaveV3AddressesProvider.contract.UnpackLog(event, "AddressSet", log); err != nil {
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

// ParseAddressSet is a log parse operation binding the contract event 0x9ef0e8c8e52743bb38b83b17d9429141d494b8041ca6d616a6c77cebae9cd8b7.
//
// Solidity: event AddressSet(bytes32 indexed id, address indexed oldAddress, address indexed newAddress)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) ParseAddressSet(log types.Log) (*AaveV3AddressesProviderAddressSet, error) {
	event := new(AaveV3AddressesProviderAddressSet)
	if err := _AaveV3AddressesProvider.contract.UnpackLog(event, "AddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3AddressesProviderAddressSetAsProxyIterator is returned from FilterAddressSetAsProxy and is used to iterate over the raw logs and unpacked data for AddressSetAsProxy events raised by the AaveV3AddressesProvider contract.
type AaveV3AddressesProviderAddressSetAsProxyIterator struct {
	Event *AaveV3AddressesProviderAddressSetAsProxy // Event containing the contract specifics and raw log

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
func (it *AaveV3AddressesProviderAddressSetAsProxyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3AddressesProviderAddressSetAsProxy)
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
		it.Event = new(AaveV3AddressesProviderAddressSetAsProxy)
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
func (it *AaveV3AddressesProviderAddressSetAsProxyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3AddressesProviderAddressSetAsProxyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3AddressesProviderAddressSetAsProxy represents a AddressSetAsProxy event raised by the AaveV3AddressesProvider contract.
type AaveV3AddressesProviderAddressSetAsProxy struct {
	Id                       [32]byte
	ProxyAddress             common.Address
	OldImplementationAddress common.Address
	NewImplementationAddress common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterAddressSetAsProxy is a free log retrieval operation binding the contract event 0x3bbd45b5429b385e3fb37ad5cd1cd1435a3c8ec32196c7937597365a3fd3e99c.
//
// Solidity: event AddressSetAsProxy(bytes32 indexed id, address indexed proxyAddress, address oldImplementationAddress, address indexed newImplementationAddress)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) FilterAddressSetAsProxy(opts *bind.FilterOpts, id [][32]byte, proxyAddress []common.Address, newImplementationAddress []common.Address) (*AaveV3AddressesProviderAddressSetAsProxyIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var proxyAddressRule []interface{}
	for _, proxyAddressItem := range proxyAddress {
		proxyAddressRule = append(proxyAddressRule, proxyAddressItem)
	}

	var newImplementationAddressRule []interface{}
	for _, newImplementationAddressItem := range newImplementationAddress {
		newImplementationAddressRule = append(newImplementationAddressRule, newImplementationAddressItem)
	}

	logs, sub, err := _AaveV3AddressesProvider.contract.FilterLogs(opts, "AddressSetAsProxy", idRule, proxyAddressRule, newImplementationAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3AddressesProviderAddressSetAsProxyIterator{contract: _AaveV3AddressesProvider.contract, event: "AddressSetAsProxy", logs: logs, sub: sub}, nil
}

// WatchAddressSetAsProxy is a free log subscription operation binding the contract event 0x3bbd45b5429b385e3fb37ad5cd1cd1435a3c8ec32196c7937597365a3fd3e99c.
//
// Solidity: event AddressSetAsProxy(bytes32 indexed id, address indexed proxyAddress, address oldImplementationAddress, address indexed newImplementationAddress)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) WatchAddressSetAsProxy(opts *bind.WatchOpts, sink chan<- *AaveV3AddressesProviderAddressSetAsProxy, id [][32]byte, proxyAddress []common.Address, newImplementationAddress []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var proxyAddressRule []interface{}
	for _, proxyAddressItem := range proxyAddress {
		proxyAddressRule = append(proxyAddressRule, proxyAddressItem)
	}

	var newImplementationAddressRule []interface{}
	for _, newImplementationAddressItem := range newImplementationAddress {
		newImplementationAddressRule = append(newImplementationAddressRule, newImplementationAddressItem)
	}

	logs, sub, err := _AaveV3AddressesProvider.contract.WatchLogs(opts, "AddressSetAsProxy", idRule, proxyAddressRule, newImplementationAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3AddressesProviderAddressSetAsProxy)
				if err := _AaveV3AddressesProvider.contract.UnpackLog(event, "AddressSetAsProxy", log); err != nil {
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

// ParseAddressSetAsProxy is a log parse operation binding the contract event 0x3bbd45b5429b385e3fb37ad5cd1cd1435a3c8ec32196c7937597365a3fd3e99c.
//
// Solidity: event AddressSetAsProxy(bytes32 indexed id, address indexed proxyAddress, address oldImplementationAddress, address indexed newImplementationAddress)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) ParseAddressSetAsProxy(log types.Log) (*AaveV3AddressesProviderAddressSetAsProxy, error) {
	event := new(AaveV3AddressesProviderAddressSetAsProxy)
	if err := _AaveV3AddressesProvider.contract.UnpackLog(event, "AddressSetAsProxy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3AddressesProviderMarketIdSetIterator is returned from FilterMarketIdSet and is used to iterate over the raw logs and unpacked data for MarketIdSet events raised by the AaveV3AddressesProvider contract.
type AaveV3AddressesProviderMarketIdSetIterator struct {
	Event *AaveV3AddressesProviderMarketIdSet // Event containing the contract specifics and raw log

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
func (it *AaveV3AddressesProviderMarketIdSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3AddressesProviderMarketIdSet)
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
		it.Event = new(AaveV3AddressesProviderMarketIdSet)
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
func (it *AaveV3AddressesProviderMarketIdSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3AddressesProviderMarketIdSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3AddressesProviderMarketIdSet represents a MarketIdSet event raised by the AaveV3AddressesProvider contract.
type AaveV3AddressesProviderMarketIdSet struct {
	OldMarketId common.Hash
	NewMarketId common.Hash
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMarketIdSet is a free log retrieval operation binding the contract event 0xe685c8cdecc6030c45030fd54778812cb84ed8e4467c38294403d68ba7860823.
//
// Solidity: event MarketIdSet(string indexed oldMarketId, string indexed newMarketId)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) FilterMarketIdSet(opts *bind.FilterOpts, oldMarketId []string, newMarketId []string) (*AaveV3AddressesProviderMarketIdSetIterator, error) {

	var oldMarketIdRule []interface{}
	for _, oldMarketIdItem := range oldMarketId {
		oldMarketIdRule = append(oldMarketIdRule, oldMarketIdItem)
	}
	var newMarketIdRule []interface{}
	for _, newMarketIdItem := range newMarketId {
		newMarketIdRule = append(newMarketIdRule, newMarketIdItem)
	}

	logs, sub, err := _AaveV3AddressesProvider.contract.FilterLogs(opts, "MarketIdSet", oldMarketIdRule, newMarketIdRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3AddressesProviderMarketIdSetIterator{contract: _AaveV3AddressesProvider.contract, event: "MarketIdSet", logs: logs, sub: sub}, nil
}

// WatchMarketIdSet is a free log subscription operation binding the contract event 0xe685c8cdecc6030c45030fd54778812cb84ed8e4467c38294403d68ba7860823.
//
// Solidity: event MarketIdSet(string indexed oldMarketId, string indexed newMarketId)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) WatchMarketIdSet(opts *bind.WatchOpts, sink chan<- *AaveV3AddressesProviderMarketIdSet, oldMarketId []string, newMarketId []string) (event.Subscription, error) {

	var oldMarketIdRule []interface{}
	for _, oldMarketIdItem := range oldMarketId {
		oldMarketIdRule = append(oldMarketIdRule, oldMarketIdItem)
	}
	var newMarketIdRule []interface{}
	for _, newMarketIdItem := range newMarketId {
		newMarketIdRule = append(newMarketIdRule, newMarketIdItem)
	}

	logs, sub, err := _AaveV3AddressesProvider.contract.WatchLogs(opts, "MarketIdSet", oldMarketIdRule, newMarketIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3AddressesProviderMarketIdSet)
				if err := _AaveV3AddressesProvider.contract.UnpackLog(event, "MarketIdSet", log); err != nil {
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

// ParseMarketIdSet is a log parse operation binding the contract event 0xe685c8cdecc6030c45030fd54778812cb84ed8e4467c38294403d68ba7860823.
//
// Solidity: event MarketIdSet(string indexed oldMarketId, string indexed newMarketId)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) ParseMarketIdSet(log types.Log) (*AaveV3AddressesProviderMarketIdSet, error) {
	event := new(AaveV3AddressesProviderMarketIdSet)
	if err := _AaveV3AddressesProvider.contract.UnpackLog(event, "MarketIdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3AddressesProviderOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AaveV3AddressesProvider contract.
type AaveV3AddressesProviderOwnershipTransferredIterator struct {
	Event *AaveV3AddressesProviderOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AaveV3AddressesProviderOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3AddressesProviderOwnershipTransferred)
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
		it.Event = new(AaveV3AddressesProviderOwnershipTransferred)
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
func (it *AaveV3AddressesProviderOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3AddressesProviderOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3AddressesProviderOwnershipTransferred represents a OwnershipTransferred event raised by the AaveV3AddressesProvider contract.
type AaveV3AddressesProviderOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AaveV3AddressesProviderOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AaveV3AddressesProvider.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3AddressesProviderOwnershipTransferredIterator{contract: _AaveV3AddressesProvider.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AaveV3AddressesProviderOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AaveV3AddressesProvider.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3AddressesProviderOwnershipTransferred)
				if err := _AaveV3AddressesProvider.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) ParseOwnershipTransferred(log types.Log) (*AaveV3AddressesProviderOwnershipTransferred, error) {
	event := new(AaveV3AddressesProviderOwnershipTransferred)
	if err := _AaveV3AddressesProvider.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3AddressesProviderPoolConfiguratorUpdatedIterator is returned from FilterPoolConfiguratorUpdated and is used to iterate over the raw logs and unpacked data for PoolConfiguratorUpdated events raised by the AaveV3AddressesProvider contract.
type AaveV3AddressesProviderPoolConfiguratorUpdatedIterator struct {
	Event *AaveV3AddressesProviderPoolConfiguratorUpdated // Event containing the contract specifics and raw log

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
func (it *AaveV3AddressesProviderPoolConfiguratorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3AddressesProviderPoolConfiguratorUpdated)
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
		it.Event = new(AaveV3AddressesProviderPoolConfiguratorUpdated)
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
func (it *AaveV3AddressesProviderPoolConfiguratorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3AddressesProviderPoolConfiguratorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3AddressesProviderPoolConfiguratorUpdated represents a PoolConfiguratorUpdated event raised by the AaveV3AddressesProvider contract.
type AaveV3AddressesProviderPoolConfiguratorUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPoolConfiguratorUpdated is a free log retrieval operation binding the contract event 0x8932892569eba59c8382a089d9b732d1f49272878775235761a2a6b0309cd465.
//
// Solidity: event PoolConfiguratorUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) FilterPoolConfiguratorUpdated(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*AaveV3AddressesProviderPoolConfiguratorUpdatedIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveV3AddressesProvider.contract.FilterLogs(opts, "PoolConfiguratorUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3AddressesProviderPoolConfiguratorUpdatedIterator{contract: _AaveV3AddressesProvider.contract, event: "PoolConfiguratorUpdated", logs: logs, sub: sub}, nil
}

// WatchPoolConfiguratorUpdated is a free log subscription operation binding the contract event 0x8932892569eba59c8382a089d9b732d1f49272878775235761a2a6b0309cd465.
//
// Solidity: event PoolConfiguratorUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) WatchPoolConfiguratorUpdated(opts *bind.WatchOpts, sink chan<- *AaveV3AddressesProviderPoolConfiguratorUpdated, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveV3AddressesProvider.contract.WatchLogs(opts, "PoolConfiguratorUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3AddressesProviderPoolConfiguratorUpdated)
				if err := _AaveV3AddressesProvider.contract.UnpackLog(event, "PoolConfiguratorUpdated", log); err != nil {
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

// ParsePoolConfiguratorUpdated is a log parse operation binding the contract event 0x8932892569eba59c8382a089d9b732d1f49272878775235761a2a6b0309cd465.
//
// Solidity: event PoolConfiguratorUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) ParsePoolConfiguratorUpdated(log types.Log) (*AaveV3AddressesProviderPoolConfiguratorUpdated, error) {
	event := new(AaveV3AddressesProviderPoolConfiguratorUpdated)
	if err := _AaveV3AddressesProvider.contract.UnpackLog(event, "PoolConfiguratorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3AddressesProviderPoolDataProviderUpdatedIterator is returned from FilterPoolDataProviderUpdated and is used to iterate over the raw logs and unpacked data for PoolDataProviderUpdated events raised by the AaveV3AddressesProvider contract.
type AaveV3AddressesProviderPoolDataProviderUpdatedIterator struct {
	Event *AaveV3AddressesProviderPoolDataProviderUpdated // Event containing the contract specifics and raw log

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
func (it *AaveV3AddressesProviderPoolDataProviderUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3AddressesProviderPoolDataProviderUpdated)
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
		it.Event = new(AaveV3AddressesProviderPoolDataProviderUpdated)
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
func (it *AaveV3AddressesProviderPoolDataProviderUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3AddressesProviderPoolDataProviderUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3AddressesProviderPoolDataProviderUpdated represents a PoolDataProviderUpdated event raised by the AaveV3AddressesProvider contract.
type AaveV3AddressesProviderPoolDataProviderUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPoolDataProviderUpdated is a free log retrieval operation binding the contract event 0xc853974cfbf81487a14a23565917bee63f527853bcb5fa54f2ae1cdf8a38356d.
//
// Solidity: event PoolDataProviderUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) FilterPoolDataProviderUpdated(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*AaveV3AddressesProviderPoolDataProviderUpdatedIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveV3AddressesProvider.contract.FilterLogs(opts, "PoolDataProviderUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3AddressesProviderPoolDataProviderUpdatedIterator{contract: _AaveV3AddressesProvider.contract, event: "PoolDataProviderUpdated", logs: logs, sub: sub}, nil
}

// WatchPoolDataProviderUpdated is a free log subscription operation binding the contract event 0xc853974cfbf81487a14a23565917bee63f527853bcb5fa54f2ae1cdf8a38356d.
//
// Solidity: event PoolDataProviderUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) WatchPoolDataProviderUpdated(opts *bind.WatchOpts, sink chan<- *AaveV3AddressesProviderPoolDataProviderUpdated, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveV3AddressesProvider.contract.WatchLogs(opts, "PoolDataProviderUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3AddressesProviderPoolDataProviderUpdated)
				if err := _AaveV3AddressesProvider.contract.UnpackLog(event, "PoolDataProviderUpdated", log); err != nil {
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

// ParsePoolDataProviderUpdated is a log parse operation binding the contract event 0xc853974cfbf81487a14a23565917bee63f527853bcb5fa54f2ae1cdf8a38356d.
//
// Solidity: event PoolDataProviderUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) ParsePoolDataProviderUpdated(log types.Log) (*AaveV3AddressesProviderPoolDataProviderUpdated, error) {
	event := new(AaveV3AddressesProviderPoolDataProviderUpdated)
	if err := _AaveV3AddressesProvider.contract.UnpackLog(event, "PoolDataProviderUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3AddressesProviderPoolUpdatedIterator is returned from FilterPoolUpdated and is used to iterate over the raw logs and unpacked data for PoolUpdated events raised by the AaveV3AddressesProvider contract.
type AaveV3AddressesProviderPoolUpdatedIterator struct {
	Event *AaveV3AddressesProviderPoolUpdated // Event containing the contract specifics and raw log

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
func (it *AaveV3AddressesProviderPoolUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3AddressesProviderPoolUpdated)
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
		it.Event = new(AaveV3AddressesProviderPoolUpdated)
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
func (it *AaveV3AddressesProviderPoolUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3AddressesProviderPoolUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3AddressesProviderPoolUpdated represents a PoolUpdated event raised by the AaveV3AddressesProvider contract.
type AaveV3AddressesProviderPoolUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPoolUpdated is a free log retrieval operation binding the contract event 0x90affc163f1a2dfedcd36aa02ed992eeeba8100a4014f0b4cdc20ea265a66627.
//
// Solidity: event PoolUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) FilterPoolUpdated(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*AaveV3AddressesProviderPoolUpdatedIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveV3AddressesProvider.contract.FilterLogs(opts, "PoolUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3AddressesProviderPoolUpdatedIterator{contract: _AaveV3AddressesProvider.contract, event: "PoolUpdated", logs: logs, sub: sub}, nil
}

// WatchPoolUpdated is a free log subscription operation binding the contract event 0x90affc163f1a2dfedcd36aa02ed992eeeba8100a4014f0b4cdc20ea265a66627.
//
// Solidity: event PoolUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) WatchPoolUpdated(opts *bind.WatchOpts, sink chan<- *AaveV3AddressesProviderPoolUpdated, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveV3AddressesProvider.contract.WatchLogs(opts, "PoolUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3AddressesProviderPoolUpdated)
				if err := _AaveV3AddressesProvider.contract.UnpackLog(event, "PoolUpdated", log); err != nil {
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

// ParsePoolUpdated is a log parse operation binding the contract event 0x90affc163f1a2dfedcd36aa02ed992eeeba8100a4014f0b4cdc20ea265a66627.
//
// Solidity: event PoolUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) ParsePoolUpdated(log types.Log) (*AaveV3AddressesProviderPoolUpdated, error) {
	event := new(AaveV3AddressesProviderPoolUpdated)
	if err := _AaveV3AddressesProvider.contract.UnpackLog(event, "PoolUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3AddressesProviderPriceOracleSentinelUpdatedIterator is returned from FilterPriceOracleSentinelUpdated and is used to iterate over the raw logs and unpacked data for PriceOracleSentinelUpdated events raised by the AaveV3AddressesProvider contract.
type AaveV3AddressesProviderPriceOracleSentinelUpdatedIterator struct {
	Event *AaveV3AddressesProviderPriceOracleSentinelUpdated // Event containing the contract specifics and raw log

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
func (it *AaveV3AddressesProviderPriceOracleSentinelUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3AddressesProviderPriceOracleSentinelUpdated)
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
		it.Event = new(AaveV3AddressesProviderPriceOracleSentinelUpdated)
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
func (it *AaveV3AddressesProviderPriceOracleSentinelUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3AddressesProviderPriceOracleSentinelUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3AddressesProviderPriceOracleSentinelUpdated represents a PriceOracleSentinelUpdated event raised by the AaveV3AddressesProvider contract.
type AaveV3AddressesProviderPriceOracleSentinelUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPriceOracleSentinelUpdated is a free log retrieval operation binding the contract event 0x5326514eeca90494a14bedabcff812a0e683029ee85d1e23824d44fd14cd6ae7.
//
// Solidity: event PriceOracleSentinelUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) FilterPriceOracleSentinelUpdated(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*AaveV3AddressesProviderPriceOracleSentinelUpdatedIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveV3AddressesProvider.contract.FilterLogs(opts, "PriceOracleSentinelUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3AddressesProviderPriceOracleSentinelUpdatedIterator{contract: _AaveV3AddressesProvider.contract, event: "PriceOracleSentinelUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceOracleSentinelUpdated is a free log subscription operation binding the contract event 0x5326514eeca90494a14bedabcff812a0e683029ee85d1e23824d44fd14cd6ae7.
//
// Solidity: event PriceOracleSentinelUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) WatchPriceOracleSentinelUpdated(opts *bind.WatchOpts, sink chan<- *AaveV3AddressesProviderPriceOracleSentinelUpdated, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveV3AddressesProvider.contract.WatchLogs(opts, "PriceOracleSentinelUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3AddressesProviderPriceOracleSentinelUpdated)
				if err := _AaveV3AddressesProvider.contract.UnpackLog(event, "PriceOracleSentinelUpdated", log); err != nil {
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

// ParsePriceOracleSentinelUpdated is a log parse operation binding the contract event 0x5326514eeca90494a14bedabcff812a0e683029ee85d1e23824d44fd14cd6ae7.
//
// Solidity: event PriceOracleSentinelUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) ParsePriceOracleSentinelUpdated(log types.Log) (*AaveV3AddressesProviderPriceOracleSentinelUpdated, error) {
	event := new(AaveV3AddressesProviderPriceOracleSentinelUpdated)
	if err := _AaveV3AddressesProvider.contract.UnpackLog(event, "PriceOracleSentinelUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3AddressesProviderPriceOracleUpdatedIterator is returned from FilterPriceOracleUpdated and is used to iterate over the raw logs and unpacked data for PriceOracleUpdated events raised by the AaveV3AddressesProvider contract.
type AaveV3AddressesProviderPriceOracleUpdatedIterator struct {
	Event *AaveV3AddressesProviderPriceOracleUpdated // Event containing the contract specifics and raw log

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
func (it *AaveV3AddressesProviderPriceOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3AddressesProviderPriceOracleUpdated)
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
		it.Event = new(AaveV3AddressesProviderPriceOracleUpdated)
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
func (it *AaveV3AddressesProviderPriceOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3AddressesProviderPriceOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3AddressesProviderPriceOracleUpdated represents a PriceOracleUpdated event raised by the AaveV3AddressesProvider contract.
type AaveV3AddressesProviderPriceOracleUpdated struct {
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPriceOracleUpdated is a free log retrieval operation binding the contract event 0x56b5f80d8cac1479698aa7d01605fd6111e90b15fc4d2b377417f46034876cbd.
//
// Solidity: event PriceOracleUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) FilterPriceOracleUpdated(opts *bind.FilterOpts, oldAddress []common.Address, newAddress []common.Address) (*AaveV3AddressesProviderPriceOracleUpdatedIterator, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveV3AddressesProvider.contract.FilterLogs(opts, "PriceOracleUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3AddressesProviderPriceOracleUpdatedIterator{contract: _AaveV3AddressesProvider.contract, event: "PriceOracleUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceOracleUpdated is a free log subscription operation binding the contract event 0x56b5f80d8cac1479698aa7d01605fd6111e90b15fc4d2b377417f46034876cbd.
//
// Solidity: event PriceOracleUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) WatchPriceOracleUpdated(opts *bind.WatchOpts, sink chan<- *AaveV3AddressesProviderPriceOracleUpdated, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _AaveV3AddressesProvider.contract.WatchLogs(opts, "PriceOracleUpdated", oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3AddressesProviderPriceOracleUpdated)
				if err := _AaveV3AddressesProvider.contract.UnpackLog(event, "PriceOracleUpdated", log); err != nil {
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

// ParsePriceOracleUpdated is a log parse operation binding the contract event 0x56b5f80d8cac1479698aa7d01605fd6111e90b15fc4d2b377417f46034876cbd.
//
// Solidity: event PriceOracleUpdated(address indexed oldAddress, address indexed newAddress)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) ParsePriceOracleUpdated(log types.Log) (*AaveV3AddressesProviderPriceOracleUpdated, error) {
	event := new(AaveV3AddressesProviderPriceOracleUpdated)
	if err := _AaveV3AddressesProvider.contract.UnpackLog(event, "PriceOracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AaveV3AddressesProviderProxyCreatedIterator is returned from FilterProxyCreated and is used to iterate over the raw logs and unpacked data for ProxyCreated events raised by the AaveV3AddressesProvider contract.
type AaveV3AddressesProviderProxyCreatedIterator struct {
	Event *AaveV3AddressesProviderProxyCreated // Event containing the contract specifics and raw log

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
func (it *AaveV3AddressesProviderProxyCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AaveV3AddressesProviderProxyCreated)
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
		it.Event = new(AaveV3AddressesProviderProxyCreated)
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
func (it *AaveV3AddressesProviderProxyCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AaveV3AddressesProviderProxyCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AaveV3AddressesProviderProxyCreated represents a ProxyCreated event raised by the AaveV3AddressesProvider contract.
type AaveV3AddressesProviderProxyCreated struct {
	Id                    [32]byte
	ProxyAddress          common.Address
	ImplementationAddress common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterProxyCreated is a free log retrieval operation binding the contract event 0x4a465a9bd819d9662563c1e11ae958f8109e437e7f4bf1c6ef0b9a7b3f35d478.
//
// Solidity: event ProxyCreated(bytes32 indexed id, address indexed proxyAddress, address indexed implementationAddress)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) FilterProxyCreated(opts *bind.FilterOpts, id [][32]byte, proxyAddress []common.Address, implementationAddress []common.Address) (*AaveV3AddressesProviderProxyCreatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var proxyAddressRule []interface{}
	for _, proxyAddressItem := range proxyAddress {
		proxyAddressRule = append(proxyAddressRule, proxyAddressItem)
	}
	var implementationAddressRule []interface{}
	for _, implementationAddressItem := range implementationAddress {
		implementationAddressRule = append(implementationAddressRule, implementationAddressItem)
	}

	logs, sub, err := _AaveV3AddressesProvider.contract.FilterLogs(opts, "ProxyCreated", idRule, proxyAddressRule, implementationAddressRule)
	if err != nil {
		return nil, err
	}
	return &AaveV3AddressesProviderProxyCreatedIterator{contract: _AaveV3AddressesProvider.contract, event: "ProxyCreated", logs: logs, sub: sub}, nil
}

// WatchProxyCreated is a free log subscription operation binding the contract event 0x4a465a9bd819d9662563c1e11ae958f8109e437e7f4bf1c6ef0b9a7b3f35d478.
//
// Solidity: event ProxyCreated(bytes32 indexed id, address indexed proxyAddress, address indexed implementationAddress)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) WatchProxyCreated(opts *bind.WatchOpts, sink chan<- *AaveV3AddressesProviderProxyCreated, id [][32]byte, proxyAddress []common.Address, implementationAddress []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var proxyAddressRule []interface{}
	for _, proxyAddressItem := range proxyAddress {
		proxyAddressRule = append(proxyAddressRule, proxyAddressItem)
	}
	var implementationAddressRule []interface{}
	for _, implementationAddressItem := range implementationAddress {
		implementationAddressRule = append(implementationAddressRule, implementationAddressItem)
	}

	logs, sub, err := _AaveV3AddressesProvider.contract.WatchLogs(opts, "ProxyCreated", idRule, proxyAddressRule, implementationAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AaveV3AddressesProviderProxyCreated)
				if err := _AaveV3AddressesProvider.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
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

// ParseProxyCreated is a log parse operation binding the contract event 0x4a465a9bd819d9662563c1e11ae958f8109e437e7f4bf1c6ef0b9a7b3f35d478.
//
// Solidity: event ProxyCreated(bytes32 indexed id, address indexed proxyAddress, address indexed implementationAddress)
func (_AaveV3AddressesProvider *AaveV3AddressesProviderFilterer) ParseProxyCreated(log types.Log) (*AaveV3AddressesProviderProxyCreated, error) {
	event := new(AaveV3AddressesProviderProxyCreated)
	if err := _AaveV3AddressesProvider.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
