package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"github.com/gdamore/tcell/v2"
)

const (
	LENGTH = 9
	HEIGHT = 9
)

type Block struct {
	IsMine bool
	Value  int64
}

func generateGrid() [][]Block {
	grid := make([][]Block, LENGTH)
	return grid
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
