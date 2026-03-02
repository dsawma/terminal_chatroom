package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/dsawma/Sentinel-/internal/auth"
	"github.com/dsawma/Sentinel-/internal/database"
	"github.com/dsawma/terminal_chatroom/internal/chatlogic"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main(){

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
	
	userName, err := auth.Login(ctx, q)
	if err != nil{
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



}