package player

import (
	"tic-tac-toe/internal/board"
	"tic-tac-toe/internal/model"
)

type Player interface {
	ChooseConstraint(b *board.Board) model.Line
	MakeMove(b *board.Board, line model.Line) model.Move
	Symbol() board.Cell
}
