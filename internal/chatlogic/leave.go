package chatlogic

import (
	"errors"
	"fmt"
)

func (cs *ChatState) CommandLeave(words string) (Message, error) {
	if cs.isPaused() {
		return Message{}, errors.New("chat is paused")
	}

	goodbyeMsg := Message{
		Username: "System",
		Body: fmt.Sprintf("%s has left the chat", cs.Chatter.Username),
		RoomName: cs.CurrentRoomName,
	}
	return goodbyeMsg, nil
}

