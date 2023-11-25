package chatroom

import (
	"log"
	"sync"
	"time"

	"github.com/gofiber/contrib/websocket"
	"go.mongodb.org/mongo-driver/mongo"
)

type client struct {
	isClosing bool
	mu        sync.Mutex
	username  string
}

type Message struct {
	Username string    `bson:"username"`
	Text     string    `bson:"text"`
	Time     time.Time `bson:"time"`
}

var clients = make(map[*websocket.Conn]*client)
var register = make(chan *websocket.Conn)
var broadcast = make(chan string)
var unregister = make(chan *websocket.Conn)

func runChatRoom() {
	for {
		select {
		case connection := <-register:
			clients[connection] = &client{}
			log.Println("connection registered")

		case message := <-broadcast:
			log.Println("message received:", message)
			// Send the message to all clients
			for connection, c := range clients {
				go func(connection *websocket.Conn, c *client) { // send to each client in parallel so we don't block on a slow client
					c.mu.Lock()
					defer c.mu.Unlock()
					if c.isClosing {
						return
					}
					if err := connection.WriteMessage(websocket.TextMessage, []byte(message)); err != nil {
						c.isClosing = true
						log.Println("write error:", err)

						connection.WriteMessage(websocket.CloseMessage, []byte{})
						connection.Close()
						unregister <- connection
					}
				}(connection, c)
			}

		case connection := <-unregister:
			// Remove the client from the hub
			delete(clients, connection)

			log.Println("connection unregistered")
		}
	}
}

func SaveMessage(db *mongo.Client, username, text string) error {
	// collection := db.Database("mydb").Collection("messages")

	// message := Message{
	// 	Username: username,
	// 	Text:     text,
	// 	Time:     time.Now(),
	// }

	// _, err := collection.InsertOne(context.Background(), message)
	// if err != nil {
	// 	fmt.Printf("Error inserting message into MongoDB: %v\n", err)
	// 	return err
	// }

	// fmt.Println("Message saved to MongoDB:", message)
	return nil
}

func handleWebSocketConnection(c *websocket.Conn) {
	// When the function returns, unregister the client and close the connection
	defer func() {
		unregister <- c
		c.Close()
	}()

	// Register the client
	register <- c

	for {
		messageType, message, err := c.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Println("read error:", err)
			}

			return // Calls the deferred function, i.e. closes the connection on error
		}

		if messageType == websocket.TextMessage {
			// Broadcast the received message
			broadcast <- string(message)
		} else {
			log.Println("websocket message received of type", messageType)
		}
	}
}
