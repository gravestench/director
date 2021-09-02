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

type tweenConfig struct {
	duration    time.Duration
	delay       time.Duration
	repeatCount int
	ease        func(complete float64) float64
	onStart     func()
	onComplete  func()
	onRepeat    func()
	onUpdate    func(complete float64)
}


func NewBuilder() *Builder {
	tb := &Builder{}
	tb.cfg = &tweenConfig{}

	tb.Time(defaultDuration)
	tb.Ease(defaultEase)
	tb.OnStart(func() {})
	tb.OnComplete(func() {})
	tb.OnUpdate(func(float64) {})

	return tb
}

type Builder struct {
	cfg *tweenConfig
}

func (tb *Builder) Build() *Tween {
	return &Tween{
		tweenConfig: tb.cfg,
	}
}

func (tb *Builder) Time(dt time.Duration) *Builder {
	tb.cfg.duration = dt

	return tb
}

func (tb *Builder) Ease(args ...interface{}) *Builder {
	var easeFn func(float64) float64

	if len(args) >= 2 {
		if params, ok := args[1].([]float64); ok {
			easeFn = getEaseFn(args[0], params)
		}
	} else if len(args) == 1 {
		easeFn = getEaseFn(args[0], nil)
	} else {
		easeFn = getEaseFn(defaultEase, nil)
	}

	tb.cfg.ease = easeFn

	return tb
}

func (tb *Builder) OnStart(fn func()) *Builder {
	tb.cfg.onStart = fn

	return tb
}

func (tb *Builder) OnComplete(fn func()) *Builder {
	tb.cfg.onComplete = fn

	return tb
}

func (tb *Builder) OnRepeat(fn func()) *Builder {
	tb.cfg.onRepeat = fn

	return tb
}

func (tb *Builder) OnUpdate(fn func(float64)) *Builder {
	tb.cfg.onUpdate = fn

	return tb
}

func (tb *Builder) Delay(dt time.Duration) *Builder {
	tb.cfg.delay = dt

	return tb
}

func (tb *Builder) Repeat(count int) *Builder {
	tb.cfg.repeatCount = count

	return tb
}

func getEaseFn(ease interface{}, params []float64) func(float64) float64 {
	switch e := ease.(type) {
	case string:
		provider, found := easing.EaseMap[e]
		if found {
			return provider.New(params)
		}
	case func(float64) float64:
		return e
	}

	return easing.EaseMap[easing.Default].New(params)
}
