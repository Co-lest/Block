package modules

import "time"

type Block struct {
	Index        int
	Timestamp    string
	Data         string
	PrevHash     string
	Hash         string
	Nonce        int
}

var Blockchain []Block

func CreateGenesisBlock() Block {
	return Block{
		Index:     0,
		Timestamp: time.Now().String(),
		Data:      "Genesis Block",
		PrevHash:  "",
		Hash:      "",
		Nonce:     0,
	}
}
