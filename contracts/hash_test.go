package contracts

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"testing"

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
		solsha3.String(nil),
		solsha3.Uint256(big.NewInt(0)),
	)
	fmt.Println("0x" + common.Bytes2Hex(hash))
}

func TestHash1(t *testing.T) {
	hash := solsha3.SoliditySHA3(
		solsha3.Uint256(big.NewInt(1)),
	)
	fmt.Printf("%x", hash)
}
