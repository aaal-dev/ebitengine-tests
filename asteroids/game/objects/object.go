package objects

import "github.com/hajimehoshi/ebiten/v2"

type Object interface {
	Draw(*ebiten.Image)
	Update() error
}
