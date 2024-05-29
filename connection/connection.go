package connection

import (
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy-so/config"
	"arbitrage-bot-v2-flattened-contract-solc-abigen-deploy-so/structs"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient" //nolint:goimports
)

var RPC Connection

type Connection struct {
	Client *ethclient.Client
}

func SetUp() {
	fmt.Println("Connection...")
	//client, err := ethclient.Dial("https://" + config.Map[config.Active].NetType + ".infura.io/v3/"+ structs.Credentials.InfuraAPIKey)
	client, err := ethclient.Dial("wss://" + config.Map[config.Active].NetType + ".infura.io/ws/v3/" + structs.Credentials.InfuraAPIKey)
	if err != nil {
		log.Fatal(err)
	}
	RPC.Client = client
}
