package protocols

import (
	"fmt"
	"levyieldx/cmd/protocols/aavev3"
	"levyieldx/cmd/protocols/schema"

	"github.com/ethereum/go-ethereum/core/types"
)

// Interface/handler for different protocols (Aave, Compound)

type YieldProtocol interface {
	// Connects to the protocol on the specified chain
	Connect(chain string) error
	// Get all markets
	GetMarkets() ([]*schema.ProtocolChain, error)
	// Returns the transactions required to execute the strategy step.
	GetTransactions(wallet string, step *schema.StrategyStep) ([]*types.Transaction, error)
}

func GetYieldProtocol(protocol string) (YieldProtocol, error) {
	switch protocol {
	case "aavev3":
		return aavev3.NewAaveV3Protocol(), nil
	default:
		return nil, fmt.Errorf("unknown protocol: %s", protocol)
	}
}
