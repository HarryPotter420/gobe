package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/harrypotter420/gobe/model"
	"github.com/harrypotter420/gobe/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	userRepository *repository.UserRepository
	secretKey      = []byte("233kopow2323dfxc") // Change this to a secure secret key
)

func init() {
	// Initialize the user repository
	// Make sure to create a connection to your user database
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(nil, clientOptions)
	if err != nil {
		panic(err)
	}

	// Check the connection
	err = client.Ping(nil, nil)
	if err != nil {
		panic(err)
	}

	// Initialize the user repository with the database connection
	userRepository = repository.NewUserRepository(client.Database("mydb"))
}

func RegisterUser(c *fiber.Ctx) error {
	var user model.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}

	// Validate user input (e.g., check if required fields are present)
	if user.Username == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Username and password are required")
	}

	// Register the user
	err = userRepository.RegisterUser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}

	// Generate and send JWT
	token, err := generateJWT(user)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
	}

	return c.JSON(fiber.Map{"token": token})
}

func LoginUser(c *fiber.Ctx) error {
	var user model.User
	err := c.BodyParser(&user)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}

	// Validate user input
	if user.Username == "" || user.Password == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Username and password are required")
	}

	// Authenticate the user
	authenticatedUser, err := userRepository.GetUserByUsernameAndPassword(user.Username, user.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid credentials")
	}

	// Generate and send JWT
	token, err := generateJWT(*authenticatedUser)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Internal server error")
	}

	return c.JSON(fiber.Map{"token": token})
}

func generateJWT(user model.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.Id
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix() // Token expires in 24 hours

	// Sign the token with the secret key
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
