package chatlogic

import (
	"time"
)

type Chatter struct{
	Username string
	Messages []Message
}

type Message struct {
    Username string
    Body     string
    RoomName string
    SentAt   time.Time
}