package setup

import (
    "os"
    "testing"
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

func TestChooseModeValid(t *testing.T) {
    withStdin("1\n", func() {
        mode, err := ChooseMode()
        if err != nil || mode != "human" {
            t.Fatalf("expected human mode, got %s, err %v", mode, err)
        }
    })
    withStdin("2\n", func() {
        mode, err := ChooseMode()
        if err != nil || mode != "bot" {
            t.Fatalf("expected bot mode, got %s, err %v", mode, err)
        }
    })
}

func TestChooseModeInvalid(t *testing.T) {
    withStdin("9\n", func() {
        _, err := ChooseMode()
        if err == nil {
            t.Error("expected error for invalid choice")
        }
    })
}

func TestChooseSide(t *testing.T) {
    withStdin("1\n", func() {
        mark, err := ChooseSide()
        if err != nil || mark != model.X {
            t.Fatalf("expected X, got %v, err %v", mark, err)
        }
    })
    withStdin("2\n", func() {
        mark, err := ChooseSide()
        if err != nil || mark != model.O {
            t.Fatalf("expected O, got %v, err %v", mark, err)
        }
    })
}

func TestChooseSideInvalid(t *testing.T) {
    withStdin("0\n", func() {
        _, err := ChooseSide()
        if err == nil {
            t.Error("expected error for invalid side")
        }
    })
}

func TestGetPlayersHuman(t *testing.T) {
    withStdin("Alice\nBob\n", func() {
        pX, pO := GetPlayers("human")
        if pX.GetName() != "Alice" || pX.GetMark() != model.X {
            t.Errorf("unexpected player X: %v", pX)
        }
        if pO.GetName() != "Bob" || pO.GetMark() != model.O {
            t.Errorf("unexpected player O: %v", pO)
        }
    })
}

func TestGetPlayersBot(t *testing.T) {
    withStdin("1\n", func() {
        pH, pB := GetPlayers("bot")
        if pH.GetMark() != model.X {
            t.Errorf("expected human X got %v", pH.GetMark())
        }
        if pB.GetMark() != model.O {
            t.Errorf("expected bot O got %v", pB.GetMark())
        }
    })
}

func TestGetPlayersInvalid(t *testing.T) {
    p1, p2 := GetPlayers("nonsense")
    if p1 != nil || p2 != nil {
        t.Error("expected nil players for invalid mode")
    }
}
