package main

import (
	"context"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	lastBlockNum, e := client.BlockByNumber(context.Background(), nil)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Printf("last block nun time:%d\n", lastBlockNum.Time())      // 1527211625
	fmt.Printf("last block hash:%s\n", lastBlockNum.Hash().String()) // 0x9e8751ebb5069389b855bba72d94902cc385042661498a415979b7b6ee9ba4b9
	trs := lastBlockNum.Transactions()
	fmt.Printf("last block transactions count:%d\n", len(trs)) // 144
	for _, tr := range trs {
		access := tr.AccessList()

		for _, acc := range access {
			fmt.Println(acc.Address.Hex())
		}
	}
	// count, err := client.TransactionCount(context.Background(), lastBlockNum.Hash())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(count) // 144
}

//curl https://cloudflare-eth.com -H 'Content-Type: application/json' -X POST --data '{"jsonrpc":"2.0","method":"eth_blockNumber","params":[],"id":1}'
