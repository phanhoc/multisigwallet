package contracts

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/miguelmota/go-solidity-sha3"
	"log"
	"math/big"
	"testing"
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
	fmt.Println("gasPrice", gasPrice.Mul(gasPrice, big.NewInt(100000000)))
	// deployAddr := common.HexToAddress("0x84DA03a28495cFbEe6100Df44e8aEFBd3809C43c")
	toAddr := common.HexToAddress("0x0dCE90d5E4C5dB60c9a9c788CDe1e9468B8Ac99B")
	accKey_1, err := crypto.HexToECDSA("36b791066cdb3add0b6cbbf3ebc11f821a870f5360bb8aa54472be92ff5d5ff9")
	if err != nil {
		log.Fatal(err)
	}
	accKey_2, err := crypto.HexToECDSA("414b00a5739d608d406cf31282513102608688a5189b1a515694917a0a4e7af2")
	if err != nil {
		log.Fatal(err)
	}

	signature1 := make([]byte, 65)
	signature2 := make([]byte, 65)
	value := big.NewInt(100000000000000000)

	hash := solsha3.SoliditySHA3(
		solsha3.Address("0x0dCE90d5E4C5dB60c9a9c788CDe1e9468B8Ac99B"),
		solsha3.Uint256(big.NewInt(100000000000000000)),
		solsha3.String("0x00"),
		solsha3.Uint256(big.NewInt(0)),
	)

	// tx := types.NewTransaction(uint64(0), toAddr, value, uint64(688254), gasPrice, nil)

	signTx_1, err := crypto.Sign(hash, accKey_1)
	if err != nil {
		log.Fatal(err)
	}
	copy(signature1[:64], signTx_1)
	signature1[64] = signTx_1[64] + 27
	// signTx_1, err := tx.WithSignature(types.NewEIP155Signer(big.NewInt(3)), sig)
	// V, R, S := signTx_1.RawSignatureValues()
	// r, s := R.Bytes(), S.Bytes()
	// v := byte(V.Uint64() - 27)
	// copy(signature1[32-len(r):32], r)
	// copy(signature1[64-len(s):64], s)
	// signature1[64] = v

	signTx_2, err := crypto.Sign(hash, accKey_2)
	if err != nil {
		log.Fatal(err)
	}
	copy(signature2[:64], signTx_2)
	signature2[64] = signTx_2[64] + 27
	// signTx_2, err := tx.WithSignature(types.NewEIP155Signer(big.NewInt(3)), sig)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// V, R, S = signTx_2.RawSignatureValues()
	// r, s = R.Bytes(), S.Bytes()
	// v = byte(V.Uint64() - 27)
	// copy(signature2[32-len(r):32], r)
	// copy(signature2[64-len(s):64], s)
	// signature2[64] = v
	fmt.Println(fmt.Sprintf("signature1 0x%s", common.Bytes2Hex(signTx_1)))
	fmt.Println(fmt.Sprintf("signature2 0x%s", common.Bytes2Hex(signTx_2)))

	// ////////////////
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

	// gasPrice, err = client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice.Mul(gasPrice, big.NewInt(1000000))

	deployAddr := common.HexToAddress("0xCaBD6f693D362d6e3dA8133D928B7424Ad1A925d")

	instance, err := NewMultisigWallet(deployAddr, client)
	if err != nil {
		log.Fatal(err)
	}
	result, err := instance.Execute(auth, []byte(common.Bytes2Hex(signature1)), []byte(common.Bytes2Hex(signature2)), toAddr, value, nil)
	// result, err := instance.Execute(auth,)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result org:", result.Hash().String())

}
