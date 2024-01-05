package core

import (
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Timer struct {
	currentTicks int
	targetTicks  int
}

func NewTimer(duration time.Duration) *Timer {
	return &Timer{
		currentTicks: 0,
		targetTicks:  int(duration.Milliseconds()) * ebiten.TPS() / 1000,
	}
}

func (timer *Timer) Update() {
	if timer.currentTicks < timer.targetTicks {
		timer.currentTicks++
	}
}

func (timer *Timer) IsReady() bool {
	return timer.currentTicks >= timer.targetTicks
}

func (timer *Timer) Reset() {
	timer.currentTicks = 0
}
