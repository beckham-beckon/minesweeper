package ui

import "log"

const (
	TITLE  = "M I N E S W E E P E R"
	EASY   = "E A S Y"
	MEDIUM = "M E D I U M"
	HARD   = "H A R D"
	QUIT   = "Q U I T"
)

var MENU_ITEMS = []string{EASY, MEDIUM, HARD, QUIT}

func (u *UIManager) renderCenter(s string, y int) {
	x := (u.ScreenWidth - len(s)) / 2
	for _, r := range s {
		u.Screen.SetContent(x, y, rune(r), nil, TitleStyle)
		x++
	}
}

func (u *UIManager) DrawMenu() {
	u.Screen.Clear()
	log.Printf("in draw menu: %v, %v", u.ScreenWidth, u.ScreenHeight)
	y_title := (u.ScreenHeight-len(MENU_ITEMS))/2 - 2
	u.renderCenter(TITLE, y_title)
	y_menu_items := y_title + 2
	for _, item := range MENU_ITEMS {
		if u.ScreenWidth%2 == 0 {
			item = item + " "
		} else {
			item = " " + item
		}
		u.renderCenter(item, y_menu_items)
		y_menu_items++
	}
}
