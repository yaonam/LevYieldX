package schema

import "math/big"

type ProtocolChain struct {
	Protocol      string        `json:"protocol"`
	Chain         string        `json:"chain"`
	SupplyMarkets []*MarketInfo `json:"lendingSpecs"`
	BorrowMarkets []*MarketInfo `json:"borrowingSpecs"`
}

type MarketInfo struct {
	Protocol   string                 `json:"protocol"`
	Chain      string                 `json:"chain"`
	Token      string                 `json:"token"`
	Decimals   *big.Int               `json:"decimals"`
	LTV        *big.Int               `json:"ltv"`        // In basis points, 0 if cannot be collateral
	PriceInUSD *big.Int               `json:"priceInUSD"` // How much USD is required to purchase 1 ether unit, with 8 decimals
	Params     map[string]interface{} `json:"params"`     // For protocol specific parameters
}

type BridgeRequest struct {
	Origin      string   `json:"origin"`
	Destination string   `json:"destination"`
	TokenIn     string   `json:"tokenIn"`
	TokenOut    string   `json:"tokenOut"`
	AmountIn    *big.Int `json:"amountIn"`
}

type StrategyStep struct {
	Market   *MarketInfo `json:"market"`
	IsSupply bool        `json:"isSupply"`
	APY      *big.Int    `json:"apy"`
	Amount   *big.Int    `json:"amount"`
}

type Strategy struct {
	Steps []*StrategyStep `json:"steps"`
	APY   *big.Int        `json:"apy"`
}
