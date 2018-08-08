package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"fmt"
	"context"
	mw "./contracts"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	client, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("cf437654cd2a2faec8982deafbcbfbb17314b0f80bbfd952af5f0270812d968c")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	deployAddr := common.HexToAddress("0x84DA03a28495cFbEe6100Df44e8aEFBd3809C43c")

	instance, err := mw.NewMultisigWallet(deployAddr, client)
	if err != nil {
		log.Fatal(err)
	}
	result, err := instance.Nonce(&bind.CallOpts{})
	// result, err := instance.Execute(auth,)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result org:", result)
}
