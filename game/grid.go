package game

import (
	"math/rand"

	c "example.com/minesweeper/common"
)

var (
	Grid       [][]int
	Unexplored [][]int
	Init       = true
)

func InitGrids() {
  Init = false
	Grid = make([][]int, c.Length)
	Unexplored = make([][]int, c.Length)

	for i := 0; i < c.Length; i++ {
		Grid[i] = make([]int, c.Breadth)
		Unexplored[i] = make([]int, c.Breadth)
		for j := 0; j < c.Breadth; j++ {
			Grid[i][j] = 0
			Unexplored[i][j] = 10
		}
	}
	GenerateMines()
}

func GenerateMines() {
	generateRandomCoords := func() (int, int) {
		return rand.Intn(c.Length), rand.Intn(c.Breadth)
	}

	for i := 0; i < c.Mines; i++ {
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
			if 0 <= newX && newX < c.Length && 0 <= newY && newY < c.Breadth {
				Grid[newX][newY]++
			}
		}
	}
}
