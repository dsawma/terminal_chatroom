package chatlogic

import "github.com/dsawma/terminal_chatroom/internal/database"

type Chatter struct{
	Username string
	Messages []database.Message
}