package tttoe

import (
	"fmt"
	"strings"
)

const _s = "s" // IA symbol
const _a = "*" // empty or

var winPatterns = []string{
	_s + _s + _s + _a + _a + _a + _a + _a + _a,
	_a + _a + _a + _s + _s + _s + _a + _a + _a,
	_a + _a + _a + _a + _a + _a + _s + _s + _s,
	_s + _a + _a + _s + _a + _a + _s + _a + _a,
	_a + _s + _a + _a + _s + _a + _a + _s + _a,
	_a + _a + _s + _a + _a + _s + _a + _a + _s,
	_a + _a + _s + _a + _s + _a + _s + _a + _a,
}

type AI struct {
	Symbol    string
	tree      *AINode
	winCases  []*AINode
	lostCases []*AINode
	drawCases []*AINode
}

type AINode struct {
	ai        *AI
	parent    *AINode
	depth     int
	flatStage string
	winner    string
	playedX   int
	playedY   int
}

func NewAI() *AI {
	return &AI{
		Symbol:    Symbols[Player2],
		winCases:  []*AINode{},
		lostCases: []*AINode{},
	}
}

func NewAINode(parent *AINode, flatStage string) *AINode {
	return &AINode{
		ai:        parent.ai,
		parent:    parent,
		depth:     parent.depth + 1,
		flatStage: flatStage,
	}
}

func (n *AINode) buildTree() {
	n.winner = n.getWinner()
	if n.winner != Nobody {
		if n.winner == Player1 {
			n.ai.winCases = append(n.ai.winCases, n)
		} else {
			n.ai.lostCases = append(n.ai.lostCases, n)
		}
		return
	}
	if strings.Index(n.flatStage, Symbols[Empty]) == 1 {
		n.ai.drawCases = append(n.ai.drawCases, n)
		return
	}
	n.computeChildren()
}

func (n *AINode) computeChildren() {
	player := n.turnOf()
	x := 0
	for i := 0; i < 9; i += 1 {
		cell := string(n.flatStage[i])
		if cell != Symbols[Empty] {
			continue
		}
		newStage := n.flatStage[:i] + Symbols[player] + n.flatStage[i+1:]
		subNode := NewAINode(n, newStage)
		subNode.playedX = x
		subNode.playedY = i % 3
		x += 1
		if x == 3 {
			x = 0
		}
		subNode.buildTree()
	}
}

func (n *AINode) turnOf() string {
	nbP1Symbols := n.nbPlayerSymbols(Player1)
	nbP2Symbols := n.nbPlayerSymbols(Player2)
	if nbP1Symbols == nbP2Symbols {
		return Player1
	}
	return Player2
}

func (n *AINode) nbPlayerSymbols(player string) int {
	nbOcc := 0
	for i := 0; i < 9; i += 1 {
		cell := string(n.flatStage[i])
		if strings.Compare(cell, Symbols[player]) == 0 {
			nbOcc += 1
		}
	}
	return nbOcc
}

func (n *AINode) isWinner(player string) bool {
	opponentSymbol := n.getOpponentSymbol(player)
	for _, winPattern := range winPatterns {
		curPattern := strings.Replace(n.flatStage, Symbols[player], _s, -1)
		curPattern = strings.Replace(curPattern, opponentSymbol, _a, -1)
		curPattern = strings.Replace(curPattern, Symbols[Empty], _a, -1)
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

func backToParent(current *AINode) *AINode {
	fmt.Print("\tback to:")
	fmt.Println(current)
	if current.parent != nil && current.depth > 2 {
		return backToParent(current.parent)
	}
	return current
}

// The Old way
func (ai *AI) playSomeWhere(stage *Stage) (int, int) {
	rows, cols, _ := stage.Split()
	//TODO fix this loop, we don't check all cells
	for i := 0; i < 3; i += 1 {
		x := strings.Index(rows[i], Symbols[Empty])
		if x > -1 {
			if strings.Contains(rows[i], ai.Symbol) {
				return x, i
			}
		}
		y := strings.Index(cols[i], Symbols[Empty])
		if y > -1 {
			if strings.Contains(cols[i], ai.Symbol) {
				return i, y
			}
		}
	}
	//TODO check diagonals
	ai.tree = &AINode{ai: ai, depth: 0, flatStage: strings.Join(rows, "")}
	ai.tree.buildTree()
	if len(ai.winCases) > 0 {
		fmt.Println("Play a win case")
		fmt.Println(ai.winCases[0])
		parent := backToParent(ai.winCases[0])
		fmt.Println(parent)
		return parent.playedX, parent.playedY
	}
	fmt.Println("No win case")
	if len(ai.drawCases) > 0 {
		fmt.Println("Play a draw case")
		return 0, 0
	}
	fmt.Println("Play a lost case")
	return 0, 0
}

func (ai *AI) oneStepToWin(player string, stage *Stage) (bool, int, int) {
	symbol := Symbols[player]
	expectations := []string{
		Symbols[Empty] + symbol + symbol,
		symbol + Symbols[Empty] + symbol,
		symbol + symbol + Symbols[Empty],
	}
	for i := 0; i < 3; i += 1 {
		onX := stage.cells[i][0] + stage.cells[i][1] + stage.cells[i][2]
		onY := stage.cells[0][i] + stage.cells[1][i] + stage.cells[2][i]
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

func (ai *AI) NextPlay(stage *Stage) (int, int) {
	ai.tree = nil // deference all the nodes
	if canWin, x, y := ai.oneStepToWin(Player2, stage); canWin {
		return x, y
	}
	if canWin, x, y := ai.oneStepToWin(Player1, stage); canWin {
		return x, y
	}
	return ai.playSomeWhere(stage)
}
