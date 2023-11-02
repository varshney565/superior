package helper

func TakeTxns(out <-chan string, done <-chan bool) {
	for {
		select {
		case <-done:
			return
		case <-out:
			return
		}
	}
}
