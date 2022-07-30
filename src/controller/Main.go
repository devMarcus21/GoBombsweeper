package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"

	"github.com/devMarcus21/GoBombsweeper/src/gameService"
)

type createGameRequestBody struct {
	Row       string `json:"row"`
	Col       string `json:"col"`
	BombCount string `json:"bombCount"`
}

type makeMoveRequestBody struct {
	GameId string `json:"gameId"`
	Row    int    `json:"row"`
	Col    int    `json:"col"`
}

var service *gameService.GameService = gameService.CreateGameService()

func main() {

	r := gin.Default()
	r.Use(cors.Default())

	r.GET("/status", func(c *gin.Context) {
		//c.Header("Access-Control-Allow-Origin", "*")
		//c.Header("Access-Control-Allow-Credentials", "true")
		//c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		//c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")
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
	/*c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Credentials", "true")
	c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")*/
	/*c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")*/
	requestBody := createGameRequestBody{}
	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid params",
			"error":   err.Error(),
		})
		return
	}

	fmt.Println(requestBody.Row)
	fmt.Println(requestBody.Col)
	fmt.Println(requestBody.BombCount)

	var gameId string

	parseInt := func(s string) int {
		ret, _ := strconv.Atoi(s)

		return ret
	}

	if err, id := service.CreateNewGoBombsweeperGame(parseInt(requestBody.Row), parseInt(requestBody.Col), parseInt(requestBody.BombCount)); err != nil {
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
	//c.Header("Access-Control-Allow-Origin", "*")
	//c.Header("Access-Control-Allow-Credentials", "true")
	//c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	//c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")
	gameId := c.Param("gameId")

	if err, gameData := service.GetGameDataById(gameId); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Game could not be retrieved",
			"error":   err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"id":       gameId,
			"board":    gameData.Board,
			"gameover": gameData.Gameover,
			"gameWon":  gameData.GameWon,
		})
	}
}

func MakeMoveOnBoardById(c *gin.Context) {
	//c.Header("Access-Control-Allow-Origin", "*")
	//c.Header("Access-Control-Allow-Credentials", "true")
	//c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
	//c.Header("Access-Control-Allow-Methods", "POST,HEAD,PATCH, OPTIONS, GET, PUT")
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
