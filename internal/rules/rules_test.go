package rules

import (
    "testing"
    "tic-tac-toe/internal/board"
    "tic-tac-toe/internal/model"
)

// helper to set a mark
func setMarks(b *board.Board, coords ...[3]int) {
    // each coordinate: row,col,mark(0 for X,1 for O)
    for _, c := range coords {
        m := model.X
        if c[2] == 1 {
            m = model.O
        }
        b.Grid[c[0]][c[1]] = m
    }
}

func TestCheckWinHorizontal(t *testing.T) {
    b := board.NewBoard()
    setMarks(b, [3]int{0, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 2, 0})
    if !CheckWin(b, model.X) {
        t.Error("expected horizontal win for X")
    }
}

func TestCheckWinVertical(t *testing.T) {
    b := board.NewBoard()
    setMarks(b, [3]int{0, 1, 1}, [3]int{1, 1, 1}, [3]int{2, 1, 1})
    if !CheckWin(b, model.O) {
        t.Error("expected vertical win for O")
    }
}

func TestCheckWinDiagonal(t *testing.T) {
    b := board.NewBoard()
    setMarks(b, [3]int{0, 0, 0}, [3]int{1, 1, 0}, [3]int{2, 2, 0})
    if !CheckWin(b, model.X) {
        t.Error("expected diagonal win for X")
    }
}

func TestCheckWinNoWin(t *testing.T) {
    b := board.NewBoard()
    setMarks(b, [3]int{0, 0, 0}, [3]int{0, 1, 1}, [3]int{1, 0, 1})
    if CheckWin(b, model.X) {
        t.Error("did not expect win for X")
    }
    if CheckWin(b, model.O) {
        t.Error("did not expect win for O")
    }
}

func TestIsFull(t *testing.T) {
    b := board.NewBoard()
    // fill entirely with X; only X can win
    for r := 0; r < board.BoardSize; r++ {
        for c := 0; c < board.BoardSize; c++ {
            b.Grid[r][c] = model.X
        }
    }
    // in this case CheckWin(X)=true, CheckWin(O)=false so IsFull should be false
    if IsFull(b) {
        t.Error("expected IsFull to return false when only one side can win")
    }
    // clear and make only one win, board still not full by logic
    b = board.NewBoard()
    setMarks(b, [3]int{0, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 2, 0})
    if IsFull(b) {
        t.Error("board not full but IsFull returned true")
    }
}

func TestIsWinMove(t *testing.T) {
    b := board.NewBoard()
    // put two X in a row and test winning move
    setMarks(b, [3]int{0, 0, 0}, [3]int{0, 1, 0})
    move := model.Move{Row: 0, Col: 2}
    if !IsWinMove(b, move, model.X) {
        t.Error("expected win move for X")
    }

    // occupied cell should return false
    if IsWinMove(b, model.Move{Row: 0, Col: 0}, model.X) {
        t.Error("occupied cell should not be a win move")
    }

    // non-winning move
    if IsWinMove(b, model.Move{Row: 1, Col: 1}, model.X) {
        t.Error("non-winning move should return false")
    }
}
