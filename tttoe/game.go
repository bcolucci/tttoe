package tttoe

import (
	"errors"
	"time"
)

const Empty = "none"
const Cross = "cross"
const Circle = "circle"

const Player1 = Cross
const Player2 = Circle
const Nobody = ""

var Symbols = map[string]string{
	Empty:  " ",
	Cross:  "X",
	Circle: "O",
}

var PlayerAlreadyPlayed error = errors.New("Player already played.")

func CreateInitialState() State {
	return State{Stage: NewStage(), Events: []Event{}, CreatedAt: time.Now()}
}

func CheckTurn(player string, stage *Stage) error {
	nbCrossOcc := stage.NbSymbolOcc(Symbols[Cross])
	nbCircleOcc := stage.NbSymbolOcc(Symbols[Circle])
	if player == Player1 {
		if nbCrossOcc == nbCircleOcc+1 {
			return PlayerAlreadyPlayed
		}
	} else {
		if nbCircleOcc == nbCrossOcc {
			return PlayerAlreadyPlayed
		}
	}
	return nil
}

func GetWinner(stage *Stage) string {
	if stage.LineOfSymbol(Symbols[Player1]) {
		return Player1
	}
	if stage.LineOfSymbol(Symbols[Player2]) {
		return Player2
	}
	return Nobody
}

func Reduce(state State, event Event) (State, error) {
	nextState := state
	commonUpdates := func() {
		nextState.Events = append(nextState.Events, event)
		nextState.UpdatedAt = time.Now()
	}
	switch event.Name {
	case PlayAction:
		player := event.Data["player"].(string)
		x := event.Data["x"].(int)
		y := event.Data["y"].(int)
		if err := nextState.Stage.CheckCoordinates(x, y); err != nil {
			return nextState, err
		}
		CheckTurn(player, &nextState.Stage)
		nextState.Stage.cells[x][y] = Symbols[player]
		nextState.Winner = GetWinner(&nextState.Stage)
		commonUpdates()
	default:
		break
	}
	return nextState, nil
}
