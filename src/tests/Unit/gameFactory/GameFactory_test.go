package GameFactoryTest

import (
	"testing"

	"github.com/devMarcus21/GoBombsweeper/src/game"
	"github.com/devMarcus21/GoBombsweeper/src/gameFactory"
	"github.com/devMarcus21/GoBombsweeper/src/internalErrors"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateAGame(t *testing.T) {
	expectedRowSize := 8
	expectedColSize := 8
	expectedBombCount := 2

	err, gm := gameFactory.CreateNewGame(expectedRowSize, expectedColSize, expectedBombCount)

	assert.Nil(t, err, "No error returned")
	assert.NotNil(t, gm, "IGame returned not nil")

	_, IGameIsGame := gm.(*game.Game)
	rowSize, colSize := gm.GetBoardDimensions()

	// TODO add validation for bomb insertions

	assert.True(t, IGameIsGame, "Should return IGame that is a game")
	assert.Equal(t, expectedRowSize, rowSize, "Should be equal row sizes")
	assert.Equal(t, expectedColSize, colSize, "Should be equal col sizes")
	assert.False(t, gm.HasGameFinished(), "Should not be finished")
	assert.False(t, gm.GameWon(), "Should not be won")
}

func TestShouldNotCreateAGame_InvalidRowSize(t *testing.T) {
	row := -1
	col := 8

	err, game := gameFactory.CreateNewGame(row, col, 0)

	assert.Nil(t, game, "No game returned")
	assert.NotNil(t, err, "Should return error")
	assert.Equal(t, internalErrors.BuildInvalidRowSize(row), err, "Should return invalid row size error")
}

func TestShouldNotCreateAGame_InvalidColumnSize(t *testing.T) {
	row := 8
	col := -1

	err, game := gameFactory.CreateNewGame(row, col, 0)

	assert.Nil(t, game, "No game returned")
	assert.NotNil(t, err, "Should return error")
	assert.Equal(t, internalErrors.BuildInvalidColumnSize(col), err, "Should return invalid column size error")
}

func TestShouldNotCreateAGame_BombCountTooSmall(t *testing.T) {
	row := 8
	col := 8
	bombCount := -1

	err, game := gameFactory.CreateNewGame(row, col, bombCount)

	assert.Nil(t, game, "No game returned")
	assert.NotNil(t, err, "Should return error")
	assert.Equal(t, internalErrors.BuildBombCountToSmall(), err, "Should return bomb size too small error")
}

func TestShouldNotCreateAGame_BombCountTooLarge(t *testing.T) {
	row := 2
	col := 2
	bombCount := 3

	err, game := gameFactory.CreateNewGame(row, col, bombCount)

	assert.Nil(t, game, "No game returned")
	assert.NotNil(t, err, "Should return error")
	assert.Equal(t, internalErrors.BuildBombCountToLarge(bombCount), err, "Should return bomb size too large error")
}

func TestShouldNotCreateAGame_BombCountTooLargeNonEvenArea(t *testing.T) {
	row := 5
	col := 3
	bombCount := 8

	err, game := gameFactory.CreateNewGame(row, col, bombCount)

	assert.Nil(t, game, "No game returned")
	assert.NotNil(t, err, "Should return error")
	assert.Equal(t, internalErrors.BuildBombCountToLarge(bombCount), err, "Should return bomb size too large error")
}