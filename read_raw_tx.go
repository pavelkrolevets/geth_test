package main

import (
	"encoding/hex"
	//"log"

	//"github.com/c3systems/vendor.bak/github.com/davecgh/go-spew/spew"
	"github.com/astra-x/go-ethereum/core/types"
	//"github.com/ethereum/go-ethereum/ethclient"
	"github.com/astra-x/go-ethereum/rlp"
	"fmt"
)

func main() {
	//client, err := ethclient.Dial("http://127.0.0.1:8545")
	//	if err != nil {
	//		log.Fatal(err)
	//}

	rawTx := "f86680098261a882520894d08a05283ad35600ab448c08db31a7c3797c831982271080820bf9a072a4e3a32fa832375b8176c63e30e9b14678b58be1b09bba6d83f4def6fba0d4a03c3c02d4b44c79142aa273bd21e98a878d924991e31c2f92703592cb33874ef2"
	right_rawTx := "f86480258082520894d08a05283ad35600ab448c08db31a7c3797c831982271080820bf9a0c39bd3cbd42c2f0a0d1c8f53d0c66b6a432a85c79dec667ab5294bbddf335e93a046aef5c02f517dbe71ce5a8447f4a6ffe9eb27332ca6d2d3555a275e7d1785f2"
	fmt.Println(len(rawTx), len(right_rawTx))
	var tx *types.Transaction

	rawTxBytes, err := hex.DecodeString(rawTx)
	rlp.DecodeBytes(rawTxBytes, &tx)

	fmt.Println(tx, err)
}
