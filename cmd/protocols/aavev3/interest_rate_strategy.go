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

// DataTypesCalculateInterestRatesParams is an auto generated low-level Go binding around an user-defined struct.
type DataTypesCalculateInterestRatesParams struct {
	Unbacked                *big.Int
	LiquidityAdded          *big.Int
	LiquidityTaken          *big.Int
	TotalStableDebt         *big.Int
	TotalVariableDebt       *big.Int
	AverageStableBorrowRate *big.Int
	ReserveFactor           *big.Int
	Reserve                 common.Address
	AToken                  common.Address
}

// InterestRateStrategyMetaData contains all meta data concerning the InterestRateStrategy contract.
var InterestRateStrategyMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIPoolAddressesProvider\",\"name\":\"provider\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"optimalUsageRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseVariableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableRateSlope1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableRateSlope2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableRateSlope1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableRateSlope2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"baseStableRateOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableRateExcessOffset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"optimalStableToTotalDebtRatio\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ADDRESSES_PROVIDER\",\"outputs\":[{\"internalType\":\"contractIPoolAddressesProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_EXCESS_STABLE_TO_TOTAL_DEBT_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_EXCESS_USAGE_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPTIMAL_STABLE_TO_TOTAL_DEBT_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"OPTIMAL_USAGE_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"unbacked\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityAdded\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityTaken\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"averageStableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveFactor\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"aToken\",\"type\":\"address\"}],\"internalType\":\"structDataTypes.CalculateInterestRatesParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"calculateInterestRates\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBaseStableBorrowRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBaseVariableBorrowRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxVariableBorrowRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStableRateExcessOffset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStableRateSlope1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStableRateSlope2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVariableRateSlope1\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVariableRateSlope2\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// InterestRateStrategyABI is the input ABI used to generate the binding from.
// Deprecated: Use InterestRateStrategyMetaData.ABI instead.
var InterestRateStrategyABI = InterestRateStrategyMetaData.ABI

// InterestRateStrategy is an auto generated Go binding around an Ethereum contract.
type InterestRateStrategy struct {
	InterestRateStrategyCaller     // Read-only binding to the contract
	InterestRateStrategyTransactor // Write-only binding to the contract
	InterestRateStrategyFilterer   // Log filterer for contract events
}

// InterestRateStrategyCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterestRateStrategyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterestRateStrategyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterestRateStrategyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterestRateStrategyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterestRateStrategyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterestRateStrategySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterestRateStrategySession struct {
	Contract     *InterestRateStrategy // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// InterestRateStrategyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterestRateStrategyCallerSession struct {
	Contract *InterestRateStrategyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// InterestRateStrategyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterestRateStrategyTransactorSession struct {
	Contract     *InterestRateStrategyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// InterestRateStrategyRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterestRateStrategyRaw struct {
	Contract *InterestRateStrategy // Generic contract binding to access the raw methods on
}

// InterestRateStrategyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterestRateStrategyCallerRaw struct {
	Contract *InterestRateStrategyCaller // Generic read-only contract binding to access the raw methods on
}

// InterestRateStrategyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterestRateStrategyTransactorRaw struct {
	Contract *InterestRateStrategyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterestRateStrategy creates a new instance of InterestRateStrategy, bound to a specific deployed contract.
func NewInterestRateStrategy(address common.Address, backend bind.ContractBackend) (*InterestRateStrategy, error) {
	contract, err := bindInterestRateStrategy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterestRateStrategy{InterestRateStrategyCaller: InterestRateStrategyCaller{contract: contract}, InterestRateStrategyTransactor: InterestRateStrategyTransactor{contract: contract}, InterestRateStrategyFilterer: InterestRateStrategyFilterer{contract: contract}}, nil
}

// NewInterestRateStrategyCaller creates a new read-only instance of InterestRateStrategy, bound to a specific deployed contract.
func NewInterestRateStrategyCaller(address common.Address, caller bind.ContractCaller) (*InterestRateStrategyCaller, error) {
	contract, err := bindInterestRateStrategy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterestRateStrategyCaller{contract: contract}, nil
}

// NewInterestRateStrategyTransactor creates a new write-only instance of InterestRateStrategy, bound to a specific deployed contract.
func NewInterestRateStrategyTransactor(address common.Address, transactor bind.ContractTransactor) (*InterestRateStrategyTransactor, error) {
	contract, err := bindInterestRateStrategy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterestRateStrategyTransactor{contract: contract}, nil
}

// NewInterestRateStrategyFilterer creates a new log filterer instance of InterestRateStrategy, bound to a specific deployed contract.
func NewInterestRateStrategyFilterer(address common.Address, filterer bind.ContractFilterer) (*InterestRateStrategyFilterer, error) {
	contract, err := bindInterestRateStrategy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterestRateStrategyFilterer{contract: contract}, nil
}

// bindInterestRateStrategy binds a generic wrapper to an already deployed contract.
func bindInterestRateStrategy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := InterestRateStrategyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterestRateStrategy *InterestRateStrategyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterestRateStrategy.Contract.InterestRateStrategyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterestRateStrategy *InterestRateStrategyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterestRateStrategy.Contract.InterestRateStrategyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterestRateStrategy *InterestRateStrategyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterestRateStrategy.Contract.InterestRateStrategyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterestRateStrategy *InterestRateStrategyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterestRateStrategy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterestRateStrategy *InterestRateStrategyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterestRateStrategy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterestRateStrategy *InterestRateStrategyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterestRateStrategy.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_InterestRateStrategy *InterestRateStrategyCaller) ADDRESSESPROVIDER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _InterestRateStrategy.contract.Call(opts, &out, "ADDRESSES_PROVIDER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_InterestRateStrategy *InterestRateStrategySession) ADDRESSESPROVIDER() (common.Address, error) {
	return _InterestRateStrategy.Contract.ADDRESSESPROVIDER(&_InterestRateStrategy.CallOpts)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_InterestRateStrategy *InterestRateStrategyCallerSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _InterestRateStrategy.Contract.ADDRESSESPROVIDER(&_InterestRateStrategy.CallOpts)
}

// MAXEXCESSSTABLETOTOTALDEBTRATIO is a free data retrieval call binding the contract method 0xfe5fd698.
//
// Solidity: function MAX_EXCESS_STABLE_TO_TOTAL_DEBT_RATIO() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategyCaller) MAXEXCESSSTABLETOTOTALDEBTRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateStrategy.contract.Call(opts, &out, "MAX_EXCESS_STABLE_TO_TOTAL_DEBT_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXEXCESSSTABLETOTOTALDEBTRATIO is a free data retrieval call binding the contract method 0xfe5fd698.
//
// Solidity: function MAX_EXCESS_STABLE_TO_TOTAL_DEBT_RATIO() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategySession) MAXEXCESSSTABLETOTOTALDEBTRATIO() (*big.Int, error) {
	return _InterestRateStrategy.Contract.MAXEXCESSSTABLETOTOTALDEBTRATIO(&_InterestRateStrategy.CallOpts)
}

// MAXEXCESSSTABLETOTOTALDEBTRATIO is a free data retrieval call binding the contract method 0xfe5fd698.
//
// Solidity: function MAX_EXCESS_STABLE_TO_TOTAL_DEBT_RATIO() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategyCallerSession) MAXEXCESSSTABLETOTOTALDEBTRATIO() (*big.Int, error) {
	return _InterestRateStrategy.Contract.MAXEXCESSSTABLETOTOTALDEBTRATIO(&_InterestRateStrategy.CallOpts)
}

// MAXEXCESSUSAGERATIO is a free data retrieval call binding the contract method 0xa9c622f8.
//
// Solidity: function MAX_EXCESS_USAGE_RATIO() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategyCaller) MAXEXCESSUSAGERATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateStrategy.contract.Call(opts, &out, "MAX_EXCESS_USAGE_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXEXCESSUSAGERATIO is a free data retrieval call binding the contract method 0xa9c622f8.
//
// Solidity: function MAX_EXCESS_USAGE_RATIO() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategySession) MAXEXCESSUSAGERATIO() (*big.Int, error) {
	return _InterestRateStrategy.Contract.MAXEXCESSUSAGERATIO(&_InterestRateStrategy.CallOpts)
}

// MAXEXCESSUSAGERATIO is a free data retrieval call binding the contract method 0xa9c622f8.
//
// Solidity: function MAX_EXCESS_USAGE_RATIO() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategyCallerSession) MAXEXCESSUSAGERATIO() (*big.Int, error) {
	return _InterestRateStrategy.Contract.MAXEXCESSUSAGERATIO(&_InterestRateStrategy.CallOpts)
}

// OPTIMALSTABLETOTOTALDEBTRATIO is a free data retrieval call binding the contract method 0x6fb92589.
//
// Solidity: function OPTIMAL_STABLE_TO_TOTAL_DEBT_RATIO() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategyCaller) OPTIMALSTABLETOTOTALDEBTRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateStrategy.contract.Call(opts, &out, "OPTIMAL_STABLE_TO_TOTAL_DEBT_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OPTIMALSTABLETOTOTALDEBTRATIO is a free data retrieval call binding the contract method 0x6fb92589.
//
// Solidity: function OPTIMAL_STABLE_TO_TOTAL_DEBT_RATIO() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategySession) OPTIMALSTABLETOTOTALDEBTRATIO() (*big.Int, error) {
	return _InterestRateStrategy.Contract.OPTIMALSTABLETOTOTALDEBTRATIO(&_InterestRateStrategy.CallOpts)
}

// OPTIMALSTABLETOTOTALDEBTRATIO is a free data retrieval call binding the contract method 0x6fb92589.
//
// Solidity: function OPTIMAL_STABLE_TO_TOTAL_DEBT_RATIO() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategyCallerSession) OPTIMALSTABLETOTOTALDEBTRATIO() (*big.Int, error) {
	return _InterestRateStrategy.Contract.OPTIMALSTABLETOTOTALDEBTRATIO(&_InterestRateStrategy.CallOpts)
}

// OPTIMALUSAGERATIO is a free data retrieval call binding the contract method 0x54c365c6.
//
// Solidity: function OPTIMAL_USAGE_RATIO() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategyCaller) OPTIMALUSAGERATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateStrategy.contract.Call(opts, &out, "OPTIMAL_USAGE_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OPTIMALUSAGERATIO is a free data retrieval call binding the contract method 0x54c365c6.
//
// Solidity: function OPTIMAL_USAGE_RATIO() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategySession) OPTIMALUSAGERATIO() (*big.Int, error) {
	return _InterestRateStrategy.Contract.OPTIMALUSAGERATIO(&_InterestRateStrategy.CallOpts)
}

// OPTIMALUSAGERATIO is a free data retrieval call binding the contract method 0x54c365c6.
//
// Solidity: function OPTIMAL_USAGE_RATIO() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategyCallerSession) OPTIMALUSAGERATIO() (*big.Int, error) {
	return _InterestRateStrategy.Contract.OPTIMALUSAGERATIO(&_InterestRateStrategy.CallOpts)
}

// CalculateInterestRates is a free data retrieval call binding the contract method 0xa5898709.
//
// Solidity: function calculateInterestRates((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address) params) view returns(uint256, uint256, uint256)
func (_InterestRateStrategy *InterestRateStrategyCaller) CalculateInterestRates(opts *bind.CallOpts, params DataTypesCalculateInterestRatesParams) (*big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _InterestRateStrategy.contract.Call(opts, &out, "calculateInterestRates", params)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return out0, out1, out2, err

}

// CalculateInterestRates is a free data retrieval call binding the contract method 0xa5898709.
//
// Solidity: function calculateInterestRates((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address) params) view returns(uint256, uint256, uint256)
func (_InterestRateStrategy *InterestRateStrategySession) CalculateInterestRates(params DataTypesCalculateInterestRatesParams) (*big.Int, *big.Int, *big.Int, error) {
	return _InterestRateStrategy.Contract.CalculateInterestRates(&_InterestRateStrategy.CallOpts, params)
}

// CalculateInterestRates is a free data retrieval call binding the contract method 0xa5898709.
//
// Solidity: function calculateInterestRates((uint256,uint256,uint256,uint256,uint256,uint256,uint256,address,address) params) view returns(uint256, uint256, uint256)
func (_InterestRateStrategy *InterestRateStrategyCallerSession) CalculateInterestRates(params DataTypesCalculateInterestRatesParams) (*big.Int, *big.Int, *big.Int, error) {
	return _InterestRateStrategy.Contract.CalculateInterestRates(&_InterestRateStrategy.CallOpts, params)
}

// GetBaseStableBorrowRate is a free data retrieval call binding the contract method 0xacd78686.
//
// Solidity: function getBaseStableBorrowRate() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategyCaller) GetBaseStableBorrowRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateStrategy.contract.Call(opts, &out, "getBaseStableBorrowRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBaseStableBorrowRate is a free data retrieval call binding the contract method 0xacd78686.
//
// Solidity: function getBaseStableBorrowRate() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategySession) GetBaseStableBorrowRate() (*big.Int, error) {
	return _InterestRateStrategy.Contract.GetBaseStableBorrowRate(&_InterestRateStrategy.CallOpts)
}

// GetBaseStableBorrowRate is a free data retrieval call binding the contract method 0xacd78686.
//
// Solidity: function getBaseStableBorrowRate() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategyCallerSession) GetBaseStableBorrowRate() (*big.Int, error) {
	return _InterestRateStrategy.Contract.GetBaseStableBorrowRate(&_InterestRateStrategy.CallOpts)
}

// GetBaseVariableBorrowRate is a free data retrieval call binding the contract method 0x34762ca5.
//
// Solidity: function getBaseVariableBorrowRate() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategyCaller) GetBaseVariableBorrowRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateStrategy.contract.Call(opts, &out, "getBaseVariableBorrowRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBaseVariableBorrowRate is a free data retrieval call binding the contract method 0x34762ca5.
//
// Solidity: function getBaseVariableBorrowRate() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategySession) GetBaseVariableBorrowRate() (*big.Int, error) {
	return _InterestRateStrategy.Contract.GetBaseVariableBorrowRate(&_InterestRateStrategy.CallOpts)
}

// GetBaseVariableBorrowRate is a free data retrieval call binding the contract method 0x34762ca5.
//
// Solidity: function getBaseVariableBorrowRate() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategyCallerSession) GetBaseVariableBorrowRate() (*big.Int, error) {
	return _InterestRateStrategy.Contract.GetBaseVariableBorrowRate(&_InterestRateStrategy.CallOpts)
}

// GetMaxVariableBorrowRate is a free data retrieval call binding the contract method 0x80031e37.
//
// Solidity: function getMaxVariableBorrowRate() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategyCaller) GetMaxVariableBorrowRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateStrategy.contract.Call(opts, &out, "getMaxVariableBorrowRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxVariableBorrowRate is a free data retrieval call binding the contract method 0x80031e37.
//
// Solidity: function getMaxVariableBorrowRate() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategySession) GetMaxVariableBorrowRate() (*big.Int, error) {
	return _InterestRateStrategy.Contract.GetMaxVariableBorrowRate(&_InterestRateStrategy.CallOpts)
}

// GetMaxVariableBorrowRate is a free data retrieval call binding the contract method 0x80031e37.
//
// Solidity: function getMaxVariableBorrowRate() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategyCallerSession) GetMaxVariableBorrowRate() (*big.Int, error) {
	return _InterestRateStrategy.Contract.GetMaxVariableBorrowRate(&_InterestRateStrategy.CallOpts)
}

// GetStableRateExcessOffset is a free data retrieval call binding the contract method 0xbc626908.
//
// Solidity: function getStableRateExcessOffset() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategyCaller) GetStableRateExcessOffset(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateStrategy.contract.Call(opts, &out, "getStableRateExcessOffset")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStableRateExcessOffset is a free data retrieval call binding the contract method 0xbc626908.
//
// Solidity: function getStableRateExcessOffset() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategySession) GetStableRateExcessOffset() (*big.Int, error) {
	return _InterestRateStrategy.Contract.GetStableRateExcessOffset(&_InterestRateStrategy.CallOpts)
}

// GetStableRateExcessOffset is a free data retrieval call binding the contract method 0xbc626908.
//
// Solidity: function getStableRateExcessOffset() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategyCallerSession) GetStableRateExcessOffset() (*big.Int, error) {
	return _InterestRateStrategy.Contract.GetStableRateExcessOffset(&_InterestRateStrategy.CallOpts)
}

// GetStableRateSlope1 is a free data retrieval call binding the contract method 0xd5cd7391.
//
// Solidity: function getStableRateSlope1() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategyCaller) GetStableRateSlope1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateStrategy.contract.Call(opts, &out, "getStableRateSlope1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStableRateSlope1 is a free data retrieval call binding the contract method 0xd5cd7391.
//
// Solidity: function getStableRateSlope1() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategySession) GetStableRateSlope1() (*big.Int, error) {
	return _InterestRateStrategy.Contract.GetStableRateSlope1(&_InterestRateStrategy.CallOpts)
}

// GetStableRateSlope1 is a free data retrieval call binding the contract method 0xd5cd7391.
//
// Solidity: function getStableRateSlope1() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategyCallerSession) GetStableRateSlope1() (*big.Int, error) {
	return _InterestRateStrategy.Contract.GetStableRateSlope1(&_InterestRateStrategy.CallOpts)
}

// GetStableRateSlope2 is a free data retrieval call binding the contract method 0x14e32da4.
//
// Solidity: function getStableRateSlope2() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategyCaller) GetStableRateSlope2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateStrategy.contract.Call(opts, &out, "getStableRateSlope2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStableRateSlope2 is a free data retrieval call binding the contract method 0x14e32da4.
//
// Solidity: function getStableRateSlope2() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategySession) GetStableRateSlope2() (*big.Int, error) {
	return _InterestRateStrategy.Contract.GetStableRateSlope2(&_InterestRateStrategy.CallOpts)
}

// GetStableRateSlope2 is a free data retrieval call binding the contract method 0x14e32da4.
//
// Solidity: function getStableRateSlope2() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategyCallerSession) GetStableRateSlope2() (*big.Int, error) {
	return _InterestRateStrategy.Contract.GetStableRateSlope2(&_InterestRateStrategy.CallOpts)
}

// GetVariableRateSlope1 is a free data retrieval call binding the contract method 0x0b3429a2.
//
// Solidity: function getVariableRateSlope1() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategyCaller) GetVariableRateSlope1(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateStrategy.contract.Call(opts, &out, "getVariableRateSlope1")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVariableRateSlope1 is a free data retrieval call binding the contract method 0x0b3429a2.
//
// Solidity: function getVariableRateSlope1() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategySession) GetVariableRateSlope1() (*big.Int, error) {
	return _InterestRateStrategy.Contract.GetVariableRateSlope1(&_InterestRateStrategy.CallOpts)
}

// GetVariableRateSlope1 is a free data retrieval call binding the contract method 0x0b3429a2.
//
// Solidity: function getVariableRateSlope1() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategyCallerSession) GetVariableRateSlope1() (*big.Int, error) {
	return _InterestRateStrategy.Contract.GetVariableRateSlope1(&_InterestRateStrategy.CallOpts)
}

// GetVariableRateSlope2 is a free data retrieval call binding the contract method 0xf4202409.
//
// Solidity: function getVariableRateSlope2() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategyCaller) GetVariableRateSlope2(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _InterestRateStrategy.contract.Call(opts, &out, "getVariableRateSlope2")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVariableRateSlope2 is a free data retrieval call binding the contract method 0xf4202409.
//
// Solidity: function getVariableRateSlope2() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategySession) GetVariableRateSlope2() (*big.Int, error) {
	return _InterestRateStrategy.Contract.GetVariableRateSlope2(&_InterestRateStrategy.CallOpts)
}

// GetVariableRateSlope2 is a free data retrieval call binding the contract method 0xf4202409.
//
// Solidity: function getVariableRateSlope2() view returns(uint256)
func (_InterestRateStrategy *InterestRateStrategyCallerSession) GetVariableRateSlope2() (*big.Int, error) {
	return _InterestRateStrategy.Contract.GetVariableRateSlope2(&_InterestRateStrategy.CallOpts)
}
