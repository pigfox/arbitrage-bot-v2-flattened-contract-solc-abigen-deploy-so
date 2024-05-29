package main

import (
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy-so/cbase"
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy-so/compile"
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy-so/config"
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy-so/connection"
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy-so/deploy"
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy-so/structs"
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy-so/verify"
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy-so/wallet"
	"fmt"
)

var contractName string = "Base" //file name w/o _flat or .sol

func setUp() {
	fmt.Println("Setting up...")
	structs.SetUp()
	wallet.SetUp()
	config.SetUp(config.Sepolia)
	connection.SetUp()
}

func main() {
	setUp()
	compile.Run(contractName)
	deploy.Run()
	verify.Run(contractName)
	cbase.Run()
}
