package aavev3

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	"levyieldx/cmd/protocols/schema"
	"levyieldx/cmd/utils"
)

type AaveV3 struct {
	chain                    string
	cl                       *ethclient.Client
	chainID                  *big.Int
	addressesProviderAddress common.Address
	poolAddress              common.Address
	poolContract             *AaveV3Pool
	uiPoolDataProviderCaller *AaveV3UIPoolDataProviderCaller
	wethGatewayTransactor    *WETHGatewayTransactor
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

	// Fetch chainid
	chainid, err := cl.ChainID(context.Background())
	if err != nil {
		log.Printf("Failed to fetch chainid: %v", err)
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
	// a.chainID = chainid
	a.addressesProviderAddress = addressesProviderAddress
	a.poolAddress = lendingPoolAddress
	a.poolContract = poolContract
	a.uiPoolDataProviderCaller = uiPoolDataProviderCaller
	a.wethGatewayTransactor = wethGatewayTransactor
	log.Printf("%v connected to %v (chainid: %v, pool: %v)", AaveV3Name, a.chain, a.chainID, lendingPoolAddress)
	return nil
}

func (a *AaveV3) GetMarkets() ([]*schema.ProtocolChain, error) {
	return nil, nil
}

func (a *AaveV3) GetTransactions(wallet string, step *schema.StrategyStep) ([]*types.Transaction, error)
