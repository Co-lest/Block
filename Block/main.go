package main

import (
	"fmt"
	"blockchain/genesis"
	"blockchain/generate"
)



func main() {
	genesisBlock := genesis.CreateGenesisBlock()
	genesisBlock.Hash, genesisBlock.Nonce = generate.ProofOfWork(genesisBlock, 2)
	genesis.Blockchain = append(genesis.Blockchain, genesisBlock)

	fmt.Println("Genesis Block created:")
	fmt.Println(genesisBlock)

	difficulty := 2

	for i := 1; i <= 5; i++ {
		newData := fmt.Sprintf("Block %d Data", i)
		newBlock := generate.GenerateBlock(genesis.Blockchain[len(genesis.Blockchain)-1], newData, difficulty)
		generate.AddBlock(newBlock, difficulty)
		fmt.Printf("Block %d added: %v\n", newBlock.Index, newBlock)
	}

	fmt.Println("\nBlockchain:")
	for _, block := range genesis.Blockchain {
		fmt.Printf("Index: %d\nTimestamp: %s\nData: %s\nPrevHash: %s\nHash: %s\nNonce: %d\n\n",
			block.Index, block.Timestamp, block.Data, block.PrevHash, block.Hash, block.Nonce)
	}
}

