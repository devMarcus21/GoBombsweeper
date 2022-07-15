package Game

import (
	"testing"

	"github.com/devMarcus21/GoBombsweeper/src/game"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateANewSpace(t *testing.T) {
	expectedAdjacent := 0

	space := game.CreateSpace()

	assert.Equal(t, expectedAdjacent, space.GetAdjacentBombs(), "Adjacent bomb count should be equal")
	assert.False(t, space.IsRevealed(), "Space should not be revealed and return false")
}

func TestShouldRevealSpace(t *testing.T) {
	space := game.CreateSpace()

	assert.False(t, space.IsRevealed(), "Space should not be yet revealed")

	err := space.ShowSpace()

	assert.Nil(t, err, "Should be no error")
	assert.True(t, space.IsRevealed(), "Space should be revealed")
}

func TestShowSpace(t *testing.T) {
	space := game.CreateSpace()
	err := space.ShowSpace()

	assert.Equal(t, nil, err, "No error returned")
	assert.True(t, space.IsRevealed(), "Space should be revealed and return false")
}

func TestShouldIncrementAdjacentBombsSpaceA(t *testing.T) {
	space := game.CreateSpace()

	assert.Equal(t, 0, space.GetAdjacentBombs(), "Should be no adjacentBombs")

	space.IncrementAdjacentBombs()

	assert.Equal(t, 1, space.GetAdjacentBombs(), "Should be 1 adjacentBomb")
}