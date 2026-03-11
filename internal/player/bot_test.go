package player

import (
    "math/rand"
    "testing"
    "tic-tac-toe/internal/board"
    "tic-tac-toe/internal/model"
)

func TestBotChooseLineEmptyBoard(t *testing.T) {
    rand.Seed(1)
    b := NewBot(model.X)
    gameBoard := board.NewBoard()

    line, err := b.ChooseLine(gameBoard)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if line.Index < 0 || line.Index >= board.BoardSize {
        t.Fatalf("line index out of range: %v", line)
    }
}

func TestBotChooseLineNoOptions(t *testing.T) {
    gameBoard := board.NewBoard()
    for r := 0; r < board.BoardSize; r++ {
        for c := 0; c < board.BoardSize; c++ {
            gameBoard.Grid[r][c] = model.X
        }
    }
    b := NewBot(model.X)
    _, err := b.ChooseLine(gameBoard)
    if err == nil {
        t.Error("expected error when no lines available")
    }
}

func TestBotMakeMoveWinning(t *testing.T) {
    rand.Seed(2)
    gameBoard := board.NewBoard()
    gameBoard.Grid[0][0] = model.X
    gameBoard.Grid[0][2] = model.X

    bot := NewBot(model.X)
    move, err := bot.MakeMove(gameBoard, model.Line{Type: model.Row, Index: 0})
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if move.Row != 0 || move.Col != 1 {
        t.Errorf("expected winning move {0,1}, got %v", move)
    }
}

func TestBotMakeMoveRandom(t *testing.T) {
    rand.Seed(3)
    gameBoard := board.NewBoard()
    bot := NewBot(model.O)
    move, err := bot.MakeMove(gameBoard, model.Line{Type: model.Row, Index: 1})
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if move.Row != 1 {
        t.Errorf("expected row 1, got %v", move)
    }
    if move.Col < 0 || move.Col >= board.BoardSize {
        t.Errorf("col out of range: %v", move.Col)
    }
}

func TestBotMakeMoveNoMoves(t *testing.T) {
    rand.Seed(4)
    gameBoard := board.NewBoard()
    for r := 0; r < board.BoardSize; r++ {
        gameBoard.Grid[r][2] = model.X
    }
    bot := NewBot(model.O)
    _, err := bot.MakeMove(gameBoard, model.Line{Type: model.Column, Index: 2})
    if err == nil {
        t.Error("expected error when allowed line has no empty cells")
    }
}
