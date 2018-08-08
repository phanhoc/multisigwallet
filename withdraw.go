package main

import (
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"context"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"github.com/ethereum/go-ethereum/crypto"
	"fmt"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	mw "./contracts"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	client, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil {
		log.Fatal(err)
	}
	deployAddr := common.HexToAddress("0x84DA03a28495cFbEe6100Df44e8aEFBd3809C43c")
	toAddr := common.HexToAddress("0x0dCE90d5E4C5dB60c9a9c788CDe1e9468B8Ac99B")
	// /////
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
	nonceFee, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonceFee))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(688254) // in units
	auth.GasPrice = gasPrice

	instance, err := mw.NewMultisigWallet(deployAddr, client)
	if err != nil {
		log.Fatal(err)
	}
	nonce, err := instance.Nonce(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("nonce: %d", nonce)
	// ////

	accKey_1, err := crypto.HexToECDSA("36b791066cdb3add0b6cbbf3ebc11f821a870f5360bb8aa54472be92ff5d5ff9")
	if err != nil {
		log.Fatal(err)
	}
	accKey_2, err := crypto.HexToECDSA("414b00a5739d608d406cf31282513102608688a5189b1a515694917a0a4e7af2")
	if err != nil {
		log.Fatal(err)
	}
	//
	// publicKey := privateKey.Public()
	// publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	// if !ok {
	// 	log.Fatal("error casting public key to ECDSA")
	// }

	// nonce, err := client.PendingNonceAt(context.Background(), deployAddr)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("nonce: %d", nonce)
	sigV := make([]uint8, 0, 3)
	sigR := [][32]byte{}
	sigS := [][32]byte{}
	value := big.NewInt(100000000000000000)
	tx := types.NewTransaction(nonce.Uint64(), toAddr, value, 4, big.NewInt(688254), nil)

	signTx_1, err := types.SignTx(tx, types.HomesteadSigner{}, accKey_1)
	if err != nil {
		log.Fatal(err)
	}
	v, r, s := signTx_1.RawSignatureValues()
	sigV = append(sigV, uint8(v.Uint64()))
	var temp1 [32]byte
	var temp2 [32]byte
	var temp3 [32]byte
	var temp4 [32]byte
	copy(temp1[:], []byte("0x"+common.Bytes2Hex(r.Bytes())))
	fmt.Printf("temp1", "0x"+common.Bytes2Hex(r.Bytes()))
	sigR = append(sigR, temp1)
	copy(temp2[:], []byte("0x"+common.Bytes2Hex(s.Bytes())))
	fmt.Printf("temp2", spew.Sdump(temp2))
	sigS = append(sigS, temp2)
	fmt.Printf("sigV", spew.Sdump(sigV))
	fmt.Printf("sigR", spew.Sdump(sigR))
	fmt.Printf("sigS", spew.Sdump(sigS))
	signTx_2, err := types.SignTx(tx, types.HomesteadSigner{}, accKey_2)
	if err != nil {
		log.Fatal(err)
	}
	v, r, s = signTx_2.RawSignatureValues()
	sigV = append(sigV, uint8(v.Uint64()))
	copy(temp3[:], []byte("0x"+common.Bytes2Hex(r.Bytes())))
	fmt.Printf("temp3", spew.Sdump(temp3))
	sigR = append(sigR, temp3)
	copy(temp4[:], []byte("0x"+common.Bytes2Hex(s.Bytes())))
	fmt.Printf("temp4", spew.Sdump(temp4))
	sigS = append(sigS, temp4)
	fmt.Printf("sigV", spew.Sdump(sigV))
	fmt.Printf("sigR", spew.Sdump(sigR))
	fmt.Printf("sigS", spew.Sdump(sigS))

	// ---------------
	result, err := instance.Execute(auth, sigV, sigR, sigS, toAddr, value, nil)
	// result, err := instance.Execute(auth,)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result org:", result.Hash().String())
}
