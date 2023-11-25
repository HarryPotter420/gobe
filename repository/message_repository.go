package repository

import (
	"github.com/harrypotter420/gobe/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type MessageRepository struct {
	collection *mongo.Collection
}

func NewMessageRepository(db *mongo.Database) *MessageRepository {
	return &MessageRepository{
		collection: db.Collection("messages"),
	}
}

func (repo *MessageRepository) GetAllMessages() ([]model.Message, error) {
	cursor, err := repo.collection.Find(nil, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(nil)

	var messages []model.Message
	if err := cursor.All(nil, &messages); err != nil {
		return nil, err
	}

	return messages, nil
}
