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

type CompoundV3Stats struct {
	TotalSupply     *big.Int
	TotalBorrows    *big.Int
	SupplyBase      *big.Int
	BorrowBase      *big.Int
	SupplySlopeLow  *big.Int
	BorrowSlopeLow  *big.Int
	SupplySlopeHigh *big.Int
	BorrowSlopeHigh *big.Int
}

const CompoundV3Name = "compoundv3"

var gasLimit = map[string]uint64{
	"ethereum": 200000,
	"arbitrum": 1500000,
	"base":     200000,
}

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

var decimals = map[string]uint8{
	"ETH":    18,
	"USDC.e": 8,
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

/*
Seconds Per Year = 60 * 60 * 24 * 365
Utilization = getUtilization()
Supply Rate = getSupplyRate(Utilization)
Supply APR = Supply Rate / (10 ^ 18) * Seconds Per Year * 100
*/
func (c *CompoundV3) getBaseStats() (*CompoundV3Stats, error) {
	// Pack calls
	cometABI, err := CometMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get comet abi: %v", err)
	}
	calls := make([]transactions.Multicall3Call3, 8)
	methods := []string{"totalSupply", "totalBorrow", "supplyPerSecondInterestRateBase", "borrowPerSecondInterestRateBase", "supplyPerSecondInterestRateSlopeLow", "borrowPerSecondInterestRateSlopeLow", "supplyPerSecondInterestRateSlopeHigh", "borrowPerSecondInterestRateSlopeHigh"}
	for i, method := range methods {
		data, err := cometABI.Pack(method)
		if err != nil {
			return nil, fmt.Errorf("failed to pack %v calldata: %v", method, err)
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

	// Unpack
	type ReturnData struct {
		Data *big.Int
	}
	results := make([]*big.Int, 8)
	for i, response := range *responses {
		returnData := new(ReturnData)
		if err := cometABI.UnpackIntoInterface(returnData, methods[i], response.ReturnData); err != nil {
			return nil, fmt.Errorf("failed to unpack %v: %v", methods[i], err)
		}
		results[i] = returnData.Data
	}

	return &CompoundV3Stats{
		TotalSupply:     results[0],
		TotalBorrows:    results[1],
		SupplyBase:      results[2],
		BorrowBase:      results[3],
		SupplySlopeLow:  results[4],
		BorrowSlopeLow:  results[5],
		SupplySlopeHigh: results[6],
		BorrowSlopeHigh: results[7],
	}, nil
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

		// Fill in LTV and APY for collateral tokens
		supplyMarkets := make([]*schema.MarketInfo, numAssets+1)
		for i, assetInfo := range config.AssetConfigs {
			symbol, err := utils.ConvertAddressToSymbol(c.chain, assetInfo.Asset.Hex())
			if err != nil {
				return nil, fmt.Errorf("failed to convert symbol: %v", err)
			}
			// Has LTV, no APY
			ltv := big.NewInt(int64(assetInfo.BorrowCollateralFactor))
			ltv.Quo(ltv, big.NewInt(1e14))
			supplyCap := new(big.Int).Sub(assetInfo.SupplyCap, amountsSupplied[i])
			supplyMarkets[i] = &schema.MarketInfo{
				Protocol:   CompoundV3Name,
				Chain:      c.chain,
				Token:      symbol,
				Decimals:   big.NewInt(int64(assetInfo.Decimals)),
				LTV:        ltv,
				PriceInUSD: prices[i],
				Params: map[string]interface{}{
					"baseAsset":          chainAsset,
					"supplyCapRemaining": supplyCap,
					"totalSupply":        amountsSupplied[i],
					"totalBorrows":       big.NewInt(0),

					"base":      big.NewInt(0),
					"slopeLow":  big.NewInt(0),
					"kink":      big.NewInt(0),
					"slopeHigh": big.NewInt(0),
				},
			}
		}

		// Base token, has APY, no LTV
		symbol, err := utils.ConvertAddressToSymbol(c.chain, config.BaseToken.Hex())
		if err != nil {
			return nil, fmt.Errorf("failed to convert base address to token: %v", err)
		}
		baseStats, err := c.getBaseStats()
		if err != nil {
			return nil, fmt.Errorf("failed to get base aprs: %v", err)
		}
		decimals := decimals[symbol]
		market := &schema.MarketInfo{
			Protocol:   CompoundV3Name,
			Chain:      c.chain,
			Token:      symbol,
			Decimals:   big.NewInt(int64(decimals)),
			LTV:        big.NewInt(0),
			PriceInUSD: prices[numAssets],
			Params: map[string]interface{}{
				"baseAsset":          chainAsset,
				"supplyCapRemaining": utils.MaxUint256,
				"totalSupply":        baseStats.TotalSupply,
				"totalBorrows":       baseStats.TotalBorrows,

				"base":      baseStats.SupplyBase,
				"slopeLow":  baseStats.SupplySlopeLow,
				"kink":      big.NewInt(int64(config.SupplyKink)),
				"slopeHigh": baseStats.SupplySlopeHigh,
			},
		}
		supplyMarkets[numAssets] = market
		borrowMarkets := []*schema.MarketInfo{{
			Protocol:   CompoundV3Name,
			Chain:      c.chain,
			Token:      symbol,
			Decimals:   big.NewInt(int64(decimals)),
			LTV:        big.NewInt(0),
			PriceInUSD: prices[numAssets],
			Params: map[string]interface{}{
				"baseAsset":          chainAsset,
				"supplyCapRemaining": utils.MaxUint256,
				"totalSupply":        baseStats.TotalSupply,
				"totalBorrows":       baseStats.TotalBorrows,

				"base":      baseStats.BorrowBase,
				"slopeLow":  baseStats.BorrowSlopeLow,
				"kink":      big.NewInt(int64(config.BorrowKink)),
				"slopeHigh": baseStats.BorrowSlopeHigh,
			}}}

		log.Printf("Fetched %v lending tokens & %v borrowing tokens", len(supplyMarkets), len(borrowMarkets))
		protocolChains[i] = &schema.ProtocolChain{
			Protocol:      CompoundV3Name,
			Chain:         c.chain,
			SupplyMarkets: supplyMarkets,
			BorrowMarkets: borrowMarkets,
		}
	}

	log.Printf("Time elapsed: %v", time.Since(startTime))

	return protocolChains, nil
}

func (c *CompoundV3) CalcAPY(m *schema.MarketInfo, amount *big.Int, isSupply bool) (*big.Int, *big.Int, error) {
	supplyCapRemaining := m.Params["supplyCapRemaining"].(*big.Int)
	totalSupply := m.Params["totalSupply"].(*big.Int)
	totalBorrows := m.Params["totalBorrows"].(*big.Int)

	base := m.Params["base"].(*big.Int)
	slopeLow := m.Params["slopeLow"].(*big.Int)
	kink := m.Params["kink"].(*big.Int)
	slopeHigh := m.Params["slopeHigh"].(*big.Int)

	// If TotalBorrows is nil, 0 APY
	if m.Params["totalBorrows"].(*big.Int) == nil {
		return big.NewInt(0), amount, nil
	}

	var actualAmount *big.Int
	availableLiquidity := new(big.Int).Sub(totalSupply, totalBorrows)
	if isSupply && amount.Cmp(supplyCapRemaining) == 1 {
		actualAmount = new(big.Int).Set(supplyCapRemaining)
	} else if !isSupply && amount.Cmp(availableLiquidity) == 1 {
		actualAmount = new(big.Int).Set(availableLiquidity)
	} else {
		actualAmount = new(big.Int).Set(amount)
	}

	// If not base market (totalBorrows is nil), 0 APY
	if totalBorrows == nil {
		return big.NewInt(0), actualAmount, nil
	}

	// Calc utilization
	supply := totalSupply
	borrows := totalBorrows
	if isSupply {
		supply = new(big.Int).Add(supply, actualAmount)
	} else {
		borrows = new(big.Int).Add(borrows, actualAmount)
	}
	utilization := new(big.Int).Div(new(big.Int).Mul(borrows, utils.ETHMantissaInt), supply)

	// Calculate rate per second
	var ratePerSecond *big.Int
	if utilization.Cmp(kink) < 1 {
		ratePerSecond = new(big.Int).Add(base, utils.ManMul(slopeLow, utilization))
	} else {
		ratePerSecond = new(big.Int).Add(base, utils.ManMul(slopeHigh, kink))
		ratePerSecond.Add(ratePerSecond, utils.ManMul(slopeHigh, new(big.Int).Sub(utilization, kink)))
	}

	// Calculate APY
	apy := new(big.Int).Mul(ratePerSecond, utils.SecPerYear)
	apy.Mul(apy, big.NewInt(1e9)) // Convert to ray

	return apy, actualAmount, nil
}

// Lends the token to the protocol
func (c *CompoundV3) GetTransactions(wallet string, step *schema.StrategyStep) ([]*types.Transaction, error) {
	m := step.Market
	baseAsset := m.Params["baseAsset"].(string)

	// Connect to comet
	chainAsset := step.Market.Chain
	if baseAsset != "" {
		chainAsset += ":" + baseAsset
	}
	if err := c.connectComet(chainAsset); err != nil {
		return nil, fmt.Errorf("failed to connect to comet: %v", err)
	}

	walletAddress := common.HexToAddress(wallet)
	address, err := utils.ConvertSymbolToAddress(step.Market.Chain, step.Market.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to convert symbol to address: %v", err)
	}
	tokenAddress := common.HexToAddress(address)

	cometABI, err := CometMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get comet abi: %v", err)
	}

	var txs []*types.Transaction
	var method string
	if step.IsSupply {
		// Handle approvals
		approvalTx, err := transactions.GetApprovalTxIfNeeded(c.cl, step.Market.Chain, tokenAddress, walletAddress, c.cometAddress, step.Amount)
		if err != nil {
			return nil, fmt.Errorf("failed to get approval tx: %v", err)
		}
		txs = append(txs, approvalTx)
		method = "supply"
	} else {
		method = "withdraw"
	}
	txData, err := cometABI.Pack(method, tokenAddress, step.Amount)
	if err != nil {
		return nil, fmt.Errorf("failed to pack withdraw calldata: %v", err)
	}
	tx := types.NewTx(&types.DynamicFeeTx{
		To:   &c.cometAddress,
		Data: txData,
		Gas:  gasLimit[step.Market.Chain],
	})
	txs = append(txs, tx)

	return txs, nil
}
