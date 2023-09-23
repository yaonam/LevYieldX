package aavev3

import (
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"levyieldx/cmd/protocols/schema"
	"levyieldx/cmd/transactions"
	"levyieldx/cmd/utils"
)

type AaveV3 struct {
	chain                    string
	cl                       *ethclient.Client
	addressesProviderAddress common.Address
	poolAddress              common.Address
	poolContract             *AaveV3Pool
	uiPoolDataProviderCaller *AaveV3UIPoolDataProviderCaller
	wethGatewayTransactor    *WETHGatewayTransactor
}

type ReserveData struct {
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
}

const AaveV3Name = "aavev3"

var aavev3AddressesProviders = map[string]string{
	"ethereum": "0x2f39d218133AFaB8F2B819B1066c7E434Ad94E9e",
	"arbitrum": "0xa97684ead0e402dC232d5A977953DF7ECBaB3CDb",
	"base":     "0xe20fCBdBfFC4Dd138cE8b2E6FBb6CB49777ad64D",
}
var uiPoolDataProviders = map[string]string{
	"ethereum": "0x91c0eA31b49B69Ea18607702c5d9aC360bf3dE7d",
	"arbitrum": "0x145dE30c929a065582da84Cf96F88460dB9745A7",
	"base":     "0x174446a6741300cD2E7C1b1A636Fee99c8F83502",
}
var poolDataProviders = map[string]string{
	"ethereum": "0x7B4EB56E7CD4b454BA8ff71E4518426369a138a3",
	"arbitrum": "0x69FA688f1Dc47d4B5d8029D5a35FB7a548310654",
	"base":     "0x2d8A3C5677189723C4cB8873CfC9C8976FDF38Ac",
}

var wethGatewayAddresses = map[string]string{}

func NewAaveV3Protocol() *AaveV3 {
	return &AaveV3{}
}

func (a *AaveV3) Connect(chain string) error {
	// Setup the client
	rpcEndpoint := utils.ChainConfigs[chain].RPCEndpoint
	cl, err := ethclient.Dial(rpcEndpoint)
	if err != nil {
		log.Printf("Failed to connect to the %v client: %v", chain, err)
		return err
	}

	// Instantiate AddressesProvider
	addressesProviderAddress := common.HexToAddress(aavev3AddressesProviders[chain])
	addressesProviderCaller, err := NewAaveV3AddressesProviderCaller(addressesProviderAddress, cl)
	if err != nil {
		log.Printf("Failed to instantiate AddressesProvider: %v", err)
		return err
	}

	// Fetch lending pool address
	lendingPoolAddress, err := addressesProviderCaller.GetPool(nil)
	if err != nil {
		log.Printf("Failed to get lending pool address: %v", err)
		return err
	}

	// Instantiate lending pool
	poolContract, err := NewAaveV3Pool(lendingPoolAddress, cl)
	if err != nil {
		return fmt.Errorf("failed to instantiate lending pool: %v", err)
	}

	// Instantiate UIPoolDataProvider
	uiPoolDataProviderAddress := common.HexToAddress(uiPoolDataProviders[chain])
	uiPoolDataProviderCaller, err := NewAaveV3UIPoolDataProviderCaller(uiPoolDataProviderAddress, cl)
	if err != nil {
		return fmt.Errorf("failed to instantiate UIPoolDataProvider: %v", err)
	}

	// Instantiate WETHGateway
	wethGatewayAddress := common.HexToAddress(wethGatewayAddresses[chain])
	wethGatewayTransactor, err := NewWETHGatewayTransactor(wethGatewayAddress, cl)
	if err != nil {
		log.Printf("Failed to instantiate WETHGateway: %v", err)
		return err
	}

	a.chain = chain
	a.cl = cl
	a.addressesProviderAddress = addressesProviderAddress
	a.poolAddress = lendingPoolAddress
	a.poolContract = poolContract
	a.uiPoolDataProviderCaller = uiPoolDataProviderCaller
	a.wethGatewayTransactor = wethGatewayTransactor
	log.Printf("%v connected to %v (chainid: %v, pool: %v)", AaveV3Name, a.chain, a.chainID, lendingPoolAddress)
	return nil
}

func (a *AaveV3) getRatios(aggReserveData []IUiPoolDataProviderV3AggregatedReserveData, aggReserveDataLength int) ([]*big.Int, []*big.Int, error) {
	// Fetch OptimalStableToTotalDebtRatio and StableRateExcessOffset
	optimalStableToTotalDebtRatios := make([]*big.Int, aggReserveDataLength)
	stableRateExcessOffsets := make([]*big.Int, aggReserveDataLength)
	interestRateABI, err := InterestRateStrategyMetaData.GetAbi()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch interest rate strategy abi: %v", err)
	}
	// Pack call data
	calls := make([]transactions.Multicall3Call3, 2*aggReserveDataLength)
	debtRatioData, err := interestRateABI.Pack("OPTIMAL_STABLE_TO_TOTAL_DEBT_RATIO")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to pack OPTIMAL_STABLE_TO_TOTAL_DEBT_RATIO calldata: %v", err)
	}
	maxExcessData, err := interestRateABI.Pack("MAX_EXCESS_USAGE_RATIO")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to pack MAX_EXCESS_USAGE_RATIO calldata: %v", err)
	}
	for i, reserveData := range aggReserveData {
		calls[i] = transactions.Multicall3Call3{
			Target:   reserveData.InterestRateStrategyAddress,
			CallData: debtRatioData,
		}
		calls[i+aggReserveDataLength] = transactions.Multicall3Call3{
			Target:   reserveData.InterestRateStrategyAddress,
			CallData: maxExcessData,
		}
	}
	// Make multicall
	responses, err := transactions.HandleMulticall(a.cl, &calls)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to handle multicall: %v", err)
	}
	// Unpack results
	type ReturnData struct {
		Data *big.Int
	}
	for i, response := range *responses {
		returnData := new(ReturnData)
		if i < aggReserveDataLength {
			err = interestRateABI.UnpackIntoInterface(returnData, "OPTIMAL_STABLE_TO_TOTAL_DEBT_RATIO", response.ReturnData)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to unpack OPTIMAL_STABLE_TO_TOTAL_DEBT_RATIO: %v", err)
			}
			optimalStableToTotalDebtRatios[i] = returnData.Data
		} else {
			err = interestRateABI.UnpackIntoInterface(returnData, "MAX_EXCESS_USAGE_RATIO", response.ReturnData)
			if err != nil {
				return nil, nil, fmt.Errorf("failed to unpack MAX_EXCESS_USAGE_RATIO: %v", err)
			}
			stableRateExcessOffsets[i-aggReserveDataLength] = returnData.Data
		}
	}

	return optimalStableToTotalDebtRatios, stableRateExcessOffsets, nil
}

func (a *AaveV3) getReserveDatas(aggReserveData []IUiPoolDataProviderV3AggregatedReserveData, aggReserveDataLength int) ([]*ReserveData, error) {
	// Pack data
	calls := make([]transactions.Multicall3Call3, aggReserveDataLength)
	poolDataProviderABI, err := PoolDataProviderMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch pool data provider abi: %v", err)
	}
	poolDataProviderAddress := common.HexToAddress(poolDataProviders[a.chain])
	for i, reserveData := range aggReserveData {
		data, err := poolDataProviderABI.Pack("getReserveData", reserveData.UnderlyingAsset)
		if err != nil {
			return nil, fmt.Errorf("failed to pack getReserveData calldata: %v", err)
		}
		calls[i] = transactions.Multicall3Call3{
			Target:   poolDataProviderAddress,
			CallData: data,
		}
	}
	// Make multicall
	responses, err := transactions.HandleMulticall(a.cl, &calls)
	if err != nil {
		return nil, fmt.Errorf("failed to handle multicall: %v", err)
	}
	// Unpack results
	reserveDatas := make([]*ReserveData, aggReserveDataLength)
	for i, response := range *responses {
		reserveData := make([]*big.Int, 12)
		err = poolDataProviderABI.UnpackIntoInterface(&reserveData, "getReserveData", response.ReturnData)
		if err != nil {
			return nil, fmt.Errorf("failed to unpack getReserveData: %v", err)
		}
		reserveDatas[i] = &ReserveData{
			TotalStableDebt:     reserveData[3],
			TotalVariableDebt:   reserveData[4],
			VariableBorrowIndex: reserveData[10],
		}
	}

	return reserveDatas, nil
}

func (a *AaveV3) GetMarkets() ([]*schema.ProtocolChain, error) {
	log.Printf("Fetching market data for %v...", a.chain)
	startTime := time.Now()

	// Fetch reserve data for all tokens
	aggReserveData, _, err := a.uiPoolDataProviderCaller.GetReservesData(nil, a.addressesProviderAddress)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch reserve data: %v", err)
	}
	aggReserveDataLength := len(aggReserveData)

	// Fetch OptimalStableToTotalDebtRatio and StableRateExcessOffset
	optimalStableToTotalDebtRatios, stableRateExcessOffsets, err := a.getRatios(aggReserveData, aggReserveDataLength)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch ratios: %v", err)
	}

	// Fetch reserve data
	reserveDatas, err := a.getReserveDatas(aggReserveData, aggReserveDataLength)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch reserve datas: %v", err)
	}

	// Filter out results for specified symbols
	var supplyMarkets []*schema.MarketInfo
	var borrowMarkets []*schema.MarketInfo
	for i, reserveData := range aggReserveData {
		if reserveData.IsPaused {
			continue
		}

		var supplyCap *big.Int
		if reserveData.SupplyCap.Cmp(big.NewInt(0)) == 0 { // Infinite cap
			supplyCap = utils.MaxUint256
		} else {
			decimals := new(big.Int).Exp(big.NewInt(10), reserveData.Decimals, nil)
			amountSupplied := new(big.Int).Add(reserveData.TotalScaledVariableDebt, reserveData.AvailableLiquidity)
			supplyCap = new(big.Int).Mul(reserveData.SupplyCap, decimals)
			supplyCap.Sub(supplyCap, amountSupplied)
		}

		// Check mapping for USDC.e
		symbol, err := utils.ConvertAddressToSymbol(a.chain, reserveData.UnderlyingAsset.Hex())
		if err != nil {
			return nil, fmt.Errorf("failed to convert address to symbol: %v", err)
		}

		market := &schema.MarketInfo{
			Protocol:   AaveV3Name,
			Chain:      a.chain,
			Token:      symbol,
			Decimals:   reserveData.Decimals,
			LTV:        reserveData.BaseLTVasCollateral,
			PriceInUSD: reserveData.PriceInMarketReferenceCurrency,
			Params: map[string]interface{}{
				"supplyCapRemaining": supplyCap,
				"totalVariableDebt":  reserveDatas[i].TotalVariableDebt,
				"totalStableDebt":    reserveDatas[i].TotalStableDebt,

				"reserveFactor":          reserveData.ReserveFactor,
				"availableLiquidity":     reserveData.AvailableLiquidity,
				"averageStableRate":      reserveData.AverageStableRate,
				"stableRateSlope1":       reserveData.StableRateSlope1,
				"stableRateSlope2":       reserveData.StableRateSlope2,
				"variableRateSlope1":     reserveData.VariableRateSlope1,
				"variableRateSlope2":     reserveData.VariableRateSlope2,
				"baseStableBorrowRate":   reserveData.BaseStableBorrowRate,
				"baseVariableBorrowRate": reserveData.BaseVariableBorrowRate,
				"optimalRatio":           reserveData.OptimalUsageRatio,
				"unbacked":               reserveData.Unbacked,

				"optimalStableToTotalDebtRatio": optimalStableToTotalDebtRatios[i],
				"stableRateExcessOffset":        stableRateExcessOffsets[i],
			},
		}
		supplyMarkets = append(supplyMarkets, market)
		borrowMarkets = append(borrowMarkets, market)
	}

	log.Printf("Fetched %v lending tokens & %v borrowing tokens", len(supplyMarkets), len(borrowMarkets))
	log.Printf("Time elapsed: %v", time.Since(startTime))

	return []*schema.ProtocolChain{{
		Protocol:      AaveV3Name,
		Chain:         a.chain,
		SupplyMarkets: supplyMarkets,
		BorrowMarkets: borrowMarkets,
	}}, nil
}

func (a *AaveV3) GetTransactions(wallet string, step *schema.StrategyStep) ([]*types.Transaction, error) {
	return nil, nil
}
