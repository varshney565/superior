package controller

import (
	"superior/helper"

	"github.com/gofiber/fiber/v2"
)

func TransactionLogic(c *fiber.Ctx) error {
	//// take the transactions from the memepool

	// transactions, err := helper.GetMempoolTxns()
	// if err != nil {
	// 	return c.Status(500).JSON(fiber.Map{
	// 		"message": err,
	// 	})
	// }
	// return c.Status(200).JSON(transactions)
	out, done := helper.GetMempoolTxns()
	txn := helper.TakeTxns(&out, &done)
	//// take ips after sending a ping pong message

	//create groups

	//select admins

	//brodcast the block
}
