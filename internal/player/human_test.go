package player

import (
    "os"
    "testing"
    "tic-tac-toe/internal/board"
    "tic-tac-toe/internal/model"
)

func withStdin(input string, fn func()) {
    orig := os.Stdin
    defer func() { os.Stdin = orig }()
    r, w, _ := os.Pipe()
    w.WriteString(input)
    w.Close()
    os.Stdin = r
    fn()
}

func TestHumanChooseLineValid(t *testing.T) {
    h := NewHuman("foo", model.X)
    b := board.NewBoard()

    withStdin("row 1\n", func() {
        line, err := h.ChooseLine(b)
        if err != nil {
            t.Fatalf("unexpected error: %v", err)
        }
        if line.Type != model.Row || line.Index != 0 {
            t.Errorf("expected row 0, got %v", line)
        }
    })
}

func TestHumanChooseLineInvalidFormat(t *testing.T) {
    h := NewHuman("foo", model.X)
    b := board.NewBoard()

    withStdin("badinput\n", func() {
        _, err := h.ChooseLine(b)
        if err == nil {
            t.Error("expected error for invalid format")
        }
    })
}

func TestHumanMakeMoveValid(t *testing.T) {
    h := NewHuman("foo", model.X)
    b := board.NewBoard()

    withStdin("2 3\n", func() {
        move, err := h.MakeMove(b, model.Line{Type: model.Row, Index: 1})
        if err != nil {
            t.Fatalf("unexpected error: %v", err)
        }
        if move.Row != 1 || move.Col != 2 {
            t.Errorf("expected {1,2}, got %v", move)
        }
    })
}

func TestHumanMakeMoveWrongLine(t *testing.T) {
    h := NewHuman("foo", model.X)
    b := board.NewBoard()

    withStdin("1 1\n", func() {
        _, err := h.MakeMove(b, model.Line{Type: model.Row, Index: 1})
        if err == nil {
            t.Error("expected error because row does not match allowed line")
        }
    })
}

func TestHumanMakeMoveOccupied(t *testing.T) {
    h := NewHuman("foo", model.X)
    b := board.NewBoard()
    b.Grid[1][1] = model.O

    withStdin("2 2\n", func() {
        _, err := h.MakeMove(b, model.Line{Type: model.Row, Index: 1})
        if err == nil {
            t.Error("expected error for occupied cell")
        }
    })
}
