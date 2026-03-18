package game

import (
	"tic-tac-toe/internal/board"
	"tic-tac-toe/internal/model"
	"tic-tac-toe/internal/player"
)

type Game struct {
	State *GameState

	players [2]player.Player
}

func NewGame(p1, p2 player.Player, size int) *Game {
	gameState := &GameState{
		Board:         board.NewBoard(size),
		CurrentPlayer: 1,
		Constraint:    model.Line{},
		Turn:          0,
		GameOver:      false,
		Winner:        board.Empty,
	}
	return &Game{
		State:   gameState,
		players: [2]player.Player{p1, p2},
	}
}

func (g *Game) switchPlayer() {
	g.State.CurrentPlayer = (g.State.CurrentPlayer + 1) % 2
}

func (g *Game) updateTurn() {
	g.State.Turn++
}
