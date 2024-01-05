package game

import (
	"asteroids/game/core"
	"asteroids/game/scenes"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	scene scenes.Scene
}

func NewGame() (*Game, error) {
	scene, err := scenes.CreateMainMenuScene()
	if err != nil {
		return nil, err
	}

	return &Game{
		scene: scene,
	}, nil
}

func (game *Game) Close() error {
	return nil
}

func (game *Game) Update() error {
	err := game.scene.Update()

	return err
}

func (game *Game) Draw(screen *ebiten.Image) {
	game.scene.Draw(screen)
}

func (game *Game) Layout(outsideWidth int, outsideHeight int) (int, int) {
	return core.ScreenWidth, core.ScreenHeight
}
