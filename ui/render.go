package ui

func DrawGrid(s tcell.Screen) {
	style := tcell.StyleDefault
	x1, y1 := X_OFFSET, Y_OFFSET
	x2, y2 := RENDER_LENGTH, RENDER_BREADTH

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
		if (col+X_OFFSET)%4 == 0 {
			s.SetContent(col, y1, tcell.RuneTTee, nil, style)
			s.SetContent(col, y2, tcell.RuneBTee, nil, style)
		}
	}
	for row := y1 + 1; row < y2; row++ {
		s.SetContent(x1, row, tcell.RuneVLine, nil, style)
		s.SetContent(x2, row, tcell.RuneVLine, nil, style)
		if (row+Y_OFFSET)%2 == 0 {
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
	x1, y1 := X_OFFSET+2, Y_OFFSET+1
	x2, y2 := RENDER_LENGTH+2, RENDER_BREADTH+1
	i, j := 0, 0
	for row := y1; row < y2; row = row + 2 {
		i = 0
		for col := x1; col < x2; col = col + 4 {
			r := ' '
			style := tcell.StyleDefault
			if grid[i][j] < 0 {
				r = MINERUNE
				style = mineStyle
			} else if grid[i][j] > 0 {
				r = rune('0' + grid[i][j])
				style = numberStyle
				if grid[i][j] == 10 {
					r = EMPTYBOXRUNE
					style = tcell.StyleDefault
				}
			}
			s.SetContent(col, row, r, nil, style)
			i++
		}
		j++
	}
	r := SMILEYRUNE
	style := numberStyle
	if GAME_OVER {
		r = FROWNRUNE
		style = mineStyle
	}
	s.SetContent(SCREEN_WIDTH/2, Y_OFFSET-2, r, nil, style)
}

func RenderGrid(s tcell.Screen, grid [][]int) {
	x1, y1 := X_OFFSET+2, Y_OFFSET+1
	x2, y2 := RENDER_LENGTH+2, RENDER_BREADTH+1
	i, j := 0, 0
	for row := y1; row < y2; row = row + 2 {
		i = 0
		for col := x1; col < x2; col = col + 4 {
			r := ' '
			style := tcell.StyleDefault
			if grid[i][j] < 0 {
				r = MINERUNE
				style = mineStyle
			} else if grid[i][j] > 0 {
				r = rune('0' + grid[i][j])
				style = numberStyle
				if grid[i][j] == 10 {
					r = EMPTYBOXRUNE
					style = tcell.StyleDefault
				}
			}
			s.SetContent(col, row, r, nil, style)
			i++
		}
		j++
	}
	r := SMILEYRUNE
	style := numberStyle
	if GAME_OVER {
		r = FROWNRUNE
		style = mineStyle
	}
	s.SetContent(SCREEN_WIDTH/2, Y_OFFSET-2, r, nil, style)
}

