package GameTests

import (
	"testing"

	"github.com/devMarcus21/GoBombsweeper/src/game"
	"github.com/stretchr/testify/assert"
	"github.com/devMarcus21/GoBombsweeper/src/internalErrors"
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

func TestShouldAddBombToBoard(t *testing.T) {
	gm := game.CreateGame(8, 8)
	row := 1
	col := 1
	
	err, result := gm.AddBomb(row, col)
	getErr, space := gm.GetSpaceState(row, col)
	_, ISpaceIsBombSpace := space.(*game.BombSpace)

	assert.Nil(t, err, "No error returned")
	assert.True(t, result, "Bomb was added")

	assert.Nil(t, getErr, "No error when getting space state")
	assert.True(t, ISpaceIsBombSpace, "Space is now a bomb space")
}

func TestShouldNotAddBombToBoard_SpaceAlreadyBomb(t *testing.T) {
	game := game.CreateGame(8, 8)
	expectedError := internalErrors.BuildSpaceAlreadyHasBomb()

	game.AddBomb(1,1)
	err, result := game.AddBomb(1,1)

	assert.NotNil(t, err, "Error returned")
	assert.Equal(t, expectedError, err, "Errors are equal")
	assert.False(t, result, "Bomb was not added")
}

func TestShouldNotAddBombToBoard_RowIndexLess0(t *testing.T) {
	game := game.CreateGame(8, 8)
	row := -1
	col := 0
	expectedError := internalErrors.BuildInvalidRowIndex(row)

	err, result := game.AddBomb(row, col)

	assert.NotNil(t, err, "Error returned")
	assert.Equal(t, expectedError, err, "Errors are equal")
	assert.False(t, result, "Bomb was not added")
}

func TestShouldNotAddBombToBoard_RowIndexOutOfBounds(t *testing.T) {
	game := game.CreateGame(8, 8)
	row := 9
	col := 0
	expectedError := internalErrors.BuildInvalidRowIndex(row)

	err, result := game.AddBomb(row, col)

	assert.NotNil(t, err, "Error returned")
	assert.Equal(t, expectedError, err, "Errors are equal")
	assert.False(t, result, "Bomb was not added")
}

func TestShouldNotAddBombToBoard_ColIndexLess0(t *testing.T) {
	game := game.CreateGame(8, 8)
	row := 0
	col := -1
	expectedError := internalErrors.BuildInvalidColumnIndex(col)

	err, result := game.AddBomb(row, col)

	assert.NotNil(t, err, "Error returned")
	assert.Equal(t, expectedError, err, "Errors are equal")
	assert.False(t, result, "Bomb was not added")
}

func TestShouldNotAddBombToBoard_ColIndexOutOfBounds(t *testing.T) {
	game := game.CreateGame(8, 8)
	row := 0
	col := 9
	expectedError := internalErrors.BuildInvalidColumnIndex(col)

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
	_, ISpaceIsBombSpace := space.(*game.BombSpace)

	assert.Nil(t, err, "No error returned")
	// verify the state of the space
	assert.True(t, ISpaceIsBombSpace, "ISpace returned is a BombSpace")
	assert.False(t, space.IsRevealed(), "Space should not be revealed")
	assert.Equal(t, 0, space.GetAdjacentBombs(), "Space should have no adjacent bombs")
}

func TestShouldNotGetSpaceStateAtIndex_RowIndexLess0(t *testing.T) {
	game := game.CreateGame(8, 8)
	row := -1
	col := 0
	expectedError := internalErrors.BuildInvalidRowIndex(row)

	err, space := game.GetSpaceState(row, col)

	assert.Nil(t, space, "No ISpace was returned")
	assert.NotNil(t, err, "Error is not nil")
	assert.Equal(t, expectedError, err, "Errors are equal")
}

func TestShouldNotGetSpaceStateAtIndex_RowIndexOutOfBounds(t *testing.T) {
	game := game.CreateGame(8, 8)
	row := 9
	col := 0
	expectedError := internalErrors.BuildInvalidRowIndex(row)

	err, space := game.GetSpaceState(row, col)

	assert.Nil(t, space, "No ISpace was returned")
	assert.NotNil(t, err, "Error is not nil")
	assert.Equal(t, expectedError, err, "Errors are equal")
}

func TestShouldNotGetSpaceStateAtIndex_ColIndexLess0(t *testing.T) {
	game := game.CreateGame(8, 8)
	row := 0
	col := -1
	expectedError := internalErrors.BuildInvalidColumnIndex(col)

	err, space := game.GetSpaceState(row, col)

	assert.Nil(t, space, "No ISpace was returned")
	assert.NotNil(t, err, "Error is not nil")
	assert.Equal(t, expectedError, err, "Errors are equal")
}

func TestShouldNotGetSpaceStateAtIndex_ColIndexOutOfBounds(t *testing.T) {
	game := game.CreateGame(8, 8)
	row := 0
	col := 9
	expectedError := internalErrors.BuildInvalidColumnIndex(col)

	err, space := game.GetSpaceState(row, col)

	assert.Nil(t, space, "No ISpace was returned")
	assert.NotNil(t, err, "Error is not nil")
	assert.Equal(t, expectedError, err, "Errors are equal")
}

func TestShouldIncrementAdjacentBombsAtIndex(t *testing.T) {
	game := game.CreateGame(8, 8)
	row := 1
	col := 1

	_, spaceBefore := game.GetSpaceState(row, col)
	assert.Equal(t, 0, spaceBefore.GetAdjacentBombs(), "adjacentBombs should be 0")

	err, result := game.IncrementAdjacentBombsAtIndex(row, row)
	_, spaceAfter := game.GetSpaceState(row, col)

	assert.Nil(t, err, "Should not return error")
	assert.True(t, result, "Should return true")
	assert.Equal(t, 1, spaceAfter.GetAdjacentBombs(), "adjacentBombs should have been incremented")
}

func TestShouldSelectBombAtIndex(t *testing.T) {
	game := game.CreateGame(8, 8)
	row := 1
	col := 1

	_, spaceBefore := game.GetSpaceState(row, col)

	assert.False(t, spaceBefore.IsRevealed(), "Should not be revealed yet")
	err, result := game.SelectBombAtIndex(row, col)

	_, spaceAfter := game.GetSpaceState(row, col)

	assert.Nil(t, err, "Should not return error")
	assert.True(t, result, "Should return true")
	assert.True(t, spaceAfter.IsRevealed(), "Should reveal space")
}

func TestShouldNotSelectBombAtIndex_SpaceSelectedAlready(t *testing.T) {
	game := game.CreateGame(8, 8)
	row := 1
	col := 1

	game.SelectBombAtIndex(row, col)
	_, spaceBefore := game.GetSpaceState(row, col)

	assert.True(t, spaceBefore.IsRevealed(), "Should be revealed")
	err, result := game.SelectBombAtIndex(row, col)

	_, spaceAfter := game.GetSpaceState(row, col)

	assert.Nil(t, err, "Should not return error")
	assert.False(t, result, "Should return False")
	assert.True(t, spaceAfter.IsRevealed(), "Should still be revealed space")
}

func TestShouldNotSelectBombAtIndex_BombSpace(t *testing.T) {
	game := game.CreateGame(8, 8)
	row := 1
	col := 1
	expectedError := internalErrors.BuildBombSpaceSelected()
	game.AddBomb(row, col) // Add bomb space at index

	_, spaceBefore := game.GetSpaceState(row, col)

	assert.False(t, spaceBefore.IsRevealed(), "Should not be revealed yet")
	err, result := game.SelectBombAtIndex(row, col)

	_, spaceAfter := game.GetSpaceState(row, col)

	assert.NotNil(t, err, "Should return error")
	assert.Equal(t, expectedError, err, "Error should be expected")
	assert.False(t, result, "Should return False")
	assert.True(t, spaceAfter.IsRevealed(), "Should reveal space")
}