package chatlogic

import (
	"time"

	"github.com/dsawma/terminal_chatroom/internal/database"
)

type Chatter struct{
	Username string
	Messages []database.Message
}

type Message struct {
    Username string
    Body     string
    RoomName string
    SentAt   time.Time
}