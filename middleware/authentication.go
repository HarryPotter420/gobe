package middleware

import (
	"github.com/harrypotter420/gobe/repository"

	"github.com/gofiber/fiber/v2"
)

var userRepository *repository.UserRepository

func init() {
	// Initialize the user repository (similar to the message repository)
	// Make sure to create a connection to your user database
	// userRepository = repository.NewUserRepository(...)
}

func AuthenticationMiddleware(c *fiber.Ctx) error {
	// Extract credentials from the request (e.g., from headers or request body)
	username := c.Get("username")
	password := c.Get("password")

	// Validate credentials
	user, err := userRepository.GetUserByUsernameAndPassword(username, password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid credentials")
	}

	// You can set user information in the context for further processing
	c.Locals("user", user)

	return c.Next()
}
