package game

import (
	"fmt"
	"tic-tac-toe/internal/board"
	"tic-tac-toe/internal/model"
	"tic-tac-toe/internal/player"
	"tic-tac-toe/internal/rules"
)

type Game struct {
	Board         *board.Board
	players       [2]player.Player
	currentPlayer int
	allowedLine   model.Line
	turns         int
}

func NewGame() *Game {
	return &Game{
		Board:         board.NewBoard(),
		currentPlayer: 0,
		turns:         board.BoardSize * board.BoardSize,
	}
}

func (g *Game) switchPlayer() {
	g.currentPlayer = (g.currentPlayer + 1) % 2
}

func (g *Game) applyAllowedLine(line model.Line) error {
	if rules.IsMoveAllowed(g.Board, line) {
		g.allowedLine = line
		return nil
	}
	return fmt.Errorf("В данной линии нет свободных ячеек. Выберите другую линию.")
}
