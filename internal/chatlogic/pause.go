package chatlogic

import (
	"fmt"

	"github.com/dsawma/terminal_chatroom/internal/routing"
)

func (cs *ChatState) HandlePause(ps routing.PauseState) {
	defer fmt.Println("------------------------")
	fmt.Println()
	if ps.IsPaused {
		fmt.Println("==== Pause Detected ====")
		cs.pauseChat()
	} else {
		fmt.Println("==== Resume Detected ====")
		cs.resumeChat()
	}
}