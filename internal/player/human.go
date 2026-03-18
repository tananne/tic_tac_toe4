package player

import (
	"tic-tac-toe/internal/board"
	"tic-tac-toe/internal/model"
	"tic-tac-toe/internal/ui"
)

type Human struct {
	symbol board.Cell
	ui     ui.UI
}

func NewHumanPlayer(symbol board.Cell, ui ui.UI) Player {
	return &Human{
		symbol: symbol,
		ui:     ui,
	}
}

func (h *Human) ChooseConstraint(b *board.Board) model.Line {
	return h.ui.AskConstraint(h.symbol, b)
}

func (h *Human) MakeMove(b *board.Board, line model.Line) model.Move {
	return h.ui.AskMove(h.symbol, b, line)
}

func (h *Human) Symbol() board.Cell {
	return h.symbol
}
