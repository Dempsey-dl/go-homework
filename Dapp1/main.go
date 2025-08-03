package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var Client *ethclient.Client
var To common.Address
var Pubaddress common.Address
var PriKey *ecdsa.PrivateKey

func main() {

	client, err := ethclient.Dial("https://sepolia.infura.io/v3/3cd6670397b442a3b321cdb81cb3074d")
	if err != nil {
		log.Fatal(err)
	}
	Client = client
	blockNumber := big.NewInt(8703931)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Hash().Hex())
	fmt.Println(block.Time())
	fmt.Println(block.GasLimit())
	fmt.Println(block.Nonce())
	fmt.Println(len(block.Transactions()))

	To = common.HexToAddress("0x2735b9fd652C1b1286bf3656248f957f7595EFB4")
	priKey, err := crypto.HexToECDSA("cc4d1984ff43567efeb9edb41f9ee1658c6f1016e7c62a98eba7e9738a83057c")
	PriKey = priKey
	if err != nil {
		log.Fatal(err)
	}
	PubKey := PriKey.Public()
	pubKey, ok := PubKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("Cannot assert types")
	}

	Pubaddress = crypto.PubkeyToAddress(*pubKey)

	fmt.Println(Pubaddress)
	test1()
	test2()

}
