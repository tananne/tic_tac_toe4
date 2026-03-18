package board

import "fmt"

type Board struct {
	Size  int
	Cells [][]Cell
}

func NewBoard(size int) *Board {
	cells := make([][]Cell, size)
	for i := range cells {
		cells[i] = make([]Cell, size)
	}

	return &Board{
		Size:  size,
		Cells: cells,
	}
}

func (b *Board) IsEmpty(row, col int) bool {
	return b.Cells[row][col] == Empty
}

func (b *Board) Place(row, col int, cell Cell) error {
	if b.IsEmpty(row, col) {
		b.Cells[row][col] = cell
		return nil
	}

	return fmt.Errorf("Данная ячейка уже занята")
}
