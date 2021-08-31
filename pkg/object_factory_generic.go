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
	s.Components.UUID.Add(e)

	return e
}

func (factory *genericFactory) visibleEntity(s *Scene) akara.EID {
	e := factory.entity(s)

	// it will always set its parent to the scene's root scene graph node
	node := s.Components.SceneGraphNode.Add(e)
	node.SetParent(&s.Graph)

	// a visible entity always has a position, rotation, and scale
	s.Components.Transform.Add(e)

	// a visible entity will always have an origin point
	// which is relative to its display dimensions
	s.Components.Origin.Add(e)

	return e
}
