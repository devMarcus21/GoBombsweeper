package gameService

import (
	"github.com/devMarcus21/Go-Stack/src/stack"
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

	return floodFillBoardTillStackIsEmpty(game, stack)
}

func floodFillBoardTillStackIsEmpty(game game.IGame, stack *stack.Stack[coordinates]) (error, bool) {
	for stack.Len() > 0 {
		_, curr := stack.Pop()

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

func addSurroundingPointsToStack(stack *stack.Stack[coordinates], row int, col int) {
	stack.Push(coordinates{row, col-1})
	stack.Push(coordinates{row, col+1})

	stack.Push(coordinates{row-1, col-1})
	stack.Push(coordinates{row-1, col})
	stack.Push(coordinates{row-1, col+1})

	stack.Push(coordinates{row+1, col-1})
	stack.Push(coordinates{row+1, col})
	stack.Push(coordinates{row+1, col+1})
}

