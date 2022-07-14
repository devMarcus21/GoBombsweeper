package Game

import (
	"testing"

	"github.com/devMarcus21/GoBombsweeper/src/game"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateASmallNewGame(t *testing.T) {
	row := 8
	col := 8
	expectedBoard := buildBoard(row, col)

	game := game.CreateGame(row, col)

	assert.Equal(t, expectedBoard, game.GetBoardState(), "Board state should be equal")
	assert.False(t, game.GameWon(), "Win state should be equal")
	assert.False(t, game.HasGameFinished(), "Finished state should be equal")
}

func buildBoard(row int, col int) [][]game.ISpace {
	board := make([][]game.ISpace, row)

	for r := range board {
		boardRow := make([]game.ISpace, col)
		for c := range boardRow {
			boardRow[c] = game.CreateSpace()
		}
		board[r] = boardRow
	}

	return board
}
