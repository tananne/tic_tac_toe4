package player

import (
	"tic-tac-toe/internal/board"
	"tic-tac-toe/internal/model"
)

type Player interface {
	ChooseLine() (model.Line, error)
	MakeMove(board *board.Board, allowedLine model.Line) (model.Move, error)
	GetMark() rune
	GetName() string
}
