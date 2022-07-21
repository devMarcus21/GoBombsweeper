package gameService

import (
	"github.com/devMarcus21/GoBombsweeper/src/game"
	"github.com/devMarcus21/GoBombsweeper/src/gameFactory"
	"github.com/google/uuid"
)

type GameService struct {
	currentGames map[string] game.IGame
}


func CreateGameService() *GameService{
	return &GameService{make(map[string]game.IGame)}
}

func (service *GameService) CreateNewGoBombsweeperGame(row int, col int, bombCount int) (error, string) {
	gameId := uuid.New()
	// service.currentGames[gameId] = gameFactory.CreateNewGame(row, col)
	err, gm := gameFactory.CreateNewGame(row, col, bombCount)

	if err == nil {
		service.currentGames[gameId.String()] = gm
		return nil, gameId.String()
	}

	return err, ""
}