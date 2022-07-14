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

func TestShowSpace(t *testing.T) {
	space := game.CreateSpace()
	err := space.ShowSpace()

	assert.Equal(t, nil, err, "No error returned")
	assert.True(t, space.IsRevealed(), "Space should be revealed and return false")
}
