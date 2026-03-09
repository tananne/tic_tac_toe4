package player

import (
	"tic-tac-toe/internal/board"
	"tic-tac-toe/internal/model"

	//"strings"
	//"strconv"
	"fmt"
)

type Bot struct {
	mark rune
}

func (b *Bot) ChooseLine() (model.Line, error) {
	// Implementation for bot to choose a line
	return model.Line{}, fmt.Errorf("Not implemented")
}

func (b *Bot) MakeMove(board *board.Board, allowedLine model.Line) model.Move {
	//add logic for bot to make a move based on the allowed line and current board state
	return model.Move{}
}

func (b *Bot) GetMark() rune {
	return b.mark
}

func (b *Bot) GetName() string {
	return "Bot"
}
