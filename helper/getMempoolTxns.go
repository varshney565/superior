package helper

import (
	"encoding/json"
	"fmt"
	"superior/model"

	"github.com/tecbot/gorocksdb"
)

// func GetMempoolTxns() ([]string, error) {
// 	var transactions []string
// 	// Open a RocksDB database
// 	opts := gorocksdb.NewDefaultOptions()
// 	defer opts.Destroy()
// 	opts.SetCreateIfMissing(true)
// 	db, err := gorocksdb.OpenDb(opts, "database/transaction")
// 	if err != nil {
// 		fmt.Println("Error opening database:", err)
// 		return nil, err
// 	}
// 	defer db.Close()
// 	// Reading data
// 	readOpts := gorocksdb.NewDefaultReadOptions()
// 	defer readOpts.Destroy()
// 	// Iterating through the database
// 	iter := db.NewIterator(readOpts)
// 	defer iter.Close()
// 	cnt := 0
// 	for iter.SeekToFirst(); iter.Valid(); iter.Next() {
// 		key := iter.Key()
// 		value := iter.Value()
// 		transaction := model.Transaction{}
// 		if err := json.Unmarshal(value.Data(), &transaction); err != nil {
// 			fmt.Println("Error deserializing data", err)
// 			return nil, err
// 		}
// 		transactions = append(transactions, string(key.Data()))
// 		cnt++
// 		if cnt == 20 {
// 			break
// 		}
// 	}
// 	return transactions, nil
// }

func GetMempoolTxns() (<-chan string, <-chan bool) {
	out := make(chan string)
	done := make(chan bool)
	go func() {
		// Open a RocksDB database
		opts := gorocksdb.NewDefaultOptions()
		defer opts.Destroy()
		opts.SetCreateIfMissing(true)
		db, err := gorocksdb.OpenDb(opts, "database/transaction")
		if err != nil {
			fmt.Println("Error opening database:", err)
			done <- false
			close(done)
			close(out)
			return
		}
		defer db.Close()
		// Reading data
		readOpts := gorocksdb.NewDefaultReadOptions()
		defer readOpts.Destroy()
		iter := db.NewIterator(readOpts)
		defer iter.Close()
		cnt := 0
		for iter.SeekToFirst(); iter.Valid(); iter.Next() {
			key := iter.Key()
			value := iter.Value()
			transaction := model.Transaction{}
			if err := json.Unmarshal(value.Data(), &transaction); err != nil {
				fmt.Println("Error deserializing data", err)
				done <- false
				close(done)
				close(out)
				return
			}
			out <- string(key.Data())
			cnt++
			if cnt == 20 {
				break
			}
		}
		done <- true
		close(done)
		close(out)
	}()
	return out, done
}