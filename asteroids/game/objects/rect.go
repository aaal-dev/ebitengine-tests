package objects

import (
	"asteroids/game/core"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func NewRect(
	position core.Vector[float32],
	size core.Size[float32],
	color color.Color,
) (*Rect, error) {
	return &Rect{
		X:      position.X,
		Y:      position.Y,
		Width:  size.Width,
		Height: size.Height,
		Color:  color,
	}, nil
}

type Rect struct {
	X      float32
	Y      float32
	Width  float32
	Height float32
	Color  color.Color
}

func (rect *Rect) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(
		screen,
		rect.X,
		rect.Y,
		rect.Width,
		rect.Height,
		rect.Color,
		false,
	)
}

func (rect *Rect) Update() error {
	return nil
}
