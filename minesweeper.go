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

var (
	mineStyle   = tcell.StyleDefault.Foreground(tcell.ColorRed)
	numberStyle = tcell.StyleDefault.Foreground(tcell.ColorYellow)
	unExplored  = make([][]int, LENGTH)
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

func drawGrid(s tcell.Screen) {
	style := tcell.StyleDefault
	x1, y1 := 0, 0
	x2, y2 := 4*LENGTH, 2*BREADTH

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
		if col%4 == 0 {
			s.SetContent(col, y1, tcell.RuneTTee, nil, style)
			s.SetContent(col, y2, tcell.RuneBTee, nil, style)
		}
	}
	for row := y1 + 1; row < y2; row++ {
		s.SetContent(x1, row, tcell.RuneVLine, nil, style)
		s.SetContent(x2, row, tcell.RuneVLine, nil, style)
		if row%2 == 0 {
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
	x1, y1 := 2, 1
	x2, y2 := 4*LENGTH, 2*BREADTH
	i, j := 0, 0
	for row := y1; row < y2; row = row + 2 {
		i = 0
		for col := x1; col < x2; col = col + 4 {
			r := ' '
			style := tcell.StyleDefault
			if grid[i][j] < 0 {
				r = '*'
				style = mineStyle
			} else if grid[i][j] > 0 {
				r = rune('0' + grid[i][j])
				style = numberStyle
				if grid[i][j] == 10 {
					r = '\u2610'
					style = tcell.StyleDefault
				}
			}
			s.SetContent(col, row, r, nil, style)
			i++
		}
		j++
	}
}

func explore(grid [][]int, i int, j int, isNumber bool) {
	if i >= LENGTH || j >= BREADTH || i < 0 || j < 0 {
		return
	}
	if isNumber {
		if grid[i][j] != 0 {
			return
		}
		unExplored[i][j] = 0
		isNumber = false
	}
	if grid[i][j] >= 0 {
		unExplored[i][j] = grid[i][j]
		isNumber = true
		if grid[i][j] == 0 {
			isNumber = false
		}
	}
	explore(grid, i+1, j, isNumber)
	explore(grid, i, j+1, isNumber)
	explore(grid, i-1, j, isNumber)
	explore(grid, i, j-1, isNumber)
}

func main() {
	file, err := os.OpenFile("log.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	logger := log.New(file, "", log.LstdFlags)

	grid := generateGrid()

	for i := range LENGTH {
		unExplored[i] = make([]int, BREADTH)
		for j := range BREADTH {
			unExplored[i][j] = 10
		}
	}

	s, err := tcell.NewScreen()
	if err != nil {
		logger.Fatalf("Error Creating new screen: %v", err)
	}
	if err := s.Init(); err != nil {
		logger.Fatalf("Error initiating new Screen: %v", err)
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
				if x < 4*LENGTH && y < 2*BREADTH && c == '\u2610' {
					logger.Println(x, y, c)
					i := x / 4
					j := y / 2
					logger.Println(i, j, grid[i][j])
					if grid[i][j] < 0 {
						renderGrid(s, grid)
						break
					}
					if grid[i][j] > 0 {
						unExplored[i][j] = grid[i][j]
						// s.SetContent(x, y, rune('0'+grid[i][j]), nil, numberStyle)
						renderGrid(s, unExplored)
						break
					}
					explore(grid, i, j, false)
					renderGrid(s, unExplored)
				}
			}
		}

	}
}
