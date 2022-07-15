package game

type IGame interface {
	AddBomb(int, int) (error, bool)
	HasGameFinished() bool
	GameWon() bool
	GetBoardState() [][]ISpace
	GetBoardDimensions() (int, int)
	IncrementAdjacentBombsAtIndex(int, int) (error, bool)
	GetSpaceState(int, int) (error, ISpace)
	SelectBombAtIndex(int, int) (error, bool)
}
