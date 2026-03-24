package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dsawma/terminal_chatroom/internal/auth"
	"github.com/dsawma/terminal_chatroom/internal/pubsub"
	"github.com/dsawma/terminal_chatroom/internal/routing"
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)


func main() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	connectStr := os.Getenv("RABBIT_URL")
	if connectStr == "" {
		log.Fatal("RABBIT_URL is missing")
	}
	fmt.Println("Starting Chat server...")
	connection, err := amqp.Dial(connectStr)
	if err != nil {
		log.Fatalf("could not create connection: %v", err)
	}
	defer connection.Close()
	channel, err := connection.Channel()
	if err != nil {
		log.Fatalf("could not create channel: %v", err)
	}
	fmt.Println("Connection Successful")

	err = pubsub.DeclareExchange(channel, routing.ExchangeChatTopic, "topic")
	if err != nil {
		log.Fatalf("could not declare topic exchange: %v", err)
	}

	err = pubsub.DeclareExchange(channel, routing.ExchangeChatDirect, "direct")
	if err != nil {
		log.Fatalf("could not declare direct exchange: %v", err)
	}

	fmt.Println("Exchange created successfully!")


	err = pubsub.SubscribeGob(connection, routing.ExchangeChatTopic,"chat_logs" , routing.ChatLogSlug + ".*", pubsub.DurableQueue, handlerLogs())
	if err != nil {
		log.Fatalf("could not make queue: %v", err)
	}

	fmt.Println("Commands:")
	fmt.Println("* pause")
	fmt.Println("* resume")
	fmt.Println("* quit")
	for {
		fmt.Print("> ")
		strSlice := auth.GetInput()
		if len(strSlice) == 0 {
			continue
		}
		switch strSlice[0] {
		case "pause":
			fmt.Println("sending a pause message")
			err = pubsub.PublishGob(channel, routing.ExchangeChatDirect, routing.PauseKey,routing.PauseState{IsPaused:true,}, ) 
			if err != nil {
				log.Fatalf("could not publish Gob: %v", err)
			}
		case "resume":
			fmt.Println("sending a resume message")
			err = pubsub.PublishGob(channel, routing.ExchangeChatDirect, routing.PauseKey,routing.PauseState{IsPaused:false,}, ) 
			if err != nil {
				log.Fatalf("could not publish Gob: %v", err)
			}
		case "quit":
			fmt.Println("exiting program")
			return
		default:
			fmt.Println("dont understand command")
		}
	}

}