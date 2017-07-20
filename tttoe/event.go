package tttoe

import "time"

const PlayAction = "play"

type Event struct {
	Name      string
	Data      map[string]interface{}
	CreatedAt time.Time
}

func NewPlayEvent(player string, x, y int) Event {
	return Event{
		Name: PlayAction,
		Data: map[string]interface{}{
			"player": player,
			"x":      x,
			"y":      y,
		},
		CreatedAt: time.Now()}
}
