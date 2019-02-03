package main

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"time"
)

// HashBytes defines size in bytes of output of our hash function
const HashBytes = 32

// NonceBytes defines number of bytes to use in our nonce
const NonceBytes = 4

// TransactionOutput defines one of multiple outputs of a transaction
type TransactionOutput struct {
	value  float64
	pubkey [HashBytes]byte
}

// TransactionInput ...
type TransactionInput struct {
	number [HashBytes]byte
	output TransactionOutput
}

// Transaction defines a transaction in the bc
type Transaction struct {
	number  [HashBytes]byte
	inputs  []TransactionInput
	outputs []TransactionOutput
	sig     [HashBytes]byte
}

// Block represents a block in the blockchain
type Block struct {
	tx       Transaction
	prevHash [HashBytes]byte
	nonce    [NonceBytes]byte
	pow      [HashBytes]byte
}

func main() {
	var magic [HashBytes]byte
	var magicNonce [NonceBytes]byte
	genesisBlock := Block{prevHash: magic, nonce: magicNonce, pow: magic}
	genesisBlock.tx = Transaction{number: magic}

	fmt.Printf("tx: %x\n", genesisBlock.tx.sig)

	rand.Seed(time.Now().UnixNano())

	h := sha256.New()
	data := []byte("hello worldx\n")

	h.Write(data)
	fmt.Printf("original hash: %x\n", h.Sum(nil))

	nonce := findNonce(data)

	h.Write(nonce)
	hash := h.Sum(nil)
	fmt.Printf("nonce: %x\nHash with nonce:%x\n", nonce, hash)

}

func findNonce(data []byte) []byte {
	hasher := sha256.New()
	minVal := byte(0x07) // Optimization: Assume only first byte is compared
	nonce := make([]byte, 4)
	for {
		fmt.Print(".")
		rand.Read(nonce)
		hasher.Reset()
		hasher.Write(data)
		hasher.Write(nonce)
		hash := hasher.Sum(nil)
		if hash[0] <= minVal {
			fmt.Print("\n")
			return nonce
		}
	}
}
