package main

import (
	"log"
	"os"

	"github.com/gdamore/tcell/v2"
)

const (
	LENGTH       = 9
	BREADTH      = 9
	MINES        = 10
	FLAGRUNE     = '\u2691'
	EMPTYBOXRUNE = '\u2610'
	MINERUNE     = '\u2739'
	SMILEYRUNE   = '\u263A'
	FROWNRUNE    = '\u2639'
)

var (
	mineStyle      = tcell.StyleDefault.Foreground(tcell.ColorRed)
	numberStyle    = tcell.StyleDefault.Foreground(tcell.ColorYellow)
	unExplored     = make([][]int, LENGTH)
	grid           = make([][]int, LENGTH)
	exploreQ       = &CoordQ{}
	X_OFFSET       = 10
	Y_OFFSET       = 5
	RENDER_LENGTH  = 4*LENGTH + X_OFFSET
	RENDER_BREADTH = 2*BREADTH + Y_OFFSET
	GAME_OVER      = false
	SCREEN_WIDTH   = 0
	SCREEN_HEIGHT  = 0
)

func handleResize(s tcell.Screen) {
	s.Clear()
	SCREEN_WIDTH, SCREEN_HEIGHT = s.Size()
	X_OFFSET = (SCREEN_WIDTH / 2) - 2*LENGTH
	Y_OFFSET = (SCREEN_HEIGHT / 2) - BREADTH
	RENDER_LENGTH = 4*LENGTH + X_OFFSET
	RENDER_BREADTH = 2*BREADTH + Y_OFFSET
}

func main() {
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

	generateGrids()
	handleResize(s)

	for {
		s.Show()

		ev := s.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			handleResize(s)
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
						GAME_OVER = true
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
