package board

import (
    "testing"
    "tic-tac-toe/internal/model"
)

func TestNewBoard(t *testing.T) {
    b := NewBoard()
    if b == nil {
        t.Fatal("NewBoard returned nil")
    }

    for r := 0; r < BoardSize; r++ {
        for c := 0; c < BoardSize; c++ {
            if b.Grid[r][c] != model.Empty {
                t.Errorf("expected cell %d,%d to be empty, got %v", r, c, b.Grid[r][c])
            }
        }
    }
}

func TestIsEmptyCellAndPlaceMark(t *testing.T) {
    b := NewBoard()
    if !b.IsEmptyCell(0, 0) {
        t.Error("expected (0,0) to be empty")
    }

    ok, err := b.PlaceMark(model.Move{Row: 0, Col: 0}, model.X)
    if err != nil || !ok {
        t.Fatalf("failed to place mark: %v", err)
    }

    if b.IsEmptyCell(0, 0) {
        t.Error("expected (0,0) not to be empty after placing")
    }

    // placing again should fail
    if _, err := b.PlaceMark(model.Move{Row: 0, Col: 0}, model.O); err == nil {
        t.Error("expected error when placing on occupied cell")
    }
}

func TestLineHasEmptyCells(t *testing.T) {
    b := NewBoard()

    // initially every row and column has empty cells
    for i := 0; i < BoardSize; i++ {
        if !b.LineHasEmptyCells(model.Line{Type: model.Row, Index: i}) {
            t.Errorf("expected row %d to have empty cells", i)
        }
        if !b.LineHasEmptyCells(model.Line{Type: model.Column, Index: i}) {
            t.Errorf("expected column %d to have empty cells", i)
        }
    }

    // fill row 1 completely
    for c := 0; c < BoardSize; c++ {
        b.Grid[1][c] = model.X
    }

    if b.LineHasEmptyCells(model.Line{Type: model.Row, Index: 1}) {
        t.Error("expected row 1 to have no empty cells")
    }

    // fill column 2 completely
    for r := 0; r < BoardSize; r++ {
        b.Grid[r][2] = model.O
    }

    if b.LineHasEmptyCells(model.Line{Type: model.Column, Index: 2}) {
        t.Error("expected column 2 to have no empty cells")
    }

    // other lines still have empties
    if !b.LineHasEmptyCells(model.Line{Type: model.Row, Index: 0}) {
        t.Error("expected row 0 to still have empty cells")
    }
    if !b.LineHasEmptyCells(model.Line{Type: model.Column, Index: 0}) {
        t.Error("expected column 0 to still have empty cells")
    }
}
