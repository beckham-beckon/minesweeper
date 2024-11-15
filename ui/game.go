package ui

import (
	"example.com/minesweeper/common"
	"example.com/minesweeper/game"
	"github.com/gdamore/tcell/v2"
)

func (u *UIManager) RenderGame() {
	u.Screen.Clear()

	u.DrawGrid()

	if game.Init {
		game.InitUnexplored()
	}

	switch u.ScreenType {
	case common.GAME:
		r := SMILEYRUNE
		u.Screen.SetContent(u.ScreenWidth/2, u.YOffset-1, r, nil, GridStyle)
		u.PopulateGrid(game.Unexplored)
	case common.GAMEOVER:
        u.RenderGameOver()
	}
}

func (u *UIManager) DrawGrid() {
	x1, y1 := u.XOffset, u.YOffset
	x2, y2 := u.XFinish, u.YFinish

	for col := x1; col < x2; col = col + 4 {
		for row := y1; row <= y2; row++ {
			u.Screen.SetContent(col, row, tcell.RuneVLine, nil, GridStyle)
		}
	}

	for row := y1; row <= y2; row = row + 2 {
		for col := x1; col <= x2; col++ {
			u.Screen.SetContent(col, row, tcell.RuneHLine, nil, GridStyle)
		}
	}

	for col := x1; col < x2; col++ {
		u.Screen.SetContent(col, y1, tcell.RuneHLine, nil, GridStyle)
		u.Screen.SetContent(col, y2, tcell.RuneHLine, nil, GridStyle)
	}

	for col := x1; col < x2; col = col + 4 {
		u.Screen.SetContent(col, y1, tcell.RuneTTee, nil, GridStyle)
		u.Screen.SetContent(col, y2, tcell.RuneBTee, nil, GridStyle)
	}

	for row := y1 + 1; row < y2; row++ {
		u.Screen.SetContent(x1, row, tcell.RuneVLine, nil, GridStyle)
		u.Screen.SetContent(x2, row, tcell.RuneVLine, nil, GridStyle)
		if (row+u.YOffset)%2 == 0 {
			u.Screen.SetContent(x1, row, tcell.RuneLTee, nil, GridStyle)
			u.Screen.SetContent(x2, row, tcell.RuneRTee, nil, GridStyle)
		}
	}

	for row := y1 + 2; row <= y2-2; row = row + 2 {
		for col := x1 + 4; col <= x2-2; col = col + 4 {
			u.Screen.SetContent(col, row, tcell.RunePlus, nil, GridStyle)
		}
	}

	u.Screen.SetContent(x1, y1, tcell.RuneULCorner, nil, GridStyle)
	u.Screen.SetContent(x2, y1, tcell.RuneURCorner, nil, GridStyle)
	u.Screen.SetContent(x1, y2, tcell.RuneLLCorner, nil, GridStyle)
	u.Screen.SetContent(x2, y2, tcell.RuneLRCorner, nil, GridStyle)
}

func (u *UIManager) PopulateGrid(grid [][]int) {
	/*
	   Coordinate (XOffset, YOffest) starts with the grid lines
	   Populate numbers from the next coordinate; for
	   x -> XOffset + 2
	   y -> YOffest + 1
	*/
	x1, y1 := u.XOffset+2, u.YOffset+1
	x2, y2 := u.XFinish+2, u.YFinish+1
	i, j := 0, 0
	for row := y1; row < y2; row = row + 2 {
		i = 0
		for col := x1; col < x2; col = col + 4 {
			r := ' '
			style := tcell.StyleDefault
			if grid[i][j] < 0 {
				r = MINERUNE
				style = MineStyle
			} else if grid[i][j] > 0 {
				r = rune('0' + grid[i][j])
				style = NumberStyle
				if grid[i][j] == 10 {
					r = EMPTYBOXRUNE
					style = GridStyle
				}
			}
			u.Screen.SetContent(col, row, r, nil, style)
			i++
		}
		j++
	}
}
