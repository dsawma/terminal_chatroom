package routing

import "time"

type PlayingState struct {
	IsPaused bool
}

type ChatLog struct {
	CurrentTime time.Time
	Message     string
	Username    string
}
