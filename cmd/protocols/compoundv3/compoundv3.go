package compoundv3

import (
	"fmt"
	"levyieldx/cmd/protocols/schema"
	"levyieldx/cmd/utils"
	"log"

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

func (c *CompoundV3) GetMarkets() ([]*schema.ProtocolChain, error) {
	return nil, nil
}

func (c *CompoundV3) GetTransactions(wallet string, step *schema.StrategyStep) ([]*types.Transaction, error) {
	return nil, nil
}
