package main

import (
	"tic-tac-toe/internal/game"
)

func main() {
	g := game.NewGame()
	g.Board.Display()
}
