package chatlogic

import (
	"sync"
)

type ChatState struct {
	Chatter Chatter
	CurrentRoomName string
 	Paused bool
	Mu     sync.RWMutex
}

func NewChatState(username string, roomName string) *ChatState {
	return &ChatState{
		Chatter: Chatter{
			Username: username,
			Messages:    []Message{},
		},
		CurrentRoomName: roomName,
		Paused: false,
		Mu:     sync.RWMutex{},
	}
}

func (cs *ChatState) resumeChat() {
	cs.Mu.Lock()
	defer cs.Mu.Unlock()
	cs.Paused = false
}

func (cs *ChatState) pauseChat() {
	cs.Mu.Lock()
	defer cs.Mu.Unlock()
	cs.Paused = true
}

func (cs *ChatState) isPaused() bool {
	cs.Mu.RLock()
	defer cs.Mu.RUnlock()
	return cs.Paused
}