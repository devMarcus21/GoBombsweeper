package game

import "github.com/devMarcus21/GoBombsweeper/src/internalErrors"

// Implements ISpace interface used to represent bomb space on a board
type BombSpace struct {
	revealed bool
}

func CreateBombSpace() *BombSpace {
	return &BombSpace{false}
}

func (space *BombSpace) ShowSpace() error {
	space.revealed = true
	return internalErrors.BuildBombSpaceSelected()
}

func (space *BombSpace) IsRevealed() bool {
	return space.revealed
}

func (space *BombSpace) SetAdjacent(adjCount int) bool {
	return false
}

func (space *BombSpace) GetAdjacentBombs() int {
	return 0
}

func (space *BombSpace) IncrementAdjacentBombs() bool {
	return false
}

func (space BombSpace) String() string {
	return "#"
}

func (space BombSpace) IsBombspace() bool {
	return true
}
