package main

import (
	"fmt"
	"log"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main(){
	fmt.Println("Starting Chat client...")

	connectStr := "amqp://guest:guest@localhost:5672/"
	connection, err := amqp.Dial(connectStr) 
	if err != nil {
		log.Fatalf("Could not create connection: %v", err)
		
	}
}