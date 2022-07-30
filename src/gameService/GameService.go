package gameService

import (
	"fmt"

	"github.com/devMarcus21/GoBombsweeper/src/game"
	"github.com/devMarcus21/GoBombsweeper/src/gameFactory"
	"github.com/devMarcus21/GoBombsweeper/src/internalErrors"
	"github.com/google/uuid"
)

type GameService struct {
	currentGames map[string]game.IGame
}

type GameDataResponse struct {
	Board    [][]any
	Gameover bool
	GameWon  bool
}

type SpaceData struct {
	AdjacentBombs int
	Revealed      bool
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

func (service GameService) GetGameStateById(id string) (error, [][]game.ISpace) {
	game, gameFound := service.currentGames[id]
	if !gameFound {
		return internalErrors.BuildGameIdDoesNotExist(id), nil
	}

	return nil, game.GetBoardState()
}

func (service *GameService) GetGameDataById(id string) (error, GameDataResponse) {
	resp := GameDataResponse{}
	gm, gameFound := service.currentGames[id]
	if !gameFound {
		return internalErrors.BuildGameIdDoesNotExist(id), resp
	}

	board := gm.GetBoardState()

	boardOfAny := make([][]any, len(board))

	allSpacesFound := true

	// Bad logic need to fix
	for row := range board {
		boardOfAny[row] = make([]any, len(board[row]))

		for col := range board[row] {
			if board[row][col].IsRevealed() {
				boardOfAny[row][col] = SpaceData{board[row][col].GetAdjacentBombs(), board[row][col].IsRevealed()}
			} else {
				if fmt.Sprintf("%T", board[row][col]) == fmt.Sprintf("%T", &game.Space{}) {
					allSpacesFound = false
				}
				boardOfAny[row][col] = SpaceData{0, board[row][col].IsRevealed()}
			}
		}
	}

	if allSpacesFound {
		gm.SetGameWon()
	}
	resp.GameWon = gm.GameWon()
	resp.Gameover = gm.HasGameFinished()
	resp.Board = boardOfAny

	return nil, resp
}
