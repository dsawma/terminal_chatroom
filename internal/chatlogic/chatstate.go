package chatlogic

import (
	"sync"

	"github.com/dsawma/terminal_chatroom/internal/database"
	"github.com/google/uuid"
)

type ChatState struct {
	Chatter Chatter
	CurrentRoomID uuid.UUID
 	Paused bool
	mu     *sync.RWMutex
}

func NewChatState(username string, roomID uuid.UUID) *ChatState {
	return &ChatState{
		Chatter: Chatter{
			Username: username,
			Messages:    []database.Message{},
		},
		CurrentRoomID: roomID,
		Paused: false,
		mu:     &sync.RWMutex{},
	}
}

func (gs *ChatState) resumeChat() {
	gs.mu.Lock()
	defer gs.mu.Unlock()
	gs.Paused = false
}

func (gs *ChatState) pauseChat() {
	gs.mu.Lock()
	defer gs.mu.Unlock()
	gs.Paused = true
}

func (gs *ChatState) isPaused() bool {
	gs.mu.RLock()
	defer gs.mu.RUnlock()
	return gs.Paused
}