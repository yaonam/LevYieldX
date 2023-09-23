package transactions

import (
	"context"
	"fmt"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var MulticallAddress = common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11")

func HandleMulticall(e *ethclient.Client, calls *[]Multicall3Call3) (*[]Multicall3Result, error) {
	// Pack aggregated calldata
	multicallABI, _ := Multicall3MetaData.GetAbi()
	aggData, err := multicallABI.Pack("aggregate3", *calls)
	if err != nil {
		return nil, fmt.Errorf("failed to pack multicall: %v", err)
	}

	// Call Multicall contract
	response, err := e.CallContract(context.Background(), ethereum.CallMsg{
		To:   &MulticallAddress,
		Data: aggData,
	}, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to call multicall: %v", err)
	}

	// Unpack into Multicall3Result
	responses := new([]Multicall3Result)
	if err := multicallABI.UnpackIntoInterface(responses, "aggregate3", response); err != nil {
		return nil, fmt.Errorf("failed to unpack into results: %v", err)
	}

	return responses, nil
}
