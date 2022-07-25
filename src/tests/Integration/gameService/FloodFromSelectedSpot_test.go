package gameServiceTests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/devMarcus21/GoBombsweeper/src/game"
	"github.com/devMarcus21/GoBombsweeper/src/gameFactory"
	"github.com/devMarcus21/GoBombsweeper/src/gameService"
)

func TestShouldCreateAGameAndFillPath(t *testing.T) {
	rowSize := 8
	colSize := 8
	bombCount := 4
	err, gm := gameFactory.CreateNewGame(rowSize, colSize, bombCount)
	assert.Nil(t, err, "No error returned")

	fmt.Println(gm.GetBoardState()) // print inital board state
	fmt.Println()

	printBoardView(gm, rowSize, colSize)
	fmt.Println()

	err, res := gameService.FloodFromSelectedSpot(gm, 1, 1) // make 1st move
	assert.True(t, res, "Operation successful")
	assert.Nil(t, err, "No error returned")

	printBoardView(gm, rowSize, colSize) // print board state after 1st move was made
	fmt.Println()

	err, res = gameService.FloodFromSelectedSpot(gm, 1, 4) // make 2nd move
	assert.True(t, res, "Operation successful") // chance these could fail after selecting an already revealed spot will need more verbose end to end testing but for now this is good
	assert.Nil(t, err, "No error returned")

	printBoardView(gm, rowSize, colSize) // print board state after 2nd move was made
	fmt.Println()
}

func printBoardView(game game.IGame, rowSize int, colSize int) {
	for r := 0; r < rowSize; r++ {
		fmt.Print("[")
		for c := 0; c < colSize; c++ {
			_, space := game.GetSpaceState(r, c)
			if space.IsRevealed() {
				fmt.Print("+")
			} else {
				fmt.Print("-")
			}
		}
		fmt.Print("]\n")
	}
}