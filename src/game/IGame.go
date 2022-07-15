package game

type IGame interface {
	AddBomb(int, int) (error, bool)
	HasGameFinished() bool
	GameWon() bool
	GetBoardState() [][]ISpace
	GetBoardDimensions() (int, int)
}
