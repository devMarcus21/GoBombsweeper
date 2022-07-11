package game

type Game struct {
	board [][]ISpace
	done  bool
	won   bool
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

	return &Game{board, false, false}
}

func (game *Game) AddBomb(row int, col int) bool {
	game.board[row][col] = CreateBombSpace()
	return true
}
