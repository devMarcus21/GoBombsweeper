package game

type Space struct {
	adjacentBombs int
	revealed      bool
}

func CreateSpace() Space {
	return Space{0, false}
}

func (space Space) ShowSpace() error {
	space.revealed = true
	return nil
}

func (space Space) IsRevealed() bool {
	return space.revealed
}

func (space Space) SetAdjacent(adjCount int) bool {
	space.adjacentBombs = adjCount
	return true
}
