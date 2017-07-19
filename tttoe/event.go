package tttoe

import "time"

const PlayAction = "play"

type Event struct {
	Name      string
	Data      map[string]interface{}
	CreatedAt time.Time
}

func NewPlayEvent(player string, y, x int) Event {
	return Event{
		Name: PlayAction,
		Data: map[string]interface{}{
			"player": player,
			"y":      y,
			"x":      x,
		},
		CreatedAt: time.Now()}
}
