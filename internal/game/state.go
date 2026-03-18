package game

import (
	"tic-tac-toe/internal/board"
	"tic-tac-toe/internal/model"
)

type GameState struct {
	Board         *board.Board
	CurrentPlayer int
	Constraint    model.Line
	Turn          int
	GameOver      bool
	Winner        board.Cell
}
