package structs

import (
	"log"
	"math/big"
	"os"

	"github.com/joho/godotenv"
)

var OnChainContract DeployedContract
var Credentials CredentialValues

type DeployedContract struct {
	Address  string
	TxHash   string
	TxCost   *big.Int
	Gas      uint64
	GasPrice *big.Int
}

type CredentialValues struct {
	InfuraAPIKey      string
	SepoliaAddress    string
	SepoliaPrivateKey string
}

func SetUp() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Get environment variables
	infuraAPIKey := "INFURA_API_KEY"
	Credentials.InfuraAPIKey = os.Getenv(infuraAPIKey)
	if len(Credentials.InfuraAPIKey) == 0 {
		panic(infuraAPIKey + " not set")
	}

	sepoliaAddress := "SEPOLIA_ADDRESS"
	Credentials.SepoliaAddress = os.Getenv(sepoliaAddress)
	if len(Credentials.SepoliaAddress) == 0 {
		panic(sepoliaAddress + " not set")
	}

	sepoliaPrivateKey := "SEPOLIA_PRIVATE_KEY"
	Credentials.SepoliaPrivateKey = os.Getenv(sepoliaPrivateKey)
	if len(Credentials.SepoliaPrivateKey) == 0 {
		panic(sepoliaPrivateKey + " not set")
	}
}
