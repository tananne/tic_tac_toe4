package model

type Move struct {
	Row int
	Col int
}

type LineType int

const (
	ErrInvalidLine LineType = iota
	Row
	Column
)

type Line struct {
	Type  LineType
	Index int
}

type Mark rune

const (
	X     Mark = 'X'
	O     Mark = 'O'
	Empty Mark = ' '
)
