package service

import (
	"github.com/harrypotter420/gobe/repository"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var messageRepo *repository.MessageRepository

func init() {
	// Initialize the message repository
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(nil, clientOptions)
	if err != nil {
		panic(err)
	}
	messageRepo = repository.NewMessageRepository(client.Database("mydb"))
}

func GetMessages(c *fiber.Ctx) error {
	// Fetch messages from the repository
	messages, err := messageRepo.GetAllMessages()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Error fetching messages")
	}

	return c.JSON(messages)
}
