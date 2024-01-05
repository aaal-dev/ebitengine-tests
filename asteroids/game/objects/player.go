package objects

import (
	"asteroids/assets"
	"asteroids/game/core"
	"image/color"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
)

const (
	startRotation = math.Pi / 2
)

var (
	//factor        float64
	moveSpeed     float64
	rotationSpeed float64
)

func NewPlayer(sprite *ebiten.Image, position core.Vector[float64]) *Player {
	bounds := sprite.Bounds()

	size := core.Vector[float64]{
		X: float64(bounds.Dx()),
		Y: float64(bounds.Dy()),
	}

	anchor := core.Vector[float64]{
		X: size.X * 0.5,
		Y: size.Y * 0.5,
	}

	position.X -= anchor.X
	position.Y -= anchor.Y

	return &Player{
		Position:      position,
		Sprite:        sprite,
		Anchor:        anchor,
		Size:          size,
		Rotation:      startRotation,
		ShootCooldown: core.NewTimer(1 * time.Second),
	}
}

type Player struct {
	Position      core.Vector[float64]
	PositionDelta core.Vector[float64]
	Anchor        core.Vector[float64]
	Size          core.Vector[float64]
	Rotation      float64
	Sprite        *ebiten.Image
	ShootCooldown *core.Timer
}

func (player *Player) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}

	options.GeoM.Translate(-player.Anchor.X, -player.Anchor.Y)
	options.GeoM.Rotate(player.Rotation - startRotation)
	options.GeoM.Translate(player.Anchor.X, player.Anchor.Y)

	options.GeoM.Translate(player.Position.X, player.Position.Y)
	screen.DrawImage(player.Sprite, options)

	if player.Position.X < 0 {
		options.GeoM.Translate(
			player.Position.X+float64(core.ScreenWidth),
			player.Position.Y,
		)
		screen.DrawImage(player.Sprite, options)
	}

	if (player.Position.X + player.Size.X) > float64(core.ScreenWidth) {
		options.GeoM.Translate(
			player.Position.X-float64(core.ScreenWidth),
			player.Position.Y,
		)
		screen.DrawImage(player.Sprite, options)
	}

	if player.Position.Y < 0 {
		options.GeoM.Translate(
			player.Position.X,
			player.Position.Y+float64(core.ScreenHeight),
		)
		screen.DrawImage(player.Sprite, options)
	}

	if (player.Position.Y + player.Size.Y) > float64(core.ScreenHeight) {
		options.GeoM.Translate(
			player.Position.X,
			player.Position.Y-float64(core.ScreenHeight),
		)
		screen.DrawImage(player.Sprite, options)
	}

	text.Draw(
		screen,
		"String",
		assets.NormalFont,
		int(player.Position.X),
		int(player.Position.Y),
		color.White,
	)
}

func (player *Player) Update() error {
	//if player.positionDelta.X != 0 && player.positionDelta.Y != 0 {
	//	factor = moveSpeed / player.positionDelta.Magnitude()
	//	player.positionDelta.X *= factor
	//	player.positionDelta.Y *= factor
	//}

	player.Position.X += player.PositionDelta.X
	player.Position.Y += player.PositionDelta.Y

	if (player.Position.X - player.Size.X) < 0 {
		player.Position.X = float64(core.ScreenWidth) - player.Size.X
	}

	if player.Position.X > float64(core.ScreenWidth) {
		player.Position.X = 0.0
	}

	if (player.Position.Y - player.Size.Y) < 0 {
		player.Position.Y = float64(core.ScreenHeight) - player.Size.Y
	}

	if player.Position.Y > float64(core.ScreenHeight) {
		player.Position.Y = 0.0
	}

	player.PositionDelta.X = 0
	player.PositionDelta.Y = 0

	return nil
}

func (player *Player) CalculateSpeed() {
	moveSpeed = float64(300 / ebiten.TPS())
	rotationSpeed = math.Pi / float64(ebiten.TPS())
}

func (player *Player) MoveUp() {
	player.PositionDelta.Y = -moveSpeed
}

func (player *Player) MoveDown() {
	player.PositionDelta.Y = moveSpeed
}

func (player *Player) MoveLeft() {
	player.PositionDelta.X = -moveSpeed
}

func (player *Player) MoveRight() {
	player.PositionDelta.X = moveSpeed
}

func (player *Player) MoveForward() {
	player.PositionDelta = core.Vector[float64]{
		X: (-moveSpeed) * math.Cos(player.Rotation),
		Y: (-moveSpeed) * math.Sin(player.Rotation),
	}
}

func (player *Player) MoveBackward() {
	player.PositionDelta = core.Vector[float64]{
		X: moveSpeed * math.Cos(player.Rotation),
		Y: moveSpeed * math.Sin(player.Rotation),
	}
}

func (player *Player) StrafeLeft() {

}

func (player *Player) StrafeRight() {

}

func (player *Player) RotateClockwise() {
	player.Rotation += rotationSpeed
}

func (player *Player) RotateCounterclockwise() {
	player.Rotation -= rotationSpeed
}
