package main

import (
	"bufio"
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/dsawma/terminal_chatroom/internal/auth"
	"github.com/dsawma/terminal_chatroom/internal/chatlogic"
	"github.com/dsawma/terminal_chatroom/internal/database"
	"github.com/dsawma/terminal_chatroom/internal/pubsub"
	"github.com/dsawma/terminal_chatroom/internal/routing"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {

	err := godotenv.Load("../../.env")
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
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}

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
	HelloMsg := chatlogic.Message{
		Username: "System",
		Body: fmt.Sprintf("%s has joined the chat", chatState.Chatter.Username),
		RoomName: chatState.CurrentRoomName,
	}
	err = pubsub.PublishGob(ch, routing.ExchangeChatTopic, "sends_msg.*", HelloMsg)
	if err != nil {
		log.Printf("could not make queue: %v", err)
	}
	fmt.Println("To exit chat: /quit")
	fmt.Println("To join a different chatroom: /join [roomName]")
	fmt.Println("-------------------------------------------------")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan(){
			break
		}
		input := scanner.Text()
		if strings.HasPrefix(input, "/"){
			parts := strings.Split(input, " ")
			operation := parts[0]
			switch operation {
			case "/quit":
				goodbyeMsg, err := chatState.CommandLeave()
				if err != nil {
					fmt.Printf("cannot send message: %v", err)
				}
				err =pubsub.PublishGob(ch, routing.ExchangeChatTopic, "sends_msg.*", goodbyeMsg)
				if err != nil {
					log.Printf("could not make queue: %v", err)
				}
				fmt.Println("Goodbye!")
			case "/join":
				if len(parts) == 2{
					newRoom,err := q.GetRoomByRoomName(ctx, parts[1])
					if err == nil{
						fmt.Println("Room not Found")
					}
					goodbyeMsg, err := chatState.CommandLeave() 
					if err != nil {
						fmt.Printf("cannot send message: %v", err)
					}
					err = pubsub.PublishGob(ch, routing.ExchangeChatTopic, "sends_msg.*", goodbyeMsg)
					if err != nil {
						log.Printf("could not make queue: %v", err)
					}
					chatState.CurrentRoomName = newRoom
					HelloMsg := chatlogic.Message{
						Username: "System",
						Body: fmt.Sprintf("%s has joined the chat", chatState.Chatter.Username),
						RoomName: chatState.CurrentRoomName,
					}
					err = pubsub.PublishGob(ch, routing.ExchangeChatTopic, "sends_msg.*", HelloMsg)
					if err != nil {
						log.Printf("could not make queue: %v", err)
					}
				} else{
					fmt.Println("Incorrect join command")
				}
			default: 
				fmt.Println("Unknown command")
				fmt.Println("To exit chat: /quit")
				fmt.Println("To join a different chatroom: /join [roomName]")
			}
			continue

		}else{
			newMsg, err := chatState.CommandMessage(input) 
			if err != nil {
				fmt.Printf("cannot send message: %v", err)
			}
			err = pubsub.PublishGob(ch, routing.ExchangeChatTopic, "sends_msg.*", newMsg)
			if err != nil {
				log.Printf("could not make queue: %v", err)
			}
		}

	}

}
