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

// AaveProtocolDataProviderTokenData is an auto generated low-level Go binding around an user-defined struct.
type AaveProtocolDataProviderTokenData struct {
	Symbol       string
	TokenAddress common.Address
}

// PoolDataProviderMetaData contains all meta data concerning the PoolDataProvider contract.
var PoolDataProviderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPoolAddressesProvider\",\"name\":\"addressesProvider\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ADDRESSES_PROVIDER\",\"outputs\":[{\"internalType\":\"contractIPoolAddressesProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getATokenTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllATokens\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"internalType\":\"structAaveProtocolDataProvider.TokenData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllReservesTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"internalType\":\"structAaveProtocolDataProvider.TokenData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getDebtCeiling\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDebtCeilingDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getInterestRateStrategyAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"irStrategyAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getLiquidationProtocolFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isPaused\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveCaps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowCap\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supplyCap\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveConfigurationData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ltv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationBonus\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveFactor\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"borrowingEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"stableBorrowRateEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isFrozen\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"unbacked\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"accruedToTreasuryScaled\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalAToken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"averageStableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableBorrowIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveEModeCategory\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveTokensAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"aTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stableDebtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"variableDebtTokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getSiloedBorrowing\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getTotalDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getUnbackedMintCap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserReserveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"currentATokenBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principalStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scaledVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"stableRateLastUpdated\",\"type\":\"uint40\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// PoolDataProviderABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolDataProviderMetaData.ABI instead.
var PoolDataProviderABI = PoolDataProviderMetaData.ABI

// PoolDataProvider is an auto generated Go binding around an Ethereum contract.
type PoolDataProvider struct {
	PoolDataProviderCaller     // Read-only binding to the contract
	PoolDataProviderTransactor // Write-only binding to the contract
	PoolDataProviderFilterer   // Log filterer for contract events
}

// PoolDataProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolDataProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolDataProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolDataProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolDataProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolDataProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolDataProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolDataProviderSession struct {
	Contract     *PoolDataProvider // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolDataProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolDataProviderCallerSession struct {
	Contract *PoolDataProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// PoolDataProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolDataProviderTransactorSession struct {
	Contract     *PoolDataProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// PoolDataProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolDataProviderRaw struct {
	Contract *PoolDataProvider // Generic contract binding to access the raw methods on
}

// PoolDataProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolDataProviderCallerRaw struct {
	Contract *PoolDataProviderCaller // Generic read-only contract binding to access the raw methods on
}

// PoolDataProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolDataProviderTransactorRaw struct {
	Contract *PoolDataProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolDataProvider creates a new instance of PoolDataProvider, bound to a specific deployed contract.
func NewPoolDataProvider(address common.Address, backend bind.ContractBackend) (*PoolDataProvider, error) {
	contract, err := bindPoolDataProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PoolDataProvider{PoolDataProviderCaller: PoolDataProviderCaller{contract: contract}, PoolDataProviderTransactor: PoolDataProviderTransactor{contract: contract}, PoolDataProviderFilterer: PoolDataProviderFilterer{contract: contract}}, nil
}

// NewPoolDataProviderCaller creates a new read-only instance of PoolDataProvider, bound to a specific deployed contract.
func NewPoolDataProviderCaller(address common.Address, caller bind.ContractCaller) (*PoolDataProviderCaller, error) {
	contract, err := bindPoolDataProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolDataProviderCaller{contract: contract}, nil
}

// NewPoolDataProviderTransactor creates a new write-only instance of PoolDataProvider, bound to a specific deployed contract.
func NewPoolDataProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolDataProviderTransactor, error) {
	contract, err := bindPoolDataProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolDataProviderTransactor{contract: contract}, nil
}

// NewPoolDataProviderFilterer creates a new log filterer instance of PoolDataProvider, bound to a specific deployed contract.
func NewPoolDataProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolDataProviderFilterer, error) {
	contract, err := bindPoolDataProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolDataProviderFilterer{contract: contract}, nil
}

// bindPoolDataProvider binds a generic wrapper to an already deployed contract.
func bindPoolDataProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoolDataProviderMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolDataProvider *PoolDataProviderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolDataProvider.Contract.PoolDataProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolDataProvider *PoolDataProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolDataProvider.Contract.PoolDataProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolDataProvider *PoolDataProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolDataProvider.Contract.PoolDataProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PoolDataProvider *PoolDataProviderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PoolDataProvider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PoolDataProvider *PoolDataProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PoolDataProvider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PoolDataProvider *PoolDataProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PoolDataProvider.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_PoolDataProvider *PoolDataProviderCaller) ADDRESSESPROVIDER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PoolDataProvider.contract.Call(opts, &out, "ADDRESSES_PROVIDER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_PoolDataProvider *PoolDataProviderSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _PoolDataProvider.Contract.ADDRESSESPROVIDER(&_PoolDataProvider.CallOpts)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_PoolDataProvider *PoolDataProviderCallerSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _PoolDataProvider.Contract.ADDRESSESPROVIDER(&_PoolDataProvider.CallOpts)
}

// GetATokenTotalSupply is a free data retrieval call binding the contract method 0x51460e25.
//
// Solidity: function getATokenTotalSupply(address asset) view returns(uint256)
func (_PoolDataProvider *PoolDataProviderCaller) GetATokenTotalSupply(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PoolDataProvider.contract.Call(opts, &out, "getATokenTotalSupply", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetATokenTotalSupply is a free data retrieval call binding the contract method 0x51460e25.
//
// Solidity: function getATokenTotalSupply(address asset) view returns(uint256)
func (_PoolDataProvider *PoolDataProviderSession) GetATokenTotalSupply(asset common.Address) (*big.Int, error) {
	return _PoolDataProvider.Contract.GetATokenTotalSupply(&_PoolDataProvider.CallOpts, asset)
}

// GetATokenTotalSupply is a free data retrieval call binding the contract method 0x51460e25.
//
// Solidity: function getATokenTotalSupply(address asset) view returns(uint256)
func (_PoolDataProvider *PoolDataProviderCallerSession) GetATokenTotalSupply(asset common.Address) (*big.Int, error) {
	return _PoolDataProvider.Contract.GetATokenTotalSupply(&_PoolDataProvider.CallOpts, asset)
}

// GetAllATokens is a free data retrieval call binding the contract method 0xf561ae41.
//
// Solidity: function getAllATokens() view returns((string,address)[])
func (_PoolDataProvider *PoolDataProviderCaller) GetAllATokens(opts *bind.CallOpts) ([]AaveProtocolDataProviderTokenData, error) {
	var out []interface{}
	err := _PoolDataProvider.contract.Call(opts, &out, "getAllATokens")

	if err != nil {
		return *new([]AaveProtocolDataProviderTokenData), err
	}

	out0 := *abi.ConvertType(out[0], new([]AaveProtocolDataProviderTokenData)).(*[]AaveProtocolDataProviderTokenData)

	return out0, err

}

// GetAllATokens is a free data retrieval call binding the contract method 0xf561ae41.
//
// Solidity: function getAllATokens() view returns((string,address)[])
func (_PoolDataProvider *PoolDataProviderSession) GetAllATokens() ([]AaveProtocolDataProviderTokenData, error) {
	return _PoolDataProvider.Contract.GetAllATokens(&_PoolDataProvider.CallOpts)
}

// GetAllATokens is a free data retrieval call binding the contract method 0xf561ae41.
//
// Solidity: function getAllATokens() view returns((string,address)[])
func (_PoolDataProvider *PoolDataProviderCallerSession) GetAllATokens() ([]AaveProtocolDataProviderTokenData, error) {
	return _PoolDataProvider.Contract.GetAllATokens(&_PoolDataProvider.CallOpts)
}

// GetAllReservesTokens is a free data retrieval call binding the contract method 0xb316ff89.
//
// Solidity: function getAllReservesTokens() view returns((string,address)[])
func (_PoolDataProvider *PoolDataProviderCaller) GetAllReservesTokens(opts *bind.CallOpts) ([]AaveProtocolDataProviderTokenData, error) {
	var out []interface{}
	err := _PoolDataProvider.contract.Call(opts, &out, "getAllReservesTokens")

	if err != nil {
		return *new([]AaveProtocolDataProviderTokenData), err
	}

	out0 := *abi.ConvertType(out[0], new([]AaveProtocolDataProviderTokenData)).(*[]AaveProtocolDataProviderTokenData)

	return out0, err

}

// GetAllReservesTokens is a free data retrieval call binding the contract method 0xb316ff89.
//
// Solidity: function getAllReservesTokens() view returns((string,address)[])
func (_PoolDataProvider *PoolDataProviderSession) GetAllReservesTokens() ([]AaveProtocolDataProviderTokenData, error) {
	return _PoolDataProvider.Contract.GetAllReservesTokens(&_PoolDataProvider.CallOpts)
}

// GetAllReservesTokens is a free data retrieval call binding the contract method 0xb316ff89.
//
// Solidity: function getAllReservesTokens() view returns((string,address)[])
func (_PoolDataProvider *PoolDataProviderCallerSession) GetAllReservesTokens() ([]AaveProtocolDataProviderTokenData, error) {
	return _PoolDataProvider.Contract.GetAllReservesTokens(&_PoolDataProvider.CallOpts)
}

// GetDebtCeiling is a free data retrieval call binding the contract method 0x3c798109.
//
// Solidity: function getDebtCeiling(address asset) view returns(uint256)
func (_PoolDataProvider *PoolDataProviderCaller) GetDebtCeiling(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PoolDataProvider.contract.Call(opts, &out, "getDebtCeiling", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDebtCeiling is a free data retrieval call binding the contract method 0x3c798109.
//
// Solidity: function getDebtCeiling(address asset) view returns(uint256)
func (_PoolDataProvider *PoolDataProviderSession) GetDebtCeiling(asset common.Address) (*big.Int, error) {
	return _PoolDataProvider.Contract.GetDebtCeiling(&_PoolDataProvider.CallOpts, asset)
}

// GetDebtCeiling is a free data retrieval call binding the contract method 0x3c798109.
//
// Solidity: function getDebtCeiling(address asset) view returns(uint256)
func (_PoolDataProvider *PoolDataProviderCallerSession) GetDebtCeiling(asset common.Address) (*big.Int, error) {
	return _PoolDataProvider.Contract.GetDebtCeiling(&_PoolDataProvider.CallOpts, asset)
}

// GetDebtCeilingDecimals is a free data retrieval call binding the contract method 0x69b169e1.
//
// Solidity: function getDebtCeilingDecimals() pure returns(uint256)
func (_PoolDataProvider *PoolDataProviderCaller) GetDebtCeilingDecimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PoolDataProvider.contract.Call(opts, &out, "getDebtCeilingDecimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDebtCeilingDecimals is a free data retrieval call binding the contract method 0x69b169e1.
//
// Solidity: function getDebtCeilingDecimals() pure returns(uint256)
func (_PoolDataProvider *PoolDataProviderSession) GetDebtCeilingDecimals() (*big.Int, error) {
	return _PoolDataProvider.Contract.GetDebtCeilingDecimals(&_PoolDataProvider.CallOpts)
}

// GetDebtCeilingDecimals is a free data retrieval call binding the contract method 0x69b169e1.
//
// Solidity: function getDebtCeilingDecimals() pure returns(uint256)
func (_PoolDataProvider *PoolDataProviderCallerSession) GetDebtCeilingDecimals() (*big.Int, error) {
	return _PoolDataProvider.Contract.GetDebtCeilingDecimals(&_PoolDataProvider.CallOpts)
}

// GetInterestRateStrategyAddress is a free data retrieval call binding the contract method 0x6744362a.
//
// Solidity: function getInterestRateStrategyAddress(address asset) view returns(address irStrategyAddress)
func (_PoolDataProvider *PoolDataProviderCaller) GetInterestRateStrategyAddress(opts *bind.CallOpts, asset common.Address) (common.Address, error) {
	var out []interface{}
	err := _PoolDataProvider.contract.Call(opts, &out, "getInterestRateStrategyAddress", asset)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInterestRateStrategyAddress is a free data retrieval call binding the contract method 0x6744362a.
//
// Solidity: function getInterestRateStrategyAddress(address asset) view returns(address irStrategyAddress)
func (_PoolDataProvider *PoolDataProviderSession) GetInterestRateStrategyAddress(asset common.Address) (common.Address, error) {
	return _PoolDataProvider.Contract.GetInterestRateStrategyAddress(&_PoolDataProvider.CallOpts, asset)
}

// GetInterestRateStrategyAddress is a free data retrieval call binding the contract method 0x6744362a.
//
// Solidity: function getInterestRateStrategyAddress(address asset) view returns(address irStrategyAddress)
func (_PoolDataProvider *PoolDataProviderCallerSession) GetInterestRateStrategyAddress(asset common.Address) (common.Address, error) {
	return _PoolDataProvider.Contract.GetInterestRateStrategyAddress(&_PoolDataProvider.CallOpts, asset)
}

// GetLiquidationProtocolFee is a free data retrieval call binding the contract method 0x3cb8a622.
//
// Solidity: function getLiquidationProtocolFee(address asset) view returns(uint256)
func (_PoolDataProvider *PoolDataProviderCaller) GetLiquidationProtocolFee(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PoolDataProvider.contract.Call(opts, &out, "getLiquidationProtocolFee", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLiquidationProtocolFee is a free data retrieval call binding the contract method 0x3cb8a622.
//
// Solidity: function getLiquidationProtocolFee(address asset) view returns(uint256)
func (_PoolDataProvider *PoolDataProviderSession) GetLiquidationProtocolFee(asset common.Address) (*big.Int, error) {
	return _PoolDataProvider.Contract.GetLiquidationProtocolFee(&_PoolDataProvider.CallOpts, asset)
}

// GetLiquidationProtocolFee is a free data retrieval call binding the contract method 0x3cb8a622.
//
// Solidity: function getLiquidationProtocolFee(address asset) view returns(uint256)
func (_PoolDataProvider *PoolDataProviderCallerSession) GetLiquidationProtocolFee(asset common.Address) (*big.Int, error) {
	return _PoolDataProvider.Contract.GetLiquidationProtocolFee(&_PoolDataProvider.CallOpts, asset)
}

// GetPaused is a free data retrieval call binding the contract method 0xb55d9904.
//
// Solidity: function getPaused(address asset) view returns(bool isPaused)
func (_PoolDataProvider *PoolDataProviderCaller) GetPaused(opts *bind.CallOpts, asset common.Address) (bool, error) {
	var out []interface{}
	err := _PoolDataProvider.contract.Call(opts, &out, "getPaused", asset)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetPaused is a free data retrieval call binding the contract method 0xb55d9904.
//
// Solidity: function getPaused(address asset) view returns(bool isPaused)
func (_PoolDataProvider *PoolDataProviderSession) GetPaused(asset common.Address) (bool, error) {
	return _PoolDataProvider.Contract.GetPaused(&_PoolDataProvider.CallOpts, asset)
}

// GetPaused is a free data retrieval call binding the contract method 0xb55d9904.
//
// Solidity: function getPaused(address asset) view returns(bool isPaused)
func (_PoolDataProvider *PoolDataProviderCallerSession) GetPaused(asset common.Address) (bool, error) {
	return _PoolDataProvider.Contract.GetPaused(&_PoolDataProvider.CallOpts, asset)
}

// GetReserveCaps is a free data retrieval call binding the contract method 0x46fbe558.
//
// Solidity: function getReserveCaps(address asset) view returns(uint256 borrowCap, uint256 supplyCap)
func (_PoolDataProvider *PoolDataProviderCaller) GetReserveCaps(opts *bind.CallOpts, asset common.Address) (struct {
	BorrowCap *big.Int
	SupplyCap *big.Int
}, error) {
	var out []interface{}
	err := _PoolDataProvider.contract.Call(opts, &out, "getReserveCaps", asset)

	outstruct := new(struct {
		BorrowCap *big.Int
		SupplyCap *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BorrowCap = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SupplyCap = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetReserveCaps is a free data retrieval call binding the contract method 0x46fbe558.
//
// Solidity: function getReserveCaps(address asset) view returns(uint256 borrowCap, uint256 supplyCap)
func (_PoolDataProvider *PoolDataProviderSession) GetReserveCaps(asset common.Address) (struct {
	BorrowCap *big.Int
	SupplyCap *big.Int
}, error) {
	return _PoolDataProvider.Contract.GetReserveCaps(&_PoolDataProvider.CallOpts, asset)
}

// GetReserveCaps is a free data retrieval call binding the contract method 0x46fbe558.
//
// Solidity: function getReserveCaps(address asset) view returns(uint256 borrowCap, uint256 supplyCap)
func (_PoolDataProvider *PoolDataProviderCallerSession) GetReserveCaps(asset common.Address) (struct {
	BorrowCap *big.Int
	SupplyCap *big.Int
}, error) {
	return _PoolDataProvider.Contract.GetReserveCaps(&_PoolDataProvider.CallOpts, asset)
}

// GetReserveConfigurationData is a free data retrieval call binding the contract method 0x3e150141.
//
// Solidity: function getReserveConfigurationData(address asset) view returns(uint256 decimals, uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus, uint256 reserveFactor, bool usageAsCollateralEnabled, bool borrowingEnabled, bool stableBorrowRateEnabled, bool isActive, bool isFrozen)
func (_PoolDataProvider *PoolDataProviderCaller) GetReserveConfigurationData(opts *bind.CallOpts, asset common.Address) (struct {
	Decimals                 *big.Int
	Ltv                      *big.Int
	LiquidationThreshold     *big.Int
	LiquidationBonus         *big.Int
	ReserveFactor            *big.Int
	UsageAsCollateralEnabled bool
	BorrowingEnabled         bool
	StableBorrowRateEnabled  bool
	IsActive                 bool
	IsFrozen                 bool
}, error) {
	var out []interface{}
	err := _PoolDataProvider.contract.Call(opts, &out, "getReserveConfigurationData", asset)

	outstruct := new(struct {
		Decimals                 *big.Int
		Ltv                      *big.Int
		LiquidationThreshold     *big.Int
		LiquidationBonus         *big.Int
		ReserveFactor            *big.Int
		UsageAsCollateralEnabled bool
		BorrowingEnabled         bool
		StableBorrowRateEnabled  bool
		IsActive                 bool
		IsFrozen                 bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Decimals = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Ltv = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LiquidationThreshold = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LiquidationBonus = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ReserveFactor = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.UsageAsCollateralEnabled = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.BorrowingEnabled = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.StableBorrowRateEnabled = *abi.ConvertType(out[7], new(bool)).(*bool)
	outstruct.IsActive = *abi.ConvertType(out[8], new(bool)).(*bool)
	outstruct.IsFrozen = *abi.ConvertType(out[9], new(bool)).(*bool)

	return *outstruct, err

}

// GetReserveConfigurationData is a free data retrieval call binding the contract method 0x3e150141.
//
// Solidity: function getReserveConfigurationData(address asset) view returns(uint256 decimals, uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus, uint256 reserveFactor, bool usageAsCollateralEnabled, bool borrowingEnabled, bool stableBorrowRateEnabled, bool isActive, bool isFrozen)
func (_PoolDataProvider *PoolDataProviderSession) GetReserveConfigurationData(asset common.Address) (struct {
	Decimals                 *big.Int
	Ltv                      *big.Int
	LiquidationThreshold     *big.Int
	LiquidationBonus         *big.Int
	ReserveFactor            *big.Int
	UsageAsCollateralEnabled bool
	BorrowingEnabled         bool
	StableBorrowRateEnabled  bool
	IsActive                 bool
	IsFrozen                 bool
}, error) {
	return _PoolDataProvider.Contract.GetReserveConfigurationData(&_PoolDataProvider.CallOpts, asset)
}

// GetReserveConfigurationData is a free data retrieval call binding the contract method 0x3e150141.
//
// Solidity: function getReserveConfigurationData(address asset) view returns(uint256 decimals, uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus, uint256 reserveFactor, bool usageAsCollateralEnabled, bool borrowingEnabled, bool stableBorrowRateEnabled, bool isActive, bool isFrozen)
func (_PoolDataProvider *PoolDataProviderCallerSession) GetReserveConfigurationData(asset common.Address) (struct {
	Decimals                 *big.Int
	Ltv                      *big.Int
	LiquidationThreshold     *big.Int
	LiquidationBonus         *big.Int
	ReserveFactor            *big.Int
	UsageAsCollateralEnabled bool
	BorrowingEnabled         bool
	StableBorrowRateEnabled  bool
	IsActive                 bool
	IsFrozen                 bool
}, error) {
	return _PoolDataProvider.Contract.GetReserveConfigurationData(&_PoolDataProvider.CallOpts, asset)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(uint256 unbacked, uint256 accruedToTreasuryScaled, uint256 totalAToken, uint256 totalStableDebt, uint256 totalVariableDebt, uint256 liquidityRate, uint256 variableBorrowRate, uint256 stableBorrowRate, uint256 averageStableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex, uint40 lastUpdateTimestamp)
func (_PoolDataProvider *PoolDataProviderCaller) GetReserveData(opts *bind.CallOpts, asset common.Address) (struct {
	Unbacked                *big.Int
	AccruedToTreasuryScaled *big.Int
	TotalAToken             *big.Int
	TotalStableDebt         *big.Int
	TotalVariableDebt       *big.Int
	LiquidityRate           *big.Int
	VariableBorrowRate      *big.Int
	StableBorrowRate        *big.Int
	AverageStableBorrowRate *big.Int
	LiquidityIndex          *big.Int
	VariableBorrowIndex     *big.Int
	LastUpdateTimestamp     *big.Int
}, error) {
	var out []interface{}
	err := _PoolDataProvider.contract.Call(opts, &out, "getReserveData", asset)

	outstruct := new(struct {
		Unbacked                *big.Int
		AccruedToTreasuryScaled *big.Int
		TotalAToken             *big.Int
		TotalStableDebt         *big.Int
		TotalVariableDebt       *big.Int
		LiquidityRate           *big.Int
		VariableBorrowRate      *big.Int
		StableBorrowRate        *big.Int
		AverageStableBorrowRate *big.Int
		LiquidityIndex          *big.Int
		VariableBorrowIndex     *big.Int
		LastUpdateTimestamp     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Unbacked = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AccruedToTreasuryScaled = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalAToken = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TotalStableDebt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.TotalVariableDebt = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.LiquidityRate = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.VariableBorrowRate = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.StableBorrowRate = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.AverageStableBorrowRate = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.LiquidityIndex = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)
	outstruct.VariableBorrowIndex = *abi.ConvertType(out[10], new(*big.Int)).(**big.Int)
	outstruct.LastUpdateTimestamp = *abi.ConvertType(out[11], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(uint256 unbacked, uint256 accruedToTreasuryScaled, uint256 totalAToken, uint256 totalStableDebt, uint256 totalVariableDebt, uint256 liquidityRate, uint256 variableBorrowRate, uint256 stableBorrowRate, uint256 averageStableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex, uint40 lastUpdateTimestamp)
func (_PoolDataProvider *PoolDataProviderSession) GetReserveData(asset common.Address) (struct {
	Unbacked                *big.Int
	AccruedToTreasuryScaled *big.Int
	TotalAToken             *big.Int
	TotalStableDebt         *big.Int
	TotalVariableDebt       *big.Int
	LiquidityRate           *big.Int
	VariableBorrowRate      *big.Int
	StableBorrowRate        *big.Int
	AverageStableBorrowRate *big.Int
	LiquidityIndex          *big.Int
	VariableBorrowIndex     *big.Int
	LastUpdateTimestamp     *big.Int
}, error) {
	return _PoolDataProvider.Contract.GetReserveData(&_PoolDataProvider.CallOpts, asset)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(uint256 unbacked, uint256 accruedToTreasuryScaled, uint256 totalAToken, uint256 totalStableDebt, uint256 totalVariableDebt, uint256 liquidityRate, uint256 variableBorrowRate, uint256 stableBorrowRate, uint256 averageStableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex, uint40 lastUpdateTimestamp)
func (_PoolDataProvider *PoolDataProviderCallerSession) GetReserveData(asset common.Address) (struct {
	Unbacked                *big.Int
	AccruedToTreasuryScaled *big.Int
	TotalAToken             *big.Int
	TotalStableDebt         *big.Int
	TotalVariableDebt       *big.Int
	LiquidityRate           *big.Int
	VariableBorrowRate      *big.Int
	StableBorrowRate        *big.Int
	AverageStableBorrowRate *big.Int
	LiquidityIndex          *big.Int
	VariableBorrowIndex     *big.Int
	LastUpdateTimestamp     *big.Int
}, error) {
	return _PoolDataProvider.Contract.GetReserveData(&_PoolDataProvider.CallOpts, asset)
}

// GetReserveEModeCategory is a free data retrieval call binding the contract method 0x163a0f20.
//
// Solidity: function getReserveEModeCategory(address asset) view returns(uint256)
func (_PoolDataProvider *PoolDataProviderCaller) GetReserveEModeCategory(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PoolDataProvider.contract.Call(opts, &out, "getReserveEModeCategory", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveEModeCategory is a free data retrieval call binding the contract method 0x163a0f20.
//
// Solidity: function getReserveEModeCategory(address asset) view returns(uint256)
func (_PoolDataProvider *PoolDataProviderSession) GetReserveEModeCategory(asset common.Address) (*big.Int, error) {
	return _PoolDataProvider.Contract.GetReserveEModeCategory(&_PoolDataProvider.CallOpts, asset)
}

// GetReserveEModeCategory is a free data retrieval call binding the contract method 0x163a0f20.
//
// Solidity: function getReserveEModeCategory(address asset) view returns(uint256)
func (_PoolDataProvider *PoolDataProviderCallerSession) GetReserveEModeCategory(asset common.Address) (*big.Int, error) {
	return _PoolDataProvider.Contract.GetReserveEModeCategory(&_PoolDataProvider.CallOpts, asset)
}

// GetReserveTokensAddresses is a free data retrieval call binding the contract method 0xd2493b6c.
//
// Solidity: function getReserveTokensAddresses(address asset) view returns(address aTokenAddress, address stableDebtTokenAddress, address variableDebtTokenAddress)
func (_PoolDataProvider *PoolDataProviderCaller) GetReserveTokensAddresses(opts *bind.CallOpts, asset common.Address) (struct {
	ATokenAddress            common.Address
	StableDebtTokenAddress   common.Address
	VariableDebtTokenAddress common.Address
}, error) {
	var out []interface{}
	err := _PoolDataProvider.contract.Call(opts, &out, "getReserveTokensAddresses", asset)

	outstruct := new(struct {
		ATokenAddress            common.Address
		StableDebtTokenAddress   common.Address
		VariableDebtTokenAddress common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ATokenAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.StableDebtTokenAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.VariableDebtTokenAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetReserveTokensAddresses is a free data retrieval call binding the contract method 0xd2493b6c.
//
// Solidity: function getReserveTokensAddresses(address asset) view returns(address aTokenAddress, address stableDebtTokenAddress, address variableDebtTokenAddress)
func (_PoolDataProvider *PoolDataProviderSession) GetReserveTokensAddresses(asset common.Address) (struct {
	ATokenAddress            common.Address
	StableDebtTokenAddress   common.Address
	VariableDebtTokenAddress common.Address
}, error) {
	return _PoolDataProvider.Contract.GetReserveTokensAddresses(&_PoolDataProvider.CallOpts, asset)
}

// GetReserveTokensAddresses is a free data retrieval call binding the contract method 0xd2493b6c.
//
// Solidity: function getReserveTokensAddresses(address asset) view returns(address aTokenAddress, address stableDebtTokenAddress, address variableDebtTokenAddress)
func (_PoolDataProvider *PoolDataProviderCallerSession) GetReserveTokensAddresses(asset common.Address) (struct {
	ATokenAddress            common.Address
	StableDebtTokenAddress   common.Address
	VariableDebtTokenAddress common.Address
}, error) {
	return _PoolDataProvider.Contract.GetReserveTokensAddresses(&_PoolDataProvider.CallOpts, asset)
}

// GetSiloedBorrowing is a free data retrieval call binding the contract method 0xfcf40a62.
//
// Solidity: function getSiloedBorrowing(address asset) view returns(bool)
func (_PoolDataProvider *PoolDataProviderCaller) GetSiloedBorrowing(opts *bind.CallOpts, asset common.Address) (bool, error) {
	var out []interface{}
	err := _PoolDataProvider.contract.Call(opts, &out, "getSiloedBorrowing", asset)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetSiloedBorrowing is a free data retrieval call binding the contract method 0xfcf40a62.
//
// Solidity: function getSiloedBorrowing(address asset) view returns(bool)
func (_PoolDataProvider *PoolDataProviderSession) GetSiloedBorrowing(asset common.Address) (bool, error) {
	return _PoolDataProvider.Contract.GetSiloedBorrowing(&_PoolDataProvider.CallOpts, asset)
}

// GetSiloedBorrowing is a free data retrieval call binding the contract method 0xfcf40a62.
//
// Solidity: function getSiloedBorrowing(address asset) view returns(bool)
func (_PoolDataProvider *PoolDataProviderCallerSession) GetSiloedBorrowing(asset common.Address) (bool, error) {
	return _PoolDataProvider.Contract.GetSiloedBorrowing(&_PoolDataProvider.CallOpts, asset)
}

// GetTotalDebt is a free data retrieval call binding the contract method 0x4d44ac4f.
//
// Solidity: function getTotalDebt(address asset) view returns(uint256)
func (_PoolDataProvider *PoolDataProviderCaller) GetTotalDebt(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PoolDataProvider.contract.Call(opts, &out, "getTotalDebt", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalDebt is a free data retrieval call binding the contract method 0x4d44ac4f.
//
// Solidity: function getTotalDebt(address asset) view returns(uint256)
func (_PoolDataProvider *PoolDataProviderSession) GetTotalDebt(asset common.Address) (*big.Int, error) {
	return _PoolDataProvider.Contract.GetTotalDebt(&_PoolDataProvider.CallOpts, asset)
}

// GetTotalDebt is a free data retrieval call binding the contract method 0x4d44ac4f.
//
// Solidity: function getTotalDebt(address asset) view returns(uint256)
func (_PoolDataProvider *PoolDataProviderCallerSession) GetTotalDebt(asset common.Address) (*big.Int, error) {
	return _PoolDataProvider.Contract.GetTotalDebt(&_PoolDataProvider.CallOpts, asset)
}

// GetUnbackedMintCap is a free data retrieval call binding the contract method 0x7ba1ae36.
//
// Solidity: function getUnbackedMintCap(address asset) view returns(uint256)
func (_PoolDataProvider *PoolDataProviderCaller) GetUnbackedMintCap(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PoolDataProvider.contract.Call(opts, &out, "getUnbackedMintCap", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnbackedMintCap is a free data retrieval call binding the contract method 0x7ba1ae36.
//
// Solidity: function getUnbackedMintCap(address asset) view returns(uint256)
func (_PoolDataProvider *PoolDataProviderSession) GetUnbackedMintCap(asset common.Address) (*big.Int, error) {
	return _PoolDataProvider.Contract.GetUnbackedMintCap(&_PoolDataProvider.CallOpts, asset)
}

// GetUnbackedMintCap is a free data retrieval call binding the contract method 0x7ba1ae36.
//
// Solidity: function getUnbackedMintCap(address asset) view returns(uint256)
func (_PoolDataProvider *PoolDataProviderCallerSession) GetUnbackedMintCap(asset common.Address) (*big.Int, error) {
	return _PoolDataProvider.Contract.GetUnbackedMintCap(&_PoolDataProvider.CallOpts, asset)
}

// GetUserReserveData is a free data retrieval call binding the contract method 0x28dd2d01.
//
// Solidity: function getUserReserveData(address asset, address user) view returns(uint256 currentATokenBalance, uint256 currentStableDebt, uint256 currentVariableDebt, uint256 principalStableDebt, uint256 scaledVariableDebt, uint256 stableBorrowRate, uint256 liquidityRate, uint40 stableRateLastUpdated, bool usageAsCollateralEnabled)
func (_PoolDataProvider *PoolDataProviderCaller) GetUserReserveData(opts *bind.CallOpts, asset common.Address, user common.Address) (struct {
	CurrentATokenBalance     *big.Int
	CurrentStableDebt        *big.Int
	CurrentVariableDebt      *big.Int
	PrincipalStableDebt      *big.Int
	ScaledVariableDebt       *big.Int
	StableBorrowRate         *big.Int
	LiquidityRate            *big.Int
	StableRateLastUpdated    *big.Int
	UsageAsCollateralEnabled bool
}, error) {
	var out []interface{}
	err := _PoolDataProvider.contract.Call(opts, &out, "getUserReserveData", asset, user)

	outstruct := new(struct {
		CurrentATokenBalance     *big.Int
		CurrentStableDebt        *big.Int
		CurrentVariableDebt      *big.Int
		PrincipalStableDebt      *big.Int
		ScaledVariableDebt       *big.Int
		StableBorrowRate         *big.Int
		LiquidityRate            *big.Int
		StableRateLastUpdated    *big.Int
		UsageAsCollateralEnabled bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CurrentATokenBalance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CurrentStableDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CurrentVariableDebt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.PrincipalStableDebt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ScaledVariableDebt = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.StableBorrowRate = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.LiquidityRate = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.StableRateLastUpdated = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.UsageAsCollateralEnabled = *abi.ConvertType(out[8], new(bool)).(*bool)

	return *outstruct, err

}

// GetUserReserveData is a free data retrieval call binding the contract method 0x28dd2d01.
//
// Solidity: function getUserReserveData(address asset, address user) view returns(uint256 currentATokenBalance, uint256 currentStableDebt, uint256 currentVariableDebt, uint256 principalStableDebt, uint256 scaledVariableDebt, uint256 stableBorrowRate, uint256 liquidityRate, uint40 stableRateLastUpdated, bool usageAsCollateralEnabled)
func (_PoolDataProvider *PoolDataProviderSession) GetUserReserveData(asset common.Address, user common.Address) (struct {
	CurrentATokenBalance     *big.Int
	CurrentStableDebt        *big.Int
	CurrentVariableDebt      *big.Int
	PrincipalStableDebt      *big.Int
	ScaledVariableDebt       *big.Int
	StableBorrowRate         *big.Int
	LiquidityRate            *big.Int
	StableRateLastUpdated    *big.Int
	UsageAsCollateralEnabled bool
}, error) {
	return _PoolDataProvider.Contract.GetUserReserveData(&_PoolDataProvider.CallOpts, asset, user)
}

// GetUserReserveData is a free data retrieval call binding the contract method 0x28dd2d01.
//
// Solidity: function getUserReserveData(address asset, address user) view returns(uint256 currentATokenBalance, uint256 currentStableDebt, uint256 currentVariableDebt, uint256 principalStableDebt, uint256 scaledVariableDebt, uint256 stableBorrowRate, uint256 liquidityRate, uint40 stableRateLastUpdated, bool usageAsCollateralEnabled)
func (_PoolDataProvider *PoolDataProviderCallerSession) GetUserReserveData(asset common.Address, user common.Address) (struct {
	CurrentATokenBalance     *big.Int
	CurrentStableDebt        *big.Int
	CurrentVariableDebt      *big.Int
	PrincipalStableDebt      *big.Int
	ScaledVariableDebt       *big.Int
	StableBorrowRate         *big.Int
	LiquidityRate            *big.Int
	StableRateLastUpdated    *big.Int
	UsageAsCollateralEnabled bool
}, error) {
	return _PoolDataProvider.Contract.GetUserReserveData(&_PoolDataProvider.CallOpts, asset, user)
}
