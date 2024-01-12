package main

import (
	"context"
	"fmt"
	"github.com/NullpointerW/ethereum-wallet-tool/pkg/tx"
	"github.com/NullpointerW/ethereum-wallet-tool/pkg/util"
	"github.com/ethereum/go-ethereum/ethclient"
	"strings"
)

func main() {
	edp := "https://opbnb-mainnet-rpc.bnbchain.org"
	rpc, err := ethclient.Dial(edp)
	if err != nil {
		fmt.Printf("Failed to connect to the Ethereum client: %v", err)
	}
	pks, _ := util.LoadLineString(".PKS")
	mpk := pks[0]
	mpk = strings.TrimPrefix(mpk, "0x")
	contract := "0xb342e7d33b806544609370271a8d074313b7bc30"
	id, _ := rpc.ChainID(context.Background())
	fmt.Println(id)
	txHash, err := tx.Transfer(mpk, contract, "0", Buffer.Get(), rpc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(txHash)

}
