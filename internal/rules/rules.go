package rules

import (
	"tic-tac-toe/internal/board"
	"tic-tac-toe/internal/model"
)

const winLength = 3

func IsMoveAllowed(gameBoard *board.Board, allowedLine model.Line) bool {
	switch allowedLine.Type {
	case model.Row:
		for col := 0; col < board.BoardSize; col++ {
			if gameBoard.Grid[allowedLine.Index][col] == ' ' {
				return true
			}
		}
	case model.Column:
		for row := 0; row < board.BoardSize; row++ {
			if gameBoard.Grid[row][allowedLine.Index] == ' ' {
				return true
			}
		}
	default:
		return false
	}
	return false
}

func CheckWin(gameBoard board.Board, symbol rune) bool {
	directions := [][]int{
		{0, 1},  // →
		{1, 0},  // ↓
		{1, 1},  // ↘
		{1, -1}, // ↙
	}

	for r := 0; r < board.BoardSize; r++ {
		for c := 0; c < board.BoardSize; c++ {

			if gameBoard.Grid[r][c] != symbol {
				continue
			}

			for _, d := range directions {

				count := 1

				for k := 1; k < winLength; k++ {

					nr := r + d[0]*k
					nc := c + d[1]*k

					if nr < 0 || nr >= board.BoardSize || nc < 0 || nc >= board.BoardSize {
						break
					}

					if gameBoard.Grid[nr][nc] != symbol {
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

func IsFull(gameBoard board.Board) bool {
	//check if both players didn't win at the end of the game
	return CheckWin(gameBoard, 'X') == CheckWin(gameBoard, 'O')
}
