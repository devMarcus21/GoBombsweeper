package game

import "errors"

type BombSpace struct {
	revealed bool
}

func CreateBombSpace() BombSpace {
	return BombSpace{false}
}

func (space BombSpace) ShowSpace() error {
	space.revealed = true
	return errors.New("Bomb space selected")
}

func (space BombSpace) IsRevealed() bool {
	return space.revealed
}

func (space BombSpace) SetAdjacent(adjCount int) bool {
	return false
}
