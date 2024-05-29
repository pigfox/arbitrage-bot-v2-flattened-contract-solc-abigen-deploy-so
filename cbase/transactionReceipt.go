package cbase

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
)

type TransactionReceipt struct {
	Type              uint64         `json:"type"`
	Status            uint64         `json:"status"`
	CumulativeGasUsed uint64         `json:"cumulativeGasUsed"`
	Bloom             types.Bloom    `json:"bloom"`
	Logs              []*types.Log   `json:"logs"`
	TransactionHash   common.Hash    `json:"transactionHash"`
	ContractAddress   common.Address `json:"contractAddress"`
	GasUsed           uint64         `json:"gasUsed"`
	BlockHash         common.Hash    `json:"blockHash"`
	BlockNumber       *big.Int       `json:"blockNumber"`
	TransactionIndex  uint           `json:"transactionIndex"`
}
