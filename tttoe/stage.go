package tttoe

import (
	"errors"
	"strings"
)

var InvalidCoordinate error = errors.New("Coordinate should be in [0, 2] interval.")
var NoneEmptyCell error = errors.New("You're trying to play on a none empty cell.")

type Stage struct {
	cells [][]string
}

func NewStage() Stage {
	stage := Stage{}
	stage.cells = make([][]string, 3, 3)
	for x := 0; x < 3; x += 1 {
		stage.cells[x] = make([]string, 3)
		for y := 0; y < 3; y += 1 {
			stage.cells[x][y] = Symbols[Empty]
		}
	}
	return stage
}

func (s *Stage) ToString() string {
	stageStr := ""
	for y := 0; y < 3; y += 1 {
		values := []string{}
		for x := 0; x < 3; x += 1 {
			values = append(values, "["+s.cells[x][y]+"]")
		}
		stageStr += strings.Join(values, "") + "\n"
	}
	return stageStr
}

func (s *Stage) NbPlayerSymbols(player string) int {
	nbOcc := 0
	for x := 0; x < 3; x += 1 {
		for y := 0; y < 3; y += 1 {
			if s.cells[x][y] == Symbols[player] {
				nbOcc += 1
			}
		}
	}
	return nbOcc
}

func (s *Stage) FoundPlayerLine(player string) bool {
	expected := Symbols[player] + Symbols[player] + Symbols[player]
	for i := 0; i < 3; i += 1 {
		if s.cells[0][i]+s.cells[1][i]+s.cells[2][i] == expected {
			return true
		}
		if s.cells[i][0]+s.cells[i][1]+s.cells[i][2] == expected {
			return true
		}
	}
	return (s.cells[0][0]+s.cells[1][1]+s.cells[2][2] == expected) ||
		(s.cells[0][2]+s.cells[1][1]+s.cells[2][0] == expected)
}

func (s *Stage) Split() ([]string, []string, []string) {
	rows := make([]string, 3)
	cols := make([]string, 3)
	for i := 0; i < 3; i += 1 {
		rows[i] = s.cells[i][0] + s.cells[i][1] + s.cells[i][2]
		cols[i] = s.cells[0][i] + s.cells[1][i] + s.cells[2][i]
	}
	diagonal1 := s.cells[0][0] + s.cells[1][1] + s.cells[2][2]
	diagonal2 := s.cells[0][2] + s.cells[1][1] + s.cells[2][0]
	return rows, cols, []string{diagonal1, diagonal2}
}

func (s *Stage) CheckCoordinates(x, y int) error {
	if (x < 0 || x > 2) || (y < 0 || y > 2) {
		return InvalidCoordinate
	}
	if s.cells[x][y] != Symbols[Empty] {
		return NoneEmptyCell
	}
	return nil
}
