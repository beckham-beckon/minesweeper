package common

const (
	MENU     = "MENU"
	GAME     = "GAME"
	GAMEOVER = "GAMEOVER"
)

var (
	Length  int
	Breadth int
  Mines   int
)

type Coord struct {
	X int
	Y int
}

