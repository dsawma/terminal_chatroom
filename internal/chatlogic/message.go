package chatlogic

import "errors"

func (cs *ChatState) CommandMessage(words string) (Message, error) {
	if cs.isPaused() {
		return Message{}, errors.New("chat is paused")
	}
	newMsg := Message {
		Username: cs.Chatter.Username,
		Body: words,
		RoomName: cs.CurrentRoomName,
	}
	return newMsg, nil
}