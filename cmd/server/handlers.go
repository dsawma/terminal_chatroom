package main

import (
	"fmt"
	"github.com/dsawma/terminal_chatroom/internal/pubsub"
	"github.com/dsawma/terminal_chatroom/internal/routing"
)

func handlerLogs() func(gamelog routing.ChatLog) pubsub.AckType {
	return func(gamelog routing.ChatLog) pubsub.AckType {
		defer fmt.Print("> ")

		err := chatlogic.WriteLog(gamelog)
		if err != nil {
			fmt.Printf("error writing log: %v\n", err)
			return pubsub.NackRequeue
		}
		return pubsub.Ack
	}
}