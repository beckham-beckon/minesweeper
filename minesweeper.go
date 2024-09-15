package main

import (
	"log"
	"math/rand"
	"os"

	"github.com/gdamore/tcell/v2"
)

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

const (
	LENGTH         = 9
	BREADTH        = 9
	MINES          = 10
	FLAGRUNE       = '\u2691'
	EMPTYBOXRUNE   = '\u2610'
	MINERUNE       = '\u2739'
	X_OFFSET       = 10
	Y_OFFSET       = 5
	RENDER_LENGTH  = 4*LENGTH + X_OFFSET
	RENDER_BREADTH = 2*BREADTH + Y_OFFSET
)

var (
	mineStyle   = tcell.StyleDefault.Foreground(tcell.ColorRed)
	numberStyle = tcell.StyleDefault.Foreground(tcell.ColorYellow)
	unExplored  = make([][]int, LENGTH)
	grid        = make([][]int, LENGTH)
	exploreQ    = &CoordQ{}
)

func generateGrids() {
	for i := 0; i < LENGTH; i++ {
		grid[i] = make([]int, BREADTH)
		for j := 0; j < BREADTH; j++ {
			grid[i][j] = 0
		}
	}

	generateCoords := func() (int, int) {
		x := rand.Intn(LENGTH)
		y := rand.Intn(BREADTH)
		return x, y
	}

	for range MINES {
		var X, Y int
		for {
			x, y := generateCoords()
			if grid[x][y] >= 0 {
				grid[x][y] = -9 // Max mines nearby can be 8
				X = x
				Y = y
				break
			}
		}

		for i := -1; i < 2; i++ {
			for j := -1; j < 2; j++ {
				if i == 0 && j == 0 {
					continue
				}
				new_x := X + i
				new_y := Y + j

				if 0 <= new_x && new_x < LENGTH && 0 <= new_y && new_y < BREADTH {
					grid[new_x][new_y]++
				}
			}
		}
	}

	for i := 0; i < LENGTH; i++ {
		unExplored[i] = make([]int, BREADTH)
		for j := 0; j < BREADTH; j++ {
			unExplored[i][j] = 10
		}
	}
}

func drawGrid(s tcell.Screen) {
	style := tcell.StyleDefault
	x1, y1 := X_OFFSET, Y_OFFSET
	x2, y2 := RENDER_LENGTH, RENDER_BREADTH

	for col := x1; col < x2; col = col + 4 {
		for row := y1; row <= y2; row++ {
			s.SetContent(col, row, tcell.RuneVLine, nil, style)
		}
	}

	for row := y1; row <= y2; row = row + 2 {
		for col := x1; col <= x2; col++ {
			s.SetContent(col, row, tcell.RuneHLine, nil, style)
		}
	}

	for col := x1; col < x2; col++ {
		s.SetContent(col, y1, tcell.RuneHLine, nil, style)
		s.SetContent(col, y2, tcell.RuneHLine, nil, style)
		if (col+X_OFFSET)%4 == 0 {
			s.SetContent(col, y1, tcell.RuneTTee, nil, style)
			s.SetContent(col, y2, tcell.RuneBTee, nil, style)
		}
	}
	for row := y1 + 1; row < y2; row++ {
		s.SetContent(x1, row, tcell.RuneVLine, nil, style)
		s.SetContent(x2, row, tcell.RuneVLine, nil, style)
		if (row+Y_OFFSET)%2 == 0 {
			s.SetContent(x1, row, tcell.RuneLTee, nil, style)
			s.SetContent(x2, row, tcell.RuneRTee, nil, style)
		}
	}

	for row := y1 + 2; row <= y2-2; row = row + 2 {
		for col := x1 + 4; col <= x2-2; col = col + 4 {
			s.SetContent(col, row, tcell.RunePlus, nil, style)
		}
	}

	s.SetContent(x1, y1, tcell.RuneULCorner, nil, style)
	s.SetContent(x2, y1, tcell.RuneURCorner, nil, style)
	s.SetContent(x1, y2, tcell.RuneLLCorner, nil, style)
	s.SetContent(x2, y2, tcell.RuneLRCorner, nil, style)
}

func renderGrid(s tcell.Screen, grid [][]int) {
	x1, y1 := X_OFFSET+2, Y_OFFSET+1
	x2, y2 := RENDER_LENGTH+2, RENDER_BREADTH+1
	i, j := 0, 0
	for row := y1; row < y2; row = row + 2 {
		i = 0
		for col := x1; col < x2; col = col + 4 {
			r := ' '
			style := tcell.StyleDefault
			if grid[i][j] < 0 {
				r = MINERUNE
				style = mineStyle
			} else if grid[i][j] > 0 {
				r = rune('0' + grid[i][j])
				style = numberStyle
				if grid[i][j] == 10 {
					r = EMPTYBOXRUNE
					style = tcell.StyleDefault
				}
			}
			s.SetContent(col, row, r, nil, style)
			i++
		}
		j++
	}
}

func explore() {
	for len(exploreQ.Coords) > 0 {
		c := exploreQ.Dequeue()
		i, j := c.X, c.Y
		if i >= LENGTH || j >= BREADTH || i < 0 || j < 0 || unExplored[i][j] != 10 {
			continue
		}
		unExplored[i][j] = grid[i][j]
		if grid[i][j] == 0 {
			exploreQ.Enqueue(Coord{X: i + 1, Y: j})
			exploreQ.Enqueue(Coord{X: i, Y: j + 1})
			exploreQ.Enqueue(Coord{X: i - 1, Y: j})
			exploreQ.Enqueue(Coord{X: i, Y: j - 1})
			exploreQ.Enqueue(Coord{X: i + 1, Y: j + 1})
			exploreQ.Enqueue(Coord{X: i - 1, Y: j + 1})
			exploreQ.Enqueue(Coord{X: i + 1, Y: j - 1})
			exploreQ.Enqueue(Coord{X: i - 1, Y: j - 1})
		}
	}
}

func main() {
	generateGrids()

	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("Error Creating new screen: %v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("Error initiating new Screen: %v", err)
	}
	s.EnableMouse()
	quit := func() {
		s.Fini()
		os.Exit(0)
	}

	drawGrid(s)
	renderGrid(s, unExplored)

	for {
		s.Show()

		ev := s.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape {
				quit()
			}
		case *tcell.EventMouse:
			x, y := ev.Position()
			switch ev.Buttons() {
			case tcell.Button1:
				c, _, _, _ := s.GetContent(x, y)
				if x < RENDER_LENGTH && y < RENDER_BREADTH && (c == EMPTYBOXRUNE || c == FLAGRUNE) {
					i := (x - X_OFFSET) / 4
					j := (y - Y_OFFSET) / 2
					if grid[i][j] < 0 {
						renderGrid(s, grid)
						break
					}
					if grid[i][j] > 0 {
						unExplored[i][j] = grid[i][j]
						s.SetContent(x, y, rune('0'+grid[i][j]), nil, numberStyle)
						break
					}
					exploreQ.Enqueue(Coord{X: i, Y: j})
					explore()
					renderGrid(s, unExplored)
				}
			case tcell.Button2:
				c, _, _, _ := s.GetContent(x, y)
				if x < RENDER_LENGTH && y < RENDER_BREADTH && (c == EMPTYBOXRUNE || c == FLAGRUNE) {
					s.SetContent(x, y, FLAGRUNE, nil, mineStyle)
				}
			}
		}

	}
}
