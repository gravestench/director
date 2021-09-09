package tween

import "time"

type mode int

const (
	playingForward mode = iota
	paused
	finished
)

type Tween struct {
	*tweenConfig
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
	return float64((t.elapsed - t.delay).Milliseconds()) / float64(t.duration.Milliseconds())
}

func (t *Tween) Update(dt time.Duration) *Tween {
	if t.mode == paused || t.mode == finished {
		return t
	}

	t.elapsed += dt

	if t.elapsed > t.duration {
		t.elapsed %= t.duration
		if t.repeatCount > 0 {
			t.repeatCount--
		} else {
			t.onComplete()
			t.elapsed = t.delay + t.duration
			t.Stop()
		}
	}

	if t.elapsed < t.delay {
		return t
	}

	if t.elapsed < t.duration && t.onUpdate != nil {
		t.onUpdate(t.ease(t.Complete()))
	}

	return t
}
