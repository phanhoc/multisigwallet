package contracts

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"testing"

	"bytes"
	"crypto/ecdsa"
	"github.com/miguelmota/go-solidity-sha3"
)

func TestHash(t *testing.T) {
	toAddr := common.HexToAddress("0x0dCE90d5E4C5dB60c9a9c788CDe1e9468B8Ac99B")
	accKey_1, err := crypto.HexToECDSA("36b791066cdb3add0b6cbbf3ebc11f821a870f5360bb8aa54472be92ff5d5ff9")
	if err != nil {
		log.Fatal(err)
	}
	value := big.NewInt(100000000000000000)
	tx := types.NewTransaction(uint64(0), toAddr, value, uint64(688254), big.NewInt(10), nil)
	signTx_1, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(3)), accKey_1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(signTx_1.Hash().String())
}

func TestNewHash(t *testing.T) {
	hash := solsha3.SoliditySHA3(
		solsha3.Address("0x0dCE90d5E4C5dB60c9a9c788CDe1e9468B8Ac99B"),
		solsha3.Uint256(big.NewInt(100000000000000000)),
		// solsha3.String(nil),
		// solsha3.Uint256(big.NewInt(0)),
	)
	fmt.Println("0x" + common.Bytes2Hex(hash))
	accKey_1, err := crypto.HexToECDSA("36b791066cdb3add0b6cbbf3ebc11f821a870f5360bb8aa54472be92ff5d5ff9")
	if err != nil {
		log.Fatal(err)
	}
	publicKey := accKey_1.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)

	signTx_1, err := crypto.Sign(hash, accKey_1)
	if err != nil {
		log.Fatal(err)
	}
	signTx_1[64] = signTx_1[64]
	sigPublicKey, err := crypto.Ecrecover(hash, signTx_1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(common.Bytes2Hex(sigPublicKey))
	matches := bytes.Equal(sigPublicKey, publicKeyBytes)
	pubKey, _ := crypto.UnmarshalPubkey(sigPublicKey)
	fmt.Println(matches)
	fmt.Println(crypto.PubkeyToAddress(*pubKey).String())
}

func TestHash1(t *testing.T) {
	hash := solsha3.SoliditySHA3(
		solsha3.Uint256(big.NewInt(1)),
	)
	fmt.Printf("%x", hash)
}
