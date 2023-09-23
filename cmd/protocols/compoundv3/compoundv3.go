package compoundv3

import (
	"fmt"
	"levyieldx/cmd/protocols/schema"
	"levyieldx/cmd/transactions"
	"levyieldx/cmd/utils"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type CompoundV3 struct {
	chain          string
	cl             *ethclient.Client
	configAddress  common.Address
	configContract *Configurator
	cometAddress   common.Address
	cometContract  *Comet
}

const CompoundV3Name = "compoundv3"

var compv3ConfigAddresses = map[string]string{
	"ethereum": "0x316f9708bB98af7dA9c68C1C3b5e79039cD336E3",
	"arbitrum": "0xb21b06D71c75973babdE35b49fFDAc3F82Ad3775",
	"base":     "0x45939657d1CA34A8FA39A924B71D28Fe8431e581",
}

var baseAssets = map[string][]string{
	"ethereum": {"usdc", "weth"},
	"base":     {"usdbc", "weth"},
}

var compv3CometAddresses = map[string]string{
	"ethereum:usdc": "0xc3d688B66703497DAA19211EEdff47f25384cdc3",
	"ethereum:weth": "0xA17581A9E3356d9A858b789D68B4d866e593aE94",
	"arbitrum":      "0xA5EDBDD9646f8dFF606d7448e414884C7d905dCA",
	"base:usdbc":    "0x9c4ec768c28520B50860ea7a15bd7213a9fF58bf",
	"base:weth":     "0x46e6b214b524310239732D51387075E0e70970bf",
}

func NewCompoundV3Protocol() *CompoundV3 {
	return &CompoundV3{}
}

func (c *CompoundV3) Connect(chain string) error {
	// Setup the client
	rpcEndpoint := utils.ChainConfigs[chain].RPCEndpoint
	cl, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		log.Printf("Failed to connect to the %v client: %v", chain, err)
		return err
	}

	// Instantiate configurator
	configAddress, ok := compv3ConfigAddresses[chain]
	if !ok {
		return fmt.Errorf("failed to find config address for %v", chain)
	}
	c.configAddress = common.HexToAddress(configAddress)
	c.configContract, err = NewConfigurator(c.configAddress, cl)
	if err != nil {
		log.Printf("Failed to instantiate configurator: %v", err)
		return err
	}

	c.chain = chain
	c.cl = cl

	log.Printf("%v connected to %v", CompoundV3Name, c.chain)
	return nil
}

// Uses multicall to reduce RPC calls
func (c *CompoundV3) getPrices(pfs []common.Address) ([]*big.Int, error) {
	pfLength := len(pfs)
	// Aggregate calldata
	cometABI, err := CometMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get comet abi: %v", err)
	}
	calls := make([]transactions.Multicall3Call3, pfLength)
	for i, pf := range pfs {
		data, err := cometABI.Pack("getPrice", pf)
		if err != nil {
			return nil, fmt.Errorf("failed to pack get asset info calldata: %v", err)
		}
		calls[i] = transactions.Multicall3Call3{
			Target:   c.cometAddress,
			CallData: data,
		}
	}

	// Perform multicall
	responses, err := transactions.HandleMulticall(c.cl, &calls)
	if err != nil {
		return nil, fmt.Errorf("failed to multicall asset info: %v", err)
	}

	// Unpack into CometCoreAssetInfo
	type ReturnData struct {
		Data *big.Int
	}
	result := make([]*big.Int, pfLength)
	for i, response := range *responses {
		returnData := new(ReturnData)
		if err := cometABI.UnpackIntoInterface(returnData, "getPrice", response.ReturnData); err != nil {
			return nil, fmt.Errorf("failed to unpack asset info: %v", err)
		}
		// Prices have 8 decimals for compv3
		result[i] = returnData.Data
	}

	return result, nil
}

// Returns the amount supplied for each token.
func (c *CompoundV3) getTotalsCollateral(assets []CometConfigurationAssetConfig) ([]*big.Int, error) {
	// Aggregate calldata
	cometABI, err := CometMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get comet abi: %v", err)
	}
	numAssets := len(assets)
	calls := make([]transactions.Multicall3Call3, numAssets)
	for i, asset := range assets {
		data, err := cometABI.Pack("totalsCollateral", asset.Asset)
		if err != nil {
			return nil, fmt.Errorf("failed to pack totals collateral calldata: %v", err)
		}
		calls[i] = transactions.Multicall3Call3{
			Target:   c.cometAddress,
			CallData: data,
		}
	}

	// Perform multicall
	responses, err := transactions.HandleMulticall(c.cl, &calls)
	if err != nil {
		return nil, fmt.Errorf("failed to multicall asset info: %v", err)
	}

	// Unpack into CometCoreAssetInfo
	type ReturnData struct {
		TotalSupplyAsset *big.Int
		Reserved         *big.Int
	}
	result := make([]*big.Int, numAssets)
	for i, response := range *responses {
		returnData := new(ReturnData)
		if err := cometABI.UnpackIntoInterface(returnData, "totalsCollateral", response.ReturnData); err != nil {
			return nil, fmt.Errorf("failed to unpack totalsCollateral: %v", err)
		}
		result[i] = returnData.TotalSupplyAsset
	}

	return result, nil
}

func (c *CompoundV3) connectComet(chainAsset string) error {
	// Instantiate comet
	var err error
	c.cometAddress = common.HexToAddress(compv3CometAddresses[chainAsset])
	c.cometContract, err = NewComet(c.cometAddress, c.cl)
	if err != nil {
		log.Printf("Failed to instantiate comet: %v", err)
	}

	return err
}

func (c *CompoundV3) GetMarkets() ([]*schema.ProtocolChain, error) {
	log.Printf("Fetching market data for %v...", c.chain)
	startTime := time.Now()

	// If multiple base assets, connect to comet
	baseAssets, ok := baseAssets[c.chain]
	var chainAssets []string
	if ok {
		for _, baseAsset := range baseAssets {
			chainAssets = append(chainAssets, fmt.Sprintf("%v:%v", c.chain, baseAsset))
		}
	} else {
		chainAssets = []string{c.chain}
	}

	protocolChains := make([]*schema.ProtocolChain, len(chainAssets))
	for i, chainAsset := range chainAssets {
		// Connect to comet
		if err := c.connectComet(chainAsset); err != nil {
			return nil, fmt.Errorf("failed to connect to comet: %v", err)
		}

		// Get config
		config, err := c.configContract.GetConfiguration(nil, c.cometAddress)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch configuration: %v", err)
		}

		// Get all token info
		numAssets := len(config.AssetConfigs)

		// Get prices
		pfs := make([]common.Address, numAssets+1)
		for i, assetInfo := range config.AssetConfigs {
			pfs[i] = assetInfo.PriceFeed
		}
		pfs[numAssets] = config.BaseTokenPriceFeed
		prices, err := c.getPrices(pfs)
		if err != nil {
			return nil, fmt.Errorf("failed to fetch all prices: %v", err)
		}

		// Get amounts supplied
		amountsSupplied, err := c.getTotalsCollateral(config.AssetConfigs)
		if err != nil {
			return nil, fmt.Errorf("failed to get amounts supplied: %v", err)
		}
	}

	log.Printf("Time elapsed: %v", time.Since(startTime))

	return protocolChains, nil
}

func (c *CompoundV3) GetTransactions(wallet string, step *schema.StrategyStep) ([]*types.Transaction, error) {
	return nil, nil
}
