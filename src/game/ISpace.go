package game

type ISpace interface {
	ShowSpace() error
	IsRevealed() bool
	SetAdjacent(int) bool // Deprecate this method here in interface + implementations (BombSpace and Space)
	GetAdjacentBombs() int
	IncrementAdjacentBombs() bool
	String() string
	IsBombspace() bool
}
