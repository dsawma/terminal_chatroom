package main

import (
	"fmt"
	"log"

	"github.com/dsawma/terminal_chatroom/internal/auth"
	"github.com/dsawma/terminal_chatroom/internal/pubsub"
	amqp "github.com/rabbitmq/amqp091-go"
)


func main() {
	fmt.Println("Starting Chat server...")
	connectStr :=  "amqp://guest:guest@localhost:5672/"
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
	err = pubsub.SubscribeGob(connection, routing.ExchangePerilTopic,"game_logs" , routing.GameLogSlug + ".*", pubsub.DurableQueue, handlerLogs())
	if err != nil {
		log.Fatalf("could not make queue: %v", err)
	}

	fmt.Println("Commands:")
	fmt.Println("* pause")
	fmt.Println("* resume")
	fmt.Println("* quit")
	for {
			strSlice := auth.GetInput()
			if len(strSlice) == 0 {
				continue
			}
			switch strSlice[0] {
			case "pause":
				fmt.Println("sending a pause message")
				err = pubsub.PublishGob(channel, routing.ExchangePerilDirect, routing.PauseKey,routing.PlayingState{IsPaused:true,}, ) 
				if err != nil {
					log.Fatalf("could not publish JSON: %v", err)
				}
			case "resume":
				fmt.Println("sending a resume message")
				err = pubsub.PublishGob(channel, routing.ExchangePerilDirect, routing.PauseKey,routing.PlayingState{IsPaused:false,}, ) 
				if err != nil {
					log.Fatalf("could not publish JSON: %v", err)
				}
			case "quit":
				fmt.Println("exiting program")
				return
			default:
				fmt.Println("dont understand command")
			}
		}


}