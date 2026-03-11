package board

import (
	"fmt"
	"tic-tac-toe/internal/model"
)

const BoardSize = 4

type Board struct {
	Grid [BoardSize][BoardSize]model.Mark
}

func NewBoard() *Board {
	grid := [BoardSize][BoardSize]model.Mark{}
	for i := 0; i < BoardSize; i++ {
		for j := 0; j < BoardSize; j++ {
			grid[i][j] = model.Empty
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
			case b.Grid[i-1][j-1] == model.X:
				fmt.Print(string(model.X))
			case b.Grid[i-1][j-1] == model.O:
				fmt.Print(string(model.O))
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

func (b *Board) IsEmptyCell(row, col int) bool {
	return b.Grid[row][col] == model.Empty
}

func (b *Board) LineHasEmptyCells(line model.Line) bool {
	switch line.Type {
	case model.Row:
		for col := 0; col < BoardSize; col++ {
			if b.IsEmptyCell(line.Index, col) {
				return true
			}
		}
	case model.Column:
		for row := 0; row < BoardSize; row++ {
			if b.IsEmptyCell(row, line.Index) {
				return true
			}
		}
	}
	return false
}

func (b *Board) PlaceMark(move model.Move, mark model.Mark) (bool, error) {
	if b.IsEmptyCell(move.Row, move.Col) {
		b.Grid[move.Row][move.Col] = mark
		return true, nil
	}

	return false, fmt.Errorf("Данная ячейка уже занята")
}
