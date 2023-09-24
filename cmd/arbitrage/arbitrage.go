package arbitrage

import (
	"fmt"
	"levyieldx/cmd/protocols"
	"levyieldx/cmd/protocols/schema"
	"levyieldx/cmd/protocols/synapse"
	"levyieldx/cmd/utils"
	"math/big"

	"golang.org/x/exp/slices"
)

// Finds all the strategies that can be made with the given tokens using BFS.
//
// Params:
//   - pcs: List of protocol chains
//   - baseTokens: List of base token symbols
//   - maxLevels: Maximum number of levels to search
//
// Returns:
//   - map[string][][]*t.MarketInfo: Map of collateral to list of strategies
func GetAllStrats(pcs []*schema.ProtocolChain, baseTokens []string, maxLevels int) [][]*schema.MarketInfo {
	var result [][]*schema.MarketInfo
	// Mapping of last token to list of strategies
	marketToStrats := make(map[string][][]*schema.MarketInfo)
	// BFS
	for level := 0; level < maxLevels; level++ {
		newStrats := make(map[string][][]*schema.MarketInfo)
		for _, pc := range pcs {
			// Add supply-borrow pair to existing strategies
			for _, supplyMarket := range pc.SupplyMarkets {
				supplySymbol := utils.CommonSymbol(supplyMarket.Token)
				if level == 0 && !slices.Contains(baseTokens, supplySymbol) {
					continue // Skip if first level and not base token
				}
				if level == 0 { // Add directly to result
					result = append(result, []*schema.MarketInfo{supplyMarket})
				} else {
					for _, strat := range marketToStrats[supplySymbol] {
						if !(utils.PCString(strat[len(strat)-1]) == utils.PCString(supplyMarket)) {
							newStrat := make([]*schema.MarketInfo, len(strat)+1)
							copy(newStrat[:len(strat)], strat)
							newStrat[len(strat)] = supplyMarket
							result = append(result, newStrat)
						}
					}
				}

				for _, borrowMarket := range pc.BorrowMarkets {
					if borrowMarket.Token == supplyMarket.Token {
						continue // Skip if borrowing/supplying same token
					}
					borrowSymbol := utils.CommonSymbol(borrowMarket.Token)

					if level == 0 { // Add directly to result
						marketToStrats[borrowSymbol] = append(marketToStrats[borrowSymbol], []*schema.MarketInfo{supplyMarket, borrowMarket})
					} else { // Add on to existing
						for _, strat := range marketToStrats[supplySymbol] {
							// Only if different pc
							if !(utils.PCString(strat[len(strat)-1]) == utils.PCString(supplyMarket)) {
								newStrat := make([]*schema.MarketInfo, len(strat)+2)
								copy(newStrat[:len(strat)], strat)
								copy(newStrat[len(strat):], []*schema.MarketInfo{supplyMarket, borrowMarket})
								if level < maxLevels-1 { // Only if not last level
									newStrats[borrowSymbol] = append(newStrats[borrowSymbol], newStrat)
								}
							}
						}
					}
				}
			}
		}

		// Copy new strats to result
		for token, strats := range newStrats {
			marketToStrats[token] = append(marketToStrats[token], strats...)
		}
	}

	return result
}

// Helper for calculating supply/borrow apy and amount.
//
// Params:
//   - ps: Map of protocol name to protocol
//   - market: Market to calculate
//   - amountUSD: Amount in USD, with 8 decimals
//   - isSupply: True if supply, false if borrow
//
// Returns:
//   - *t.StrategyStep: Strategy step
//   - *big.Int: Actual amount in USD, with 8 decimals
func calcStep(ps map[string]*protocols.YieldProtocol, market *schema.MarketInfo, amountUSD *big.Int, isSupply bool) (*schema.StrategyStep, *big.Int, error) {
	p, ok := ps[market.Protocol]
	if !ok {
		return nil, nil, fmt.Errorf("failed to get protocol")
	}
	decimals := new(big.Int).Exp(big.NewInt(10), market.Decimals, nil)
	amount := new(big.Int).Div(new(big.Int).Mul(amountUSD, decimals), market.PriceInUSD)
	supplyAPY, supplyAmount, err := (*p).CalcAPY(market, amount, isSupply)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get apy: %v", err)
	}
	actualAmountUSD := new(big.Int).Div(new(big.Int).Mul(supplyAmount, market.PriceInUSD), decimals)
	return &schema.StrategyStep{
		Market:   market,
		IsSupply: isSupply,
		APY:      supplyAPY,
		Amount:   supplyAmount,
	}, actualAmountUSD, nil
}

// Calculates the strategies' steps and total APYs, taking into account
// bridging between chains.
func CalcBridgedStrats(ps map[string]*protocols.YieldProtocol, strats [][]*schema.MarketInfo, initialUSD, safety *big.Int) ([]*schema.Strategy, error) {
	result := make([]*schema.Strategy, len(strats))
	currentAmountUSDs := make([]*big.Int, len(strats))
	// Calc 1st supply
	for j, strat := range strats {
		market := strat[0]
		step, actualUSD, err := calcStep(ps, market, initialUSD, true)
		if err != nil {
			return nil, fmt.Errorf("failed to calc first step: %v", err)
		}
		result[j] = &schema.Strategy{
			Steps: []*schema.StrategyStep{step},
			APY:   step.APY,
		}
		// Adjust for LTV and safety factor
		ltv := utils.BasisMul(market.LTV, safety)
		currentAmountUSDs[j] = utils.BasisMul(actualUSD, ltv) // Percentage
	}

	// Interatively go through each borrow-supply pair, taking into account bridge
	for i := 1; true; i += 2 {
		// Get amounts for all next borrows
		borrowSteps := make([]*schema.StrategyStep, len(strats))
		for j, strat := range strats {
			if len(strat) <= i {
				continue
			}
			borrowMarket := strat[i]
			step, actualUSD, err := calcStep(ps, borrowMarket, currentAmountUSDs[j], false)
			if err != nil {
				return nil, fmt.Errorf("failed to calc borrow step: %v", err)
			}
			borrowSteps[j] = step
			// Update amount if lower
			if actualUSD.Cmp(currentAmountUSDs[j]) == -1 {
				currentAmountUSDs[j] = actualUSD
			}
		}

		// Get amounts for all next supplies
		supplySteps := make([]*schema.StrategyStep, len(strats))
		for j, strat := range strats {
			if len(strat) <= i {
				continue
			}
			supplyMarket := strat[i+1]
			step, actualUSD, err := calcStep(ps, supplyMarket, currentAmountUSDs[j], true)
			if err != nil {
				return nil, fmt.Errorf("failed to calc supply step: %v", err)
			}
			supplySteps[j] = step

			// Recalc borrow if supply amount lower
			if actualUSD.Cmp(currentAmountUSDs[j]) == -1 {
				currentAmountUSDs[j] = actualUSD
			}
			// Ignoring new borrow actualUSD because it should be equal to currentAmountUSDs[j]
			borrowStep, _, err := calcStep(ps, strat[i], currentAmountUSDs[j], false)
			if err != nil {
				return nil, fmt.Errorf("failed to calc borrow step: %v", err)
			}
			borrowSteps[j] = borrowStep
		}

		// Calc bridges if any
		// Queue requests
		indices := make([]int, 0)
		requests := make([]*schema.BridgeRequest, 0)
		for j, supplyStep := range supplySteps {
			if supplyStep == nil {
				continue
			}
			borrowStep := borrowSteps[j]
			if supplyStep.Market.Chain != borrowStep.Market.Chain {
				indices = append(indices, j)
				requests = append(requests, &schema.BridgeRequest{
					Origin:      borrowStep.Market.Chain,
					Destination: supplyStep.Market.Chain,
					TokenIn:     borrowStep.Market.Token,
					TokenOut:    supplyStep.Market.Token,
					AmountIn:    borrowStep.Amount,
				})
			}
		}
		// Call bridges
		var bridge protocols.BridgeProtocol = synapse.NewSynapse()
		bridgeSteps, bridgeAmounts, err := bridge.GetSteps(requests)
		if err != nil {
			return nil, fmt.Errorf("failed to get bridge markets: %v", err)
		}
		// Recalc supplies if bridge slippage
		for k, bridgeAmount := range bridgeAmounts {
			if bridgeAmount == nil {
				continue
			}
			j := indices[k]
			supplyStep := supplySteps[j]
			decimals := new(big.Int).Exp(big.NewInt(10), supplyStep.Market.Decimals, nil)
			amountOutUSD := new(big.Int).Div(new(big.Int).Mul(bridgeAmount, supplyStep.Market.PriceInUSD), decimals)
			if amountOutUSD.Cmp(currentAmountUSDs[j]) == -1 {
				newSupplyStep, _, err := calcStep(ps, supplyStep.Market, amountOutUSD, false)
				if err != nil {
					return nil, fmt.Errorf("failed to recalculate supply step: %v", err)
				}
				supplySteps[j] = newSupplyStep
			}
		}

		// If borrows empty, break
		borrowsEmpty := true

		// Add to result
		for j, borrowStep := range borrowSteps {
			if borrowStep != nil {
				result[j].Steps = append(result[j].Steps, borrowStep)
			}
		}
		for k, j := range indices {
			if bridgeSteps[k] != nil {
				result[j].Steps = append(result[j].Steps, bridgeSteps[k])
			}
		}
		for j, supplyStep := range supplySteps {
			if supplyStep != nil {
				result[j].Steps = append(result[j].Steps, supplyStep)
			}
			borrowsEmpty = borrowsEmpty && supplyStep == nil
		}

		if borrowsEmpty {
			break
		}
	}

	return result, nil
}

// Compares the two strategies' APYs.
// Returns true if a is larger than b, false otherwise.
func apyMore(a, b *schema.Strategy) bool {
	return a.APY.Cmp(b.APY) == 1
}

// Sorts the strategies by APY in descending order.
func SortStrategies(strats []*schema.Strategy) {
	slices.SortFunc(strats, apyMore)
}
