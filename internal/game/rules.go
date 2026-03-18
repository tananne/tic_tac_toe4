package game

import (
	"tic-tac-toe/internal/board"
)

// func IsMoveAllowed(gameBoard *board.Board, allowedLine model.Line) bool {
// 	return gameBoard.LineHasEmptyCells(allowedLine)
// }

func IsWin(b *board.Board, symbol board.Cell, winLength int) bool {
	directions := [][]int{
		{0, 1},  // →
		{1, 0},  // ↓
		{1, 1},  // ↘
		{1, -1}, // ↙
	}

	for r := 0; r < b.Size; r++ {
		for c := 0; c < b.Size; c++ {

			if b.Cells[r][c] != symbol {
				continue
			}

			for _, d := range directions {

				count := 1

				for k := 1; k < winLength; k++ {

					nr := r + d[0]*k
					nc := c + d[1]*k

					if nr < 0 || nr >= b.Size || nc < 0 || nc >= b.Size {
						break
					}

					if b.Cells[nr][nc] != symbol {
						break
					}

					count++
				}

				if count == winLength {
					return true
				}
			}
		}
	}

	return false
}

func IsFull(b *board.Board, winLength int) bool {
	return IsWin(b, board.X, winLength) == IsWin(b, board.O, winLength)
}
