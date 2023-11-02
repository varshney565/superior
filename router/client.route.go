package router

import (
	"superior/controller"
	"superior/middleware"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes func
func SetupClientRoutes(app *fiber.App) {
	client := app.Group("/client")
	client.Post("/add", middleware.Clientcheck, controller.AddClient)
	client.Get("/get", controller.GetClients)
}
