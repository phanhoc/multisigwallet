package main

import (
	"log"
	"github.com/ethereum/go-ethereum/crypto"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"math/big"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"context"
	mw "./contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/davecgh/go-spew/spew"
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

	owners := make([]common.Address, 0, 3)
	owners = append(owners, common.HexToAddress("0x7CE414E6B43027d893A6AE88a11DF350F4dB4C3d"))
	owners = append(owners, common.HexToAddress("0xdAFCd8743DaeaFF8AE4ee3C4F096849601Da52BF"))
	owners = append(owners, common.HexToAddress("0x6Cc44cacF69FeA9d439A37fb969E2bD118C5ec56"))
	fmt.Println("owners:", spew.Sdump(owners))
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(688254) // in units
	auth.GasPrice = gasPrice

	address, transaction, _, err := mw.DeployMultisigWallet(
		auth,
		client,
		owners,
		big.NewInt(2), )
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fmt.Sprintf("Transaction hash: %s, Contract address: %s", transaction.Hash().String(), address.String()))
}
