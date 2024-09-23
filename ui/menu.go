package ui

const TITLE = "M I N E S W E E P E R"

var MENU_ITEMS = map[string]string{
	"DIFFICULTY1": "EASY",
	"DIFFICULTY2": "MEDIUM",
	"DIFFICULTY3": "HARD",
	"QUIT":        "QUIT",
}

func (ui *UIManager) renderCenter(s string, y int) {
	x := (ui.ScreenWidth - len(s)) / 2
	for _, r := range []rune(s) {
		ui.Screen.SetContent(x, y, r, nil, titleStyle)
		x++
	}
}

func (ui *UIManager) DrawMenu() {
  ui.Screen.Clear()
	y_title := (ui.ScreenHeight-len(MENU_ITEMS))/2 - 2
	ui.renderCenter(TITLE, y_title)
	y_menu_items := y_title + 2
	for _, item := range MENU_ITEMS {
		ui.renderCenter(item, y_menu_items)
		y_menu_items++
	}
}
