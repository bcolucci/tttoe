package tttoe

import (
	"time"
)

type State struct {
	Stage     Stage
	Winner    string
	Events    []Event
	CreatedAt time.Time
	UpdatedAt time.Time
}
