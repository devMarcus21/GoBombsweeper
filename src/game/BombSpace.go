package game

import "errors"

// Implements ISpace interface used to represent bomb space on a board
type BombSpace struct {
	revealed bool
}

func CreateBombSpace() *BombSpace {
	return &BombSpace{false}
}

func (space *BombSpace) ShowSpace() error {
	space.revealed = true
	return errors.New("Bomb space selected")
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
