package tween

import (
	"github.com/gravestench/akara"
)

type System struct {
	akara.BaseSystem
	queue []*Tween
}

func (s *System) Name() string {
	return "Tween"
}

// New creates a new tween, but does not add it for processing.
func (s *System) New() *Tween {
	t := &Tween{}
	t.config = &config{}

	t.justStarted = true
	t.Time(defaultDuration)
	t.Ease(defaultEase)
	t.OnStart(func() {})
	t.OnComplete(func() {})
	t.OnUpdate(func(float64) {})

	return t
}

// Add the given teen to the processing queue
func (s *System) Add(t *Tween) {
	s.queue = append(s.queue, t)
}

// Remove the given tween from the queue
func (s *System) Remove(t *Tween) {
	for idx := range s.queue {
		if s.queue[idx] != t {
			continue
		}

		s.queue = append(s.queue[:idx], s.queue[idx+1:]...)

		break
	}
}

func (s *System) Update() {
	for idx := range s.queue {
		s.queue[idx].Update(s.TimeDelta)
	}
}

func (s *System) Init(_ *akara.World) {
	s.queue = make([]*Tween, 0)
}

func (s *System) IsInitialized() bool {
	return s.queue != nil
}
