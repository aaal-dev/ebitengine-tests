package objects

import (
	"asteroids/assets"
	"asteroids/game/core"
	"math"
	"math/rand"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	radius = core.ScreenWidth / 2.0
	target = core.Vector[float64]{
		X: float64(core.ScreenWidth) / 2,
		Y: float64(core.ScreenHeight) / 2,
	}
)

type Meteor struct {
	position      core.Vector[float64]
	positionDelta core.Vector[float64]
	anchor        core.Vector[float64]
	rotation      float64
	rotationSpeed float64
	velocity      float64
	sprite        *ebiten.Image
}

func NewMeteor() *Meteor {
	sprite := assets.MeteorSprites[rand.Intn(len(assets.MeteorSprites))]

	bounds := sprite.Bounds()

	anchor := core.Vector[float64]{
		X: float64(bounds.Dx()) / 2,
		Y: float64(bounds.Dy()) / 2,
	}

	angle := rand.Float64() * 2 * math.Pi

	position := core.Vector[float64]{
		X: target.X + math.Cos(angle)*float64(radius),
		Y: target.Y + math.Sin(angle)*float64(radius),
	}
	velocity := rand.Float64()*1.5 + 0.25

	direction := core.Vector[float64]{
		X: target.X - position.X,
		Y: target.Y - position.Y,
	}

	normalizedDirection := direction.Normalize()

	positionDelta := core.Vector[float64]{
		X: normalizedDirection.X * velocity,
		Y: normalizedDirection.Y * velocity,
	}

	rotationSpeed := rand.Float64()*0.04 - 0.02

	return &Meteor{
		position:      position,
		positionDelta: positionDelta,
		anchor:        anchor,
		velocity:      velocity,
		rotationSpeed: rotationSpeed,
		sprite:        sprite,
	}
}

func (meteor *Meteor) Update() error {
	meteor.position.X += meteor.positionDelta.X
	meteor.position.Y += meteor.positionDelta.Y
	meteor.rotation += meteor.rotationSpeed

	return nil
}

func (meteor *Meteor) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}

	options.GeoM.Translate(-meteor.anchor.X, -meteor.anchor.Y)
	options.GeoM.Rotate(meteor.rotation)
	options.GeoM.Translate(meteor.anchor.X, meteor.anchor.Y)

	options.GeoM.Translate(meteor.position.X, meteor.position.Y)

	screen.DrawImage(meteor.sprite, options)
}
