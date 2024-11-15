package ui

import "example.com/minesweeper/game"

const (
	GAME_OVER = "GAME  OVER"
)

func (u *UIManager) RenderGameOver() {
	r := FROWNRUNE
	u.Screen.SetContent(u.ScreenWidth/2, u.YOffset-1, r, nil, MineStyle)
	u.PopulateGrid(game.Grid)
}
