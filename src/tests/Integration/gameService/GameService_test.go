package gameServiceTests

import (
	"fmt"
	"testing"

	"github.com/devMarcus21/GoBombsweeper/src/gameService"
	"github.com/stretchr/testify/assert"
)

func TestShouldCreateNewGameService(t *testing.T) {
	service := gameService.CreateGameService()

	assert.NotNil(t, service, "Should create GameService instance")
}

func TestShouldCreateANewGame(t *testing.T) {
	service := gameService.CreateGameService()
	expectedRow := 8
	expectedCol := 8
	expectedBombCount := 2

	err, id := service.CreateNewGoBombsweeperGame(expectedRow, expectedCol, expectedBombCount)

	fmt.Println("game id: "+id)

	assert.Nil(t, err, "Should not return errors")
	assert.NotEqual(t, "", id, "Should not be empty string")
}