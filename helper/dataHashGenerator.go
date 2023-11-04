package helper

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"superior/model"
)

func GenerateDataHash(data []model.Transaction) string {
	var datahash string
	for i := 0; i < len(data); i++ {
		data, _ := json.Marshal(data[i])
		transaction_hash := sha256.Sum256([]byte(data))
		new_datahash := sha256.Sum256([]byte(datahash + hex.EncodeToString(transaction_hash[:])))
		datahash = hex.EncodeToString(new_datahash[:])
	}
	return datahash
}
