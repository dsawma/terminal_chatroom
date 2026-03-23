package chatlogic

import (
	
	"fmt"
)

func (cs *ChatState) CommandLeave() (Message, error) {
	goodbyeMsg := Message{
		Username: "System",
		Body: fmt.Sprintf("%s has left the chat", cs.Chatter.Username),
		RoomName: cs.CurrentRoomName,
	}
	return goodbyeMsg, nil
}

