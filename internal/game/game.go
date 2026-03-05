package game

type Game struct {
	board         *Board
	currentPlayer rune
}

func NewGame() *Game {
	return &Game{
		board:         NewBoard(),
		currentPlayer: 'X',
	}
}

func (g *Game) switchPlayer() {
	if g.currentPlayer == 'X' {
		g.currentPlayer = 'O'
	} else {
		g.currentPlayer = 'X'
	}
}
