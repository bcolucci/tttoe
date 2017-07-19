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
	for y := 0; y < 3; y += 1 {
		stage.cells[y] = make([]string, 3)
		for x := 0; x < 3; x += 1 {
			stage.cells[y][x] = Symbols[Empty]
		}
	}
	return stage
}

func (s *Stage) ToString() string {
	stageStr := ""
	for y := 0; y < 3; y += 1 {
		values := []string{}
		for x := 0; x < 3; x += 1 {
			values = append(values, s.cells[y][x])
		}
		stageStr += strings.Join(values, "") + "\n"
	}
	return stageStr
}

func (s *Stage) NbSymbolOcc(symbol string) int {
	nbOcc := 0
	for y := 0; y < 3; y += 1 {
		for x := 0; x < 3; x += 1 {
			if s.cells[y][x] == symbol {
				nbOcc += 1
			}
		}
	}
	return nbOcc
}

func (s *Stage) LineOfSymbol(symbol string) bool {
	expected := symbol + symbol + symbol
	for i := 0; i < 3; i += 1 {
		if s.cells[i][0]+s.cells[i][1]+s.cells[i][2] == expected {
			return true
		}
		if s.cells[0][i]+s.cells[1][i]+s.cells[2][i] == expected {
			return true
		}
	}
	return (s.cells[0][0]+s.cells[1][1]+s.cells[2][2] == expected) ||
		(s.cells[0][2]+s.cells[1][1]+s.cells[2][0] == expected)
}

func (s *Stage) OneStepToWin(symbol string) (bool, int, int) {
	expectations := []string{
		Symbols[Empty] + symbol + symbol,
		symbol + Symbols[Empty] + symbol,
		symbol + symbol + Symbols[Empty],
	}
	for i := 0; i < 3; i += 1 {
		onX := s.cells[i][0] + s.cells[i][1] + s.cells[i][2]
		onY := s.cells[0][i] + s.cells[1][i] + s.cells[2][i]
		for idx, expected := range expectations {
			if onX == expected {
				return true, i, idx
			}
			if onY == expected {
				return true, idx, i
			}
		}
	}
	return false, 0, 0
}

func (s *Stage) CheckCoordinates(x, y int) error {
	if (x < 0 || x > 2) || (y < 0 || y > 2) {
		return InvalidCoordinate
	}
	if s.cells[y][x] != Symbols[Empty] {
		return NoneEmptyCell
	}
	return nil
}
