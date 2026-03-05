package game

import (
	"fmt"
)

type Board struct {
	grid [4][4]rune
}

func NewBoard() *Board {
	grid := [4][4]rune{}
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			grid[i][j] = ' '
		}
	}
	return &Board{grid: grid}
}

func (b *Board) Display() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			switch {
			case i == 0 && j == 0:
				fmt.Print(" ")
			case i == 0:
				fmt.Print(j)
			case j == 0:
				fmt.Print(i)
			case b.grid[i-1][j-1] == 'X':
				fmt.Print("X")
			case b.grid[i-1][j-1] == 'O':
				fmt.Print("O")
			default:
				fmt.Print(" ")
			}

			if j < 4 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if i < 4 {
			fmt.Println("----------")
		}
	}
}

func (b *Board) IsEmpty(row, col int) bool {
	return b.grid[row][col] == ' '
}

func (b *Board) MakeMove(row, col int, mark rune) bool {
	row--
	col--
	if b.IsEmpty(row, col) {
		b.grid[row][col] = mark
		return true
	}
	return false
}
