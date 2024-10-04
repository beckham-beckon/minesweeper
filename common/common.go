package common

const (
	MENU     = "MENU"
	GAME     = "GAME"
	GAMEOVER = "GAMEOVER"
)

var (
	Length  = 9
	Breadth = 9
	Mines   = 10
)

type Coord struct {
	X int
	Y int
}
