package objects

import (
	"asteroids/game/core"

	"github.com/hajimehoshi/ebiten/v2"
)

type Bullet struct {
	position core.Vector[float64]
	rotation float64
}

func NewBullet(position core.Vector[float64], rotation float64) *Bullet {
	return &Bullet{
		position: position,
		rotation: rotation,
	}
}

func (bullet *Bullet) Draw(screen *ebiten.Image) {

}

func (bullet *Bullet) Update() error {
	return nil
}
