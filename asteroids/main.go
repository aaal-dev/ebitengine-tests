package main

import (
	"asteroids/game"
	"asteroids/game/core"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	os.Setenv("EBITEN_GRAPHICS_LIBRARY", "opengl")

	newGame, err := game.NewGame()
	if err != nil {
		panic(err)
	}
	defer func() {
		err = newGame.Close()
		if err != nil {
			panic(err)
		}
	}()

	ebiten.SetWindowSize(core.ScreenWidth, core.ScreenHeight)
	ebiten.SetWindowTitle(core.Title)

	err = ebiten.RunGame(newGame)
	if err != nil {
		panic(err)
	}
}
