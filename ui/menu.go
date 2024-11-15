package ui

import (
	c "example.com/minesweeper/common"
)

const (
	TITLE  = "M I N E S W E E P E R"
	EASY   = "E A S Y"
	MEDIUM = "M E D I U M"
	HARD   = "H A R D"
	CUSTOM = "C U S T O M"
	QUIT   = "Q U I T"
)

var (
	MenuItems         = []string{EASY, MEDIUM, HARD, CUSTOM, QUIT}
	MenuItemsCoords   = make(map[string]c.Coord)
	SelectorPos       = 0
	LastSelectorCoord = c.Coord{X: -1, Y: -1}
)

func (u *UIManager) RenderCenter(s string, x int, y int) {
	for _, r := range s {
		u.Screen.SetContent(x, y, rune(r), nil, TitleStyle)
		x++
	}
}

func (u *UIManager) RenderMenu() {
	u.Screen.Clear()
	y_title := (u.ScreenHeight-len(MenuItems))/2 - 2
	x_title := (u.ScreenWidth - len(TITLE)) / 2
	u.RenderCenter(TITLE, x_title, y_title)

	y_menu_item := y_title + 2
	for _, item := range MenuItems {
		x_menu_item := (u.ScreenWidth - len(item)) / 2
		menuItemCoord := c.Coord{X: x_menu_item, Y: y_menu_item}
		MenuItemsCoords[item] = menuItemCoord

		if u.ScreenWidth%2 == 0 {
			item = item + " "
		} else {
			item = " " + item
		}

		u.RenderCenter(item, x_menu_item, y_menu_item)
		y_menu_item++
	}
	u.MenuRenderSelector(0)
}

func (u *UIManager) MenuRenderSelector(move int) {
	if SelectorPos+move >= len(MenuItems) || SelectorPos+move < 0 {
		return
	}

	// Clear Last Selector
	if LastSelectorCoord.X >= 0 && LastSelectorCoord.Y >= 0 {
		u.Screen.SetContent(LastSelectorCoord.X, LastSelectorCoord.Y, ' ', nil, NumberStyle)
	}

	SelectorPos = SelectorPos + move
	item := MenuItems[SelectorPos]
	coords := MenuItemsCoords[item]

	// Render selector at (x-2, y)
	u.Screen.SetContent(coords.X-2, coords.Y, MINERUNE, nil, MineStyle)

	LastSelectorCoord.X, LastSelectorCoord.Y = coords.X-2, coords.Y
}

func (u *UIManager) MenuProcessSelect() {
	for item, coord := range MenuItemsCoords {
		if LastSelectorCoord.Y == coord.Y {
			// Setting Length and Breadth based on Game Mode
			switch item {
			case EASY:
				c.Length = 9
				c.Breadth = 9
				c.Mines = 10
			case MEDIUM:
				c.Length = 16
				c.Breadth = 16
				c.Mines = 40
			case HARD:
				c.Length = 30
				c.Breadth = 16
				c.Mines = 100
			case QUIT:
				u.Quit()
				return
			}

			u.ScreenType = c.GAME
			u.HandleResize()
		}
	}
}
