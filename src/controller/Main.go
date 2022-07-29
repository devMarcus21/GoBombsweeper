package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/devMarcus21/GoBombsweeper/src/gameService"
)

type createGameRequestBody struct {
	Row       int `json:"row"`
	Col       int `json:"col"`
	BombCount int `json:"bombCount"`
}

type makeMoveRequestBody struct {
	GameId string `json:"gameId"`
	Row    int    `json:"row"`
	Col    int    `json:"col"`
}

var service *gameService.GameService = gameService.CreateGameService()

func main() {

	r := gin.Default()

	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "active",
		})
	})

	r.POST("/game/create", CreateGame)
	r.GET("/game/:gameId", GetGameStateById)
	r.POST("/game", MakeMoveOnBoardById)

	r.Run(":3000")
}

func CreateGame(c *gin.Context) {
	requestBody := createGameRequestBody{}
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid params",
			"error":   err.Error(),
		})
		return
	}

	var gameId string

	if err, id := service.CreateNewGoBombsweeperGame(requestBody.Row, requestBody.Col, requestBody.BombCount); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Game could not be created",
			"error":   err.Error(),
		})
		return
	} else {
		gameId = id
	}

	c.JSON(http.StatusOK, gin.H{
		"id": gameId,
	})
}

func GetGameStateById(c *gin.Context) {
	gameId := c.Param("gameId")

	if err, gameData := service.GetGameDataById(gameId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Game could not be retrieved",
			"error":   err.Error(),
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"id":       gameId,
			"board":    gameData.Board,
			"gameover": gameData.Gameover,
			"gameWon":  gameData.GameWon,
		})
	}
}

func MakeMoveOnBoardById(c *gin.Context) {
	requestBody := makeMoveRequestBody{}
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid params",
			"error":   err.Error(),
		})
		return
	}

	if err, result := service.MakeMoveOnBoardById(requestBody.GameId, requestBody.Row, requestBody.Col); result {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
		return
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid move",
			"error":   err.Error(),
		})
		return
	}
}
