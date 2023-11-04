package utils

import (
	"superior/config"
	"superior/helper"
	"superior/model"
	"time"
)

func NewBlock(txns []model.Transaction, metadata []string) model.Block {
	// generate the merkle root
	merkleRoot := helper.GenerateMerkleRoot(txns)
	// generate the DataHash
	datahash := helper.GenerateDataHash(txns)
	// find the  previousHash
	Hash := helper.PreviousHash()
	// get the address of current node
	address := config.Config("HOST") + ":" + config.Config("PORT")
	// get the current time
	currentTime := time.Now().Format("2000-01-01 15:04:05")
	//generate the header
	header := model.Header{
		Merkleroot:      merkleRoot,
		Datahash:        datahash,
		Prevhash:        Hash,
		Proposeraddress: address,
		Timestamp:       currentTime,
		Height:          len(txns),
	}

	block := model.Block{
		BlockHeader: header,
		MetaData:    metadata,
	}
	return block
}
