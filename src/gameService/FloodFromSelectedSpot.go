package gameService

import (
	"github.com/devMarcus21/Go-Stack/src/stack"
	"github.com/devMarcus21/GoBombsweeper/src/game"
	"github.com/devMarcus21/GoBombsweeper/src/internalErrors"
)

type Coordinates struct {
	Row int
	Col int
}

func FloodFromSelectedSpot(game game.IGame, row int, col int) (error, bool) {
	
	err, result := game.SelectAtIndex(row, col)

	if !result {
		if err != nil {
			return err, false
		}
		return internalErrors.BuildSpaceAlreadySelected(), false
	}

	_, spot := game.GetSpaceState(row, col)

	if spot.GetAdjacentBombs() > 0 {
		return nil, true
	}

	stack := stack.New[Coordinates]()
	
	AddSurroundingPointsToStack(stack, row, col)

	return FloodFillBoardTillStackIsEmpty(game, stack)
}

func FloodFillBoardTillStackIsEmpty(game game.IGame, stack *stack.Stack[Coordinates]) (error, bool) {
	for stack.Len() > 0 {
		_, curr := stack.Pop()

		e, sp := game.GetSpaceState(curr.Row, curr.Col)

		if e != nil {
			if e.Error() == internalErrors.BuildInvalidRowIndex(curr.Row).Error() || e.Error() == internalErrors.BuildInvalidColumnIndex(curr.Col).Error() { // if space is out of bounds move on
				continue
			}

			return e, false
		}

		if sp.IsRevealed() {
			continue
		}

		if !sp.IsBombspace() {
			sp.ShowSpace()
			if sp.GetAdjacentBombs() == 0 {
				AddSurroundingPointsToStack(stack, curr.Row, curr.Col)
			}
		}
	}

	return nil, true
}

func AddSurroundingPointsToStack(stack *stack.Stack[Coordinates], row int, col int) {
	stack.Push(Coordinates{row, col-1})
	stack.Push(Coordinates{row, col+1})

	stack.Push(Coordinates{row-1, col-1})
	stack.Push(Coordinates{row-1, col})
	stack.Push(Coordinates{row-1, col+1})

	stack.Push(Coordinates{row+1, col-1})
	stack.Push(Coordinates{row+1, col})
	stack.Push(Coordinates{row+1, col+1})
}

