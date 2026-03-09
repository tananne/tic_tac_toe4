package board

import (
	"fmt"
	"tic-tac-toe/internal/model"
)

const BoardSize = 4

type Board struct {
	Grid [BoardSize][BoardSize]rune
}

func NewBoard() *Board {
	grid := [BoardSize][BoardSize]rune{}
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			grid[i][j] = ' '
		}
	}
	return &Board{Grid: grid}
}

func (b *Board) Display() {
	for i := 0; i <= BoardSize; i++ {
		for j := 0; j <= BoardSize; j++ {
			switch {
			case i == 0 && j == 0:
				fmt.Print(" ")
			case i == 0:
				fmt.Print(j)
			case j == 0:
				fmt.Print(i)
			case b.Grid[i-1][j-1] == 'X':
				fmt.Print("X")
			case b.Grid[i-1][j-1] == 'O':
				fmt.Print("O")
			default:
				fmt.Print("_")
			}

			if j < BoardSize {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

func (b *Board) IsEmpty(row, col int) bool {
	return b.Grid[row][col] == ' '
}

func (b *Board) PlaceMark(move model.Move, mark rune) (bool, error) {
	if b.IsEmpty(move.Row, move.Col) {
		b.Grid[move.Row][move.Col] = mark
		return true, nil
	}

	return false, fmt.Errorf("Данная ячейка уже занята")
}
