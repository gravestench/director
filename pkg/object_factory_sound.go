package pkg

import (
	"time"

	"github.com/gravestench/akara"

	"github.com/gravestench/director/pkg/common"
)

type soundFactory struct {
	common.EntityManager
	*common.SceneComponents
}

func (factory *soundFactory) New(s *Scene, filePath string, paused bool, volume float64, muted bool, loop bool, speedMultiplier float64) akara.EID {
	e := s.Add.generic.entity(s)

	// all sounds need Audio and FileLoadRequest
	audible := s.Components.Audible.Add(e)
	if paused {
		audible.Pause()
	}
	if muted {
		audible.Mute()
	}
	audible.SetVolume(volume)
	audible.SetLooping(loop)
	audible.SetSpeedMultiplier(speedMultiplier)

	fileLoadRequest := s.Components.FileLoadRequest.Add(e)
	fileLoadRequest.Path = filePath

	factory.EntityManager.AddEntity(e)

	return e
}

func (factory *soundFactory) update(_ *Scene, _ time.Duration) {
	if !factory.EntityManager.IsInit() {
		factory.EntityManager.Init()
	}

	factory.EntityManager.ProcessRemovalQueue()
}
