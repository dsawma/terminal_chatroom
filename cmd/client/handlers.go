package main

import (
	"fmt"

	"github.com/dsawma/terminal_chatroom/internal/chatlogic"
	"github.com/dsawma/terminal_chatroom/internal/pubsub"
	"github.com/dsawma/terminal_chatroom/internal/routing"

)

func handlerPause(cs *chatlogic.ChatState) func(routing.PauseState) pubsub.AckType{
	return func(ps routing.PauseState) pubsub.AckType{
		cs.HandlePause(ps)
		return pubsub.Ack
	}
}

func handlerChatMessage(state *chatlogic.ChatState) func(msg chatlogic.Message) pubsub.AckType {
    return func(msg chatlogic.Message) pubsub.AckType {
        state.Mu.Lock()
        defer state.Mu.Unlock()
    
        if msg.RoomName == state.CurrentRoomName {
            state.Chatter.Messages = append(state.Chatter.Messages, msg)
            fmt.Printf("[%s]: %s\n", msg.Username, msg.Body)
        }
        return pubsub.Ack
    }
}