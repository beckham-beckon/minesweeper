package main

import (
	"example.com/minesweeper/ui"
	"github.com/gdamore/tcell/v2"
)

func main() {
	UI, _ := ui.NewUIManager()
  UI.HandleResize()

	for {
		ev := UI.Screen.PollEvent()

    switch ev := ev.(type) {
		case *tcell.EventResize:
			UI.Screen.Sync()
		case *tcell.EventKey:
			if ev.Key() == tcell.KeyEscape || ev.Key() == tcell.KeyCtrlC {
				UI.Quit()
			}
		}
	}
}
