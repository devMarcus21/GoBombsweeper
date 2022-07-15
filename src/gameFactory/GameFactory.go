package gameFactory

import (
	"github.com/devMarcus21/GoBombsweeper/src/game"
	"github.com/devMarcus21/GoBombsweeper/src/internalErrors"
)

func CreateNewGame(rowSize int, colSize int, bombCount int) (error, game.IGame) {
	if rowSize < 0 {
		return internalErrors.BuildInvalidRowSize(rowSize), nil
	}
	if colSize < 0 {
		return internalErrors.BuildInvalidColumnSize(colSize), nil
	}
	if bombCount < 0 {
		return internalErrors.BuildBombCountToSmall(), nil
	}
	if bombCount > (rowSize*colSize/2) {
		return internalErrors.BuildBombCountToLarge(bombCount), nil
	}

	createdGame := game.CreateGame(rowSize, colSize)

	return nil, createdGame
}