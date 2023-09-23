package compoundv3

import (
	"levyieldx/cmd/protocols/schema"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type CompoundV3 struct {
	chain          string
	cl             *ethclient.Client
	chainID        *big.Int
	configAddress  common.Address
	configContract *Configurator
	cometAddress   common.Address
	cometContract  *Comet
}

const CompoundV3Name = "compoundv3"

func NewCompoundV3Protocol() *CompoundV3 {
	return &CompoundV3{}
}

func (c *CompoundV3) Connect(chain string) error {
	return nil
}

func (c *CompoundV3) GetMarkets() ([]*schema.ProtocolChain, error) {
	return nil, nil
}

func (c *CompoundV3) GetTransactions(wallet string, step *schema.StrategyStep) ([]*types.Transaction, error) {
	return nil, nil
}
