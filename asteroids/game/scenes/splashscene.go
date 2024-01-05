package scenes

import "github.com/hajimehoshi/ebiten/v2"

func CreateSplashScene() (*SplashScene, error) {
	return &SplashScene{}, nil
}

type SplashScene struct {
}

func (splash *SplashScene) Draw(screen *ebiten.Image) {

}

func (splash *SplashScene) Update() error {
	return nil
}
