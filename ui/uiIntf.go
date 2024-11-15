package ui

import (
	"os"

	"example.com/minesweeper/common"
	"example.com/minesweeper/game"
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
	GridStyle   = tcell.StyleDefault
)

type UIManager struct {
	Screen       tcell.Screen
	ScreenHeight int
	ScreenWidth  int
	XOffset      int
	YOffset      int
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
	ui.RenderMenu()
}

func (ui *UIManager) HandleResizeGrid() {
	ui.XOffset = (ui.ScreenWidth / 2) - 2*common.Length
	ui.YOffset = (ui.ScreenHeight / 2) - common.Breadth

	ui.XFinish = 4*common.Length + ui.XOffset
	ui.YFinish = 2*common.Breadth + ui.YOffset

	ui.RenderGame()
}

func (ui *UIManager) HandeResizeGameOver() {
	ui.HandleResizeGrid()
	ui.RenderGameOver()
}

func (ui *UIManager) HandleKeyEvent(ev *tcell.EventKey) {
	if ev.Rune() == 'q' || ev.Rune() == 'Q' {
		ui.Quit()
	}
	switch ui.ScreenType {
	case common.MENU:
		ui.HandleMenuKeyEvent(ev)
	}
}

func (ui *UIManager) HandleMenuKeyEvent(ev *tcell.EventKey) {
	switch ev.Key() {
	case tcell.KeyUp:
		ui.MenuRenderSelector(-1)
	case tcell.KeyDown:
		ui.MenuRenderSelector(1)
	case tcell.KeyEnter:
		ui.MenuProcessSelect()
	}
}

func (ui *UIManager) HandleMouseEvent(ev *tcell.EventMouse) {
	x, y := ev.Position()
	switch ev.Buttons() {
	case tcell.Button1:
		c, _, _, _ := ui.Screen.GetContent(x, y)
		if x < ui.XFinish && y < ui.YFinish && (c == EMPTYBOXRUNE || c == FLAGRUNE) {
			i := (x - ui.XOffset) / 4
			j := (y - ui.YOffset) / 2
			if game.Init {
				game.InitGrid(i, j)
				game.Init = false
			}
			if game.Grid[i][j] < 0 {
				ui.ScreenType = common.GAMEOVER
				ui.HandleResize()
				break
			}
			if game.Grid[i][j] > 0 {
				game.Unexplored[i][j] = game.Grid[i][j]
				game.CellsExplored++
				ui.Screen.SetContent(x, y, rune('0'+game.Grid[i][j]), nil, NumberStyle)
				break
			}
			game.Explore(i, j)
			if !game.CheckComplete() {
				ui.PopulateGrid(game.Unexplored)
				break
			}
			ui.ScreenType = common.GAMEOVER
			ui.HandleResize()
		}
	case tcell.Button2:
		c, _, _, _ := ui.Screen.GetContent(x, y)
		if x < ui.XFinish && y < ui.YFinish && (c == EMPTYBOXRUNE || c == FLAGRUNE) {
			ui.Screen.SetContent(x, y, FLAGRUNE, nil, MineStyle)
		}
	}
}
