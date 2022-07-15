package game

import (
	"errors"
	"strconv"
)

// Implements IGame interface
type Game struct {
	board [][]ISpace
	done  bool
	won   bool
	rowLength int
	colLength int
}

func CreateGame(row int, col int) *Game {
	board := make([][]ISpace, row)

	for r := range board {
		boardRow := make([]ISpace, col)
		for c := range boardRow {
			boardRow[c] = CreateSpace()
		}
		board[r] = boardRow
	}

	return &Game{board, false, false, row, col}
}

func (game Game) AddBomb(row int, col int) (error, bool) {
	if game.rowLength <= row || row < 0 {
		return errors.New("Row index invalid: "+strconv.Itoa(row)), false
	}
	if game.colLength <= col || col < 0 {
		return errors.New("Column index invalid: "+strconv.Itoa(col)), false
	}

	if _, ok := game.board[row][col].(*Space); ok {
		game.board[row][col] = CreateBombSpace()
		return nil, true
	}

	return errors.New("Space already has bomb"), false
}

func (game Game) HasGameFinished() bool {
	return game.done
}

func (game Game) GameWon() bool {
	return game.won
}

func (game *Game) GetBoardState() [][]ISpace {
	return game.board
}

func (game *Game) GetBoardDimensions() (int, int) {
	return game.rowLength, game.colLength
}