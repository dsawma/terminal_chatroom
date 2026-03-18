package routing

import "time"

type PauseState struct {
	IsPaused bool
}

type ChatLog struct {
	CurrentTime time.Time
	Message     string
	Username    string
}
