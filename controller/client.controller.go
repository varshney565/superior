package controller

import (
	"encoding/json"
	"fmt"

	"superior/model"
	"superior/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/tecbot/gorocksdb"
)

func AddClient(c *fiber.Ctx) error {
	// get a Id
	id := utils.GenerateId()

	// Open a RocksDB database
	opts := gorocksdb.NewDefaultOptions()
	defer opts.Destroy()
	opts.SetCreateIfMissing(true)
	db, err := gorocksdb.OpenDb(opts, "database/client")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "DB Error while opening the database !",
		})
	}
	defer db.Close()

	//Fetching the client data from the post body
	client := model.Client{}
	c.BodyParser(&client)

	// Generating the next id and Serielizing the struct
	clientJSON, err := json.Marshal(client)
	if err != nil {
		fmt.Println("Error serializing client:", err)
		return c.Status(500).JSON(fiber.Map{
			"Error": err,
		})
	}

	// Writing data to the db
	writeOpts := gorocksdb.NewDefaultWriteOptions()
	defer writeOpts.Destroy()
	err = db.Put(writeOpts, []byte(id), clientJSON)
	if err != nil {
		fmt.Println("Error writing data:", err)
		return c.Status(500).JSON(fiber.Map{
			"Error": err,
		})
	}

	// Success Response
	return c.Status(200).JSON(fiber.Map{
		"message": "Client added successfully!",
	})
}

func GetClients(c *fiber.Ctx) error {
	// Open a RocksDB database
	opts := gorocksdb.NewDefaultOptions()
	defer opts.Destroy()
	db, err := gorocksdb.OpenDb(opts, "database/client")
	if err != nil {
		fmt.Println("Error opening database:", err)
		return c.Status(500).JSON(fiber.Map{
			"message": "DB Error while opening the database!",
		})
	}
	defer db.Close()

	// Reading data
	readOpts := gorocksdb.NewDefaultReadOptions()
	defer readOpts.Destroy()

	// Iterating through the database
	iter := db.NewIterator(readOpts)
	defer iter.Close()

	var ips []struct {
		ID      string
		Address model.Client
	}

	for iter.SeekToFirst(); iter.Valid(); iter.Next() {
		key := iter.Key()
		value := iter.Value()
		client := model.Client{}
		if err := json.Unmarshal(value.Data(), &client); err != nil {
			fmt.Println("Error deserializing data", err)
			return c.Status(500).JSON(fiber.Map{
				"Err": err,
			})
		}
		ips = append(ips, struct {
			ID      string
			Address model.Client
		}{
			ID:      string(key.Data()),
			Address: client,
		})
	}

	return c.Status(200).JSON(ips)
}
