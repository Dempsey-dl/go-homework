package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

func test1() {
	fmt.Println(Pubaddress)
	Nonce, err := Client.PendingNonceAt(context.Background(), Pubaddress)
	if err != nil {
		log.Fatal(err)
	}

	gasprice, err := Client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	chainId, err := Client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(int64(math.Pow10(9)))
	tx := types.NewTransaction(Nonce, To, value, uint64(300000), gasprice, nil)
	signTx, err := types.SignTx(tx, types.NewEIP155Signer(chainId), PriKey)
	if err != nil {
		log.Fatal(err)
	}

	errr := Client.SendTransaction(context.Background(), signTx)
	if errr != nil {
		log.Fatal(errr)
	}
	fmt.Println("交易地址:", signTx.Hash().Hex())

}
