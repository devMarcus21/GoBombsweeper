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

var service *gameService.GameService = gameService.CreateGameService()

func main() {

	r := gin.Default()

	r.GET("/status", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "active",
		})
	})

	r.POST("/game/create", CreateGame)

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
