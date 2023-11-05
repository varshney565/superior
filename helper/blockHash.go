package helper

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"superior/model"
)

func GenerateBlockHash(block model.Block) string {
	var blockhash string
	blockheader, _ := json.Marshal(block.BlockHeader)
	data, _ := json.Marshal(block.MetaData)
	blockhash_bytes := sha256.Sum256([]byte(hex.EncodeToString(data) + hex.EncodeToString(blockheader)))
	blockhash = hex.EncodeToString(blockhash_bytes[:])
	return blockhash
}
