package verify

import (
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy-so/connection"
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy-so/structs"
	"context"
	"fmt"
	"log"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

func Run(contractName string) {
	fmt.Println("Contract verifying...")
	contractAddress := common.HexToAddress(structs.OnChainContract.Address)
	bytecode, err := connection.RPC.Client.CodeAt(context.Background(), contractAddress, nil) // nil is the latest block
	if err != nil {
		log.Fatalf("Failed to get contract bytecode: %v", err)
	}

	if len(bytecode) == 0 {
		log.Fatalf("No contract xcode found at the given address")
	}

	//fmt.Printf("Contract bytecode: %x\n", bytecode)

	compiledBytecode := getBin(contractName)
	//fmt.Printf("Compilded bytecode: %x\n", compiledBytecode)
	compiledBytecodeBytes := common.FromHex(compiledBytecode)

	// Compare the deployed bytecode with the compiled bytecode
	if string(bytecode) == string(compiledBytecodeBytes) {
		fmt.Println("The deployed contract bytecode matches the compiled bytecode.")
	} else {
		fmt.Println("****************************ERROR****************************************")
		fmt.Println("The deployed contract bytecode does not match the compiled bytecode.")
		fmt.Println("Length deployed contract bytecode", len(string(bytecode)), "Length compiled bytecode", len(string(compiledBytecodeBytes)))
		fmt.Println("****************************END ERROR****************************************")
	}
}

func getBin(contractName string) string {
	path, err := os.Getwd()
	if err != nil {
		log.Fatalf("Error:%s", err)
	}
	filePath := path + "/solc-output/" + fmt.Sprintf("%s.bin", contractName)

	// Read the file contents
	binData, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read bin file: %v", err)
	}

	// Return the contents as a string
	return string(binData)
}
