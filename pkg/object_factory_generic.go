package pkg

import (
	"github.com/gravestench/akara"
	"time"
)

type genericFactory struct {
	*basicComponents
}

func (factory *genericFactory) update(s *Scene, dt time.Duration) {}

func (factory *genericFactory) entity(s *Scene) akara.EID {
	e := s.Director.NewEntity()

	// a generic entity always has a UUID
	s.basicComponents.UUID.Add(e)

	return e
}

func (factory *genericFactory) visibleEntity(s *Scene) akara.EID {
	e := factory.entity(s)

	// it will always set its parent to the scene's root scene graph node
	node := s.SceneGraphNode.Add(e)
	node.SetParent(&s.Graph)

	// a visible entity always has a position, rotation, and scale
	s.Transform.Add(e)

	return e
}
