package wallet

import (
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy-so/structs"
	"fmt"
)

var Sepolia Wallet

func SetUp() {
	fmt.Println("Wallet...")
	Sepolia = newSepolia()
}

type Wallet struct {
	Address    string
	PrivateKey string
}

func newSepolia() Wallet {
	return Wallet{
		Address:    structs.Credentials.SepoliaAddress,
		PrivateKey: structs.Credentials.SepoliaPrivateKey,
	}
}
