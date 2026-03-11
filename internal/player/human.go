package player

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"tic-tac-toe/internal/board"
	"tic-tac-toe/internal/model"
)

type Human struct {
	name string
	mark model.Mark
}

func NewHuman(name string, mark model.Mark) Player {
	return &Human{name: name, mark: mark}
}

func (h *Human) ChooseLine(gameBoard *board.Board) (model.Line, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите линию: ")
	userInput, err := reader.ReadString('\n')
	if err != nil {
		return model.Line{}, err
	}
	userInput = strings.TrimSpace(userInput)
	fields := strings.Fields(userInput)

	if len(fields) != 2 {
		return model.Line{}, fmt.Errorf("Неверный формат ввода. Введите тип линии и номер через пробел.")
	}

	lineType := fields[0]
	index, err := strconv.Atoi(fields[1])
	if err != nil {
		return model.Line{}, fmt.Errorf("Неверный формат ввода. Введите корректный номер.")
	}

	index--
	if index < 0 || index >= board.BoardSize {
		return model.Line{}, fmt.Errorf("Номер линии вне диапазона.")
	}

	var line model.Line = model.Line{}

	switch lineType {
	case "row", "r", "ряд", "р":
		line = model.Line{Type: model.Row, Index: index}
	case "column", "c", "колонка", "к":
		line = model.Line{Type: model.Column, Index: index}
	default:
		return model.Line{}, fmt.Errorf("Неверный тип линии. Введите правильный тип.")
	}

	if gameBoard.LineHasEmptyCells(line) {
		return line, nil
	}

	return model.Line{}, fmt.Errorf("Выбранная линия недоступна для ходов.")
}

func (h *Human) MakeMove(gameBoard *board.Board, allowedLine model.Line) (model.Move, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Игрок " + h.name + " (" + string(h.mark) + "), введите адрес ячейки: ")
	userInput, err := reader.ReadString('\n')
	if err != nil {
		return model.Move{}, err
	}
	userInput = strings.TrimSpace(userInput)
	fields := strings.Fields(userInput)

	if len(fields) != 2 {
		return model.Move{}, fmt.Errorf("Неверный формат ввода. Введите два числа: номер строки и номер столбца.")
	}

	row, err1 := strconv.Atoi(fields[0])
	col, err2 := strconv.Atoi(fields[1])
	if err1 != nil || err2 != nil {
		return model.Move{}, fmt.Errorf("Неверный формат ввода. Убедитесь, что вы вводите числа для строки и столбца.")
	}

	if row < 1 || row > board.BoardSize || col < 1 || col > board.BoardSize {
		return model.Move{}, fmt.Errorf("Неверный адрес ячейки. Убедитесь, что вы вводите числа от 1 до %d для строки и столбца.", board.BoardSize)
	}

	if !gameBoard.IsEmptyCell(row-1, col-1) {
		return model.Move{}, fmt.Errorf("Данная ячейка уже занята. Выберите другую ячейку.")
	}

	switch {
	case allowedLine.Type == model.Row && row-1 != allowedLine.Index:
		return model.Move{}, fmt.Errorf("Вы должны ходить в строку %d.", allowedLine.Index+1)

	case allowedLine.Type == model.Column && col-1 != allowedLine.Index:
		return model.Move{}, fmt.Errorf("Вы должны ходить в столбец %d.", allowedLine.Index+1)
	default:
		return model.Move{Row: row - 1, Col: col - 1}, nil
	}
}

func (h *Human) GetMark() model.Mark {
	return h.mark
}

func (h *Human) GetName() string {
	return h.name
}
