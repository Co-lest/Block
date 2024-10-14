package modules

import (
	"fmt"
	"crypto/sha256"
	"bytes"
	"time"
)

func CalculateHash(block Block) string {
	record := fmt.Sprintf("%d%s%s%s%d", block.Index, block.Timestamp, block.Data, block.PrevHash, block.Nonce)
	h := sha256.New()
	h.Write([]byte(record))
	hashed := h.Sum(nil)
	return fmt.Sprintf("%x", hashed)
}

func ProofOfWork(block Block, difficulty int) (string, int) {
	var hash string
	nonce := 0
	prefix := bytes.Repeat([]byte("0"), difficulty)
	for {
		block.Nonce = nonce
		hash = CalculateHash(block)
		if bytes.HasPrefix([]byte(hash), prefix) {
			break
		} else {
			nonce++
		}
	}
	return hash, nonce
}

func GenerateBlock(oldBlock Block, Data string, difficulty int) Block {
	var newBlock Block

	newBlock.Index = oldBlock.Index + 1
	newBlock.Timestamp = time.Now().String()
	newBlock.Data = Data
	newBlock.PrevHash = oldBlock.Hash
	newBlock.Hash, newBlock.Nonce = ProofOfWork(newBlock, difficulty)

	return newBlock
}

func IsBlockValid(newBlock, oldBlock Block, difficulty int) bool {
	if oldBlock.Index+1 != newBlock.Index {
		return false
	}

	if oldBlock.Hash != newBlock.PrevHash {
		return false
	}

	if CalculateHash(newBlock) != newBlock.Hash {
		return false
	}

	prefix := bytes.Repeat([]byte("0"), difficulty)
	return !bytes.HasPrefix([]byte(newBlock.Hash), prefix)
}

func AddBlock(newBlock Block, difficulty int) {
	oldBlock := Blockchain[len(Blockchain)-1]
	if IsBlockValid(newBlock, oldBlock, difficulty) {
		Blockchain = append(Blockchain, newBlock)
	}
}
