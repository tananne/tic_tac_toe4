package player

import (
	"tic-tac-toe/internal/board"
	"tic-tac-toe/internal/model"
)

type Player interface {
	ChooseLine(gameBoard *board.Board) (model.Line, error)
	MakeMove(gameBoard *board.Board, allowedLine model.Line) (model.Move, error)
	GetMark() model.Mark
	GetName() string
}
