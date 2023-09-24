package synapse

import (
	"context"
	"fmt"
	"levyieldx/cmd/protocols/schema"
	"levyieldx/cmd/transactions"
	"levyieldx/cmd/utils"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type Synapse struct{}

const synapseName = "synapse"

var routers = map[string]common.Address{
	"ethereum": common.HexToAddress("0x7E7A0e201FD38d3ADAA9523Da6C109a07118C96a"),
	"arbitrum": common.HexToAddress("0x7E7A0e201FD38d3ADAA9523Da6C109a07118C96a"),
}

var usdcRouters = map[string]common.Address{
	"ethereum": common.HexToAddress("0xD359bc471554504f683fbd4f6e36848612349DDF"),
	"arbitrum": common.HexToAddress("0xD359bc471554504f683fbd4f6e36848612349DDF"),
}

var gasLimit = map[string]uint64{
	"ethereum": uint64(300000),
	"arbitrum": uint64(1200000),
}

func NewSynapse() *Synapse {
	return &Synapse{}
}

// TODO: Use multicall to reduce RPC calls
func getConnectedBridgeTokens(rs []*schema.BridgeRequest) ([][]string, []common.Address, error) {
	result := make([][]string, len(rs))
	addressOuts := make([]common.Address, len(rs))
	for i, r := range rs {
		addressOut, err := utils.ConvertSymbolToAddress(r.Destination, r.TokenOut)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to convert symbol to address: %v", err)
		}
		addressOuts[i] = common.HexToAddress(addressOut)
		cl, err := utils.Connect(r.Destination)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to connect to %v: %v", r.Destination, err)
		}
		// Check if USDC
		var routerAddress common.Address
		if r.TokenIn == "USDC" {
			routerAddress = usdcRouters[r.Destination]
		} else {
			routerAddress = routers[r.Destination]
		}
		router, err := NewRouter(routerAddress, cl)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to create router: %v", err)
		}
		bridgeTokens, err := router.GetConnectedBridgeTokens(nil, addressOuts[i])
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get connected bridge tokens: %v", err)
		}
		for _, bridgeToken := range bridgeTokens {
			result[i] = append(result[i], bridgeToken.Symbol)
		}
	}
	return result, addressOuts, nil
}

func getOriginAmountOut(rs []*schema.BridgeRequest, cbts [][]string) ([][]SwapQuery, error) {
	result := make([][]SwapQuery, len(rs))
	for i, r := range rs {
		if len(cbts[i]) == 0 {
			continue // No route, TODO: handle USDC
		}
		addressIn, err := (utils.ConvertSymbolToAddress(r.Origin, r.TokenIn))
		if err != nil {
			return nil, fmt.Errorf("failed to convert symbol to address: %v", err)
		}
		cl, err := utils.Connect(r.Origin)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to %v: %v", r.Origin, err)
		}
		// Check if USDC
		var routerAddress common.Address
		if r.TokenIn == "USDC" {
			routerAddress = usdcRouters[r.Destination]
		} else {
			routerAddress = routers[r.Destination]
		}
		router, err := NewRouter(routerAddress, cl)
		if err != nil {
			return nil, fmt.Errorf("failed to create router: %v", err)
		}
		originQueries, err := router.GetOriginAmountOut(nil, common.HexToAddress(addressIn), cbts[i], r.AmountIn)
		if err != nil {
			return nil, fmt.Errorf("failed to get origin amount out: %v", err)
		}
		result[i] = originQueries
	}
	return result, nil
}

func getDestinationAmountOut(rs []*schema.BridgeRequest, destRequests [][]DestRequest, addressOuts []common.Address) ([][]SwapQuery, error) {
	result := make([][]SwapQuery, len(rs))
	for i, r := range rs {
		if destRequests[i] == nil {
			continue
		}
		cl, err := utils.Connect(r.Origin)
		if err != nil {
			return nil, fmt.Errorf("failed to connect to %v: %v", r.Origin, err)
		}
		// Check if USDC
		var routerAddress common.Address
		if r.TokenIn == "USDC" {
			routerAddress = usdcRouters[r.Destination]
		} else {
			routerAddress = routers[r.Destination]
		}
		router, err := NewRouter(routerAddress, cl)
		if err != nil {
			return nil, fmt.Errorf("failed to create router: %v", err)
		}
		destQueries, err := router.GetDestinationAmountOut(nil, destRequests[i], addressOuts[i])
		if err != nil {
			return nil, fmt.Errorf("failed to get dest amount out: %v", err)
		}
		result[i] = destQueries
	}
	return result, nil
}

// Returns the steps and amount out for the given bridge requests.
//
// Params:
//   - rs: Bridge requests
//
// Returns:
//   - steps: Strategy steps, nil if no steps are found.
//   - amountOuts: Amount out for each step, nil if no steps are found.
func (s *Synapse) GetSteps(rs []*schema.BridgeRequest) ([]*schema.StrategyStep, []*big.Int, error) {
	// Get connected bridge tokens
	cbts, addressOuts, err := getConnectedBridgeTokens(rs)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get connected bridge tokens: %v", err)
	}

	// Get origin amount out
	originQueries, err := getOriginAmountOut(rs, cbts)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get origin amount out: %v", err)
	}

	// List queries
	destRequests := make([][]DestRequest, len(originQueries))
	for i, symbols := range cbts {
		if symbols == nil {
			continue
		}
		destRequests[i] = make([]DestRequest, len(symbols))
		for j, symbol := range symbols {
			destRequests[i][j] = DestRequest{
				Symbol:   symbol,
				AmountIn: originQueries[i][j].MinAmountOut,
			}
		}
	}

	// Get destination amount out
	destQueries, err := getDestinationAmountOut(rs, destRequests, addressOuts)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get dest amount out: %v", err)
	}

	// Pick a pair of origin and dest queries
	bestIndices := make([]int, len(destQueries))
	for i, queries := range destQueries {
		if queries == nil {
			bestIndices[i] = -1
		}
		for j, query := range queries {
			if query.MinAmountOut.Cmp(queries[bestIndices[i]].MinAmountOut) == 1 {
				bestIndices[i] = j
			}
		}
		if bestIndices[i] != -1 && originQueries[i][bestIndices[i]].MinAmountOut.Cmp(big.NewInt(0)) == 0 {
			bestIndices[i] = -1
		}
	}

	// Organize into Strat steps
	result := make([]*schema.StrategyStep, len(bestIndices))
	amountOuts := make([]*big.Int, len(bestIndices))
	for i, j := range bestIndices {
		if j == -1 {
			continue
		}
		result[i] = &schema.StrategyStep{
			Amount: rs[i].AmountIn,
			Market: &schema.MarketInfo{
				Protocol: synapseName,
				Chain:    rs[i].Origin,
				Token:    rs[i].TokenIn,
				Params: map[string]interface{}{
					"originQuery": originQueries[i][j],
					"destQuery":   destQueries[i][j],
				},
			},
		}
		amountOuts[i] = originQueries[i][j].MinAmountOut
	}

	return result, amountOuts, nil
}

func (s *Synapse) GetTransactions(wallet string, step *schema.StrategyStep) ([]*types.Transaction, error) {
	// TODO: Apply slippage/deadline
	originQueryMap := step.Market.Params["originQuery"].(map[string]interface{})
	originQuery := SwapQuery{
		SwapAdapter:  common.HexToAddress(originQueryMap["SwapAdapter"].(string)),
		TokenOut:     common.HexToAddress(originQueryMap["TokenOut"].(string)),
		MinAmountOut: big.NewInt(int64(originQueryMap["MinAmountOut"].(float64))),
		Deadline:     big.NewInt(int64(originQueryMap["Deadline"].(float64))),
		RawParams:    []byte(originQueryMap["RawParams"].(string)),
	}
	destQueryMap := step.Market.Params["destQuery"].(map[string]interface{})
	destQuery := SwapQuery{
		SwapAdapter:  common.HexToAddress(destQueryMap["SwapAdapter"].(string)),
		TokenOut:     common.HexToAddress(destQueryMap["TokenOut"].(string)),
		MinAmountOut: big.NewInt(int64(destQueryMap["MinAmountOut"].(float64))),
		Deadline:     big.NewInt(int64(destQueryMap["Deadline"].(float64))),
		RawParams:    []byte(destQueryMap["RawParams"].(string)),
	}

	// Approve/swap if needed
	cl, err := utils.Connect(step.Market.Chain)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to %v: %v", step.Market.Chain, err)
	}
	routerAddress := routers[step.Market.Chain]
	token, err := utils.ConvertSymbolToAddress(step.Market.Chain, step.Market.Token)
	if err != nil {
		return nil, fmt.Errorf("failed to convert symbol to address: %v", err)
	}
	tokenAddress := common.HexToAddress(token)
	approveTx, err := transactions.GetApprovalTxIfNeeded(cl, step.Market.Chain, tokenAddress, common.HexToAddress(wallet), routerAddress, step.Amount)
	if err != nil {
		return nil, fmt.Errorf("failed to get approval tx: %v", err)
	}

	// Bridge
	// Assume that tokenIn is not native ETH
	routerABI, err := RouterMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get router abi: %v", err)
	}
	chainID, err := cl.ChainID(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get chain ID: %v", err)
	}
	bridgeData, err := routerABI.Pack("bridge", common.HexToAddress(wallet), chainID, tokenAddress, step.Amount, originQuery, destQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to pack bridge data: %v", err)
	}
	bridgeTx := types.NewTx(&types.DynamicFeeTx{
		To:   &routerAddress,
		Data: bridgeData,
		Gas:  gasLimit[step.Market.Chain],
	})

	if approveTx != nil {
		return []*types.Transaction{approveTx, bridgeTx}, nil
	}
	return []*types.Transaction{bridgeTx}, nil
}
