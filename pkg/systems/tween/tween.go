package tween

import (
	"github.com/gravestench/director/pkg/easing"
	"time"
)

type mode int

const (
	playingForward mode = iota
	paused
	finished
)

type Tween struct {
	cfg *config
	mode
	elapsed time.Duration
}

func (t *Tween) Start() *Tween {
	t.elapsed = 0

	t.mode = playingForward

	return t
}

func (t *Tween) Stop() *Tween {
	t.mode = paused

	return t
}

func (t *Tween) Play() *Tween {
	t.mode = paused

	return t
}

func (t *Tween) Pause() *Tween {
	t.mode = paused

	return t
}

func (t *Tween) Complete() float64 {
	return float64((t.elapsed - t.cfg.delay).Milliseconds()) / float64(t.cfg.duration.Milliseconds())
}

func (t *Tween) Update(dt time.Duration) *Tween {
	if t.mode == paused || t.mode == finished {
		return t
	}

	t.elapsed += dt

	if t.elapsed > t.cfg.duration {
		t.elapsed %= t.cfg.duration
		if t.cfg.repeatCount > 0 {
			t.cfg.repeatCount--
		} else {
			t.cfg.onComplete()
			t.elapsed = t.cfg.delay + t.cfg.duration
			t.Stop()
		}
	}

	if t.elapsed < t.cfg.delay {
		return t
	}

	if t.elapsed < t.cfg.duration && t.cfg.onUpdate != nil {
		t.cfg.onUpdate(t.cfg.ease(t.Complete()))
	}

	return t
}

func (t *Tween) Time(dt time.Duration) *Tween {
	t.cfg.duration = dt

	return t
}

func (t *Tween) Ease(args ...interface{}) *Tween {
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

	t.cfg.ease = easeFn

	return t
}

func (t *Tween) OnStart(fn func()) *Tween {
	t.cfg.onStart = fn

	return t
}

func (t *Tween) OnComplete(fn func()) *Tween {
	t.cfg.onComplete = fn

	return t
}

func (t *Tween) OnRepeat(fn func()) *Tween {
	t.cfg.onRepeat = fn

	return t
}

func (t *Tween) OnUpdate(fn func(float64)) *Tween {
	t.cfg.onUpdate = fn

	return t
}

func (t *Tween) Delay(dt time.Duration) *Tween {
	t.cfg.delay = dt

	return t
}

func (t *Tween) Repeat(count int) *Tween {
	t.cfg.repeatCount = count

	return t
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
