package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("b8dc160c1f5d8609bd987faec1fed5f91feac0c6b2249cacda795cd115a317d1")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("nonce: %d\n", nonce)
	fmt.Printf("fromAddress: %x\n", fromAddress)

	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gas := uint64(21000)                     // in units
	gasTipCap, err := client.SuggestGasPrice(context.Background())

	//最大gas fee
	gasFeeCap := big.NewInt(38694000460)
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0xd0cC48E692B2139715F253c4Cbc002301F052B83")
	var data []byte
	//tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	tx := types.NewTx(&types.DynamicFeeTx{
		Nonce:     nonce,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Gas:       gas,
		To:        &toAddress,
		Value:     value,
		Data:      data,
	})

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewLondonSigner(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
