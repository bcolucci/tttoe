package tttoe

import (
	"fmt"
	"math/rand"
	"strings"
)

type AI struct {
	Symbol string
}

func NewAI(symbol string) *AI {
	return &AI{Symbol: symbol}
}

//TODO Fix it
func (ai *AI) playSomeWhere(stage *Stage) (int, int) {
	rows, cols, _ := stage.Split()
	emptyCells := [][]int{}
	for i := 0; i < 3; i += 1 {
		x := strings.Index(rows[i], Symbols[Empty])
		if x > -1 {
			fmt.Printf("Found a row with an empty cell (->%d,%d)\n", x, i)
			if strings.Contains(rows[i], ai.Symbol) {
				return x, i
			}
			fmt.Printf("Empty cell (%d,%d)\n", x, i)
			emptyCells = append(emptyCells, []int{x, i})
		}
		y := strings.Index(cols[i], Symbols[Empty])
		if y > -1 {
			fmt.Printf("Found a row with an empty cell (%d,->%d)\n", i, y)
			if strings.Contains(cols[i], ai.Symbol) {
				return i, y
			}
			fmt.Printf("Empty cell (%d,%d)\n", i, y)
			emptyCells = append(emptyCells, []int{i, y})
		}
	}
	//TODO check diagonals
	randCell := emptyCells[rand.Intn(len(emptyCells)-1)]
	return randCell[0], randCell[1]
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
