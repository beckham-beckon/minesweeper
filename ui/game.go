package ui

import (
	"example.com/minesweeper/common"
	"example.com/minesweeper/game"
	"github.com/gdamore/tcell/v2"
)

func (ui *UIManager) RenderGame() {
	r := SMILEYRUNE
	ui.Screen.SetContent(ui.ScreenWidth/2, ui.YOffset-1, r, nil, GridStyle)

	ui.DrawGrid()

	if game.Init {
		game.InitUnexplored()
	}
	if ui.ScreenType == common.GAME {
		ui.PopulateGrid(game.Unexplored)
	} else if ui.ScreenType == common.GAMEOVER {
		ui.PopulateGrid(game.Grid)
	}
}

func (ui *UIManager) DrawGrid() {
	x1, y1 := ui.XOffset, ui.YOffset
	x2, y2 := ui.XFinish, ui.YFinish

	for col := x1; col < x2; col = col + 4 {
		for row := y1; row <= y2; row++ {
			ui.Screen.SetContent(col, row, tcell.RuneVLine, nil, GridStyle)
		}
	}

	for row := y1; row <= y2; row = row + 2 {
		for col := x1; col <= x2; col++ {
			ui.Screen.SetContent(col, row, tcell.RuneHLine, nil, GridStyle)
		}
	}

	for col := x1; col < x2; col++ {
		ui.Screen.SetContent(col, y1, tcell.RuneHLine, nil, GridStyle)
		ui.Screen.SetContent(col, y2, tcell.RuneHLine, nil, GridStyle)
	}

	for col := x1; col < x2; col = col + 4 {
		ui.Screen.SetContent(col, y1, tcell.RuneTTee, nil, GridStyle)
		ui.Screen.SetContent(col, y2, tcell.RuneBTee, nil, GridStyle)
	}

	for row := y1 + 1; row < y2; row++ {
		ui.Screen.SetContent(x1, row, tcell.RuneVLine, nil, GridStyle)
		ui.Screen.SetContent(x2, row, tcell.RuneVLine, nil, GridStyle)
		if (row+ui.YOffset)%2 == 0 {
			ui.Screen.SetContent(x1, row, tcell.RuneLTee, nil, GridStyle)
			ui.Screen.SetContent(x2, row, tcell.RuneRTee, nil, GridStyle)
		}
	}

	for row := y1 + 2; row <= y2-2; row = row + 2 {
		for col := x1 + 4; col <= x2-2; col = col + 4 {
			ui.Screen.SetContent(col, row, tcell.RunePlus, nil, GridStyle)
		}
	}

	ui.Screen.SetContent(x1, y1, tcell.RuneULCorner, nil, GridStyle)
	ui.Screen.SetContent(x2, y1, tcell.RuneURCorner, nil, GridStyle)
	ui.Screen.SetContent(x1, y2, tcell.RuneLLCorner, nil, GridStyle)
	ui.Screen.SetContent(x2, y2, tcell.RuneLRCorner, nil, GridStyle)
}

func (ui *UIManager) PopulateGrid(grid [][]int) {
	/*
	   Coordinate (XOffset, YOffest) starts with the grid lines
	   Populate numbers from the next coordinate; for
	   x -> XOffset + 2
	   y -> YOffest + 1
	*/
	x1, y1 := ui.XOffset+2, ui.YOffset+1
	x2, y2 := ui.XFinish+2, ui.YFinish+1
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
			ui.Screen.SetContent(col, row, r, nil, style)
			i++
		}
		j++
	}
}
