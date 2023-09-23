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

// WETHGatewayMetaData contains all meta data concerning the WETHGateway contract.
var WETHGatewayMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"weth\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"contractIPool\",\"name\":\"pool\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRateMode\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"borrowETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"depositETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"emergencyEtherTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"emergencyTokenTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWETHAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rateMode\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"}],\"name\":\"repayETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdrawETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"permitV\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"permitR\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"permitS\",\"type\":\"bytes32\"}],\"name\":\"withdrawETHWithPermit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// WETHGatewayABI is the input ABI used to generate the binding from.
// Deprecated: Use WETHGatewayMetaData.ABI instead.
var WETHGatewayABI = WETHGatewayMetaData.ABI

// WETHGateway is an auto generated Go binding around an Ethereum contract.
type WETHGateway struct {
	WETHGatewayCaller     // Read-only binding to the contract
	WETHGatewayTransactor // Write-only binding to the contract
	WETHGatewayFilterer   // Log filterer for contract events
}

// WETHGatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type WETHGatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WETHGatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WETHGatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WETHGatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WETHGatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WETHGatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WETHGatewaySession struct {
	Contract     *WETHGateway      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WETHGatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WETHGatewayCallerSession struct {
	Contract *WETHGatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// WETHGatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WETHGatewayTransactorSession struct {
	Contract     *WETHGatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// WETHGatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type WETHGatewayRaw struct {
	Contract *WETHGateway // Generic contract binding to access the raw methods on
}

// WETHGatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WETHGatewayCallerRaw struct {
	Contract *WETHGatewayCaller // Generic read-only contract binding to access the raw methods on
}

// WETHGatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WETHGatewayTransactorRaw struct {
	Contract *WETHGatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWETHGateway creates a new instance of WETHGateway, bound to a specific deployed contract.
func NewWETHGateway(address common.Address, backend bind.ContractBackend) (*WETHGateway, error) {
	contract, err := bindWETHGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WETHGateway{WETHGatewayCaller: WETHGatewayCaller{contract: contract}, WETHGatewayTransactor: WETHGatewayTransactor{contract: contract}, WETHGatewayFilterer: WETHGatewayFilterer{contract: contract}}, nil
}

// NewWETHGatewayCaller creates a new read-only instance of WETHGateway, bound to a specific deployed contract.
func NewWETHGatewayCaller(address common.Address, caller bind.ContractCaller) (*WETHGatewayCaller, error) {
	contract, err := bindWETHGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WETHGatewayCaller{contract: contract}, nil
}

// NewWETHGatewayTransactor creates a new write-only instance of WETHGateway, bound to a specific deployed contract.
func NewWETHGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*WETHGatewayTransactor, error) {
	contract, err := bindWETHGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WETHGatewayTransactor{contract: contract}, nil
}

// NewWETHGatewayFilterer creates a new log filterer instance of WETHGateway, bound to a specific deployed contract.
func NewWETHGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*WETHGatewayFilterer, error) {
	contract, err := bindWETHGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WETHGatewayFilterer{contract: contract}, nil
}

// bindWETHGateway binds a generic wrapper to an already deployed contract.
func bindWETHGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WETHGatewayMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WETHGateway *WETHGatewayRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WETHGateway.Contract.WETHGatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WETHGateway *WETHGatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETHGateway.Contract.WETHGatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WETHGateway *WETHGatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WETHGateway.Contract.WETHGatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WETHGateway *WETHGatewayCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WETHGateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WETHGateway *WETHGatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETHGateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WETHGateway *WETHGatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WETHGateway.Contract.contract.Transact(opts, method, params...)
}

// GetWETHAddress is a free data retrieval call binding the contract method 0xaffa8817.
//
// Solidity: function getWETHAddress() view returns(address)
func (_WETHGateway *WETHGatewayCaller) GetWETHAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WETHGateway.contract.Call(opts, &out, "getWETHAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWETHAddress is a free data retrieval call binding the contract method 0xaffa8817.
//
// Solidity: function getWETHAddress() view returns(address)
func (_WETHGateway *WETHGatewaySession) GetWETHAddress() (common.Address, error) {
	return _WETHGateway.Contract.GetWETHAddress(&_WETHGateway.CallOpts)
}

// GetWETHAddress is a free data retrieval call binding the contract method 0xaffa8817.
//
// Solidity: function getWETHAddress() view returns(address)
func (_WETHGateway *WETHGatewayCallerSession) GetWETHAddress() (common.Address, error) {
	return _WETHGateway.Contract.GetWETHAddress(&_WETHGateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WETHGateway *WETHGatewayCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WETHGateway.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WETHGateway *WETHGatewaySession) Owner() (common.Address, error) {
	return _WETHGateway.Contract.Owner(&_WETHGateway.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WETHGateway *WETHGatewayCallerSession) Owner() (common.Address, error) {
	return _WETHGateway.Contract.Owner(&_WETHGateway.CallOpts)
}

// BorrowETH is a paid mutator transaction binding the contract method 0x66514c97.
//
// Solidity: function borrowETH(address , uint256 amount, uint256 interestRateMode, uint16 referralCode) returns()
func (_WETHGateway *WETHGatewayTransactor) BorrowETH(opts *bind.TransactOpts, arg0 common.Address, amount *big.Int, interestRateMode *big.Int, referralCode uint16) (*types.Transaction, error) {
	return _WETHGateway.contract.Transact(opts, "borrowETH", arg0, amount, interestRateMode, referralCode)
}

// BorrowETH is a paid mutator transaction binding the contract method 0x66514c97.
//
// Solidity: function borrowETH(address , uint256 amount, uint256 interestRateMode, uint16 referralCode) returns()
func (_WETHGateway *WETHGatewaySession) BorrowETH(arg0 common.Address, amount *big.Int, interestRateMode *big.Int, referralCode uint16) (*types.Transaction, error) {
	return _WETHGateway.Contract.BorrowETH(&_WETHGateway.TransactOpts, arg0, amount, interestRateMode, referralCode)
}

// BorrowETH is a paid mutator transaction binding the contract method 0x66514c97.
//
// Solidity: function borrowETH(address , uint256 amount, uint256 interestRateMode, uint16 referralCode) returns()
func (_WETHGateway *WETHGatewayTransactorSession) BorrowETH(arg0 common.Address, amount *big.Int, interestRateMode *big.Int, referralCode uint16) (*types.Transaction, error) {
	return _WETHGateway.Contract.BorrowETH(&_WETHGateway.TransactOpts, arg0, amount, interestRateMode, referralCode)
}

// DepositETH is a paid mutator transaction binding the contract method 0x474cf53d.
//
// Solidity: function depositETH(address , address onBehalfOf, uint16 referralCode) payable returns()
func (_WETHGateway *WETHGatewayTransactor) DepositETH(opts *bind.TransactOpts, arg0 common.Address, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _WETHGateway.contract.Transact(opts, "depositETH", arg0, onBehalfOf, referralCode)
}

// DepositETH is a paid mutator transaction binding the contract method 0x474cf53d.
//
// Solidity: function depositETH(address , address onBehalfOf, uint16 referralCode) payable returns()
func (_WETHGateway *WETHGatewaySession) DepositETH(arg0 common.Address, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _WETHGateway.Contract.DepositETH(&_WETHGateway.TransactOpts, arg0, onBehalfOf, referralCode)
}

// DepositETH is a paid mutator transaction binding the contract method 0x474cf53d.
//
// Solidity: function depositETH(address , address onBehalfOf, uint16 referralCode) payable returns()
func (_WETHGateway *WETHGatewayTransactorSession) DepositETH(arg0 common.Address, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _WETHGateway.Contract.DepositETH(&_WETHGateway.TransactOpts, arg0, onBehalfOf, referralCode)
}

// EmergencyEtherTransfer is a paid mutator transaction binding the contract method 0xeed88b8d.
//
// Solidity: function emergencyEtherTransfer(address to, uint256 amount) returns()
func (_WETHGateway *WETHGatewayTransactor) EmergencyEtherTransfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WETHGateway.contract.Transact(opts, "emergencyEtherTransfer", to, amount)
}

// EmergencyEtherTransfer is a paid mutator transaction binding the contract method 0xeed88b8d.
//
// Solidity: function emergencyEtherTransfer(address to, uint256 amount) returns()
func (_WETHGateway *WETHGatewaySession) EmergencyEtherTransfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WETHGateway.Contract.EmergencyEtherTransfer(&_WETHGateway.TransactOpts, to, amount)
}

// EmergencyEtherTransfer is a paid mutator transaction binding the contract method 0xeed88b8d.
//
// Solidity: function emergencyEtherTransfer(address to, uint256 amount) returns()
func (_WETHGateway *WETHGatewayTransactorSession) EmergencyEtherTransfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WETHGateway.Contract.EmergencyEtherTransfer(&_WETHGateway.TransactOpts, to, amount)
}

// EmergencyTokenTransfer is a paid mutator transaction binding the contract method 0xa3d5b255.
//
// Solidity: function emergencyTokenTransfer(address token, address to, uint256 amount) returns()
func (_WETHGateway *WETHGatewayTransactor) EmergencyTokenTransfer(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WETHGateway.contract.Transact(opts, "emergencyTokenTransfer", token, to, amount)
}

// EmergencyTokenTransfer is a paid mutator transaction binding the contract method 0xa3d5b255.
//
// Solidity: function emergencyTokenTransfer(address token, address to, uint256 amount) returns()
func (_WETHGateway *WETHGatewaySession) EmergencyTokenTransfer(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WETHGateway.Contract.EmergencyTokenTransfer(&_WETHGateway.TransactOpts, token, to, amount)
}

// EmergencyTokenTransfer is a paid mutator transaction binding the contract method 0xa3d5b255.
//
// Solidity: function emergencyTokenTransfer(address token, address to, uint256 amount) returns()
func (_WETHGateway *WETHGatewayTransactorSession) EmergencyTokenTransfer(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _WETHGateway.Contract.EmergencyTokenTransfer(&_WETHGateway.TransactOpts, token, to, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WETHGateway *WETHGatewayTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETHGateway.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WETHGateway *WETHGatewaySession) RenounceOwnership() (*types.Transaction, error) {
	return _WETHGateway.Contract.RenounceOwnership(&_WETHGateway.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_WETHGateway *WETHGatewayTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _WETHGateway.Contract.RenounceOwnership(&_WETHGateway.TransactOpts)
}

// RepayETH is a paid mutator transaction binding the contract method 0x02c5fcf8.
//
// Solidity: function repayETH(address , uint256 amount, uint256 rateMode, address onBehalfOf) payable returns()
func (_WETHGateway *WETHGatewayTransactor) RepayETH(opts *bind.TransactOpts, arg0 common.Address, amount *big.Int, rateMode *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _WETHGateway.contract.Transact(opts, "repayETH", arg0, amount, rateMode, onBehalfOf)
}

// RepayETH is a paid mutator transaction binding the contract method 0x02c5fcf8.
//
// Solidity: function repayETH(address , uint256 amount, uint256 rateMode, address onBehalfOf) payable returns()
func (_WETHGateway *WETHGatewaySession) RepayETH(arg0 common.Address, amount *big.Int, rateMode *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _WETHGateway.Contract.RepayETH(&_WETHGateway.TransactOpts, arg0, amount, rateMode, onBehalfOf)
}

// RepayETH is a paid mutator transaction binding the contract method 0x02c5fcf8.
//
// Solidity: function repayETH(address , uint256 amount, uint256 rateMode, address onBehalfOf) payable returns()
func (_WETHGateway *WETHGatewayTransactorSession) RepayETH(arg0 common.Address, amount *big.Int, rateMode *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _WETHGateway.Contract.RepayETH(&_WETHGateway.TransactOpts, arg0, amount, rateMode, onBehalfOf)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WETHGateway *WETHGatewayTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _WETHGateway.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WETHGateway *WETHGatewaySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WETHGateway.Contract.TransferOwnership(&_WETHGateway.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_WETHGateway *WETHGatewayTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _WETHGateway.Contract.TransferOwnership(&_WETHGateway.TransactOpts, newOwner)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x80500d20.
//
// Solidity: function withdrawETH(address , uint256 amount, address to) returns()
func (_WETHGateway *WETHGatewayTransactor) WithdrawETH(opts *bind.TransactOpts, arg0 common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _WETHGateway.contract.Transact(opts, "withdrawETH", arg0, amount, to)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x80500d20.
//
// Solidity: function withdrawETH(address , uint256 amount, address to) returns()
func (_WETHGateway *WETHGatewaySession) WithdrawETH(arg0 common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _WETHGateway.Contract.WithdrawETH(&_WETHGateway.TransactOpts, arg0, amount, to)
}

// WithdrawETH is a paid mutator transaction binding the contract method 0x80500d20.
//
// Solidity: function withdrawETH(address , uint256 amount, address to) returns()
func (_WETHGateway *WETHGatewayTransactorSession) WithdrawETH(arg0 common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _WETHGateway.Contract.WithdrawETH(&_WETHGateway.TransactOpts, arg0, amount, to)
}

// WithdrawETHWithPermit is a paid mutator transaction binding the contract method 0xd4c40b6c.
//
// Solidity: function withdrawETHWithPermit(address , uint256 amount, address to, uint256 deadline, uint8 permitV, bytes32 permitR, bytes32 permitS) returns()
func (_WETHGateway *WETHGatewayTransactor) WithdrawETHWithPermit(opts *bind.TransactOpts, arg0 common.Address, amount *big.Int, to common.Address, deadline *big.Int, permitV uint8, permitR [32]byte, permitS [32]byte) (*types.Transaction, error) {
	return _WETHGateway.contract.Transact(opts, "withdrawETHWithPermit", arg0, amount, to, deadline, permitV, permitR, permitS)
}

// WithdrawETHWithPermit is a paid mutator transaction binding the contract method 0xd4c40b6c.
//
// Solidity: function withdrawETHWithPermit(address , uint256 amount, address to, uint256 deadline, uint8 permitV, bytes32 permitR, bytes32 permitS) returns()
func (_WETHGateway *WETHGatewaySession) WithdrawETHWithPermit(arg0 common.Address, amount *big.Int, to common.Address, deadline *big.Int, permitV uint8, permitR [32]byte, permitS [32]byte) (*types.Transaction, error) {
	return _WETHGateway.Contract.WithdrawETHWithPermit(&_WETHGateway.TransactOpts, arg0, amount, to, deadline, permitV, permitR, permitS)
}

// WithdrawETHWithPermit is a paid mutator transaction binding the contract method 0xd4c40b6c.
//
// Solidity: function withdrawETHWithPermit(address , uint256 amount, address to, uint256 deadline, uint8 permitV, bytes32 permitR, bytes32 permitS) returns()
func (_WETHGateway *WETHGatewayTransactorSession) WithdrawETHWithPermit(arg0 common.Address, amount *big.Int, to common.Address, deadline *big.Int, permitV uint8, permitR [32]byte, permitS [32]byte) (*types.Transaction, error) {
	return _WETHGateway.Contract.WithdrawETHWithPermit(&_WETHGateway.TransactOpts, arg0, amount, to, deadline, permitV, permitR, permitS)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WETHGateway *WETHGatewayTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _WETHGateway.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WETHGateway *WETHGatewaySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WETHGateway.Contract.Fallback(&_WETHGateway.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_WETHGateway *WETHGatewayTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _WETHGateway.Contract.Fallback(&_WETHGateway.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WETHGateway *WETHGatewayTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETHGateway.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WETHGateway *WETHGatewaySession) Receive() (*types.Transaction, error) {
	return _WETHGateway.Contract.Receive(&_WETHGateway.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WETHGateway *WETHGatewayTransactorSession) Receive() (*types.Transaction, error) {
	return _WETHGateway.Contract.Receive(&_WETHGateway.TransactOpts)
}

// WETHGatewayOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the WETHGateway contract.
type WETHGatewayOwnershipTransferredIterator struct {
	Event *WETHGatewayOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *WETHGatewayOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WETHGatewayOwnershipTransferred)
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
		it.Event = new(WETHGatewayOwnershipTransferred)
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
func (it *WETHGatewayOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WETHGatewayOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WETHGatewayOwnershipTransferred represents a OwnershipTransferred event raised by the WETHGateway contract.
type WETHGatewayOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WETHGateway *WETHGatewayFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*WETHGatewayOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WETHGateway.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &WETHGatewayOwnershipTransferredIterator{contract: _WETHGateway.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_WETHGateway *WETHGatewayFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *WETHGatewayOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _WETHGateway.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WETHGatewayOwnershipTransferred)
				if err := _WETHGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_WETHGateway *WETHGatewayFilterer) ParseOwnershipTransferred(log types.Log) (*WETHGatewayOwnershipTransferred, error) {
	event := new(WETHGatewayOwnershipTransferred)
	if err := _WETHGateway.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
