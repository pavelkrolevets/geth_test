package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"github.com/astra-x/go-ethereum/common"
	"github.com/astra-x/go-ethereum/crypto"
	"github.com/astra-x/go-ethereum/ethclient"
	//"time"
	"math/big"
	"github.com/astra-x/go-ethereum/core/types"
)

const (
	HTTP = "http://127.0.0.1:8545"
	address_to = "0xd08a05283ad35600ab448c08db31a7c3797c8319"
	priv_key = "d81952d9449a63525e2ef643e1b4ef7be924ac5a37602f00677c9940fa20d4cf"
		)


func sendTransactions() {
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
	value := big.NewInt(10000)      // in wei
	gasLimit := big.NewInt(210000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())

	if err != nil {
		log.Fatal(err)
	}
	toAddress := common.HexToAddress(address_to)
	var data []byte
	nonce, err := client.NonceAt(context.Background(), fromAddress, nil)
	balance_start, err := client.BalanceAt(context.Background(), toAddress, nil)

	// Add 2048 digit HEX of random data for heavy io
	data = common.Hex2Bytes("c569dc1c90c446a4cd2dc2418c328e0fc761254123c81eeb9fd6fc7e3f6b4a5e15d0283dcf80febc636171ea22a79e68686b5d4b2822ac8bde7276cc4d08c6d2c787157b070bcce12949d3d9811a3ef2824c3c25f1577e5c8fa2be6f0ab4ab5c739323921e8bd5884b7cde8fb00a0a8c178bd8ded1fd020b7deba4c3ac7ba46c8767588e84a3d7d1d11601e8db1283709172b078f14f8becac39b1792aacd2962182197d4db568abb6f6d8337b023a53b033b8aa54010daa002de6f37bac8d6e0435dffd57f36ff1912ddccf93d88a26757add6eb4d95c93b41554b170833a89eb3a77277a11df27dd6abf3aea960b6d0d924989fee379fb74a7e78270789a1c5348828f98eb9f08bb91ad5836d61e5df3e8c3b1103b5c44a12b06431fb52b1c0cd91f7a52460465eb4ca447df6a901153b37cfc6d1bf5046d337b8b950d66ba38d63a0e942c8bdbbaeb4b262abadc9b553cb1c228cbaa14c62fa303b581bda2042624913c483affd5c2853ba5bec0ffc6eaaaa90c937408522f904f97149bd73044d9e97d50190f619e2ea219850891684b91757f0d0267a7b07ebd83a4cf09358291c63594398387980806e151d5bb1f69db3d48c5061cd6f72144ebdfbc739c637bfc902520d29aa2b69ad5ace4081947eee60281d6a30141b97b55e8393294765c60752b2ef53c2109d64005cb5f4240f4419c6c289d438a93a9391ebdb0641d7a9f7af3c3aca07e3b3bd21d186184892b9e4cc593778cd89aec16ddc0007d1cff2016e8011e48c7b66d294762c6a9595e43db95cfab1b5139dfd51ce16f344cfd58642964628346b137b0c2961408c2bd4822d702b34d3830f032664a2dd77f19469f10cf673bd3e7570facb97fe5dd0f5ebd37f0c6e8dea9234dcb35a31dd273f2352b35070207c0c4aa7c7e64d59ad8f8204f81e9cc0620a6c9c414d2fcc8aeb8da792a0ec60ce30f455fc4b78157b90ba4d72913659632865421ba3c4c8672fe46188e70caebd9209ec9bc07a59217d7e6c2ff0762450fa04cf5817f79543a73ca20f2b7413e0d2dafaa09833a5d87726d938fd88d96666c9aa4af05f73860748d1eb970ba0eb2a05c2cc2980784d9cf0e1a60685cf2631efb53e58c39e85fe936d1ad63afd02c9d013e27460ad7d38d6fc7d80c8d944ee21fb61a6d7dd2165ab581332eb0823f4ab26ba1567f356f6cc9d0567fa071c3d92c3b4bf15010d96dfe4d08748cd3eabad44c6872ff8e961430dea133d53c56da4b81a58ac52e1af568f7d9e001cf7cd6bfe1c3bfeb25dd78f364274b96bfd406528530e8631d81a4c4a54734f40c5c244bc0ca86398ceafbdc6dab6680050874ee2b10ba36e383943176019a3b6022075fc8586e022f8b9cb3b9a67a493cea0156c7a70bec144191f7a453d65be37ebe01909fa46d006fc696b4f3007c4b044198a30e62")

	//Start iterating
	nonce_start := nonce
	for i:=0;i<100000;i++ {

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
		//fmt.Println("tx sent: ", signedTx.Hash().Hex())
		nonce = nonce + 1

		//fmt.Println("Balance  :", balance)
		//fmt.Println(nonce)
		}
	balance_end, err := client.BalanceAt(context.Background(), toAddress, nil)
	var count *big.Int
	count.Sub(balance_end,balance_start)
	nonce_end := nonce - nonce_start
	fmt.Println(count, nonce_end)
	}


func main() {
	sendTransactions()
	//tick := time.Tick(time.Microsecond)
	//select {
	//case <-tick:
	//
	//}
}
