package game

import (
	"log"

	"example.com/minesweeper/common"
	c "example.com/minesweeper/common"
)

type CoordQ struct {
	Coords []c.Coord
}

func (Q *CoordQ) Enqueue(c c.Coord) {
	Q.Coords = append(Q.Coords, c)
}

func (Q *CoordQ) Dequeue() c.Coord {
	c := Q.Coords[0]
	Q.Coords = Q.Coords[1:]
	return c
}

var ExploreQ = &CoordQ{}

func Explore(x int, y int) {
	ExploreQ.Enqueue(c.Coord{X: x, Y: y})

	for len(ExploreQ.Coords) > 0 {
		log.Printf("Length of ExploreQ: %v", len(ExploreQ.Coords))
		coord := ExploreQ.Dequeue()
		i, j := coord.X, coord.Y
		// Boundary Conditions
		if i >= c.Length || j >= c.Breadth || i < 0 || j < 0 {
			continue
		}
		// Check if its anything else other than empty cell.
		if Unexplored[i][j] != 10 {
			continue
		}
		Unexplored[i][j] = Grid[i][j]
		CellsExplored++
		// If its an empty cell, explore further in all directions (including diagonals)
		if Grid[i][j] == 0 {
			ExploreQ.Enqueue(c.Coord{X: i + 1, Y: j})
			ExploreQ.Enqueue(c.Coord{X: i, Y: j + 1})
			ExploreQ.Enqueue(c.Coord{X: i - 1, Y: j})
			ExploreQ.Enqueue(c.Coord{X: i, Y: j - 1})
			ExploreQ.Enqueue(c.Coord{X: i + 1, Y: j + 1})
			ExploreQ.Enqueue(c.Coord{X: i - 1, Y: j + 1})
			ExploreQ.Enqueue(c.Coord{X: i + 1, Y: j - 1})
			ExploreQ.Enqueue(c.Coord{X: i - 1, Y: j - 1})
		}
	}
}

func CheckComplete() bool {
	totalCells := common.Length * common.Breadth
	if CellsExplored == totalCells-common.Mines {
		return true
	}
	return false
}
