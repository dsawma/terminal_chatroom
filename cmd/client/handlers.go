package main

import (
	"fmt"

	"github.com/dsawma/terminal_chatroom/internal/chatlogic"
	"github.com/dsawma/terminal_chatroom/internal/pubsub"
	"github.com/dsawma/terminal_chatroom/internal/routing"
	amqp "github.com/rabbitmq/amqp091-go"
)

func handlerPause(gs *chatlogic.ChatState) func(routing.PlayingState) pubsub.AckType{
	return func(ps routing.PlayingState) pubsub.AckType{
		defer fmt.Print("> ")
		gs.HandlePause(ps)
		return pubsub.Ack
	}
}

func handlerMove(cs *chatlogic.ChatState, ch *amqp.Channel) func(chatlogic.ArmyMove) pubsub.AckType{
	return func(move gamelogic.ArmyMove) pubsub.AckType{
		defer fmt.Print("> ")
		mv := cs.HandleMove(move)
		switch mv {
		case gamelogic.MoveOutComeSafe:
    		return pubsub.Ack
		case gamelogic.MoveOutcomeMakeWar:
			err := pubsub.PublishJSON(ch,routing.ExchangePerilTopic, routing.WarRecognitionsPrefix +"."+ gs.GetUsername(),gamelogic.RecognitionOfWar{
   				Attacker: move.Player,
   				Defender: gs.GetPlayerSnap(),
				},)
			if err != nil {
				log.Printf("could not make queue: %v", err)
				return pubsub.NackRequeue
			}
    		return pubsub.Ack
		case gamelogic.MoveOutcomeSamePlayer:
    		return pubsub.NackDiscard
		default:
    		fmt.Println("error: unknown move outcome")
    		return pubsub.NackDiscard
		}
	}
}