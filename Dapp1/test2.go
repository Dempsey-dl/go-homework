package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	Version "test/version"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/core/types"
)

func test2() {
	//部署合约

	nonce, err := Client.PendingNonceAt(context.Background(), Pubaddress)
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

	signBind, err := bind.NewKeyedTransactorWithChainID(PriKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	signBind.Nonce = big.NewInt(int64(nonce))
	signBind.GasLimit = uint64(3000000)
	signBind.GasPrice = gasprice
	signBind.Value = big.NewInt(0)

	contractAddress, tx, instance, err := Version.DeployVersion(signBind, Client, "V1.0")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Version Contract Address", contractAddress)
	fmt.Println("translation Hash", tx.Hash().Hex())

	receipt, err := bind.WaitMined(context.Background(), Client, tx)
	if err != nil {
		log.Fatal(err)
	}
	if receipt.Status != types.ReceiptStatusSuccessful {
		log.Fatal("Deployment failed")
	}
	fmt.Println("✅ Contract deployed at block:", receipt.BlockNumber)
	fmt.Println("📜 Final Contract Address:", contractAddress.Hex())

	opt, err := bind.NewKeyedTransactorWithChainID(PriKey, chainId)
	if err != nil {
		log.Fatal(err)
	}
	setVersionTx, err := instance.SetVersion(opt, "V2.0")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(setVersionTx.Hash().Hex())

	receipt1, err := bind.WaitMined(context.Background(), Client, setVersionTx)
	if err != nil {
		log.Fatal(err)
	}
	if receipt1.Status != types.ReceiptStatusSuccessful {
		log.Fatal("Deployment failed")
	}

	version, err := instance.GetVersion(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("version", version)

}
