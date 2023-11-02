package main

import (
	"superior/config"
	"superior/router"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()
	router.SetupClientRoutes(app)
	router.TransactionRoute(app)
	port := config.Config("PORT")
	app.Listen(":" + port)
}
