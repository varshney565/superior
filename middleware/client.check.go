package middleware

import (
	"fmt"
	"superior/model"

	"github.com/gofiber/fiber/v2"
)

func Clientcheck(c *fiber.Ctx) error {
	//get the request body
	client := model.Client{}
	if err := c.BodyParser(&client); err != nil {
		fmt.Println("Error fething data of client :", err)
		return c.Status(400).JSON(fiber.Map{
			"Error": err,
		})
	}
	return c.Next()
}
