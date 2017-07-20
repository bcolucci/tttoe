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

func (cli *CLI) Start(initialState State,
	reduce func(state State, event Event) (State, error)) {
	tick := 0
	ai := &AI{Symbol: Cross}
	state := initialState
	cli.LastState = &state
	var stateErr error
	for {
		cli.clear()
		fmt.Println("###### Play Tic-Tac-Toe ######")
		fmt.Println(state.Stage.ToString())
		if state.Winner != Nobody {
			break
		}
		var player string
		var x, y int
		if tick == 2 {
			tick = 0
		}
		if tick == 0 {
			fmt.Println(">> Where do you want to play?")
			player = Player1
			x = cli.askInt(">> Give X")
			y = cli.askInt(">> Give Y")
		} else {
			player = Player2
			x, y = ai.NextPlay(&state.Stage)
		}
		if state, stateErr = reduce(state, NewPlayEvent(player, x, y)); stateErr != nil {
			fmt.Println(">> Error: " + stateErr.Error())
		} else {
			tick += 1
		}
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
	os.Exit(0)
}

func (cli *CLI) ask(msg string) string {
	fmt.Println(msg)
	var input string
	fmt.Print("<< ")
	fmt.Scanln(&input)
	return input
}

func (cli *CLI) askInt(msg string) int {
	n, err := strconv.Atoi(cli.ask(msg))
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
	cli.Stop()
}
