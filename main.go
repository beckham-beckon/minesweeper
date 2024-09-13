package main

import (
	"log"
	"math/rand"
	"os"

	"github.com/gdamore/tcell/v2"
)

const (
	LENGTH  = 9
	BREADTH = 9
	MINES   = 10
)

func generateGrid() [][]int {
	grid := make([][]int, LENGTH)
	for i := range LENGTH {
		grid[i] = make([]int, BREADTH)
		for j := range BREADTH {
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
	return grid
}

func renderGrid(grid [][]int) {

}

func main() {
	grid := generateGrid()

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
				s.SetContent(x, y, '*', nil, tcell.StyleDefault)
			}
		}

	}
}
