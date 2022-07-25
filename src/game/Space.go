package game

import "strconv"

// Implements ISpace interface used to represent space on a board
type Space struct {
	adjacentBombs int
	revealed      bool
}

func CreateSpace() *Space {
	return &Space{0, false}
}

func (space *Space) ShowSpace() error {
	space.revealed = true
	return nil
}

func (space *Space) IsRevealed() bool {
	return space.revealed
}

func (space *Space) SetAdjacent(adjCount int) bool {
	space.adjacentBombs = adjCount
	return true
}

func (space *Space) GetAdjacentBombs() int {
	return space.adjacentBombs
}

func (space *Space) IncrementAdjacentBombs() bool {
	space.adjacentBombs++
	return true
}

func (space Space) String() string {
	return strconv.Itoa(space.adjacentBombs)
}

func (space Space) IsBombspace() bool {
	return false
}
