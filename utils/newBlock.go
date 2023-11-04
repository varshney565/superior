package utils

import (
	"superior/helper"
	"superior/model"
)

func NewBlock(txns []model.Transaction) model.Block {
	// generate the merkle root
	merkleRoot := helper.GenerateMerkleRoot(txns)
	// generate the DataHash
	datahash := helper.GenerateDataHash(txns)
	// find the  
	//generate the header
	header := model.Header {
		Merkleroot: merkleRoot,
		Datahash: datahash,
		Prevhash: ,
	}

	block := model.Block{}
}
