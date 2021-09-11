package scene

import (
	"github.com/gravestench/director/pkg/common"
	"time"
)

type genericFactory struct {
	*common.BasicComponents
}

func (factory *genericFactory) update(s *Scene, dt time.Duration) {}

func (factory *genericFactory) entity(s *Scene) common.Entity {
	e := s.Director.NewEntity()

	// a generic entity always has a UUID
	s.Components.UUID.Add(e)

	return e
}

func (factory *genericFactory) visibleEntity(s *Scene) common.Entity {
	e := factory.entity(s)

	// it will always set its parent to the scene's root scene graph node
	node := s.Components.SceneGraphNode.Add(e)
	node.SetParent(&s.Graph)

	// a visible entity always has a position, rotation, and scale
	s.Components.Transform.Add(e)

	// a visible entity will always have an origin point
	// which is relative to its display dimensions
	s.Components.Origin.Add(e)

	// a visible entity will always have an opacity
	s.Components.Opacity.Add(e)

	// a visible entity will always have a render order
	s.Components.RenderOrder.Add(e)

	// add this to the scene's render list
	s.renderList = append(s.renderList, e)

	return e
}
