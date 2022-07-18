package gameFactory

import (
	"github.com/devMarcus21/GoBombsweeper/src/game"
)

// function to add x amount of bombs to a Game board randomly
func AddBombsToBoard(game *game.Game, numBombs int, rowSize int, colSize int) bool {
	for numBombs > 0 {
		numBombs--
	}

	return true
}
