package main

import (
	"github.com/diogomene/tictactoe-ai/entities"
)

func main() {
	var board entities.Board
	board.Init()
	board.Print()
}
