package model

type Transaction struct {
	Transaction_id string
	Sender_id      string
	Receiver_id    string
	Amount         int
	Timestamp      string // You can use a Unix timestamp
}
