package gameFactory

import (
	"math/rand"
	"time"

	"github.com/devMarcus21/GoBombsweeper/src/game"
)

type void struct{}
var emptyValue void

// function to add x amount of bombs to a Game board randomly
func AddBombsToBoard(game *game.Game, numBombs int, rowSize int, colSize int) bool {
	// cache boards spots where a value exists already
	rowSet := make(map[int]void)
	colSet := make(map[int]void)
	rand.Seed(time.Now().UnixNano())

	for numBombs > 0 {

		for {
			r := rand.Intn(rowSize)
			c := rand.Intn(colSize)
			
			_, rowExists := rowSet[r]
			_, colExists := colSet[c]

			// keep iterating until we find a spot we have not been at yet
			if rowExists && colExists {
				continue
			}

			bombsAdded := AddBombsIfValid(game, r, c)

			if (bombsAdded) {
				rowSet[r] = emptyValue
				colSet[c] = emptyValue
				break
			}
		}

		numBombs--
	}

	return true
}

func AddBombsIfValid(gm *game.Game, row int, col int) bool {

	// Increments surrounding 8 spots by 1
	// Methods handles bounds checking
	fillAdjacentSpaceWithBombCount := func(refGame *game.Game, r int, c int) {
		refGame.IncrementAdjacentBombsAtIndex(r, c+1)
		refGame.IncrementAdjacentBombsAtIndex(r, c-1)
		refGame.IncrementAdjacentBombsAtIndex(r+1, c)
		refGame.IncrementAdjacentBombsAtIndex(r+1, c-1)
		refGame.IncrementAdjacentBombsAtIndex(r+1, c+1)
		refGame.IncrementAdjacentBombsAtIndex(r-1, c)
		refGame.IncrementAdjacentBombsAtIndex(r-1, c-1)
		refGame.IncrementAdjacentBombsAtIndex(r-1, c+1)
	}

	if _, result := gm.AddBomb(row, col); result {
		fillAdjacentSpaceWithBombCount(gm, row, col)
		return true
	}

	return false
}
