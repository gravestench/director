package scene

import (
	"github.com/gravestench/director/pkg/common"
	"time"
)

type layerFactory struct {
	common.EntityManager
	*common.BasicComponents
}

func (factory *layerFactory) New(s *Scene, x, y int) common.Entity {
	e := s.Add.generic.visibleEntity(s)

	trs, _ := s.Components.Transform.Get(e) // this is a component all visible entities have
	trs.Translation.Set(float64(x), float64(y), trs.Translation.Z)

	s.Components.HasChildren.Add(e)

	factory.EntityManager.AddEntity(e)

	return e
}

func (factory *layerFactory) update(s *Scene, _ time.Duration) {
	if !factory.EntityManager.IsInit() {
		factory.EntityManager.Init()
	}

	factory.EntitiesMutex.Lock()
	for e := range factory.Entities {
		hc, found := s.Components.HasChildren.Get(e)
		if !found {
			s.RemoveEntity(e)
		}

		parentNode, found := s.Components.SceneGraphNode.Get(e)
		if !found {
			s.RemoveEntity(e)
		}

		for _, child := range hc.Children {
			childNode, found := s.Components.SceneGraphNode.Get(child)
			if !found {
				continue
			}

			childNode.SetParent(parentNode.Node)
		}
	}
	factory.EntitiesMutex.Unlock()

	factory.EntityManager.ProcessRemovalQueue()
}
