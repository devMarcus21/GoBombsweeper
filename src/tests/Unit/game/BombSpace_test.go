package GameTests

import (
	"testing"

	"github.com/devMarcus21/GoBombsweeper/src/game"
	"github.com/stretchr/testify/assert"
	"github.com/devMarcus21/GoBombsweeper/src/internalErrors"
)

func TestShouldCreateBombSpace(t *testing.T) {
	space := game.CreateBombSpace()

	assert.NotNil(t, space, "Should not be nil")
}

func TestShouldShowSpaceAndReveal_ReturnsError(t *testing.T) {
	space := game.CreateBombSpace()

	assert.False(t, space.IsRevealed(), "Should not be revealed now")

	err := space.ShowSpace()

	assert.Equal(t, internalErrors.BuildBombSpaceSelected(), err, "Should return BuildBombSpaceSelected error")
	assert.True(t, space.IsRevealed(), "Should be revealed now")
}

func TestShouldReturnFalseIncrementAdjacentBombs(t *testing.T) {
	space := game.CreateBombSpace()

	result := space.IncrementAdjacentBombs()

	assert.False(t, result, "Should not set adjacent by returning false")
}
