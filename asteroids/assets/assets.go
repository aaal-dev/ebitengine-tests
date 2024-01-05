package assets

import (
	"embed"
	"image"
	_ "image/png"
	"io/fs"

	"github.com/hajimehoshi/ebiten/v2"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var (
	//go:embed *
	assets        embed.FS
	MeteorSprites = mustLoadImages("meteors/*.png")
	PlayerSprite  = mustLoadImage("player.png")
	NormalFont    = mustLoadFont("fonts/Zack and Sarah.ttf")
)

func mustLoadImage(name string) *ebiten.Image {
	file, err := assets.Open(name)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}

	return ebiten.NewImageFromImage(img)
}

func mustLoadImages(path string) []*ebiten.Image {
	matches, err := fs.Glob(assets, path)
	if err != nil {
		panic(err)
	}

	images := make([]*ebiten.Image, len(matches))
	for i, match := range matches {
		images[i] = mustLoadImage(match)
	}

	return images
}

func mustLoadFont(path string) font.Face {
	var (
		fontSize = 24
		dpi      = 96
	)

	file, err := assets.ReadFile(path)
	if err != nil {
		panic(err)
	}

	tt, err := opentype.Parse(file)
	if err != nil {
		panic(err)
	}

	font, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    float64(fontSize),
		DPI:     float64(dpi),
		Hinting: font.HintingVertical,
	})
	if err != nil {
		panic(err)
	}

	return font
}
