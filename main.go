package main

import (
	"fmt"
	"time"
)

func main() {

	bc := NewBlockchain()
	bc.AddBlock("Send 1 BTC to Ivan")
	bc.AddBlock("Send 1 more to Ivan")

	for _, block := range bc.blocks {
		fmt.Printf("Prev. hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Printf("Timestamp: %s\n", time.Unix(block.Timestamp, 0))
		fmt.Println()
	}
}
