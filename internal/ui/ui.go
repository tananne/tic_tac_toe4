package ui

import (
	"tic-tac-toe/internal/board"
	"tic-tac-toe/internal/model"
)

type UI interface {
	AskConstraint(player board.Cell, b *board.Board) model.Line
	AskMove(player board.Cell, b *board.Board, line model.Line) model.Move
	DisplayBoard(b *board.Board)
	ShowConstraint(player board.Cell, line model.Line)
	ShowMove(player board.Cell, m model.Move)
	AskGameMode() string
	AskPlayerSymbol() board.Cell
}
