package objects

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func NewGrid(width float32, height float32) (*Grid, error) {
	return &Grid{
		width:   width,
		height:  height,
		color64: &color.RGBA{A: 50},
		color32: &color.RGBA{A: 20},
	}, nil
}

type Grid struct {
	width   float32
	height  float32
	color64 color.Color
	color32 color.Color
}

func (grid *Grid) Draw(screen *ebiten.Image) {
	width := screen.Bounds().Dx()
	height := screen.Bounds().Dy()

	for y := 0; y < height; y += 32 {
		vector.StrokeLine(
			screen,
			0.0,
			float32(y),
			float32(width),
			float32(y),
			1.0,
			grid.color32,
			false,
		)
	}

	for y := 0; y < height; y += 64 {
		vector.StrokeLine(
			screen,
			0.0,
			float32(y),
			float32(width),
			float32(y),
			1.0,
			grid.color64,
			false,
		)
	}

	for x := 0; x < width; x += 32 {
		vector.StrokeLine(
			screen,
			float32(x),
			0.0,
			float32(x),
			float32(height),
			1.0,
			grid.color32,
			false,
		)
	}

	for x := 0; x < width; x += 64 {
		vector.StrokeLine(
			screen,
			float32(x),
			0.0,
			float32(x),
			float32(height),
			1.0,
			grid.color64,
			false,
		)
	}
}

func (grid *Grid) Update() error {
	return nil
}
