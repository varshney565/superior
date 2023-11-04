package helper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"superior/model"
	"time"

	"github.com/tecbot/gorocksdb"
)

func TakeIps() (<-chan string, <-chan bool) {
	out := make(chan string)
	done := make(chan bool)
	go func() {
		//// Open a RocksDB database
		opts := gorocksdb.NewDefaultOptions()
		defer opts.Destroy()
		opts.SetCreateIfMissing(true)
		db, err := gorocksdb.OpenDb(opts, "database/client")
		if err != nil {
			fmt.Println("Error opening database:", err)
			done <- false
			close(done)
			close(out)
			return
		}
		defer db.Close()
		//// Reading data
		readOpts := gorocksdb.NewDefaultReadOptions()
		defer readOpts.Destroy()
		iter := db.NewIterator(readOpts)
		defer iter.Close()
		cnt := 0
		for iter.SeekToFirst(); iter.Valid(); iter.Next() {
			value := iter.Value()
			node := model.Client{}
			if err := json.Unmarshal(value.Data(), &node); err != nil {
				fmt.Println("Error deserializing data", err)
				done <- false
				close(done)
				close(out)
				return
			}
			ip := node.IP + ":" + fmt.Sprint(node.PORT)
			url := "http://" + ip + "/alive"
			client := &http.Client{
				Timeout: 10 * time.Millisecond,
			}
			res, err := client.Head(url)
			if err != nil || res.StatusCode != 200 {
				continue
			}
			out <- ip
			cnt++
			if cnt == 5 {
				break
			}
		}
		done <- true
		close(done)
		close(out)
	}()
	return out, done
}
