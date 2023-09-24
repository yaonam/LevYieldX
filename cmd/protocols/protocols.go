package protocols

import (
	"fmt"
	"levyieldx/cmd/protocols/aavev3"
	"levyieldx/cmd/protocols/compoundv3"
	"levyieldx/cmd/protocols/schema"
	"levyieldx/cmd/protocols/synapse"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

// Interface/handler for different protocols (Aave, Compound)

type YieldProtocol interface {
	// Connects to the protocol on the specified chain
	Connect(chain string) error
	// Get all markets
	GetMarkets() ([]*schema.ProtocolChain, error)
	// Returns the APY and actual amount for the given token.
	// Actual amount is the amount that can be supplied/borrowed.
	// APY in ray.
	CalcAPY(market *schema.MarketInfo, amount *big.Int, isSupply bool) (*big.Int, *big.Int, error)
	// Returns the transactions required to execute the strategy step.
	GetTransactions(wallet string, step *schema.StrategyStep) ([]*types.Transaction, error)
}

type BridgeProtocol interface {
	// Returns the market and amount out.
	// Bundled to reduce RPC calls.
	GetSteps([]*schema.BridgeRequest) ([]*schema.StrategyStep, []*big.Int, error)

	GetTransactions(wallet string, step *schema.StrategyStep) ([]*types.Transaction, error)
}

func GetYieldProtocol(protocol string) (YieldProtocol, error) {
	switch protocol {
	case "aavev3":
		return aavev3.NewAaveV3Protocol(), nil
	case "compoundv3":
		return compoundv3.NewCompoundV3Protocol(), nil
	default:
		return nil, fmt.Errorf("unknown protocol: %s", protocol)
	}
}

func GetBridgeProtocol(protocol string) (BridgeProtocol, error) {
	switch protocol {
	case "synapse":
		return synapse.NewSynapse(), nil
	default:
		return nil, fmt.Errorf("unknown protocol: %s", protocol)
	}
}
