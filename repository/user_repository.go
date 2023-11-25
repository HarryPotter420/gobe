package repository

import (
	"context"
	"errors"

	"github.com/harrypotter420/gobe/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		collection: db.Collection("users"),
	}
}

func (repo *UserRepository) GetUserByUsernameAndPassword(username, password string) (*model.User, error) {
	filter := bson.D{{"username", username}}
	var user model.User
	err := repo.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	// Verify password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid password")
	}

	return &user, nil
}

func (repo *UserRepository) RegisterUser(user *model.User) error {
	// Check if the username already exists
	existingUser := repo.collection.FindOne(context.Background(), bson.D{{"username", user.Username}})
	if existingUser.Err() == nil {
		return errors.New("username already exists")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	// Insert the new user
	_, err = repo.collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}

	return nil
}
