package game

type Coord struct {
	X int
	Y int
}

type CoordQ struct {
	Coords []Coord
}

func (Q *CoordQ) Enqueue(c Coord) {
	Q.Coords = append(Q.Coords, c)
}

func (Q *CoordQ) Dequeue() Coord {
	c := Q.Coords[0]
	Q.Coords = Q.Coords[1:]
	return c
}

var ExploreQ = &CoordQ{}

func Explore() {
	for len(ExploreQ.Coords) > 0 {
		c := ExploreQ.Dequeue()
		i, j := c.X, c.Y
		if i >= LENGTH || j >= BREADTH || i < 0 || j < 0 || Unexplored[i][j] != 10 {
			continue
		}
		// Update Unexplore -- Explore
		Unexplored[i][j] = Grid[i][j]
		// If its an empty cell, explore further in all directions (including diagonals)
		if Grid[i][j] == 0 {
			ExploreQ.Enqueue(Coord{X: i + 1, Y: j})
			ExploreQ.Enqueue(Coord{X: i, Y: j + 1})
			ExploreQ.Enqueue(Coord{X: i - 1, Y: j})
			ExploreQ.Enqueue(Coord{X: i, Y: j - 1})
			ExploreQ.Enqueue(Coord{X: i + 1, Y: j + 1})
			ExploreQ.Enqueue(Coord{X: i - 1, Y: j + 1})
			ExploreQ.Enqueue(Coord{X: i + 1, Y: j - 1})
			ExploreQ.Enqueue(Coord{X: i - 1, Y: j - 1})
		}
	}
}
