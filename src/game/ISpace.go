package game

type ISpace interface {
	ShowSpace() error
	IsRevealed() bool
	SetAdjacent(int) bool
	GetAdjacentBombs() int
	IncrementAdjacentBombs() bool
}
