package app

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/harrypotter420/gobe/middleware"
	"github.com/harrypotter420/gobe/routes"
)

func SetupRoutes(app *fiber.App) {
	// Set up your routes here using the routes package
	routes.SetupHomeRoute(app)
	routes.SetupUserRoutes(app)
	app.Use(middleware.AuthenticationMiddleware)
	routes.SetupMessagesRoute(app)
	// Add more routes as needed
}

func StartApp() {
	app := fiber.New()
	SetupRoutes(app)

	port := 3000
	fmt.Printf("Server is running on http://localhost:%d...\n", port)

	err := app.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("Error:", err)
	}
}
