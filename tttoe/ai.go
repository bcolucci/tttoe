package tttoe

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

var winPatterns = []string{
	"111......",
	"...111...",
	"......111",
	"1..1..1..",
	".1..1..1..1",
	"..1..1..1..",
	"1...1...1",
	"..1.1.1..",
}

type AI struct {
	Symbol      string
	tree        *AINode
	nbWinCases  int
	nbLossCases int
}

type AINode struct {
	ai        *AI
	parent    *AINode
	depth     int
	flatStage string
	winner    string
}

func (n *AINode) buildTree() {
	n.winner = n.getWinner()
	if n.winner != Nobody {
		//TODO find a way to give a good move
		if n.winner == Player1 {
			n.ai.nbLossCases += 1
		} else {
			n.ai.nbWinCases += 1
		}
		return
	}
	if strings.Index(n.flatStage, Symbols[Empty]) == 1 {
		return
	}
	n.computeChildren()
}

func (n *AINode) computeChildren() {
	player := n.turnOf()
	for i := 0; i < 9; i += 1 {
		cell := string(n.flatStage[i])
		if cell != Symbols[Empty] {
			continue
		}
		newStage := n.flatStage[:i] + Symbols[player] + n.flatStage[i+1:]
		subNode := &AINode{ai: n.ai, parent: n, depth: n.depth + 1, flatStage: newStage}
		subNode.buildTree()
	}
}

func (n *AINode) turnOf() string {
	nbCross := n.nbSymbolOcc(Symbols[Cross])
	nbCircles := n.nbSymbolOcc(Symbols[Circle])
	if nbCross == nbCircles+1 {
		return Player2
	}
	return Player1
}

func (n *AINode) nbSymbolOcc(symbol string) int {
	nbOcc := 0
	for i := 0; i < 9; i += 1 {
		cell := string(n.flatStage[i])
		if strings.Compare(cell, symbol) == 0 {
			nbOcc += 1
		}
	}
	return nbOcc
}

func (n *AINode) isWinner(player string) bool {
	opponentSymbol := n.getOpponentSymbol(player)
	for _, winPattern := range winPatterns {
		curPattern := strings.Replace(n.flatStage, Symbols[player], "1", -1)
		curPattern = strings.Replace(curPattern, opponentSymbol, ".", -1)
		curPattern = strings.Replace(curPattern, Symbols[Empty], ".", -1)
		if curPattern == winPattern {
			return true
		}
	}
	return false
}

func (n *AINode) getOpponentSymbol(player string) string {
	return Symbols[n.getOpponent(player)]
}

func (n *AINode) getOpponent(player string) string {
	if player == Player1 {
		return Player2
	}
	return Player1
}

func (n *AINode) getWinner() string {
	if n.isWinner(Player1) {
		return Player1
	} else if n.isWinner(Player2) {
		return Player2
	}
	return Nobody
}

// The Old way
func (ai *AI) playSomeWhere_v1(stage *Stage) (int, int) {
	rows, cols, _ := stage.Split()
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

func (ai *AI) playSomeWhere_v2(stage *Stage) (int, int) {
	rows, _, _ := stage.Split()
	ai.nbWinCases = 0
	ai.nbLossCases = 0
	ai.tree = &AINode{ai: ai, depth: 0, flatStage: strings.Join(rows, "")}
	ai.tree.buildTree()
	fmt.Printf("nbLoss=%d\n", ai.nbLossCases)
	fmt.Printf("nbWin=%d\n", ai.nbWinCases)
	//TODO
	return 0, 0
}

func (ai *AI) playSomeWhere(stage *Stage) (int, int) {
	return ai.playSomeWhere_v1(stage)
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
