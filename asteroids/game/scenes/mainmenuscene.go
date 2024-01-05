package scenes

import (
	"asteroids/game/core"
	"asteroids/game/objects"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func CreateMainMenuScene() (*MainMenuScene, error) {
	var sceneObjects []objects.Object

	background, error := objects.NewRect(
		core.Vector[float32]{X: 0.0, Y: 0.0},
		core.Size[float32]{
			Width:  float32(core.ScreenWidth),
			Height: float32(core.ScreenHeight)},
		color.Gray{Y: 255},
	)
	if error != nil {
		return nil, error
	}
	sceneObjects = append(sceneObjects, background)

	grid, error := objects.NewGrid(32.0, 32.0)
	if error != nil {
		return nil, error
	}
	sceneObjects = append(sceneObjects, grid)

	button, error := objects.NewButtonWithText(
		"Start\n==",
		core.Vector[float64]{X: 64.0, Y: 64.0},
		core.Vector[float64]{X: 192.0, Y: 32.0},
		color.Black,
		color.Gray{Y: 127},
	)
	if error != nil {
		return nil, error
	}
	sceneObjects = append(sceneObjects, button)

	return &MainMenuScene{
		grid:        grid,
		startButton: button,
		objects:     sceneObjects,
	}, nil
}

type MainMenuScene struct {
	grid        *objects.Grid
	startButton *objects.Button
	objects     []objects.Object
}

func (scene *MainMenuScene) Draw(screen *ebiten.Image) {
	for _, object := range scene.objects {
		object.Draw(screen)
	}
}

func (scene *MainMenuScene) Update() error {
	for _, object := range scene.objects {
		error := object.Update()
		if error != nil {
			return error
		}
	}
	return nil
}
