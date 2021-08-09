package pkg

import (
	"github.com/gravestench/akara"
)

type SceneFace interface {
	akara.System
	bindsToDirector
	hasKey
}

type bindsToDirector interface {
	bind(*Director)
	unbind()
}

type hasKey interface {
	Key() string
}

type Scene struct {
	*Director
	akara.BaseSystem
	basicComponents
	key string
	Add SceneObjectFactory
}

func (s *Scene) bind(d *Director) {
	s.Director = d
	s.init(d.World)
	s.Add.scene = s
}

func (s *Scene) unbind() {
	s.Director = nil
}
