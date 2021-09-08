package tween

import (
	"time"

	"github.com/gravestench/akara"
)

type System struct {
	akara.BaseSystem
	tweens []*Tween
}

// pass in a TweenBuilder or Tween instance
func (s *System) New(args ...interface{}) {
	if !s.IsInitialized() {
		s.Init(nil)
	}

	for idx := range args {
		switch v := args[idx].(type) {
		case Builder:
			s.tweens = append(s.tweens, v.Build())
		case *Builder:
			s.tweens = append(s.tweens, v.Build())
		case *Tween:
			s.tweens = append(s.tweens, v)
		}
	}
}

func (s *System) Update(duration time.Duration) {
	for idx := range s.tweens {
		s.tweens[idx].Update(duration)
	}
}

func (s *System) Init(_ *akara.World) {
	s.tweens = make([]*Tween, 0)
}

func (s *System) IsInitialized() bool {
	return s.tweens != nil
}

