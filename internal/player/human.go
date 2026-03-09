package player

import (
	"fmt"
	"strconv"
	"strings"
	"tic-tac-toe/internal/board"
	"tic-tac-toe/internal/model"
)

type Human struct {
	name string
	mark rune
}

func (h *Human) ChooseLine() model.Line {
	var userInput string
	fmt.Print("Введите линию: ")
	fmt.Scanln(&userInput)
	fields := strings.Fields(userInput)

	if len(fields) != 2 {
		return model.Line{Type: model.ErrInvalidLine, Index: 0}
	}

	lineType := fields[0]
	index, err := strconv.Atoi(fields[1])
	if err != nil {
		return model.Line{Type: model.ErrInvalidLine, Index: -1}
	}

	switch lineType {
	case "row", "r", "ряд", "р":
		return model.Line{Type: model.Row, Index: index}
	case "column", "c", "колонка", "к":
		return model.Line{Type: model.Column, Index: index}
	default:
		return model.Line{Type: model.ErrInvalidLine, Index: -1}
	}
}

func (h *Human) MakeMove(gameBoard *board.Board, allowedLine model.Line) (model.Move, error) {
	var userInput string
	fmt.Print("Игрок " + h.name + " (" + string(h.mark) + "), введите адрес ячейки: ")
	fmt.Scanln(&userInput)
	fields := strings.Fields(userInput)

	if len(fields) != 2 {
		return model.Move{Row: -1, Col: -1}, fmt.Errorf("Неверный формат ввода. Введите два числа: номер строки и номер столбца.")
	}

	row, err1 := strconv.Atoi(fields[0])
	col, err2 := strconv.Atoi(fields[1])
	if err1 != nil || err2 != nil {
		return model.Move{Row: -1, Col: -1}, fmt.Errorf("Неверный формат ввода. Убедитесь, что вы вводите числа для строки и столбца.")
	}

	if row < 1 || row > board.BoardSize || col < 1 || col > board.BoardSize {
		return model.Move{Row: -1, Col: -1}, fmt.Errorf("Неверный адрес ячейки. Убедитесь, что вы вводите числа от 1 до %d для строки и столбца.", board.BoardSize)
	}

	if !gameBoard.IsEmpty(row-1, col-1) {
		return model.Move{Row: -1, Col: -1}, fmt.Errorf("Данная ячейка уже занята. Выберите другую ячейку.")
	}

	switch {
	case allowedLine.Type == model.Row && row-1 != allowedLine.Index:
		return model.Move{Row: -1, Col: -1}, fmt.Errorf("Вы должны ходить в строку %d.", allowedLine.Index+1)

	case allowedLine.Type == model.Column && col-1 != allowedLine.Index:
		return model.Move{Row: -1, Col: -1}, fmt.Errorf("Вы должны ходить в столбец %d.", allowedLine.Index+1)
	default:
		return model.Move{Row: row - 1, Col: col - 1}, nil
	}

}

func (h *Human) GetMark() rune {
	return h.mark
}

func (h *Human) GetName() string {
	return h.name
}
