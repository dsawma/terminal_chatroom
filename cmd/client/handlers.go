package main

import (
	"fmt"

	"github.com/dsawma/terminal_chatroom/internal/chatlogic"
	"github.com/dsawma/terminal_chatroom/internal/pubsub"
	"github.com/dsawma/terminal_chatroom/internal/routing"

)

func handlerPause(gs *chatlogic.ChatState) func(routing.PlayingState) pubsub.AckType{
	return func(ps routing.PlayingState) pubsub.AckType{
		defer fmt.Print("> ")
		gs.HandlePause(ps)
		return pubsub.Ack
	}
}

func handlerChatMessage(state *chatlogic.ChatState) func(msg chatlogic.Message) pubsub.AckType {
    return func(msg chatlogic.Message) pubsub.AckType {
        state.Mu.Lock()
        defer state.Mu.Unlock()
        
        // Only care about messages for the room we are currently in
        if msg.RoomName == state.CurrentRoomName {
            state.Chatter.Messages = append(state.Chatter.Messages, msg)
            fmt.Printf("[%s]: %s\n", msg.Username, msg.Body)
        }
        return pubsub.Ack
    }
}