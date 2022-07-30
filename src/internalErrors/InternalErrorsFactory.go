package internalErrors

import (
	"errors"
	"strconv"
)

func BuildInvalidRowIndex(index int) error {
	return errors.New("Row index invalid: " + strconv.Itoa(index))
}

func BuildInvalidColumnIndex(index int) error {
	return errors.New("Column index invalid: " + strconv.Itoa(index))
}

func BuildInvalidRowSize(size int) error {
	return errors.New("Row index invalid: " + strconv.Itoa(size))
}

func BuildInvalidColumnSize(size int) error {
	return errors.New("Column index invalid: " + strconv.Itoa(size))
}

func BuildBombCountToSmall() error {
	return errors.New("bombCount cannot be less than zero")
}

func BuildBombCountToLarge(count int) error {
	return errors.New("bombCount cannot greater than half the area of the board: " + strconv.Itoa(count))
}

func BuildSpaceAlreadyHasBomb() error {
	return errors.New("Space already has bomb")
}

func BuildBombSpaceSelected() error {
	return errors.New("Bomb space selected")
}

func BuildBombsWereNotAddedToGame() error {
	return errors.New("Bombs were not added to game")
}

func BuildSpaceAlreadySelected() error {
	return errors.New("Space already selected")
}

func BuildGameIdDoesNotExist(gameId string) error {
	return errors.New("Game id does not exist: " + gameId)
}

func BuildGameover() error {
	return errors.New("Gameover")
}

func BuildGameWon() error {
	return errors.New("Game won")
}
