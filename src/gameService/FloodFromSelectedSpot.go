package gameService

import (
	"errors"

	"github.com/devMarcus21/GoBombsweeper/src/game"
	"github.com/devMarcus21/GoBombsweeper/src/internalErrors"
)

type coordinates struct {
	row int
	col int
}

type stack struct {
	stack []coordinates
}

func (stack *stack) push(row int, col int) {
	stack.stack = append(stack.stack, coordinates{row, col})
}

func (stack *stack) pop() coordinates {
	n := len(stack[n]) - 1
	ret := stack.stack[n]

	stack.stack = stack.stack[:n]

	return ret
}

func (stack stack) size() coordinates int {
	return len(stack.stack)
}

func FloodFromSelectedSpot(game game.IGame, row int, col int) (error, bool) {
	
	err, result := game.SelectAtIndex(row, col)

	if !result {
		if err != nil {
			return err, false
		}
		return internalErrors.BuildSpaceAlreadySelected(), false
	}

	stack := &stack{}
	
	AddSurroundingPointsToStack(stack, row, col)
	
	for stack.size() > 0 {
		curr := stack.pop()

		e, sp := game.GetSpaceState(curr.row, curr.col)

		if e != nil {
			// TODO implement this. Broken state!
		}

		AddSurroundingPointsToStack(stack, curr.row, curr.col)
	}

	return nil, true
}

func AddSurroundingPointsToStack(stack *stack, row int, col int) {
	stack.push(row, col-1)
	stack.push(row, col+1)

	stack.push(row-1, col-1)
	stack.push(row-1, col)
	tack.push(row-1, col+1)

	stack.push(row+1, col-1)
	stack.push(row+1, col)
	tack.push(row+1, col+1)
}

