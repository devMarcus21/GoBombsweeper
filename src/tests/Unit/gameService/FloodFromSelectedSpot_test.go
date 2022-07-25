package gameServiceUnitTests

import (
	"testing"

	"github.com/devMarcus21/Go-Stack/src/stack"
	"github.com/stretchr/testify/assert"
	"github.com/devMarcus21/GoBombsweeper/src/gameService"
)

func TestShouldAddSurroundingPointsToStack(t *testing.T) {
	row := 1
	col := 1
	stack := stack.New[gameService.Coordinates]()
	expected := []gameService.Coordinates{
		gameService.Coordinates{row+1, col+1},
		gameService.Coordinates{row+1, col},
		gameService.Coordinates{row+1, col-1},
		gameService.Coordinates{row-1, col+1},
		gameService.Coordinates{row-1, col},
		gameService.Coordinates{row-1, col-1},
		gameService.Coordinates{row, col+1},
		gameService.Coordinates{row, col-1},
	}

	gameService.AddSurroundingPointsToStack(stack, row, col)

	assert.Equal(t, 8, stack.Len(), "Stack should not be empty")

	stackState := []gameService.Coordinates{}

	for stack.Len() > 0 {
		_, coords := stack.Pop()
		stackState = append(stackState, coords)
	}

	assert.Equal(t, expected, stackState, "Stack should be equal")
}
