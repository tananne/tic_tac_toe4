package player

import (
	"tic-tac-toe/internal/board"
	"tic-tac-toe/internal/model"
)

type AI struct {
	symbol board.Cell
}

func NewAIPlayer(symbol board.Cell) Player {
	return &AI{symbol: symbol}
}

func (ai *AI) Symbol() board.Cell {
	return ai.symbol
}

func (ai *AI) ChooseConstraint(b *board.Board) model.Line {
	worstScore := 999
	var worstConstraint model.Line

	for i := 0; i < b.Size; i++ {
		for _, t := range []model.LineType{model.Row, model.Column} {
			c := model.Line{Type: t, Index: i}

			score := ai.evaluateConstraintForOpponent(b, c)

			if score < worstScore {
				worstScore = score
				worstConstraint = c
			}
		}
	}

	return worstConstraint
}

func (ai *AI) evaluateConstraintForOpponent(b *board.Board, line model.Line) int {
	opponent := opposite(ai.symbol)

	moves := availableMovesInConstraint(b, line)

	bestScore := -1

	for _, m := range moves {
		score := evaluateMove(opponent, b, m)

		if score > bestScore {
			bestScore = score
		}
	}

	return bestScore
}

func (ai *AI) MakeMove(b *board.Board, line model.Line) model.Move {
	bestScore := -1
	var bestMove model.Move

	moves := availableMovesInConstraint(b, line)

	for _, m := range moves {
		score := evaluateMove(ai.symbol, b, m)

		if score > bestScore {
			bestScore = score
			bestMove = m
		}
	}

	return bestMove
}

func availableMovesInConstraint(b *board.Board, c model.Line) []model.Move {
	var moves []model.Move

	for i := 0; i < b.Size; i++ {
		var m model.Move

		switch c.Type {
		case model.Row:
			m = model.Move{Row: c.Index, Col: i}
		case model.Column:
			m = model.Move{Row: i, Col: c.Index}
		}

		if b.IsEmpty(m.Row, m.Col) {
			moves = append(moves, m)
		}
	}

	return moves
}

func opposite(c board.Cell) board.Cell {
	if c == board.X {
		return board.O
	}
	return board.X
}

func evaluateMove(symbol board.Cell, b *board.Board, m model.Move) int {
	b.Place(m.Row, m.Col, symbol)
	score := countPotentialLines(b, symbol)
	b.Place(m.Row, m.Col, board.Empty)
	return score
}

func countPotentialLines(b *board.Board, symbol board.Cell) int {
	//FUTURE IS HERE

	return 0

}
