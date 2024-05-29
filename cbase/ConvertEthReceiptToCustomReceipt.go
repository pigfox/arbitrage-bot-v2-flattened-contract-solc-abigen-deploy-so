package cbase

import "github.com/ethereum/go-ethereum/core/types"

func ConvertEthReceiptToCustomReceipt(ethReceipt *types.Receipt) *TransactionReceipt {
	return &TransactionReceipt{
		Type:              uint64(ethReceipt.Type),
		Status:            ethReceipt.Status,
		CumulativeGasUsed: ethReceipt.CumulativeGasUsed,
		Bloom:             ethReceipt.Bloom,
		Logs:              ethReceipt.Logs,
		TransactionHash:   ethReceipt.TxHash,
		ContractAddress:   ethReceipt.ContractAddress,
		GasUsed:           ethReceipt.GasUsed,
		BlockHash:         ethReceipt.BlockHash,
		BlockNumber:       ethReceipt.BlockNumber,
		TransactionIndex:  uint(ethReceipt.TransactionIndex),
	}
}
