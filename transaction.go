package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/astra-x/go-ethereum/common"
	"github.com/astra-x/go-ethereum/crypto"
	"github.com/astra-x/go-ethereum/ethclient"
	"github.com/astra-x/go-ethereum/core/types"
)

const (
	HTTP = "http://127.0.0.1:8545"
	address_to = "0xd08a05283ad35600ab448c08db31a7c3797c8319"
	priv_key = "d81952d9449a63525e2ef643e1b4ef7be924ac5a37602f00677c9940fa20d4cf"
		)


func sendTransaction() {
	//signedTxs := make(chan *types.Transaction,1)

	client, err := ethclient.Dial(HTTP)
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA(priv_key)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println(common.Bytes2Hex(fromAddress[:]))



	if err != nil {
		log.Fatal(err)
	}
	value := big.NewInt(1000)      // in wei
	gasLimit := big.NewInt(21000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	toAddress := common.HexToAddress(address_to)
	var data []byte
	nonce, err := client.NonceAt(context.Background(), fromAddress, nil)
	fmt.Println(nonce)
	for {

		tx := types.NewTransaction(0, nonce, toAddress, value, gasLimit, gasPrice, data)
		chainID, err := client.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
		//fmt.Println("signed tx", signedTx)
		if err != nil {
			log.Fatal(err)
		}

		err = client.SendTransaction(context.Background(), signedTx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("tx sent: ", signedTx.Hash().Hex())
		nonce=nonce+1
		fmt.Println(nonce)
	}
	balance, _ :=client.BalanceAt(context.Background(), toAddress, nil)

	fmt.Println("Balance  :", balance)
}


func main() {
	sendTransaction()
}
