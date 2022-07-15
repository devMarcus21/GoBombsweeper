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
	assert.False(t, game.GameWon(), "Win state should be false")
	assert.False(t, game.HasGameFinished(), "Finished state should be false")
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

func TestShouldGetSpaceStateAtIndex(t *testing.T) {
	gm := game.CreateGame(8, 8)

	err, space := gm.GetSpaceState(1, 1)
	_, ISpaceIsSpace := space.(*game.Space)

	assert.Nil(t, err, "No error returned")
	// verify the state of the space
	assert.True(t, ISpaceIsSpace, "ISpace returned is a Space")
	assert.False(t, space.IsRevealed(), "Space should not be revealed")
	assert.Equal(t, 0, space.GetAdjacentBombs(), "Space should have no adjacent bombs")
}

func TestShouldGetBombSpaceStateAtIndex(t *testing.T) {
	gm := game.CreateGame(8, 8)
	row := 1
	col := 1

	gm.AddBomb(row, col)
	err, space := gm.GetSpaceState(row, col)
	_, ISpaceIsSpace := space.(*game.BombSpace)

	assert.Nil(t, err, "No error returned")
	// verify the state of the space
	assert.True(t, ISpaceIsSpace, "ISpace returned is a BombSpace")
	assert.False(t, space.IsRevealed(), "Space should not be revealed")
	assert.Equal(t, 0, space.GetAdjacentBombs(), "Space should have no adjacent bombs")
}

func TestShouldNotGetSpaceStateAtIndex_RowIndexLess0(t *testing.T) {
	game := game.CreateGame(8, 8)
	row := -1
	col := 0
	expectedError := errors.New("Row index invalid: "+strconv.Itoa(row))

	err, space := game.GetSpaceState(row, col)

	assert.Nil(t, space, "No ISpace was returned")
	assert.NotNil(t, err, "Error is not nil")
	assert.Equal(t, expectedError, err, "Errors are equal")
}

func TestShouldNotGetSpaceStateAtIndex_RowIndexOutOfBounds(t *testing.T) {
	game := game.CreateGame(8, 8)
	row := 9
	col := 0
	expectedError := errors.New("Row index invalid: "+strconv.Itoa(row))

	err, space := game.GetSpaceState(row, col)

	assert.Nil(t, space, "No ISpace was returned")
	assert.NotNil(t, err, "Error is not nil")
	assert.Equal(t, expectedError, err, "Errors are equal")
}

func TestShouldNotGetSpaceStateAtIndex_ColIndexLess0(t *testing.T) {
	game := game.CreateGame(8, 8)
	row := 0
	col := -1
	expectedError := errors.New("Column index invalid: "+strconv.Itoa(col))

	err, space := game.GetSpaceState(row, col)

	assert.Nil(t, space, "No ISpace was returned")
	assert.NotNil(t, err, "Error is not nil")
	assert.Equal(t, expectedError, err, "Errors are equal")
}

func TestShouldNotGetSpaceStateAtIndex_ColIndexOutOfBounds(t *testing.T) {
	game := game.CreateGame(8, 8)
	row := 0
	col := 9
	expectedError := errors.New("Column index invalid: "+strconv.Itoa(col))

	err, space := game.GetSpaceState(row, col)

	assert.Nil(t, space, "No ISpace was returned")
	assert.NotNil(t, err, "Error is not nil")
	assert.Equal(t, expectedError, err, "Errors are equal")
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
