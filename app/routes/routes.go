package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harrypotter420/gobe/service"
)

func SetupHomeRoute(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Fiber!")
	})
}

func SetupMessagesRoute(app *fiber.App) {
	app.Get("/messages", service.GetMessages)
	// Add more message-related routes as needed
}

func SetupChatRoutes(app *fiber.App) {
	// WebSocket chat routes

}

func SetupUserRoutes(app *fiber.App) {
	// User routes
	app.Post("/register", service.RegisterUser)
	app.Post("/login", service.LoginUser)
}
