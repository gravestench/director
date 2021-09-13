package tween

import (
	"time"

	"github.com/gravestench/director/pkg/easing"
)

const (
	RepeatForever   = -1
	defaultDuration = time.Second / 2
	defaultEase     = easing.Linear
)

type config struct {
	duration    time.Duration
	delay       time.Duration
	repeatCount int
	ease        func(complete float64) float64
	onStart     func()
	onComplete  func()
	onRepeat    func()
	onUpdate    func(complete float64)
}
