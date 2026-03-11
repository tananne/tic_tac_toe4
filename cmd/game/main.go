package main

//"tic-tac-toe/internal/game"
//"fmt"

import (
	//"fmt"
	"tic-tac-toe/internal/game"
	"tic-tac-toe/internal/setup"
)

func main() {
	mode, err := setup.ChooseMode()
	if err != nil {
		panic(err)
	}
	player1, player2 := setup.GetPlayers(mode)
	if player1 == nil || player2 == nil {
		panic("Ошибка при создании игроков.")
	}
	g := game.NewGame(player1, player2)
	g.State.Board.Display()
}
