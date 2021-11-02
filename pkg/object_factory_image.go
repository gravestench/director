package pkg

import (
	"time"

	"github.com/gravestench/director/pkg/common"
)

type imageFactory struct {
	*common.SceneComponents
	common.EntityManager
}

func (factory *imageFactory) update(_ *Scene, _ time.Duration) {
	if !factory.EntityManager.IsInit() {
		factory.EntityManager.Init()
	}

	factory.EntityManager.ProcessRemovalQueue()
}

func (factory *imageFactory) New(s *Scene, uri string, x, y int) common.Entity {
	e := s.Add.generic.visibleEntity(s)

	trs, _ := s.Components.Transform.Get(e) // this is a component all visible entities have
	trs.Translation.Set(float64(x), float64(y), trs.Translation.Z)

	req := s.Components.FileLoadRequest.Add(e)
	req.Path = uri

	factory.EntityManager.AddEntity(e)

	return e
}
