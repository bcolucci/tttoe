package main

import (
	"os"
	"os/signal"
	"poc-es/tttoe"
)

var cli *tttoe.CLI

func runApp() {
	state := tttoe.CreateInitialState()
	cli.Start(state, tttoe.Compute)
}

func catchExit() {
	signals := make(chan os.Signal, 1)
	done := make(chan bool)
	signal.Notify(signals, os.Interrupt)
	go func() {
		for range signals {
			cli.Stop()
			done <- true
		}
	}()
	<-done
}

func main() {
	cli = &tttoe.CLI{}
	go runApp()
	catchExit()
}
