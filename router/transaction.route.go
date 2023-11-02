package router

import (
	"superior/controller"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func
func TransactionRoute(app *fiber.App) {
	client := app.Group("/")
	client.Get("/taketransactions", controller.TransactionLogic)
}
