package deploy

import (
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy-so/connection"
	"context"
	"log"
)

func getBlockGasLimit() uint64 {
	// Get the latest block number
	header, err := connection.RPC.Client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	blockNumber := header.Number

	// Get the latest block
	block, err := connection.RPC.Client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	return block.GasLimit()
}
