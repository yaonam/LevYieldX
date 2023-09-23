package protocols

import "github.com/ethereum/go-ethereum/core/types"

// Interface/handler for different protocols (Aave, Compound)

type YieldProtocol interface {
	// Get all markets
	GetMarkets() []*MarketInfo
	// Returns the transactions required to execute the strategy step.
	GetTransactions(wallet string, step *StrategyStep) ([]*types.Transaction, error)
}
