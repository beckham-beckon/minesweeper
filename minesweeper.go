package main

import (
	"log"
	"os"

	"example.com/minesweeper/ui"
	"github.com/gdamore/tcell/v2"
)

func main() {
	file, err := os.OpenFile("l.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	defer file.Close()

	// Set log output to the file
	log.SetOutput(file)

	UI, err := ui.NewUIManager()
	if err != nil {
		log.Fatalf("%v", err)
	}

	log.Printf("%v", UI)

	UI.Screen.Clear()
	UI.DrawMenu()

	quit := func() {
		UI.Screen.Fini()
		os.Exit(0)
	}

	for {
		UI.Screen.Show()

		ev := UI.Screen.PollEvent()

		switch ev := ev.(type) {
		case *tcell.EventResize:
			UI.HandleResize()
		case *tcell.EventKey:
			if ev.Rune() == 'Q' || ev.Rune() == 'q' {
				quit()
			}
		}
	}
}
