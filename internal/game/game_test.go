package game

import (
    "testing"
    "tic-tac-toe/internal/model"
    "tic-tac-toe/internal/player"
)

func TestNewGameInitialState(t *testing.T) {
    p1 := player.NewHuman("A", model.X)
    p2 := player.NewHuman("B", model.O)
    g := NewGame(p1, p2)

    if g.State == nil {
        t.Fatal("game state should not be nil")
    }
    if g.State.CurrentPlayer != model.X {
        t.Errorf("expected current player X, got %v", g.State.CurrentPlayer)
    }
    if g.State.Turn != 0 {
        t.Errorf("expected turn 0, got %d", g.State.Turn)
    }
    if g.State.GameOver {
        t.Error("game should not be over")
    }
    if g.PlayerX.GetMark() != model.X || g.PlayerO.GetMark() != model.O {
        t.Error("players marks incorrect")
    }
}

func TestSwitchPlayer(t *testing.T) {
    g := NewGame(player.NewHuman("A", model.X), player.NewHuman("B", model.O))
    g.switchPlayer()
    if g.State.CurrentPlayer != model.O {
        t.Errorf("expected O after switch, got %v", g.State.CurrentPlayer)
    }
    g.switchPlayer()
    if g.State.CurrentPlayer != model.X {
        t.Errorf("expected X after second switch, got %v", g.State.CurrentPlayer)
    }
}

func TestGameSimulatedPlay(t *testing.T) {
    // simple play simulate a few moves by directly manipulating board
    g := NewGame(player.NewBot(model.X), player.NewBot(model.O))
    g.State.Board.Grid[0][0] = model.X
    g.State.Turn = 1
    g.switchPlayer()
    if g.State.CurrentPlayer != model.O {
        t.Error("switchPlayer should change current player")
    }
}
