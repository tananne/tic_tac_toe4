package main

import (
	"tic-tac-toe/internal/board"
	"tic-tac-toe/internal/game"
	"tic-tac-toe/internal/player"
	"tic-tac-toe/internal/ui"
)

func main() {
	ui := ui.NewConsoleUI()

	mode := ui.AskGameMode()

	var p1, p2 player.Player

	if mode == "pvp" {
		p1 = player.NewHumanPlayer(board.X, ui)
		p2 = player.NewHumanPlayer(board.O, ui)
	} else {
		symbol := ui.AskPlayerSymbol()

		if symbol == board.X {
			p1 = player.NewHumanPlayer(board.X, ui)
			p2 = player.NewAIPlayer(board.O)
		} else {
			p1 = player.NewAIPlayer(board.X)
			p2 = player.NewHumanPlayer(board.O, ui)
		}
	}

	game := game.NewGame(p1, p2, 4)
	ui.DisplayBoard(game.State.Board)
}
