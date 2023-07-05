package eth

import (
	"log"

	"context"
	"github.com/ethereum/go-ethereum/ethclient"
)

var client *ethclient.Client

func init() {
	var err error
	client, err = ethclient.Dial("https://ethereum.blockpi.network/v1/rpc/d71c9b30348a6bc91cfc08205fc2181ce4b06edc")
	if err != nil {
		log.Fatal(err)
	}
}

func GetBlockNumber() (int64, error) {
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		return 0, err
	}

	return header.Number.Int64(), nil

}
