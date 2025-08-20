package main

import (
	"fmt"
	"math/big"

	"github.com/lmittmann/w3"
	"github.com/lmittmann/w3/module/eth"
)

func main() {
	addr := w3.A("0x0000000000000000000000000000000000000000")

	// 1. Connect to an RPC endpoint
	client, err := w3.Dial("http://mainnet.cortexlabs.ai:30089")
	if err != nil {
		// handle error
	}
	defer client.Close()

	// 2. Make a batch request
	var (
		balance *big.Int
		nonce   uint64
	)
	if err := client.Call(
		eth.Balance(addr, nil).Returns(&balance),
		eth.Nonce(addr, nil).Returns(&nonce),
	); err != nil {
		// handle error
	}

	fmt.Printf("balance: %s\nnonce: %d\n", w3.FromWei(balance, 18), nonce)
}
