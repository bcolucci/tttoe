package tttoe

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
)

type CLI struct {
	LastState *State
}

func (cli *CLI) Start(initialState State, compute func(state State, event Event) State) {
	tick := 0
	ai := NewAI(Symbols[Player2])
	state := initialState
	cli.LastState = &state
	for {
		cli.clear()
		fmt.Println("###### Play Tic-Tac-Toe ######")
		if state.Winner != Nobody {
			break
		}
		var player string
		var x, y int
		if tick == 0 {
			fmt.Println(">> Where do you want to play?")
			player = Player1
			y = cli.askInt(">> Give Y")
			x = cli.askInt(">> Give X")
			tick = 1
		} else {
			player = Player2
			y, x = ai.NextPlay(&state.Stage)
			tick = 0
		}
		state = compute(state, NewPlayEvent(player, y, x))
		fmt.Println(state.Stage.ToString())
		cli.wait()
	}
	fmt.Println(">> The winner is " + state.Winner)
	cli.analyze()
}

func (cli *CLI) Stop() {
	defer fmt.Println("Bye bye.")
	fmt.Println()
	if len(cli.LastState.Events) == 0 {
		return
	}
	fmt.Println(">> Events:")
	for idx, event := range cli.LastState.Events {
		fmt.Print("#" + strconv.Itoa(idx))
		fmt.Println(event)
	}
}

func (cli *CLI) ask(msg string) string {
	fmt.Println(msg)
	var input string
	fmt.Print("<< ")
	fmt.Scanln(&input)
	return input
}

func (cli *CLI) askInt(msg string) int {
	n, err := strconv.Atoi(cli.ask(">> Give Y"))
	if err != nil {
		panic(err)
	}
	return n
}

func (cli *CLI) clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func (cli *CLI) wait() {
	fmt.Println(">> (Press enter to continue)")
	var input string
	fmt.Scanln(&input)
}

func (cli *CLI) analyze() {
	//TODO
	fmt.Println(">> TODO: AnalyzeFlow")
}
