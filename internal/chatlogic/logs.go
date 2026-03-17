package chatlogic

import (
	"fmt"
	"log"
	"os"
	"time"
	"github.com/dsawma/terminal_chatroom/internal/routing"
)

const logsFile = "game.log"

const writeToDiskSleep = 1 * time.Second

func WriteLog(chatlog routing.ChatLog) error {
	log.Printf("received chat log...")
	time.Sleep(writeToDiskSleep)

	f, err := os.OpenFile(logsFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("could not open logs file: %v", err)
	}
	defer f.Close()

	str := fmt.Sprintf("%v %v: %v\n", chatlog.CurrentTime.Format(time.RFC3339), chatlog.Username, chatlog.Message)
	_, err = f.WriteString(str)
	if err != nil {
		return fmt.Errorf("could not write to logs file: %v", err)
	}
	return nil
}
