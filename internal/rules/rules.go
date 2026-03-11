package rules

import (
	"tic-tac-toe/internal/board"
	"tic-tac-toe/internal/model"
)

const winLength = 3

func IsMoveAllowed(gameBoard *board.Board, allowedLine model.Line) bool {
	return gameBoard.LineHasEmptyCells(allowedLine)
}

func CheckWin(gameBoard *board.Board, symbol model.Mark) bool {
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

func IsFull(gameBoard *board.Board) bool {
	//check if both players didn't win at the end of the game
	return CheckWin(gameBoard, model.X) == CheckWin(gameBoard, model.O)
}

func IsWinMove(gameBoard *board.Board, move model.Move, mark model.Mark) bool {
	if !gameBoard.IsEmptyCell(move.Row, move.Col) {
		return false
	}
	gameBoard.Grid[move.Row][move.Col] = mark

	canWin := CheckWin(gameBoard, mark)

	gameBoard.Grid[move.Row][move.Col] = model.Empty

	return canWin
}
