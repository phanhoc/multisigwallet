package contracts

import (
	"testing"
		"log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"context"
)

func TestDeployMultisigWallet(t *testing.T) {
	client, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("gasPrice", gasPrice.Uint64())
	//deployAddr := common.HexToAddress("0x84DA03a28495cFbEe6100Df44e8aEFBd3809C43c")
	toAddr := common.HexToAddress("0xfdcBc8385b0a629D6D96Af1fBFD5b04F7F37aABE")
	accKey_1, err := crypto.HexToECDSA("7b2a7f6491f593669aafa14e7a041c0a8da37c6f3c083043b896f34f9d8fd1c7")
	if err != nil {
		log.Fatal(err)
	}
	accKey_2, err := crypto.HexToECDSA("554b02366ad76b45eae0508ecf12f48787af8fd4c43cb5dad70e919faa560d8f")
	if err != nil {
		log.Fatal(err)
	}

	signature1 := make([]byte, 65)
	signature2 := make([]byte, 65)
	value := big.NewInt(100000000000000000)
	tx := types.NewTransaction(uint64(0), toAddr, value, uint64(688254), gasPrice, nil)

	signTx_1, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(3)), accKey_1)
	if err != nil {
		log.Fatal(err)
	}
	V, R, S := signTx_1.RawSignatureValues()
	r, s := R.Bytes(), S.Bytes()
	v := byte(V.Uint64() - 27)
	copy(signature1[32-len(r):32], r)
	copy(signature1[64-len(s):64], s)
	signature1[64] = v

	signTx_2, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(3)), accKey_2)
	if err != nil {
		log.Fatal(err)
	}
	V, R, S = signTx_2.RawSignatureValues()
	r, s = R.Bytes(), S.Bytes()
	v = byte(V.Uint64() - 27)
	copy(signature2[32-len(r):32], r)
	copy(signature2[64-len(s):64], s)
	signature2[64] = v
	fmt.Println(fmt.Sprintf("signature1 0x%s", common.Bytes2Hex(signature1)))
	fmt.Println(fmt.Sprintf("signature2 0x%s", common.Bytes2Hex(signature2)))
}
