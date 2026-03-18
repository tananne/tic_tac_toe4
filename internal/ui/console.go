package ui

import (
	"fmt"
	"tic-tac-toe/internal/board"
	"tic-tac-toe/internal/model"
)

type ConsoleUI struct{}

func NewConsoleUI() UI {
	return &ConsoleUI{}
}

// func (ui *ConsoleUI) clearInput() {
// 	var dummy string
// 	fmt.Scanln(&dummy)
// }

func (ui *ConsoleUI) AskConstraint(player board.Cell, b *board.Board) model.Line {
	var lineType model.LineType
	var index int

	for {
		fmt.Printf("Игрок %v, выберите тип линии (%d - строка, %d - столбец): ", string(player), model.Row, model.Column)
		_, err := fmt.Scan(&lineType)
		if err != nil || (lineType != model.Row && lineType != model.Column) {
			fmt.Println("Некорректный ввод. Попробуйте снова.")
			//ui.clearInput()
			continue
		}

		fmt.Printf("Игрок %v, введите номер линии (1-%d): ", string(player), b.Size)
		_, err = fmt.Scan(&index)
		if err != nil || index < 1 || index > b.Size {
			fmt.Println("Некорректный индекс")
			//ui.clearInput()
			continue
		}

		switch lineType {
		case model.Row:
			return model.Line{Type: model.Row, Index: index - 1}
		case model.Column:
			return model.Line{Type: model.Column, Index: index - 1}
		}
	}

}

func (ui *ConsoleUI) AskMove(player board.Cell, b *board.Board, line model.Line) model.Move {
	var pos int

	for {
		fmt.Printf("Игрок %v делает ход. Введите позицию (1-%d): ", player, b.Size)

		_, err := fmt.Scan(&pos)
		if err != nil || pos < 1 || pos > b.Size {
			fmt.Println("Ошибка ввода")
			//ui.clearInput()
			continue
		}

		switch line.Type {
		case model.Row:
			return model.Move{Row: line.Index, Col: pos - 1}
		case model.Column:
			return model.Move{Row: pos - 1, Col: line.Index}
		}
	}
}

func (ui *ConsoleUI) DisplayBoard(b *board.Board) {
	for i := 0; i <= b.Size; i++ {
		for j := 0; j <= b.Size; j++ {
			switch {
			case i == 0 && j == 0:
				fmt.Print(" ")
			case i == 0:
				fmt.Print(j)
			case j == 0:
				fmt.Print(i)
			case b.Cells[i-1][j-1] == board.X:
				fmt.Print(string(board.X))
			case b.Cells[i-1][j-1] == board.O:
				fmt.Print(string(board.O))
			default:
				fmt.Print("_")
			}

			if j < b.Size {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func (ui *ConsoleUI) ShowConstraint(player board.Cell, line model.Line) {
	switch line.Type {
	case model.Row:
		fmt.Printf("Игрок %v выбрал %d ряд\n", string(player), line.Index+1)
	case model.Column:
		fmt.Printf("Игрок %v выбрал %d колонку\n", string(player), line.Index+1)
	}
}

func (ui *ConsoleUI) ShowMove(player board.Cell, m model.Move) {
	fmt.Printf("Игрок %v делает ход: строка %d, столбец %d\n", string(player), m.Row+1, m.Col+1)
}

func (ui *ConsoleUI) AskGameMode() string {
	fmt.Println("1 - Игрок против игрока")
	fmt.Println("2 - Игрок против бота")
	var input int

	for {

		fmt.Scan(&input)
		switch input {
		case 1:
			return "human"
		case 2:
			return "bot"
		default:
			fmt.Println("Неверный выбор. Пожалуйста, выберите 1 или 2.")

		}
	}

}

func (ui *ConsoleUI) AskPlayerSymbol() board.Cell {
	fmt.Println("Вы играете против бота. Выберите сторону:")
	fmt.Println("1 - " + string(board.X))
	fmt.Println("2 - " + string(board.O))

	var input int

	for {
		fmt.Scan(&input)

		switch input {
		case 1:
			return board.X
		case 2:
			return board.O
		default:
			fmt.Println("Неверный выбор. Пожалуйста, выберите 1 или 2.")
		}
	}
}
