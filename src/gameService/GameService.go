package gameService

import (
	"github.com/devMarcus21/GoBombsweeper/src/game"
	"github.com/devMarcus21/GoBombsweeper/src/gameFactory"
	"github.com/devMarcus21/GoBombsweeper/src/internalErrors"
	"github.com/google/uuid"
)

type GameService struct {
	currentGames map[string]game.IGame
}

func CreateGameService() *GameService {
	return &GameService{make(map[string]game.IGame)}
}

func (service *GameService) CreateNewGoBombsweeperGame(row int, col int, bombCount int) (error, string) {
	gameId := uuid.New()

	err, gm := gameFactory.CreateNewGame(row, col, bombCount)

	if err == nil {
		service.currentGames[gameId.String()] = gm
		return nil, gameId.String()
	}

	return err, ""
}

func (service *GameService) MakeMoveOnBoardById(id string, row int, col int) (error, bool) {
	game, gameFound := service.currentGames[id]
	if !gameFound {
		return internalErrors.BuildGameIdDoesNotExist(id), false
	}

	return FloodFromSelectedSpot(game, row, col)
}
