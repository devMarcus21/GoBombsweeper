package Game

import (
	"errors"
	"strconv"
	"testing"

	"github.com/devMarcus21/GoBombsweeper/src/game"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateASmallNewGame(t *testing.T) {
	row := 8
	col := 8
	expectedBoard := buildBoard(row, col)

	game := game.CreateGame(row, col)

	gameRowLen, gameColLen := game.GetBoardDimensions()

	assert.Equal(t, expectedBoard, game.GetBoardState(), "Board state should be equal")
	assert.False(t, game.GameWon(), "Win state should be equal")
	assert.False(t, game.HasGameFinished(), "Finished state should be equal")
	assert.Equal(t, row, gameRowLen, "Row dimension is equal")
	assert.Equal(t, col, gameColLen, "Column dimension is equal")
}

func TestShouldAddBombToBoard(t *testing.T) {
	game := game.CreateGame(8, 8)

	err, result := game.AddBomb(1,1)

	assert.Nil(t, err, "No error returned")
	assert.True(t, result, "Bomb was added")
}

func TestShouldNotAddBombToBoard_SpaceAlreadyBomb(t *testing.T) {
	game := game.CreateGame(8, 8)
	expectedError := errors.New("Space already has bomb")

	game.AddBomb(1,1)
	err, result := game.AddBomb(1,1)

	assert.NotNil(t, err, "Error returned")
	assert.Equal(t, expectedError, err, "Erros are equal")
	assert.False(t, result, "Bomb was not added")
}

func TestShouldNotAddBombToBoard_RowIndexLess0(t *testing.T) {
	game := game.CreateGame(8, 8)
	row := -1
	col := 0
	expectedError := errors.New("Row index invalid: "+strconv.Itoa(row))

	err, result := game.AddBomb(row, col)

	assert.NotNil(t, err, "Error returned")
	assert.Equal(t, expectedError, err, "Errors are equal")
	assert.False(t, result, "Bomb was not added")
}

func TestShouldNotAddBombToBoard_RowIndexOutOfBounds(t *testing.T) {
	game := game.CreateGame(8, 8)
	row := 9
	col := 0
	expectedError := errors.New("Row index invalid: "+strconv.Itoa(row))

	err, result := game.AddBomb(row, col)

	assert.NotNil(t, err, "Error returned")
	assert.Equal(t, expectedError, err, "Errors are equal")
	assert.False(t, result, "Bomb was not added")
}

func TestShouldNotAddBombToBoard_ColIndexLess0(t *testing.T) {
	game := game.CreateGame(8, 8)
	row := 0
	col := -1
	expectedError := errors.New("Column index invalid: "+strconv.Itoa(col))

	err, result := game.AddBomb(row, col)

	assert.NotNil(t, err, "Error returned")
	assert.Equal(t, expectedError, err, "Errors are equal")
	assert.False(t, result, "Bomb was not added")
}

func TestShouldNotAddBombToBoard_ColIndexOutOfBounds(t *testing.T) {
	game := game.CreateGame(8, 8)
	row := 0
	col := 9
	expectedError := errors.New("Column index invalid: "+strconv.Itoa(col))

	err, result := game.AddBomb(row, col)

	assert.NotNil(t, err, "Error returned")
	assert.Equal(t, expectedError, err, "Errors are equal")
	assert.False(t, result, "Bomb was not added")
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
