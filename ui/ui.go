package ui

import (
	"os"

	"example.com/minesweeper/common"
	"github.com/gdamore/tcell/v2"
)

const (
	FLAGRUNE     = '\u2691'
	EMPTYBOXRUNE = '\u2610'
	MINERUNE     = '\u2739'
	SMILEYRUNE   = '\u263A'
	FROWNRUNE    = '\u2639'
)

var (
	MineStyle   = tcell.StyleDefault.Foreground(tcell.ColorRed)
	NumberStyle = tcell.StyleDefault.Foreground(tcell.ColorYellow)
	TitleStyle  = tcell.StyleDefault.Foreground(tcell.ColorPurple)
)

type UIManager struct {
	Screen       tcell.Screen
	ScreenHeight int
	ScreenWidth  int
	XOffset      int
	YOffest      int
	XFinish      int
	YFinish      int
	ScreenType   string
}

func NewUIManager() (*UIManager, error) {
	var uiManager UIManager
	newScreen, err := tcell.NewScreen()
	if err != nil {
		return nil, err
	}
	if err := newScreen.Init(); err != nil {
		return nil, err
	}

	newScreen.Clear()

	uiManager.Screen = newScreen

	uiManager.ScreenWidth, uiManager.ScreenHeight = newScreen.Size()

	uiManager.ScreenType = common.MENU

	return &uiManager, nil
}

func (ui *UIManager) Quit() {
	ui.Screen.Fini()
	os.Exit(0)
}

func (ui *UIManager) RenderScreen() {
	switch ui.ScreenType {
	case common.MENU:
		ui.HandleResize()
	}
}

func (ui *UIManager) HandleResize() {
	// Update Screen height and width
	ui.ScreenWidth, ui.ScreenHeight = ui.Screen.Size()
	switch ui.ScreenType {
	case common.MENU:
		ui.HandleResizeMenu()
	case common.GAME:
		ui.HandleResizeGrid()
	case common.GAMEOVER:
		ui.HandeResizeGameOver()
	}
}

func (ui *UIManager) HandleResizeMenu() {
	ui.DrawMenu()
}

func (ui *UIManager) HandeResizeGameOver() {
}

func (ui *UIManager) HandleResizeGrid() {
	ui.Screen.Clear()

	ui.XOffset = (ui.ScreenWidth / 2) - 2*ui.XFinish
	ui.YOffest = (ui.ScreenHeight / 2) - ui.YFinish

	ui.DrawGrid()
}

func (ui *UIManager) DrawGrid() {
}

func (ui *UIManager) PopulateGrid(grid [][]int) {
	/*
	   Coordinate (XOffset, YOffest) starts with the grid lines
	   Populate numbers from the next coordinate for
	   x -> XOffset + 2
	   y -> YOffest + 1
	*/
	x1, y1 := ui.XOffset+2, ui.YOffest+1
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
					style = tcell.StyleDefault
				}
			}
			ui.Screen.SetContent(col, row, r, nil, style)
			i++
		}
		j++
	}
}
