package main

import (
	"fmt"
	"strconv"

	"github.com/darshanman40/golang-blockchain/pkg/entity"
)

func main() {
	chain := entity.InitBlockChain()
	chain.AddBlock("First block after Genesis")
	chain.AddBlock("Second block after Genesis")
	chain.AddBlock("Third block after Genesis")

	for _, block := range chain.Blocks {
		fmt.Printf("Previous Hash: %x\n", block.PrevHash)
		fmt.Printf("Data in Block: %x\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Print("block ends \n")

		pow := entity.NewProof(block)
		fmt.Printf("PoW: %s\n", strconv.FormatBool((pow.Validate())))
		fmt.Println()
	}
}
