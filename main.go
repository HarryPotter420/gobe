package main

import (
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/harrypotter420/gobe/app"
)

var db *mongo.Client

func init() {
	// MongoDB connection setup
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/mydb")
	var err error
	db, err = mongo.Connect(nil, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = db.Ping(nil, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")
}

func main() {

	go chatroom.runChatRoom()
	app.StartApp()
}
