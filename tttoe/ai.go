package tttoe

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type AI struct {
	Symbol string
}

func NewAI(symbol string) *AI {
	return &AI{Symbol: symbol}
}

func (ai *AI) playSomeWhere(stage *Stage) (int, int) {
	rows, cols, _ := stage.Split()
	fmt.Println(rows)
	fmt.Println(cols)
	emptyCells := []string{}
	pushEmptyCell := func(x, y int) {
		cell := strconv.Itoa(x) + "" + strconv.Itoa(y)
		found := false
		for _, c := range emptyCells {
			if cell == c {
				found = true
				break
			}
		}
		if !found {
			emptyCells = append(emptyCells, cell)
			fmt.Printf("Empty cell (%d,%d)\n", x, y)
		}
	}
	//TODO fix this loop, we don't check all cells
	for i := 0; i < 3; i += 1 {
		x := strings.Index(rows[i], Symbols[Empty])
		if x > -1 {
			if strings.Contains(rows[i], ai.Symbol) {
				return x, i
			}
			pushEmptyCell(x, i)
		}
		y := strings.Index(cols[i], Symbols[Empty])
		if y > -1 {
			if strings.Contains(cols[i], ai.Symbol) {
				return i, y
			}
			pushEmptyCell(i, y)
		}
	}
	//TODO check diagonals
	randCell := emptyCells[rand.Intn(len(emptyCells)-1)]
	cellCoordinates := strings.Split(randCell, "")
	x, _ := strconv.Atoi(cellCoordinates[0])
	y, _ := strconv.Atoi(cellCoordinates[1])
	return x, y
}

func (ai *AI) NextPlay(stage *Stage) (int, int) {
	if canWin, x, y := stage.OneStepToWin(Symbols[Player2]); canWin {
		return x, y
	}
	if canWin, x, y := stage.OneStepToWin(Symbols[Player1]); canWin {
		return x, y
	}
	return ai.playSomeWhere(stage)
}
