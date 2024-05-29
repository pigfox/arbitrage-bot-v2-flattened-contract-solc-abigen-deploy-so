package cbase

import (
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy-so/api"
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy-so/config"
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy-so/connection"
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy-so/structs"
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func Run() {
	fmt.Println("Running code...")
	privateKey, err := crypto.HexToECDSA(config.Map[config.Active].UserWallet.PrivateKey)
	if err != nil {
		log.Fatalf("Invalid private key: %v", err)
	}

	chainID, err := connection.RPC.Client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Failed to fetch network ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	contractAddress := common.HexToAddress(structs.OnChainContract.Address)
	instance, err := api.NewBase(contractAddress, connection.RPC.Client)
	if err != nil {
		log.Fatalf("Failed to bind to Base contract: %v", err)
	}

	x := big.NewInt(100)
	y := big.NewInt(207)
	result, err := instance.Add(auth, x, y)
	if err != nil {
		log.Fatalf("Failed to execute Add: %v", err)
	}
	fmt.Printf("Add result: %d\n", result)

	resultConcat, err := instance.Concat(nil, "qwBBBBBhuguty")
	if err != nil {
		log.Fatalf("Failed to execute Concat: %v", err)
	}
	fmt.Printf("Concat result: %s\n", resultConcat)

	q := big.NewInt(88888999999)
	tx, err := instance.SetQ(auth, q)
	if err != nil {
		log.Fatalf("Failed to execute SetQ: %v", err)
	}
	fmt.Printf("SetQ transaction hash: %s\n", tx.Hash().Hex())

	receipt, err := waitForTransaction(connection.RPC.Client, tx.Hash())
	if err != nil {
		log.Fatalf("Error waiting for transaction: %+v", err)
	}
	customReceipt := ConvertEthReceiptToCustomReceipt(receipt)
	fmt.Printf("Transaction Receipt: %+v\n", customReceipt)

	resultQ, err := instance.GetQ(nil)
	if err != nil {
		log.Fatalf("Failed to execute GetQ: %v", err)
	}
	fmt.Printf("GetQ result: %d\n", resultQ)
}

func waitForTransaction(client *ethclient.Client, txHash common.Hash) (*types.Receipt, error) {
	var txReceipt *types.Receipt
	var err error

	ctx := context.Background()
	queryTicker := time.NewTicker(time.Second) // Defines how often to query the blockchain
	defer queryTicker.Stop()

	fmt.Println("Waiting for transaction to be mined...")

	for {
		select {
		case <-queryTicker.C:
			txReceipt, err = client.TransactionReceipt(ctx, txHash)
			if err != nil {
				if err == ethereum.NotFound {
					// Transaction not found yet, continue waiting
					continue
				}
				// An actual error occurred, return it
				return nil, err
			}
			// Transaction has been mined
			fmt.Printf("Transaction %s mined in block %d\n", txHash.Hex(), txReceipt.BlockNumber.Uint64())
			return txReceipt, nil
		case <-ctx.Done():
			return nil, ctx.Err()
		}
	}
}
