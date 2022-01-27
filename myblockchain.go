package main

import (
	"fmt"
	"time"
)

type Block struct {
	nounce       int
	previousHash string
	timestamp    int64
	transactions []string
}

func NewBlock(nounce int, previousHash string) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nounce = nounce
	b.previousHash = previousHash
	return b
}

func (b *Block) Print() {
	fmt.Printf("timestamp          %d\n", b.timestamp)
	fmt.Printf("nounce             %d\n", b.nounce)
	fmt.Printf("previous hash      %s\n", b.previousHash)
	fmt.Printf("transactions       %s\n", b.transactions)
}

func main() {
	b := NewBlock(0, "init hash")
	b.Print()
}
