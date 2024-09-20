package ui

import "github.com/gdamore/tcell/v2"

type UIManager struct {
	Screen       tcell.Screen
	ScreenHeight int
	ScreenWidth  int
	XOffset      int
	YOffest      int
	GridLength   int
	GridBreadth  int
}

func NewUIManager(gameMode string) (*UIManager, error) {
  var UIManager UIManager
  s, err := tcell.NewScreen()
  if err != nil {
    return nil, err
  }
  if err := s.Init(); err != nil {
    return nil, err
  }

  UIManager.ScreenHeight, UIManager.ScreenWidth = s.Size()
  switch gameMode {
  default: 
    UIManager.GridLength = 9
    UIManager.GridBreadth = 9
  case "MEDIUM":
    UIManager.GridLength = 16
    UIManager.GridBreadth = 16
  }

  UIManager.XOffset = (UIManager.ScreenWidth / 2) - 2*UIManager.GridLength
  UIManager.YOffest = (UIManager.ScreenHeight / 2) - UIManager.GridBreadth
  
  return &UIManager, nil
}

func(ui *UIManager) Quit() {
  ui.Screen.Fini()
}

func(ui *UIManager) HandleResize() {
  ui.Screen.Clear()
  ui.ScreenWidth, ui.ScreenHeight = ui.Screen.Size()

  ui.XOffset = (ui.ScreenWidth / 2) - 2*ui.GridLength
  ui.YOffest = (ui.ScreenHeight / 2) - ui.GridBreadth

}
