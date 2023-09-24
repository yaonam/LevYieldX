// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package synapse

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

// BridgeToken is an auto generated low-level Go binding around an user-defined struct.
type BridgeToken struct {
	Symbol string
	Token  common.Address
}

// DestRequest is an auto generated low-level Go binding around an user-defined struct.
type DestRequest struct {
	Symbol   string
	AmountIn *big.Int
}

// LocalBridgeConfigBridgeTokenConfig is an auto generated low-level Go binding around an user-defined struct.
type LocalBridgeConfigBridgeTokenConfig struct {
	Id          string
	Token       common.Address
	Decimals    *big.Int
	TokenType   uint8
	BridgeToken common.Address
	BridgeFee   *big.Int
	MinFee      *big.Int
	MaxFee      *big.Int
}

// MulticallViewResult is an auto generated low-level Go binding around an user-defined struct.
type MulticallViewResult struct {
	Success    bool
	ReturnData []byte
}

// Pool is an auto generated low-level Go binding around an user-defined struct.
type Pool struct {
	Pool    common.Address
	LpToken common.Address
	Tokens  []PoolToken
}

// PoolToken is an auto generated low-level Go binding around an user-defined struct.
type PoolToken struct {
	IsWeth bool
	Token  common.Address
}

// SwapQuery is an auto generated low-level Go binding around an user-defined struct.
type SwapQuery struct {
	SwapAdapter  common.Address
	TokenOut     common.Address
	MinAmountOut *big.Int
	Deadline     *big.Int
	RawParams    []byte
}

// RouterMetaData contains all meta data concerning the Router contract.
var RouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_synapseBridge\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner_\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"rawParams\",\"type\":\"bytes\"}],\"name\":\"adapterSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"enumLocalBridgeConfig.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"bridgeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bridgeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"}],\"name\":\"addToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasAdded\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"id\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"enumLocalBridgeConfig.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"bridgeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bridgeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"}],\"internalType\":\"structLocalBridgeConfig.BridgeTokenConfig[]\",\"name\":\"tokens\",\"type\":\"tuple[]\"}],\"name\":\"addTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allPools\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"lpToken\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isWeth\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structPoolToken[]\",\"name\":\"tokens\",\"type\":\"tuple[]\"}],\"internalType\":\"structPool[]\",\"name\":\"pools\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"swapAdapter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rawParams\",\"type\":\"bytes\"}],\"internalType\":\"structSwapQuery\",\"name\":\"originQuery\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"swapAdapter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rawParams\",\"type\":\"bytes\"}],\"internalType\":\"structSwapQuery\",\"name\":\"destQuery\",\"type\":\"tuple\"}],\"name\":\"bridge\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bridgeTokensAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"calculateAddLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"calculateBridgeFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"calculateRemoveLiquidity\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amountsOut\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexFrom\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndexTo\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"dx\",\"type\":\"uint256\"}],\"name\":\"calculateSwap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"tokenIndex\",\"type\":\"uint8\"}],\"name\":\"calculateWithdrawOneToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"config\",\"outputs\":[{\"internalType\":\"enumLocalBridgeConfig.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"bridgeToken\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"bridgeFee\",\"type\":\"uint40\"},{\"internalType\":\"uint104\",\"name\":\"minFee\",\"type\":\"uint104\"},{\"internalType\":\"uint112\",\"name\":\"maxFee\",\"type\":\"uint112\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"getAmountOut\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"swapAdapter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rawParams\",\"type\":\"bytes\"}],\"internalType\":\"structSwapQuery\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"getConnectedBridgeTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structBridgeToken[]\",\"name\":\"tokens\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"internalType\":\"structDestRequest[]\",\"name\":\"requests\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"getDestinationAmountOut\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"swapAdapter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rawParams\",\"type\":\"bytes\"}],\"internalType\":\"structSwapQuery[]\",\"name\":\"destQueries\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"string[]\",\"name\":\"tokenSymbols\",\"type\":\"string[]\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"}],\"name\":\"getOriginAmountOut\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"swapAdapter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rawParams\",\"type\":\"bytes\"}],\"internalType\":\"structSwapQuery[]\",\"name\":\"originQueries\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicallView\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"internalType\":\"structMulticallView.Result[]\",\"name\":\"callResults\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"poolInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"poolTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isWeth\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"internalType\":\"structPoolToken[]\",\"name\":\"tokens\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolsAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"removeToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"wasRemoved\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"name\":\"removeTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractISwapQuoter\",\"name\":\"_swapQuoter\",\"type\":\"address\"}],\"name\":\"setSwapQuoter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"enumLocalBridgeConfig.TokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"bridgeToken\",\"type\":\"address\"}],\"name\":\"setTokenConfig\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bridgeFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"}],\"name\":\"setTokenFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"swapAdapter\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"rawParams\",\"type\":\"bytes\"}],\"internalType\":\"structSwapQuery\",\"name\":\"query\",\"type\":\"tuple\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapQuoter\",\"outputs\":[{\"internalType\":\"contractISwapQuoter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"symbolToToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"synapseBridge\",\"outputs\":[{\"internalType\":\"contractISynapseBridge\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenToSymbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// RouterABI is the input ABI used to generate the binding from.
// Deprecated: Use RouterMetaData.ABI instead.
var RouterABI = RouterMetaData.ABI

// Router is an auto generated Go binding around an Ethereum contract.
type Router struct {
	RouterCaller     // Read-only binding to the contract
	RouterTransactor // Write-only binding to the contract
	RouterFilterer   // Log filterer for contract events
}

// RouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type RouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RouterSession struct {
	Contract     *Router           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RouterCallerSession struct {
	Contract *RouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RouterTransactorSession struct {
	Contract     *RouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type RouterRaw struct {
	Contract *Router // Generic contract binding to access the raw methods on
}

// RouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RouterCallerRaw struct {
	Contract *RouterCaller // Generic read-only contract binding to access the raw methods on
}

// RouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RouterTransactorRaw struct {
	Contract *RouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRouter creates a new instance of Router, bound to a specific deployed contract.
func NewRouter(address common.Address, backend bind.ContractBackend) (*Router, error) {
	contract, err := bindRouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Router{RouterCaller: RouterCaller{contract: contract}, RouterTransactor: RouterTransactor{contract: contract}, RouterFilterer: RouterFilterer{contract: contract}}, nil
}

// NewRouterCaller creates a new read-only instance of Router, bound to a specific deployed contract.
func NewRouterCaller(address common.Address, caller bind.ContractCaller) (*RouterCaller, error) {
	contract, err := bindRouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RouterCaller{contract: contract}, nil
}

// NewRouterTransactor creates a new write-only instance of Router, bound to a specific deployed contract.
func NewRouterTransactor(address common.Address, transactor bind.ContractTransactor) (*RouterTransactor, error) {
	contract, err := bindRouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RouterTransactor{contract: contract}, nil
}

// NewRouterFilterer creates a new log filterer instance of Router, bound to a specific deployed contract.
func NewRouterFilterer(address common.Address, filterer bind.ContractFilterer) (*RouterFilterer, error) {
	contract, err := bindRouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RouterFilterer{contract: contract}, nil
}

// bindRouter binds a generic wrapper to an already deployed contract.
func bindRouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := RouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.RouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.RouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Router *RouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Router.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Router *RouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Router *RouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Router.Contract.contract.Transact(opts, method, params...)
}

// AllPools is a free data retrieval call binding the contract method 0xc5c63e65.
//
// Solidity: function allPools() view returns((address,address,(bool,address)[])[] pools)
func (_Router *RouterCaller) AllPools(opts *bind.CallOpts) ([]Pool, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "allPools")

	if err != nil {
		return *new([]Pool), err
	}

	out0 := *abi.ConvertType(out[0], new([]Pool)).(*[]Pool)

	return out0, err

}

// AllPools is a free data retrieval call binding the contract method 0xc5c63e65.
//
// Solidity: function allPools() view returns((address,address,(bool,address)[])[] pools)
func (_Router *RouterSession) AllPools() ([]Pool, error) {
	return _Router.Contract.AllPools(&_Router.CallOpts)
}

// AllPools is a free data retrieval call binding the contract method 0xc5c63e65.
//
// Solidity: function allPools() view returns((address,address,(bool,address)[])[] pools)
func (_Router *RouterCallerSession) AllPools() ([]Pool, error) {
	return _Router.Contract.AllPools(&_Router.CallOpts)
}

// BridgeTokens is a free data retrieval call binding the contract method 0xf8a06888.
//
// Solidity: function bridgeTokens() view returns(address[] tokens)
func (_Router *RouterCaller) BridgeTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "bridgeTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// BridgeTokens is a free data retrieval call binding the contract method 0xf8a06888.
//
// Solidity: function bridgeTokens() view returns(address[] tokens)
func (_Router *RouterSession) BridgeTokens() ([]common.Address, error) {
	return _Router.Contract.BridgeTokens(&_Router.CallOpts)
}

// BridgeTokens is a free data retrieval call binding the contract method 0xf8a06888.
//
// Solidity: function bridgeTokens() view returns(address[] tokens)
func (_Router *RouterCallerSession) BridgeTokens() ([]common.Address, error) {
	return _Router.Contract.BridgeTokens(&_Router.CallOpts)
}

// BridgeTokensAmount is a free data retrieval call binding the contract method 0x5ff30c0a.
//
// Solidity: function bridgeTokensAmount() view returns(uint256 amount)
func (_Router *RouterCaller) BridgeTokensAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "bridgeTokensAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BridgeTokensAmount is a free data retrieval call binding the contract method 0x5ff30c0a.
//
// Solidity: function bridgeTokensAmount() view returns(uint256 amount)
func (_Router *RouterSession) BridgeTokensAmount() (*big.Int, error) {
	return _Router.Contract.BridgeTokensAmount(&_Router.CallOpts)
}

// BridgeTokensAmount is a free data retrieval call binding the contract method 0x5ff30c0a.
//
// Solidity: function bridgeTokensAmount() view returns(uint256 amount)
func (_Router *RouterCallerSession) BridgeTokensAmount() (*big.Int, error) {
	return _Router.Contract.BridgeTokensAmount(&_Router.CallOpts)
}

// CalculateAddLiquidity is a free data retrieval call binding the contract method 0x4d864496.
//
// Solidity: function calculateAddLiquidity(address pool, uint256[] amounts) view returns(uint256)
func (_Router *RouterCaller) CalculateAddLiquidity(opts *bind.CallOpts, pool common.Address, amounts []*big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "calculateAddLiquidity", pool, amounts)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateAddLiquidity is a free data retrieval call binding the contract method 0x4d864496.
//
// Solidity: function calculateAddLiquidity(address pool, uint256[] amounts) view returns(uint256)
func (_Router *RouterSession) CalculateAddLiquidity(pool common.Address, amounts []*big.Int) (*big.Int, error) {
	return _Router.Contract.CalculateAddLiquidity(&_Router.CallOpts, pool, amounts)
}

// CalculateAddLiquidity is a free data retrieval call binding the contract method 0x4d864496.
//
// Solidity: function calculateAddLiquidity(address pool, uint256[] amounts) view returns(uint256)
func (_Router *RouterCallerSession) CalculateAddLiquidity(pool common.Address, amounts []*big.Int) (*big.Int, error) {
	return _Router.Contract.CalculateAddLiquidity(&_Router.CallOpts, pool, amounts)
}

// CalculateBridgeFee is a free data retrieval call binding the contract method 0x58b5b777.
//
// Solidity: function calculateBridgeFee(address token, uint256 amount) view returns(uint256 feeAmount)
func (_Router *RouterCaller) CalculateBridgeFee(opts *bind.CallOpts, token common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "calculateBridgeFee", token, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateBridgeFee is a free data retrieval call binding the contract method 0x58b5b777.
//
// Solidity: function calculateBridgeFee(address token, uint256 amount) view returns(uint256 feeAmount)
func (_Router *RouterSession) CalculateBridgeFee(token common.Address, amount *big.Int) (*big.Int, error) {
	return _Router.Contract.CalculateBridgeFee(&_Router.CallOpts, token, amount)
}

// CalculateBridgeFee is a free data retrieval call binding the contract method 0x58b5b777.
//
// Solidity: function calculateBridgeFee(address token, uint256 amount) view returns(uint256 feeAmount)
func (_Router *RouterCallerSession) CalculateBridgeFee(token common.Address, amount *big.Int) (*big.Int, error) {
	return _Router.Contract.CalculateBridgeFee(&_Router.CallOpts, token, amount)
}

// CalculateRemoveLiquidity is a free data retrieval call binding the contract method 0x7c61e561.
//
// Solidity: function calculateRemoveLiquidity(address pool, uint256 amount) view returns(uint256[] amountsOut)
func (_Router *RouterCaller) CalculateRemoveLiquidity(opts *bind.CallOpts, pool common.Address, amount *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "calculateRemoveLiquidity", pool, amount)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// CalculateRemoveLiquidity is a free data retrieval call binding the contract method 0x7c61e561.
//
// Solidity: function calculateRemoveLiquidity(address pool, uint256 amount) view returns(uint256[] amountsOut)
func (_Router *RouterSession) CalculateRemoveLiquidity(pool common.Address, amount *big.Int) ([]*big.Int, error) {
	return _Router.Contract.CalculateRemoveLiquidity(&_Router.CallOpts, pool, amount)
}

// CalculateRemoveLiquidity is a free data retrieval call binding the contract method 0x7c61e561.
//
// Solidity: function calculateRemoveLiquidity(address pool, uint256 amount) view returns(uint256[] amountsOut)
func (_Router *RouterCallerSession) CalculateRemoveLiquidity(pool common.Address, amount *big.Int) ([]*big.Int, error) {
	return _Router.Contract.CalculateRemoveLiquidity(&_Router.CallOpts, pool, amount)
}

// CalculateSwap is a free data retrieval call binding the contract method 0x798af720.
//
// Solidity: function calculateSwap(address pool, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx) view returns(uint256 amountOut)
func (_Router *RouterCaller) CalculateSwap(opts *bind.CallOpts, pool common.Address, tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "calculateSwap", pool, tokenIndexFrom, tokenIndexTo, dx)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateSwap is a free data retrieval call binding the contract method 0x798af720.
//
// Solidity: function calculateSwap(address pool, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx) view returns(uint256 amountOut)
func (_Router *RouterSession) CalculateSwap(pool common.Address, tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int) (*big.Int, error) {
	return _Router.Contract.CalculateSwap(&_Router.CallOpts, pool, tokenIndexFrom, tokenIndexTo, dx)
}

// CalculateSwap is a free data retrieval call binding the contract method 0x798af720.
//
// Solidity: function calculateSwap(address pool, uint8 tokenIndexFrom, uint8 tokenIndexTo, uint256 dx) view returns(uint256 amountOut)
func (_Router *RouterCallerSession) CalculateSwap(pool common.Address, tokenIndexFrom uint8, tokenIndexTo uint8, dx *big.Int) (*big.Int, error) {
	return _Router.Contract.CalculateSwap(&_Router.CallOpts, pool, tokenIndexFrom, tokenIndexTo, dx)
}

// CalculateWithdrawOneToken is a free data retrieval call binding the contract method 0xccc1bbc1.
//
// Solidity: function calculateWithdrawOneToken(address pool, uint256 tokenAmount, uint8 tokenIndex) view returns(uint256 amountOut)
func (_Router *RouterCaller) CalculateWithdrawOneToken(opts *bind.CallOpts, pool common.Address, tokenAmount *big.Int, tokenIndex uint8) (*big.Int, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "calculateWithdrawOneToken", pool, tokenAmount, tokenIndex)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalculateWithdrawOneToken is a free data retrieval call binding the contract method 0xccc1bbc1.
//
// Solidity: function calculateWithdrawOneToken(address pool, uint256 tokenAmount, uint8 tokenIndex) view returns(uint256 amountOut)
func (_Router *RouterSession) CalculateWithdrawOneToken(pool common.Address, tokenAmount *big.Int, tokenIndex uint8) (*big.Int, error) {
	return _Router.Contract.CalculateWithdrawOneToken(&_Router.CallOpts, pool, tokenAmount, tokenIndex)
}

// CalculateWithdrawOneToken is a free data retrieval call binding the contract method 0xccc1bbc1.
//
// Solidity: function calculateWithdrawOneToken(address pool, uint256 tokenAmount, uint8 tokenIndex) view returns(uint256 amountOut)
func (_Router *RouterCallerSession) CalculateWithdrawOneToken(pool common.Address, tokenAmount *big.Int, tokenIndex uint8) (*big.Int, error) {
	return _Router.Contract.CalculateWithdrawOneToken(&_Router.CallOpts, pool, tokenAmount, tokenIndex)
}

// Config is a free data retrieval call binding the contract method 0x0e68ec95.
//
// Solidity: function config(address ) view returns(uint8 tokenType, address bridgeToken)
func (_Router *RouterCaller) Config(opts *bind.CallOpts, arg0 common.Address) (struct {
	TokenType   uint8
	BridgeToken common.Address
}, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "config", arg0)

	outstruct := new(struct {
		TokenType   uint8
		BridgeToken common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenType = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.BridgeToken = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Config is a free data retrieval call binding the contract method 0x0e68ec95.
//
// Solidity: function config(address ) view returns(uint8 tokenType, address bridgeToken)
func (_Router *RouterSession) Config(arg0 common.Address) (struct {
	TokenType   uint8
	BridgeToken common.Address
}, error) {
	return _Router.Contract.Config(&_Router.CallOpts, arg0)
}

// Config is a free data retrieval call binding the contract method 0x0e68ec95.
//
// Solidity: function config(address ) view returns(uint8 tokenType, address bridgeToken)
func (_Router *RouterCallerSession) Config(arg0 common.Address) (struct {
	TokenType   uint8
	BridgeToken common.Address
}, error) {
	return _Router.Contract.Config(&_Router.CallOpts, arg0)
}

// Fee is a free data retrieval call binding the contract method 0x6fcca69b.
//
// Solidity: function fee(address ) view returns(uint40 bridgeFee, uint104 minFee, uint112 maxFee)
func (_Router *RouterCaller) Fee(opts *bind.CallOpts, arg0 common.Address) (struct {
	BridgeFee *big.Int
	MinFee    *big.Int
	MaxFee    *big.Int
}, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "fee", arg0)

	outstruct := new(struct {
		BridgeFee *big.Int
		MinFee    *big.Int
		MaxFee    *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BridgeFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.MinFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.MaxFee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Fee is a free data retrieval call binding the contract method 0x6fcca69b.
//
// Solidity: function fee(address ) view returns(uint40 bridgeFee, uint104 minFee, uint112 maxFee)
func (_Router *RouterSession) Fee(arg0 common.Address) (struct {
	BridgeFee *big.Int
	MinFee    *big.Int
	MaxFee    *big.Int
}, error) {
	return _Router.Contract.Fee(&_Router.CallOpts, arg0)
}

// Fee is a free data retrieval call binding the contract method 0x6fcca69b.
//
// Solidity: function fee(address ) view returns(uint40 bridgeFee, uint104 minFee, uint112 maxFee)
func (_Router *RouterCallerSession) Fee(arg0 common.Address) (struct {
	BridgeFee *big.Int
	MinFee    *big.Int
	MaxFee    *big.Int
}, error) {
	return _Router.Contract.Fee(&_Router.CallOpts, arg0)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x4aa06652.
//
// Solidity: function getAmountOut(address tokenIn, address tokenOut, uint256 amountIn) view returns((address,address,uint256,uint256,bytes))
func (_Router *RouterCaller) GetAmountOut(opts *bind.CallOpts, tokenIn common.Address, tokenOut common.Address, amountIn *big.Int) (SwapQuery, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getAmountOut", tokenIn, tokenOut, amountIn)

	if err != nil {
		return *new(SwapQuery), err
	}

	out0 := *abi.ConvertType(out[0], new(SwapQuery)).(*SwapQuery)

	return out0, err

}

// GetAmountOut is a free data retrieval call binding the contract method 0x4aa06652.
//
// Solidity: function getAmountOut(address tokenIn, address tokenOut, uint256 amountIn) view returns((address,address,uint256,uint256,bytes))
func (_Router *RouterSession) GetAmountOut(tokenIn common.Address, tokenOut common.Address, amountIn *big.Int) (SwapQuery, error) {
	return _Router.Contract.GetAmountOut(&_Router.CallOpts, tokenIn, tokenOut, amountIn)
}

// GetAmountOut is a free data retrieval call binding the contract method 0x4aa06652.
//
// Solidity: function getAmountOut(address tokenIn, address tokenOut, uint256 amountIn) view returns((address,address,uint256,uint256,bytes))
func (_Router *RouterCallerSession) GetAmountOut(tokenIn common.Address, tokenOut common.Address, amountIn *big.Int) (SwapQuery, error) {
	return _Router.Contract.GetAmountOut(&_Router.CallOpts, tokenIn, tokenOut, amountIn)
}

// GetConnectedBridgeTokens is a free data retrieval call binding the contract method 0xd38e7888.
//
// Solidity: function getConnectedBridgeTokens(address tokenOut) view returns((string,address)[] tokens)
func (_Router *RouterCaller) GetConnectedBridgeTokens(opts *bind.CallOpts, tokenOut common.Address) ([]BridgeToken, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getConnectedBridgeTokens", tokenOut)

	if err != nil {
		return *new([]BridgeToken), err
	}

	out0 := *abi.ConvertType(out[0], new([]BridgeToken)).(*[]BridgeToken)

	return out0, err

}

// GetConnectedBridgeTokens is a free data retrieval call binding the contract method 0xd38e7888.
//
// Solidity: function getConnectedBridgeTokens(address tokenOut) view returns((string,address)[] tokens)
func (_Router *RouterSession) GetConnectedBridgeTokens(tokenOut common.Address) ([]BridgeToken, error) {
	return _Router.Contract.GetConnectedBridgeTokens(&_Router.CallOpts, tokenOut)
}

// GetConnectedBridgeTokens is a free data retrieval call binding the contract method 0xd38e7888.
//
// Solidity: function getConnectedBridgeTokens(address tokenOut) view returns((string,address)[] tokens)
func (_Router *RouterCallerSession) GetConnectedBridgeTokens(tokenOut common.Address) ([]BridgeToken, error) {
	return _Router.Contract.GetConnectedBridgeTokens(&_Router.CallOpts, tokenOut)
}

// GetDestinationAmountOut is a free data retrieval call binding the contract method 0x077e1199.
//
// Solidity: function getDestinationAmountOut((string,uint256)[] requests, address tokenOut) view returns((address,address,uint256,uint256,bytes)[] destQueries)
func (_Router *RouterCaller) GetDestinationAmountOut(opts *bind.CallOpts, requests []DestRequest, tokenOut common.Address) ([]SwapQuery, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getDestinationAmountOut", requests, tokenOut)

	if err != nil {
		return *new([]SwapQuery), err
	}

	out0 := *abi.ConvertType(out[0], new([]SwapQuery)).(*[]SwapQuery)

	return out0, err

}

// GetDestinationAmountOut is a free data retrieval call binding the contract method 0x077e1199.
//
// Solidity: function getDestinationAmountOut((string,uint256)[] requests, address tokenOut) view returns((address,address,uint256,uint256,bytes)[] destQueries)
func (_Router *RouterSession) GetDestinationAmountOut(requests []DestRequest, tokenOut common.Address) ([]SwapQuery, error) {
	return _Router.Contract.GetDestinationAmountOut(&_Router.CallOpts, requests, tokenOut)
}

// GetDestinationAmountOut is a free data retrieval call binding the contract method 0x077e1199.
//
// Solidity: function getDestinationAmountOut((string,uint256)[] requests, address tokenOut) view returns((address,address,uint256,uint256,bytes)[] destQueries)
func (_Router *RouterCallerSession) GetDestinationAmountOut(requests []DestRequest, tokenOut common.Address) ([]SwapQuery, error) {
	return _Router.Contract.GetDestinationAmountOut(&_Router.CallOpts, requests, tokenOut)
}

// GetOriginAmountOut is a free data retrieval call binding the contract method 0x3bc758fd.
//
// Solidity: function getOriginAmountOut(address tokenIn, string[] tokenSymbols, uint256 amountIn) view returns((address,address,uint256,uint256,bytes)[] originQueries)
func (_Router *RouterCaller) GetOriginAmountOut(opts *bind.CallOpts, tokenIn common.Address, tokenSymbols []string, amountIn *big.Int) ([]SwapQuery, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "getOriginAmountOut", tokenIn, tokenSymbols, amountIn)

	if err != nil {
		return *new([]SwapQuery), err
	}

	out0 := *abi.ConvertType(out[0], new([]SwapQuery)).(*[]SwapQuery)

	return out0, err

}

// GetOriginAmountOut is a free data retrieval call binding the contract method 0x3bc758fd.
//
// Solidity: function getOriginAmountOut(address tokenIn, string[] tokenSymbols, uint256 amountIn) view returns((address,address,uint256,uint256,bytes)[] originQueries)
func (_Router *RouterSession) GetOriginAmountOut(tokenIn common.Address, tokenSymbols []string, amountIn *big.Int) ([]SwapQuery, error) {
	return _Router.Contract.GetOriginAmountOut(&_Router.CallOpts, tokenIn, tokenSymbols, amountIn)
}

// GetOriginAmountOut is a free data retrieval call binding the contract method 0x3bc758fd.
//
// Solidity: function getOriginAmountOut(address tokenIn, string[] tokenSymbols, uint256 amountIn) view returns((address,address,uint256,uint256,bytes)[] originQueries)
func (_Router *RouterCallerSession) GetOriginAmountOut(tokenIn common.Address, tokenSymbols []string, amountIn *big.Int) ([]SwapQuery, error) {
	return _Router.Contract.GetOriginAmountOut(&_Router.CallOpts, tokenIn, tokenSymbols, amountIn)
}

// MulticallView is a free data retrieval call binding the contract method 0xa734c441.
//
// Solidity: function multicallView(bytes[] data) view returns((bool,bytes)[] callResults)
func (_Router *RouterCaller) MulticallView(opts *bind.CallOpts, data [][]byte) ([]MulticallViewResult, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "multicallView", data)

	if err != nil {
		return *new([]MulticallViewResult), err
	}

	out0 := *abi.ConvertType(out[0], new([]MulticallViewResult)).(*[]MulticallViewResult)

	return out0, err

}

// MulticallView is a free data retrieval call binding the contract method 0xa734c441.
//
// Solidity: function multicallView(bytes[] data) view returns((bool,bytes)[] callResults)
func (_Router *RouterSession) MulticallView(data [][]byte) ([]MulticallViewResult, error) {
	return _Router.Contract.MulticallView(&_Router.CallOpts, data)
}

// MulticallView is a free data retrieval call binding the contract method 0xa734c441.
//
// Solidity: function multicallView(bytes[] data) view returns((bool,bytes)[] callResults)
func (_Router *RouterCallerSession) MulticallView(data [][]byte) ([]MulticallViewResult, error) {
	return _Router.Contract.MulticallView(&_Router.CallOpts, data)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Router *RouterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Router *RouterSession) Owner() (common.Address, error) {
	return _Router.Contract.Owner(&_Router.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Router *RouterCallerSession) Owner() (common.Address, error) {
	return _Router.Contract.Owner(&_Router.CallOpts)
}

// PoolInfo is a free data retrieval call binding the contract method 0x9a7b5f11.
//
// Solidity: function poolInfo(address pool) view returns(uint256, address)
func (_Router *RouterCaller) PoolInfo(opts *bind.CallOpts, pool common.Address) (*big.Int, common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "poolInfo", pool)

	if err != nil {
		return *new(*big.Int), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// PoolInfo is a free data retrieval call binding the contract method 0x9a7b5f11.
//
// Solidity: function poolInfo(address pool) view returns(uint256, address)
func (_Router *RouterSession) PoolInfo(pool common.Address) (*big.Int, common.Address, error) {
	return _Router.Contract.PoolInfo(&_Router.CallOpts, pool)
}

// PoolInfo is a free data retrieval call binding the contract method 0x9a7b5f11.
//
// Solidity: function poolInfo(address pool) view returns(uint256, address)
func (_Router *RouterCallerSession) PoolInfo(pool common.Address) (*big.Int, common.Address, error) {
	return _Router.Contract.PoolInfo(&_Router.CallOpts, pool)
}

// PoolTokens is a free data retrieval call binding the contract method 0xa9126169.
//
// Solidity: function poolTokens(address pool) view returns((bool,address)[] tokens)
func (_Router *RouterCaller) PoolTokens(opts *bind.CallOpts, pool common.Address) ([]PoolToken, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "poolTokens", pool)

	if err != nil {
		return *new([]PoolToken), err
	}

	out0 := *abi.ConvertType(out[0], new([]PoolToken)).(*[]PoolToken)

	return out0, err

}

// PoolTokens is a free data retrieval call binding the contract method 0xa9126169.
//
// Solidity: function poolTokens(address pool) view returns((bool,address)[] tokens)
func (_Router *RouterSession) PoolTokens(pool common.Address) ([]PoolToken, error) {
	return _Router.Contract.PoolTokens(&_Router.CallOpts, pool)
}

// PoolTokens is a free data retrieval call binding the contract method 0xa9126169.
//
// Solidity: function poolTokens(address pool) view returns((bool,address)[] tokens)
func (_Router *RouterCallerSession) PoolTokens(pool common.Address) ([]PoolToken, error) {
	return _Router.Contract.PoolTokens(&_Router.CallOpts, pool)
}

// PoolsAmount is a free data retrieval call binding the contract method 0xba7d536e.
//
// Solidity: function poolsAmount() view returns(uint256 amount)
func (_Router *RouterCaller) PoolsAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "poolsAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PoolsAmount is a free data retrieval call binding the contract method 0xba7d536e.
//
// Solidity: function poolsAmount() view returns(uint256 amount)
func (_Router *RouterSession) PoolsAmount() (*big.Int, error) {
	return _Router.Contract.PoolsAmount(&_Router.CallOpts)
}

// PoolsAmount is a free data retrieval call binding the contract method 0xba7d536e.
//
// Solidity: function poolsAmount() view returns(uint256 amount)
func (_Router *RouterCallerSession) PoolsAmount() (*big.Int, error) {
	return _Router.Contract.PoolsAmount(&_Router.CallOpts)
}

// SwapQuoter is a free data retrieval call binding the contract method 0x34474c8c.
//
// Solidity: function swapQuoter() view returns(address)
func (_Router *RouterCaller) SwapQuoter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "swapQuoter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapQuoter is a free data retrieval call binding the contract method 0x34474c8c.
//
// Solidity: function swapQuoter() view returns(address)
func (_Router *RouterSession) SwapQuoter() (common.Address, error) {
	return _Router.Contract.SwapQuoter(&_Router.CallOpts)
}

// SwapQuoter is a free data retrieval call binding the contract method 0x34474c8c.
//
// Solidity: function swapQuoter() view returns(address)
func (_Router *RouterCallerSession) SwapQuoter() (common.Address, error) {
	return _Router.Contract.SwapQuoter(&_Router.CallOpts)
}

// SymbolToToken is a free data retrieval call binding the contract method 0xa5bc29c2.
//
// Solidity: function symbolToToken(string ) view returns(address)
func (_Router *RouterCaller) SymbolToToken(opts *bind.CallOpts, arg0 string) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "symbolToToken", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SymbolToToken is a free data retrieval call binding the contract method 0xa5bc29c2.
//
// Solidity: function symbolToToken(string ) view returns(address)
func (_Router *RouterSession) SymbolToToken(arg0 string) (common.Address, error) {
	return _Router.Contract.SymbolToToken(&_Router.CallOpts, arg0)
}

// SymbolToToken is a free data retrieval call binding the contract method 0xa5bc29c2.
//
// Solidity: function symbolToToken(string ) view returns(address)
func (_Router *RouterCallerSession) SymbolToToken(arg0 string) (common.Address, error) {
	return _Router.Contract.SymbolToToken(&_Router.CallOpts, arg0)
}

// SynapseBridge is a free data retrieval call binding the contract method 0xe8063377.
//
// Solidity: function synapseBridge() view returns(address)
func (_Router *RouterCaller) SynapseBridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "synapseBridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SynapseBridge is a free data retrieval call binding the contract method 0xe8063377.
//
// Solidity: function synapseBridge() view returns(address)
func (_Router *RouterSession) SynapseBridge() (common.Address, error) {
	return _Router.Contract.SynapseBridge(&_Router.CallOpts)
}

// SynapseBridge is a free data retrieval call binding the contract method 0xe8063377.
//
// Solidity: function synapseBridge() view returns(address)
func (_Router *RouterCallerSession) SynapseBridge() (common.Address, error) {
	return _Router.Contract.SynapseBridge(&_Router.CallOpts)
}

// TokenToSymbol is a free data retrieval call binding the contract method 0x0ba36121.
//
// Solidity: function tokenToSymbol(address ) view returns(string)
func (_Router *RouterCaller) TokenToSymbol(opts *bind.CallOpts, arg0 common.Address) (string, error) {
	var out []interface{}
	err := _Router.contract.Call(opts, &out, "tokenToSymbol", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenToSymbol is a free data retrieval call binding the contract method 0x0ba36121.
//
// Solidity: function tokenToSymbol(address ) view returns(string)
func (_Router *RouterSession) TokenToSymbol(arg0 common.Address) (string, error) {
	return _Router.Contract.TokenToSymbol(&_Router.CallOpts, arg0)
}

// TokenToSymbol is a free data retrieval call binding the contract method 0x0ba36121.
//
// Solidity: function tokenToSymbol(address ) view returns(string)
func (_Router *RouterCallerSession) TokenToSymbol(arg0 common.Address) (string, error) {
	return _Router.Contract.TokenToSymbol(&_Router.CallOpts, arg0)
}

// AdapterSwap is a paid mutator transaction binding the contract method 0x24a98f11.
//
// Solidity: function adapterSwap(address to, address tokenIn, uint256 amountIn, address tokenOut, bytes rawParams) payable returns(uint256 amountOut)
func (_Router *RouterTransactor) AdapterSwap(opts *bind.TransactOpts, to common.Address, tokenIn common.Address, amountIn *big.Int, tokenOut common.Address, rawParams []byte) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "adapterSwap", to, tokenIn, amountIn, tokenOut, rawParams)
}

// AdapterSwap is a paid mutator transaction binding the contract method 0x24a98f11.
//
// Solidity: function adapterSwap(address to, address tokenIn, uint256 amountIn, address tokenOut, bytes rawParams) payable returns(uint256 amountOut)
func (_Router *RouterSession) AdapterSwap(to common.Address, tokenIn common.Address, amountIn *big.Int, tokenOut common.Address, rawParams []byte) (*types.Transaction, error) {
	return _Router.Contract.AdapterSwap(&_Router.TransactOpts, to, tokenIn, amountIn, tokenOut, rawParams)
}

// AdapterSwap is a paid mutator transaction binding the contract method 0x24a98f11.
//
// Solidity: function adapterSwap(address to, address tokenIn, uint256 amountIn, address tokenOut, bytes rawParams) payable returns(uint256 amountOut)
func (_Router *RouterTransactorSession) AdapterSwap(to common.Address, tokenIn common.Address, amountIn *big.Int, tokenOut common.Address, rawParams []byte) (*types.Transaction, error) {
	return _Router.Contract.AdapterSwap(&_Router.TransactOpts, to, tokenIn, amountIn, tokenOut, rawParams)
}

// AddToken is a paid mutator transaction binding the contract method 0x530d95f6.
//
// Solidity: function addToken(string symbol, address token, uint8 tokenType, address bridgeToken, uint256 bridgeFee, uint256 minFee, uint256 maxFee) returns(bool wasAdded)
func (_Router *RouterTransactor) AddToken(opts *bind.TransactOpts, symbol string, token common.Address, tokenType uint8, bridgeToken common.Address, bridgeFee *big.Int, minFee *big.Int, maxFee *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "addToken", symbol, token, tokenType, bridgeToken, bridgeFee, minFee, maxFee)
}

// AddToken is a paid mutator transaction binding the contract method 0x530d95f6.
//
// Solidity: function addToken(string symbol, address token, uint8 tokenType, address bridgeToken, uint256 bridgeFee, uint256 minFee, uint256 maxFee) returns(bool wasAdded)
func (_Router *RouterSession) AddToken(symbol string, token common.Address, tokenType uint8, bridgeToken common.Address, bridgeFee *big.Int, minFee *big.Int, maxFee *big.Int) (*types.Transaction, error) {
	return _Router.Contract.AddToken(&_Router.TransactOpts, symbol, token, tokenType, bridgeToken, bridgeFee, minFee, maxFee)
}

// AddToken is a paid mutator transaction binding the contract method 0x530d95f6.
//
// Solidity: function addToken(string symbol, address token, uint8 tokenType, address bridgeToken, uint256 bridgeFee, uint256 minFee, uint256 maxFee) returns(bool wasAdded)
func (_Router *RouterTransactorSession) AddToken(symbol string, token common.Address, tokenType uint8, bridgeToken common.Address, bridgeFee *big.Int, minFee *big.Int, maxFee *big.Int) (*types.Transaction, error) {
	return _Router.Contract.AddToken(&_Router.TransactOpts, symbol, token, tokenType, bridgeToken, bridgeFee, minFee, maxFee)
}

// AddTokens is a paid mutator transaction binding the contract method 0xd62933e2.
//
// Solidity: function addTokens((string,address,uint256,uint8,address,uint256,uint256,uint256)[] tokens) returns()
func (_Router *RouterTransactor) AddTokens(opts *bind.TransactOpts, tokens []LocalBridgeConfigBridgeTokenConfig) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "addTokens", tokens)
}

// AddTokens is a paid mutator transaction binding the contract method 0xd62933e2.
//
// Solidity: function addTokens((string,address,uint256,uint8,address,uint256,uint256,uint256)[] tokens) returns()
func (_Router *RouterSession) AddTokens(tokens []LocalBridgeConfigBridgeTokenConfig) (*types.Transaction, error) {
	return _Router.Contract.AddTokens(&_Router.TransactOpts, tokens)
}

// AddTokens is a paid mutator transaction binding the contract method 0xd62933e2.
//
// Solidity: function addTokens((string,address,uint256,uint8,address,uint256,uint256,uint256)[] tokens) returns()
func (_Router *RouterTransactorSession) AddTokens(tokens []LocalBridgeConfigBridgeTokenConfig) (*types.Transaction, error) {
	return _Router.Contract.AddTokens(&_Router.TransactOpts, tokens)
}

// Bridge is a paid mutator transaction binding the contract method 0xc2288147.
//
// Solidity: function bridge(address to, uint256 chainId, address token, uint256 amount, (address,address,uint256,uint256,bytes) originQuery, (address,address,uint256,uint256,bytes) destQuery) payable returns()
func (_Router *RouterTransactor) Bridge(opts *bind.TransactOpts, to common.Address, chainId *big.Int, token common.Address, amount *big.Int, originQuery SwapQuery, destQuery SwapQuery) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "bridge", to, chainId, token, amount, originQuery, destQuery)
}

// Bridge is a paid mutator transaction binding the contract method 0xc2288147.
//
// Solidity: function bridge(address to, uint256 chainId, address token, uint256 amount, (address,address,uint256,uint256,bytes) originQuery, (address,address,uint256,uint256,bytes) destQuery) payable returns()
func (_Router *RouterSession) Bridge(to common.Address, chainId *big.Int, token common.Address, amount *big.Int, originQuery SwapQuery, destQuery SwapQuery) (*types.Transaction, error) {
	return _Router.Contract.Bridge(&_Router.TransactOpts, to, chainId, token, amount, originQuery, destQuery)
}

// Bridge is a paid mutator transaction binding the contract method 0xc2288147.
//
// Solidity: function bridge(address to, uint256 chainId, address token, uint256 amount, (address,address,uint256,uint256,bytes) originQuery, (address,address,uint256,uint256,bytes) destQuery) payable returns()
func (_Router *RouterTransactorSession) Bridge(to common.Address, chainId *big.Int, token common.Address, amount *big.Int, originQuery SwapQuery, destQuery SwapQuery) (*types.Transaction, error) {
	return _Router.Contract.Bridge(&_Router.TransactOpts, to, chainId, token, amount, originQuery, destQuery)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address token) returns(bool wasRemoved)
func (_Router *RouterTransactor) RemoveToken(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "removeToken", token)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address token) returns(bool wasRemoved)
func (_Router *RouterSession) RemoveToken(token common.Address) (*types.Transaction, error) {
	return _Router.Contract.RemoveToken(&_Router.TransactOpts, token)
}

// RemoveToken is a paid mutator transaction binding the contract method 0x5fa7b584.
//
// Solidity: function removeToken(address token) returns(bool wasRemoved)
func (_Router *RouterTransactorSession) RemoveToken(token common.Address) (*types.Transaction, error) {
	return _Router.Contract.RemoveToken(&_Router.TransactOpts, token)
}

// RemoveTokens is a paid mutator transaction binding the contract method 0x6c3824ef.
//
// Solidity: function removeTokens(address[] tokens) returns()
func (_Router *RouterTransactor) RemoveTokens(opts *bind.TransactOpts, tokens []common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "removeTokens", tokens)
}

// RemoveTokens is a paid mutator transaction binding the contract method 0x6c3824ef.
//
// Solidity: function removeTokens(address[] tokens) returns()
func (_Router *RouterSession) RemoveTokens(tokens []common.Address) (*types.Transaction, error) {
	return _Router.Contract.RemoveTokens(&_Router.TransactOpts, tokens)
}

// RemoveTokens is a paid mutator transaction binding the contract method 0x6c3824ef.
//
// Solidity: function removeTokens(address[] tokens) returns()
func (_Router *RouterTransactorSession) RemoveTokens(tokens []common.Address) (*types.Transaction, error) {
	return _Router.Contract.RemoveTokens(&_Router.TransactOpts, tokens)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Router *RouterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Router *RouterSession) RenounceOwnership() (*types.Transaction, error) {
	return _Router.Contract.RenounceOwnership(&_Router.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Router *RouterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Router.Contract.RenounceOwnership(&_Router.TransactOpts)
}

// SetAllowance is a paid mutator transaction binding the contract method 0xda46098c.
//
// Solidity: function setAllowance(address token, address spender, uint256 amount) returns()
func (_Router *RouterTransactor) SetAllowance(opts *bind.TransactOpts, token common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "setAllowance", token, spender, amount)
}

// SetAllowance is a paid mutator transaction binding the contract method 0xda46098c.
//
// Solidity: function setAllowance(address token, address spender, uint256 amount) returns()
func (_Router *RouterSession) SetAllowance(token common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SetAllowance(&_Router.TransactOpts, token, spender, amount)
}

// SetAllowance is a paid mutator transaction binding the contract method 0xda46098c.
//
// Solidity: function setAllowance(address token, address spender, uint256 amount) returns()
func (_Router *RouterTransactorSession) SetAllowance(token common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SetAllowance(&_Router.TransactOpts, token, spender, amount)
}

// SetSwapQuoter is a paid mutator transaction binding the contract method 0x804b3dff.
//
// Solidity: function setSwapQuoter(address _swapQuoter) returns()
func (_Router *RouterTransactor) SetSwapQuoter(opts *bind.TransactOpts, _swapQuoter common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "setSwapQuoter", _swapQuoter)
}

// SetSwapQuoter is a paid mutator transaction binding the contract method 0x804b3dff.
//
// Solidity: function setSwapQuoter(address _swapQuoter) returns()
func (_Router *RouterSession) SetSwapQuoter(_swapQuoter common.Address) (*types.Transaction, error) {
	return _Router.Contract.SetSwapQuoter(&_Router.TransactOpts, _swapQuoter)
}

// SetSwapQuoter is a paid mutator transaction binding the contract method 0x804b3dff.
//
// Solidity: function setSwapQuoter(address _swapQuoter) returns()
func (_Router *RouterTransactorSession) SetSwapQuoter(_swapQuoter common.Address) (*types.Transaction, error) {
	return _Router.Contract.SetSwapQuoter(&_Router.TransactOpts, _swapQuoter)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x10f60034.
//
// Solidity: function setTokenConfig(address token, uint8 tokenType, address bridgeToken) returns()
func (_Router *RouterTransactor) SetTokenConfig(opts *bind.TransactOpts, token common.Address, tokenType uint8, bridgeToken common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "setTokenConfig", token, tokenType, bridgeToken)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x10f60034.
//
// Solidity: function setTokenConfig(address token, uint8 tokenType, address bridgeToken) returns()
func (_Router *RouterSession) SetTokenConfig(token common.Address, tokenType uint8, bridgeToken common.Address) (*types.Transaction, error) {
	return _Router.Contract.SetTokenConfig(&_Router.TransactOpts, token, tokenType, bridgeToken)
}

// SetTokenConfig is a paid mutator transaction binding the contract method 0x10f60034.
//
// Solidity: function setTokenConfig(address token, uint8 tokenType, address bridgeToken) returns()
func (_Router *RouterTransactorSession) SetTokenConfig(token common.Address, tokenType uint8, bridgeToken common.Address) (*types.Transaction, error) {
	return _Router.Contract.SetTokenConfig(&_Router.TransactOpts, token, tokenType, bridgeToken)
}

// SetTokenFee is a paid mutator transaction binding the contract method 0x302fee96.
//
// Solidity: function setTokenFee(address token, uint256 bridgeFee, uint256 minFee, uint256 maxFee) returns()
func (_Router *RouterTransactor) SetTokenFee(opts *bind.TransactOpts, token common.Address, bridgeFee *big.Int, minFee *big.Int, maxFee *big.Int) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "setTokenFee", token, bridgeFee, minFee, maxFee)
}

// SetTokenFee is a paid mutator transaction binding the contract method 0x302fee96.
//
// Solidity: function setTokenFee(address token, uint256 bridgeFee, uint256 minFee, uint256 maxFee) returns()
func (_Router *RouterSession) SetTokenFee(token common.Address, bridgeFee *big.Int, minFee *big.Int, maxFee *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SetTokenFee(&_Router.TransactOpts, token, bridgeFee, minFee, maxFee)
}

// SetTokenFee is a paid mutator transaction binding the contract method 0x302fee96.
//
// Solidity: function setTokenFee(address token, uint256 bridgeFee, uint256 minFee, uint256 maxFee) returns()
func (_Router *RouterTransactorSession) SetTokenFee(token common.Address, bridgeFee *big.Int, minFee *big.Int, maxFee *big.Int) (*types.Transaction, error) {
	return _Router.Contract.SetTokenFee(&_Router.TransactOpts, token, bridgeFee, minFee, maxFee)
}

// Swap is a paid mutator transaction binding the contract method 0xb5d1cdd4.
//
// Solidity: function swap(address to, address token, uint256 amount, (address,address,uint256,uint256,bytes) query) payable returns(uint256 amountOut)
func (_Router *RouterTransactor) Swap(opts *bind.TransactOpts, to common.Address, token common.Address, amount *big.Int, query SwapQuery) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "swap", to, token, amount, query)
}

// Swap is a paid mutator transaction binding the contract method 0xb5d1cdd4.
//
// Solidity: function swap(address to, address token, uint256 amount, (address,address,uint256,uint256,bytes) query) payable returns(uint256 amountOut)
func (_Router *RouterSession) Swap(to common.Address, token common.Address, amount *big.Int, query SwapQuery) (*types.Transaction, error) {
	return _Router.Contract.Swap(&_Router.TransactOpts, to, token, amount, query)
}

// Swap is a paid mutator transaction binding the contract method 0xb5d1cdd4.
//
// Solidity: function swap(address to, address token, uint256 amount, (address,address,uint256,uint256,bytes) query) payable returns(uint256 amountOut)
func (_Router *RouterTransactorSession) Swap(to common.Address, token common.Address, amount *big.Int, query SwapQuery) (*types.Transaction, error) {
	return _Router.Contract.Swap(&_Router.TransactOpts, to, token, amount, query)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Router *RouterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Router.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Router *RouterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Router.Contract.TransferOwnership(&_Router.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Router *RouterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Router.Contract.TransferOwnership(&_Router.TransactOpts, newOwner)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router *RouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Router.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router *RouterSession) Receive() (*types.Transaction, error) {
	return _Router.Contract.Receive(&_Router.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Router *RouterTransactorSession) Receive() (*types.Transaction, error) {
	return _Router.Contract.Receive(&_Router.TransactOpts)
}

// RouterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Router contract.
type RouterOwnershipTransferredIterator struct {
	Event *RouterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *RouterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RouterOwnershipTransferred)
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
		it.Event = new(RouterOwnershipTransferred)
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
func (it *RouterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RouterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RouterOwnershipTransferred represents a OwnershipTransferred event raised by the Router contract.
type RouterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Router *RouterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*RouterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Router.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &RouterOwnershipTransferredIterator{contract: _Router.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Router *RouterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *RouterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Router.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RouterOwnershipTransferred)
				if err := _Router.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Router *RouterFilterer) ParseOwnershipTransferred(log types.Log) (*RouterOwnershipTransferred, error) {
	event := new(RouterOwnershipTransferred)
	if err := _Router.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
