package controller

import (
	"math/rand"
	"strconv"
	"superior/config"
	"superior/helper"
	"superior/model"
	"superior/utils"

	"github.com/gofiber/fiber/v2"
)

func TransactionLogic(c *fiber.Ctx) error {
	//// take the transactions from the memepool
	out, done, wg := helper.GetMempoolTxns()
	var txnsMetaData []string
	var txns []model.Transaction
	err := -1
	go func() {
		for {
			select {
			case res := <-done:
				if res == false {
					//error is there
					err = 1
				}
				err = 0
				return
			case fun := <-out:
				metadata, txn := fun()
				txns = append(txns, txn)
				txnsMetaData = append(txnsMetaData, metadata)
				(*wg).Done()
			}
		}
	}()

	//// take ips after sending a ping pong message
	out_, done_ := helper.TakeIps()
	var ips []string
	err_ := -1
	go func() {
		for {
			select {
			case res := <-done_:
				if res == false {
					//error is there
					err_ = 1
				}
				err_ = 0
				return
			case ip := <-out_:
				ips = append(ips, ip)
			}
		}
	}()

	//// make sure ips and txns are collected.
	for {
		if err > -1 {
			break
		}
	}
	for {
		if err_ > -1 {
			break
		}
	}
	if err == 1 || err_ == 1 {
		c.Status(500).JSON(fiber.Map{
			"message": "Internal Server Error",
		})
	}

	//// Generate Groups
	utils.GroupGenerator(&ips)
	groups, _ := strconv.Atoi(config.Config("GROUPS"))
	param, _ := strconv.Atoi(config.Config("PARAM"))

	//// select admins
	var admins []int
	for i := 0; i < param; i++ {
		random := rand.Intn(100000000) % groups
		admins = append(admins, random)
	}

	//// create the block
	utils.NewBlock(txns)
	//// brodcast the block
	for i := 0; i < groups; i++ {
		for j := 0; j < param; j++ {
			//i-th group and j-th param
			//for the jth parameter admins[j] is the admin
		}
	}
	return c.Status(200).JSON(txnsMetaData)
}
