package GameFactoryTest

import (
	"testing"
	"fmt"

	"github.com/devMarcus21/GoBombsweeper/src/game"
	"github.com/devMarcus21/GoBombsweeper/src/gameFactory"
	"github.com/stretchr/testify/assert"
)

func TestAddBombsToBoard(t *testing.T) {
	rowSize := 8
	colSize := 8
	bombCount := 2

	gm := game.CreateGame(rowSize, colSize)

	result := gameFactory.AddBombsToBoard(gm, bombCount, rowSize, colSize)
	bombsFound := checkBoardForNumberOfBombs(gm, rowSize, colSize)

	fmt.Println(gm.GetBoardState())

	assert.True(t, result, "Result should be true")
	assert.Equal(t, bombCount, bombsFound, "Correct number of bombs should've been placed")
}

func checkBoardForNumberOfBombs(game *game.Game, rowSize int, colSize int) int {
	bombsFound := 0

	for r := 0; r < rowSize; r++ {
		for c := 0; c < colSize; c++ {
			_, isRegularSpace := game.SelectAtIndex(r, c)

			if !isRegularSpace {
				bombsFound++
			}
		}
	}

	return bombsFound
}