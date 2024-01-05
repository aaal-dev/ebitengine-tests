package objects

import (
	"asteroids/assets"
	"asteroids/game/core"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
)

func NewButtonWithText(
	text string,
	position core.Vector[float64],
	size core.Vector[float64],
	fgColor color.Color,
	bgColor color.Color,
) (*Button, error) {
	bg := ebiten.NewImage(int(size.X), int(size.Y))
	bg.Fill(bgColor)

	label, err := NewText(
		text,
		position,
		assets.NormalFont,
		fgColor,
		1.0,
	)
	if err != nil {
		return nil, err
	}

	label.BackgroundColor = color.RGBA{R: 100, G: 200, B: 100, A: 160}

	return &Button{
		position:   position,
		size:       size,
		background: bg,
		text:       label,
		fgColor:    fgColor,
		bgColor:    bgColor,
	}, nil
}

type Button struct {
	position   core.Vector[float64]
	size       core.Vector[float64]
	background *ebiten.Image
	text       Text
	fgColor    color.Color
	bgColor    color.Color
}

func (button *Button) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(button.position.X, button.position.Y)

	screen.DrawImage(button.background, options)

	button.text.Draw(screen)
}

func (button *Button) Update() error {
	return nil
}
