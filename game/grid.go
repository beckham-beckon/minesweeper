package game

import "math/rand"

var (
	LENGTH  = 9
	BREADTH = 9
	MINES   = 10
)

var (
	Grid       = make([][]int, LENGTH)
	Unexplored = make([][]int, LENGTH)
)

func InitGrids() {
	for i := 0; i < LENGTH; i++ {
		Grid[i] = make([]int, BREADTH)
		Unexplored[i] = make([]int, BREADTH)
		for j := 0; j < BREADTH; j++ {
			Grid[i][j] = 0
			Unexplored[i][j] = 10
		}
	}
	GenerateMines()
}

func GenerateMines() {
	generateRandomCoords := func() (int, int) {
		return rand.Intn(LENGTH), rand.Intn(BREADTH)
	}

	for i := 0; i < MINES; i++ {
		var x, y int
		// Place mine if block is not a mine
		for {
			x, y = generateRandomCoords()
			if Grid[x][y] >= 0 {
				Grid[x][y] = -9
				break
			}
		}
		AdjustSurroundingCells(x, y)
	}
}

func AdjustSurroundingCells(x int, y int) {
	// If surrounding block is not a mine, increase its value by 1
	for i := -1; i < 2; i++ {
		for j := -1; j < 2; j++ {
			if i == 0 || j == 0 {
				continue
			}
			newX, newY := x+i, y+j
			if 0 <= newX && newX < LENGTH && 0 <= newY && newY < BREADTH {
				Grid[newX][newY]++
			}
		}
	}
}
