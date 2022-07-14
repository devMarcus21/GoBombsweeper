package game

type IGame interface {
	AddBomb(int, int) bool
	HasGameFinished() bool
	GameWon() bool
	GetBoardState() [][]ISpace
}
