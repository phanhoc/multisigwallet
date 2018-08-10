package contracts

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/miguelmota/go-solidity-sha3"
	"log"
	"math/big"
	"testing"
	"time"
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
		// solsha3.String("0x00"),
		// solsha3.Uint256(big.NewInt(0)),
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
	auth.GasPrice = gasPrice

	// deployAddr := common.HexToAddress("0xCaBD6f693D362d6e3dA8133D928B7424Ad1A925d")
	deployAddr := common.HexToAddress("0x4740CB6fBD8E544e1a9AD1cc56cF5d476281CE26") // 0x21Fca69923dB273603d9b7f402F164918C058d58

	instance, err := NewMultisigWalletV2(deployAddr, client)
	if err != nil {
		log.Fatal(err)
	}
	result, err := instance.Execute(auth, signature1, signature2, toAddr, value, nil)
	// result, err := instance.Execute(auth, []byte(common.Bytes2Hex(signature1)), []byte(common.Bytes2Hex(signature2)), toAddr, value, nil)
	// result, err := instance.Execute(auth,)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result org:", result.Hash().String())

}

func TestBroadcast(t *testing.T) {
	client, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
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

	// gasPrice, err = client.SuggestGasPrice(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	// deployAddr := common.HexToAddress("0xCaBD6f693D362d6e3dA8133D928B7424Ad1A925d")
	deployAddr := common.HexToAddress("0x4740CB6fBD8E544e1a9AD1cc56cF5d476281CE26") // 0xab2db52b1d0B562687F17f6b38d08Dc9ee4764de

	instance, err := NewMultisigWalletV2(deployAddr, client)
	if err != nil {
		log.Fatal(err)
	}
	// signTx_1 := []byte("0x003f7ba0427640a55bedbb54f0061794a6355a88076cad0c371d5b9bc25028ea46f58ca457327ef0e5b12a8ee1804d23c953280d16926448dc07048a9bb2c9eb1c")
	signTx_1 := common.Hex2Bytes("45c77f91d4f4217e205f7890c4a3767dc1f32a9eaa0b5b006d6ebd84c08b72015696f5a8ccbe0edbc34c8e4cb788ada36375d0ee2039cffbb92043774710d9291b")
	// signTx_2 := []byte("0xcc0b8e81f510980df4f1269f9b778133643c75d60df9d5039d3c5d600f2809bd1661f67f72b48d068b6e35914821ab6db91fb19d06fbf77cd114a9e4a220e5e01b")
	signTx_2 := common.Hex2Bytes("d32563c800b6551be3a71758eb45c127c760610e589bec12455ed4e9b1e5190423946ed956cc4c3c7c7342ef5e615c6347eb66b5a781d170bc0e6eb1eb06e2191c")
	toAddr := common.HexToAddress("0x7CE414E6B43027d893A6AE88a11DF350F4dB4C3d")
	value := big.NewInt(100000000000000000)
	// a := "0x" + common.Bytes2Hex([]byte("abcde35f123"))
	_, err = instance.Execute(auth, signTx_1, signTx_2, toAddr, value, []byte("abcde35f123"))
	// result, err := instance.Execute(auth, []byte(common.Bytes2Hex(signature1)), []byte(common.Bytes2Hex(signature2)), toAddr, value, nil)
	// result, err := instance.Execute(auth,)
	if err != nil {
		// log.Fatal(err)
	}

	// fmt.Println("Result org:", result.Hash().String())
	time.Sleep(100 * time.Millisecond)
}

func TestHexS(t *testing.T) {
	client, err := ethclient.Dial("HTTP://127.0.0.1:7545")
	if err != nil {
		log.Fatal(err)
	}
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	privateKey, err := crypto.HexToECDSA("4ba4558bcc05b364a514abcc4cc60352b74e7da1e2e1e5aa520a82f58f78ac51")
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
	auth.GasPrice = gasPrice

	// deployAddr := common.HexToAddress("0xCaBD6f693D362d6e3dA8133D928B7424Ad1A925d")
	deployAddr := common.HexToAddress("0x8c3249a46E0317Ad000702A3Bc64488145e0E78B") // 0xab2db52b1d0B562687F17f6b38d08Dc9ee4764de

	instance, err := NewCheckData(deployAddr, client)
	if err != nil {
		log.Fatal(err)
	}

	trans, err := instance.SetBytes(auth, []byte("test"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(trans.Hash().String())
	// fmt.Println(event.Event.Signature2)
	// fmt.Println(event.Event.Value)
	// fmt.Println(event.Event.Destination)
	// fmt.Println(event.Event.Data)
}

func TestData(t *testing.T) {
	// fmt.Println(common.Bytes2Hex([]byte("test")))
	data, _ := hex.DecodeString("74657374")
	fmt.Println(string(data))
}
