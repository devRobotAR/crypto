package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("We have a connection")
	_ = client

	account := common.HexToAddress("0x25CEF2cD72A68b481927b17Cc3b0B046dD3aB823")

	balance, err := client.BalanceAt(context.Background(), account, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Balance: %v\n", balance)

	blockNumber := big.NewInt(0)

	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Balance At Block %v: %v\n", blockNumber, balanceAt)

	fbalance := new(big.Float)
	fbalance.SetString(balanceAt.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))

	fmt.Printf("Etherum Balance At Block %v: %v\n", blockNumber, ethValue)

	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)

	fmt.Printf("Pending Balance In Account: %v\n", pendingBalance)

	privateKey, err := crypto.GenerateKey()

	if err != nil {
		log.Fatal(err)
	}

	privateKeyBytes := crypto.FromECDSA(privateKey)

	fmt.Printf("Private Key Generated (Bytes): %v\n", hexutil.Encode(privateKeyBytes)[2:])

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Printf("Public Key Generated (Bytes): %v\n", hexutil.Encode(publicKeyBytes)[4:])

	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Printf("Public Address Generated: %v\n", address)
}
