package gameService

import (
	"errors"

	"github.com/devMarcus21/Go-Stack"
	"github.com/devMarcus21/GoBombsweeper/src/game"
	"github.com/devMarcus21/GoBombsweeper/src/internalErrors"
)

type coordinates struct {
	row int
	col int
}

func FloodFromSelectedSpot(game game.IGame, row int, col int) (error, bool) {
	
	err, result := game.SelectAtIndex(row, col)

	if !result {
		if err != nil {
			return err, false
		}
		return internalErrors.BuildSpaceAlreadySelected(), false
	}

	stack := stack.New[coordinates]()
	
	addSurroundingPointsToStack(stack, row, col)

	return floodFillBoardTillStackIsEmpty(stack)
}

func floodFillBoardTillStackIsEmpty(stack *stack.Stack) (error, bool) {
	for stack.size() > 0 {
		curr := stack.Pop()

		e, sp := game.GetSpaceState(curr.row, curr.col)

		if e != nil {
			if e == internalErrors.BuildInvalidRowIndex(curr.row) || e == internalErrors.BuildInvalidColumnIndex(curr.col) { // if space is out of bounds move on
				continue
			}

			return e, false
		}

		if sp.IsRevealed() {
			continue
		}

		if !sp.IsBombspace() {
			sp.ShowSpace()
			addSurroundingPointsToStack(stack, curr.row, curr.col)
		}
	}

	return nil, true
}

func addSurroundingPointsToStack(stack *stack.Stack, row int, col int) {
	stack.Push(row, col-1)
	stack.Push(row, col+1)

	stack.Push(row-1, col-1)
	stack.Push(row-1, col)
	tack.Push(row-1, col+1)

	stack.Push(row+1, col-1)
	stack.Push(row+1, col)
	tack.Push(row+1, col+1)
}

