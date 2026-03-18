package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/dsawma/terminal_chatroom/internal/auth"
	"github.com/dsawma/terminal_chatroom/internal/database"
	"github.com/dsawma/terminal_chatroom/internal/pubsub"
	"github.com/dsawma/terminal_chatroom/internal/chatlogic"
	"github.com/dsawma/terminal_chatroom/internal/routing"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL is missing")
	}
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	q := database.New(db)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	userName, err := auth.Login(ctx, q)
	if err != nil {
		log.Fatalf("Could not find User: %v", err)
	}

	fmt.Println("Starting Chat client...")
	connectStr := "amqp://guest:guest@localhost:5672/"
	connection, err := amqp.Dial(connectStr)
	if err != nil {
		log.Fatalf("Could not create connection: %v", err)
	}
	defer connection.Close()

	fmt.Println("Connection Successful")
	ch, err := connection.Channel()
	if err != nil {
		log.Fatalf("could not open channel: %v", err)
	}
	defer ch.Close()
	roomName,err := auth.JoinRoom(ctx,q)
	if err != nil {
		log.Fatalf("Could not find User: %v", err)
	}

	chatState := chatlogic.NewChatState(userName, roomName)
	err = pubsub.SubscribeGob(connection, routing.ExchangeChatDirect, "pause." +userName, routing.PauseKey,pubsub.TransientQueue, handlerPause(chatState) )
	if err != nil {
		log.Fatalf("could not make queue: %v", err)
	}

	err = pubsub.SubscribeGob(connection, routing.ExchangeChatTopic, "chat_queue." + userName, "sends_msg.*",pubsub.TransientQueue, handlerChatMessage(chatState) )
	if err != nil {
		log.Fatalf("could not make queue: %v", err)
	}
	

}
