package player

import (
	"tic-tac-toe/internal/board"
	"tic-tac-toe/internal/model"
	"tic-tac-toe/internal/rules"

	//"strings"
	//"strconv"
	"fmt"
	"math/rand"
)

type Bot struct {
	mark model.Mark
}

func NewBot(mark model.Mark) Player {
	return &Bot{mark: mark}
}

func (b *Bot) ChooseLine(gameBoard *board.Board) (model.Line, error) {
	lines := []model.Line{}
	opponentMark := model.X
	if b.mark == model.X {
		opponentMark = model.O
	}

	for i := 0; i < board.BoardSize; i++ {
		if gameBoard.LineHasEmptyCells(model.Line{Type: model.Row, Index: i}) {
			lines = append(lines, model.Line{Type: model.Row, Index: i})
		}
		if gameBoard.LineHasEmptyCells(model.Line{Type: model.Column, Index: i}) {
			lines = append(lines, model.Line{Type: model.Column, Index: i})
		}
	}

	if len(lines) == 0 {
		return model.Line{}, fmt.Errorf("Нет доступных линий для ходов.")
	}

	unwinnableOptions := []model.Line{}

	for _, line := range lines {
		canWin := false
		for i := 0; i < board.BoardSize; i++ {
			var move model.Move
			if line.Type == model.Row {
				move = model.Move{Row: line.Index, Col: i}
			} else {
				move = model.Move{Row: i, Col: line.Index}
			}

			if !gameBoard.IsEmptyCell(move.Row, move.Col) {
				continue
			}

			if rules.IsWinMove(gameBoard, move, opponentMark) {
				canWin = true
				break
			}
		}

		if !canWin {
			unwinnableOptions = append(unwinnableOptions, line)
		}
	}

	if len(unwinnableOptions) > 0 {
		randomIndex := rand.Intn(len(unwinnableOptions))
		switch unwinnableOptions[randomIndex].Type {
		case model.Row:
			fmt.Println("Бот выбрал строку ", unwinnableOptions[randomIndex].Index+1)
		case model.Column:
			fmt.Println("Бот выбрал колонку ", unwinnableOptions[randomIndex].Index+1)
		}
		return unwinnableOptions[randomIndex], nil
	}

	randomIndex := rand.Intn(len(lines))

	switch lines[randomIndex].Type {
	case model.Row:
		fmt.Println("Бот выбрал строку ", lines[randomIndex].Index+1)
	case model.Column:
		fmt.Println("Бот выбрал колонку ", lines[randomIndex].Index+1)
	}
	return lines[randomIndex], nil
}

func (b *Bot) MakeMove(gameBoard *board.Board, allowedLine model.Line) (model.Move, error) {
	moves := []model.Move{}
	switch allowedLine.Type {
	case model.Row:
		for col := 0; col < board.BoardSize; col++ {
			if gameBoard.IsEmptyCell(allowedLine.Index, col) {
				if rules.IsWinMove(gameBoard, model.Move{Row: allowedLine.Index, Col: col}, b.GetMark()) {
					return model.Move{Row: allowedLine.Index, Col: col}, nil
				}
				moves = append(moves, model.Move{Row: allowedLine.Index, Col: col})
			}
		}
	case model.Column:
		for row := 0; row < board.BoardSize; row++ {
			if gameBoard.IsEmptyCell(row, allowedLine.Index) {
				if rules.IsWinMove(gameBoard, model.Move{Row: row, Col: allowedLine.Index}, b.GetMark()) {
					return model.Move{Row: row, Col: allowedLine.Index}, nil
				}
				moves = append(moves, model.Move{Row: row, Col: allowedLine.Index})
			}
		}
	default:
		return model.Move{}, fmt.Errorf("Для данной линии нет доступных ходов.")
	}

	if len(moves) == 0 {
		return model.Move{}, fmt.Errorf("Нет доступных ходов в данной линии.")
	}

	randomIndex := rand.Intn(len(moves))
	return moves[randomIndex], nil
}

func (b *Bot) GetMark() model.Mark {
	return b.mark
}

func (b *Bot) GetName() string {
	return "Bot"
}
