package config

import (
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy-so/wallet"
	"fmt"
	"math/big"
)

const Sepolia string = "sepolia"

var Active string
var Map map[string]Config

func SetUp(current string) {
	fmt.Println("Config...")
	Map = make(map[string]Config)
	Map[Sepolia] = Config{
		UserWallet: wallet.Sepolia,
		NetType:    Sepolia,
		ChainID:    big.NewInt(11155111), //nolint:gomnd
	}

	Active = current
}

type Config struct {
	UserWallet wallet.Wallet
	NetType    string
	ChainID    *big.Int
}
