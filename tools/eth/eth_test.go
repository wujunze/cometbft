package eth

import "testing"

func TestGetBlockNumber(t *testing.T) {

	blockNumber, err := GetBlockNumber()

	if err != nil {
		t.Error("Error getting block number")
	}

	if blockNumber <= 0 {
		t.Error("Block number is not greater than 0")
	}

	t.Log("Block number is", blockNumber)
}
