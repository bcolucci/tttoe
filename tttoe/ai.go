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
			if strings.Contains(rows[i], ai.Symbol) {
				return i, x
			}
			fmt.Printf("Empty cell %d,%d\n", i, x)
			emptyCells = append(emptyCells, []int{i, x})
		}
		y := strings.Index(cols[i], Symbols[Empty])
		if y > -1 {
			if strings.Contains(cols[i], ai.Symbol) {
				return y, i
			}
			fmt.Printf("Empty cell %d,%d\n", y, i)
			emptyCells = append(emptyCells, []int{y, i})
		}
	}
	//TODO check diagonals
	randCell := emptyCells[rand.Intn(len(emptyCells))]
	return randCell[0], randCell[1]
}

func (ai *AI) NextPlay(stage *Stage) (int, int) {
	if canWin, y, x := stage.OneStepToWin(Symbols[Player2]); canWin {
		return y, x
	}
	if canWin, y, x := stage.OneStepToWin(Symbols[Player1]); canWin {
		return y, x
	}
	return ai.playSomeWhere(stage)
}
