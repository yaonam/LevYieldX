package transactions

import (
	"context"
	"fmt"
	"math/big"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var MulticallAddress = common.HexToAddress("0xcA11bde05977b3631167028862bE2a173976CA11")

var approveGasLimits = map[string]uint64{
	"ethereum": 50000,
	"arbitrum": 500000,
	"base":     40000,
}

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

func GetApprovalTxIfNeeded(cl *ethclient.Client, chain string, token, owner, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	tokenContract, err := NewERC20Permit(token, cl)
	if err != nil {
		return nil, fmt.Errorf("failed to get token contract: %v", err)
	}
	tokenABI, err := ERC20PermitMetaData.GetAbi()
	if err != nil {
		return nil, fmt.Errorf("failed to get token abi: %v", err)
	}

	// Get allowance
	allowance, err := tokenContract.Allowance(nil, owner, spender)
	if err != nil {
		return nil, fmt.Errorf("failed to get allowance: %v", err)
	}

	// Approve if needed
	if allowance.Cmp(amount) == -1 {
		data, err := tokenABI.Pack("approve", spender, amount)
		if err != nil {
			return nil, fmt.Errorf("failed to pack approve: %v", err)
		}
		inner := &types.DynamicFeeTx{
			To:   &token,
			Data: data,
			Gas:  approveGasLimits[chain],
		}
		tx := types.NewTx(inner)
		return tx, nil
	}
	return nil, nil
}
