package tttoe

type AI struct {
	Symbol string
}

func NewAI(symbol string) *AI {
	return &AI{Symbol: symbol}
}

func (ai *AI) bestPlay(stage *Stage) (int, int) {
	//TODO
	return 0, 0
}

func (ai *AI) NextPlay(stage *Stage) (int, int) {
	if canWin, y, x := stage.OneStepToWin(Symbols[Player2]); canWin {
		return y, x
	}
	if canWin, y, x := stage.OneStepToWin(Symbols[Player1]); canWin {
		return y, x
	}
	return ai.bestPlay(stage)
}
