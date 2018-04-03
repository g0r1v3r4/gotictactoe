package main

import (
	"fmt"
	"os"

	"cloudnautics.com/tictactoe/game"
)

func main() {
	_, err := game.Loop()
	if err != nil {
		fmt.Println("game crashing")
		os.Exit(1)
	}

}
