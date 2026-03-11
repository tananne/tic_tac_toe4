package game

import (
	//"fmt"

	"tic-tac-toe/internal/board"
	"tic-tac-toe/internal/model"
	"tic-tac-toe/internal/player"
	//"tic-tac-toe/internal/rules"
)

type GameState struct {
	Board         *board.Board
	CurrentPlayer model.Mark
	AllowedLine   model.Line
	Turn          int
	GameOver      bool
	Winner        model.Mark
}

type Game struct {
	State *GameState

	PlayerX player.Player
	PlayerO player.Player
}

func NewGame(p1, p2 player.Player) *Game {
	gameState := &GameState{
		Board:         board.NewBoard(),
		CurrentPlayer: model.X,
		Turn:          0,
		GameOver:      false,
		Winner:        model.Empty,
	}
	return &Game{
		State:   gameState,
		PlayerX: p1,
		PlayerO: p2,
	}
}

func (g *Game) switchPlayer() {
	if g.State.CurrentPlayer == model.X {
		g.State.CurrentPlayer = model.O
	} else {
		g.State.CurrentPlayer = model.X
	}
}

// func (g *Game) applyAllowedLine(line model.Line) error {
// 	if rules.IsMoveAllowed(g.Board, line) {
// 		g.allowedLine = line
// 		return nil
// 	}
// 	return fmt.Errorf("В данной линии нет свободных ячеек. Выберите другую линию.")
// }
