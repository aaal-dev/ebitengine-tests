package scenes

import (
	"asteroids/assets"
	"asteroids/game/core"
	"asteroids/game/objects"
	"math"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	meteorSpawnTime = 1 * time.Second
)

func CreateMainGameScene() (*MainGameScene, error) {
	return &MainGameScene{
		meteorSpawnTimer: core.NewTimer(meteorSpawnTime),
		player: objects.NewPlayer(
			assets.PlayerSprite,
			core.Vector[float64]{
				X: float64(core.ScreenWidth) / 2,
				Y: float64(core.ScreenHeight) / 2,
			},
		),
	}, nil
}

type MainGameScene struct {
	//attackTimer      *Timer
	player           *objects.Player
	meteorSpawnTimer *core.Timer
	meteors          []objects.Object
	bullets          []objects.Object
}

func (scene *MainGameScene) Draw(screen *ebiten.Image) {
	scene.player.Draw(screen)

	for _, meteor := range scene.meteors {
		meteor.Draw(screen)
	}
}

func (scene *MainGameScene) Update() error {
	scene.player.CalculateSpeed()

	scene.player.ShootCooldown.Update()
	if scene.player.ShootCooldown.IsReady() {
		if ebiten.IsKeyPressed(ebiten.KeySpace) {
			scene.player.ShootCooldown.Reset()

			offset := 50.0

			bulletSpawnPosition := core.Vector[float64]{
				X: scene.player.Position.X +
					scene.player.Anchor.X +
					math.Sin(scene.player.Rotation)*
						offset,
				Y: scene.player.Anchor.Y + math.Cos(scene.player.Rotation)*(-offset),
			}

			bullet := objects.NewBullet(bulletSpawnPosition, 0.0)
			scene.bullets = append(scene.bullets, bullet)
		}
	}

	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		//game.player.MoveDown()
		scene.player.MoveBackward()
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		//game.player.MoveUp()
		scene.player.MoveForward()
	}

	if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		//game.player.MoveLeft()
		scene.player.RotateCounterclockwise()
	}

	if ebiten.IsKeyPressed(ebiten.KeyRight) {
		//game.player.MoveRight()
		scene.player.RotateClockwise()
	}

	scene.player.Update()

	//game.attackTimer.Update()
	//if game.attackTimer.IsReady() {
	//	game.attackTimer.Reset()
	//}

	if len(scene.meteors) < 10 {
		scene.meteorSpawnTimer.Update()
		if scene.meteorSpawnTimer.IsReady() {
			scene.meteorSpawnTimer.Reset()

			meteor := objects.NewMeteor()
			scene.meteors = append(scene.meteors, meteor)
		}
	}

	for _, meteor := range scene.meteors {
		meteor.Update()
	}

	return nil
}
