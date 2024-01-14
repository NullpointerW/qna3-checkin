package main

import (
	"context"
	"fmt"
	"github.com/NullpointerW/ethereum-wallet-tool/pkg/util"
	"github.com/ethereum/go-ethereum/ethclient"
	"qna3-checkin/contract"
	"qna3-checkin/contract/api"
	"strings"
	"sync"
)

func checkin(pk string, rpc *ethclient.Client, api *api.Client) error {
	txHash, err := contract.CallCheckin(pk, rpc)
	if err != nil {
		return err
	}
	err = api.Checkin(txHash.String(), "bnb")
	if err != nil {
		return err
	}
	fmt.Println("checkin ok")
	return nil
}

func claim(pk string, rpc *ethclient.Client, api *api.Client) error {
	hid, amt, nonce, sign, err := api.ClaimStatus()
	if err != nil {
		return err
	}
	sign = strings.TrimPrefix(sign, "0x")
	txHash, err := contract.CallClaim(pk, amt, nonce, sign, rpc)
	if err != nil {
		return err
	}
	err = api.ClaimAll(hid, txHash.String())
	if err != nil {
		return err
	}
	fmt.Println("claim ok")
	return nil
}

func main() {
	//edp := "https://opbnb-mainnet-rpc.bnbchain.org"
	edp := "https://bsc-dataseed1.binance.org/"
	rpc, err := ethclient.Dial(edp)
	if err != nil {
		fmt.Printf("Failed to connect to the rpc: %v", err)
		return
	}
	id, _ := rpc.ChainID(context.Background())
	fmt.Println("chain id", id)
	pks, _ := util.LoadLineString(".PKS")
	var wg = new(sync.WaitGroup)
	for _, pk := range pks {
		pk = strings.TrimPrefix(pk, "0x")
		wg.Add(1)
		go work(pk, rpc, wg)
	}
	wg.Wait()
	fmt.Println("process fin")
}
func work(pk string, rpc *ethclient.Client, wg *sync.WaitGroup) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("checkin failed,Recovered from panic:", r)
			wg.Done()
		}
	}()
	defer wg.Done()
	c, err := api.NewClient(pk)
	if err != nil {
		fmt.Println(pk, "checkin failed:", err)
		return
	}
	err = checkin(pk, rpc, c)
	if err != nil {
		fmt.Println(pk, "checkin failed:", err)
		return
	}
	err = claim(pk, rpc, c)
	if err != nil {
		fmt.Println(pk, "claim failed:", err)
		return
	}
}
