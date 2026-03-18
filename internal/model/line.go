package model

type LineType int

const (
	InvalidLine LineType = iota
	Row
	Column
)

type Line struct {
	Type  LineType
	Index int
}
