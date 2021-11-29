package scene_graph_node

import (
	"github.com/gravestench/akara"
)

// ComponentFactory is a wrapper for the generic component factory that returns SceneGraphNode component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a SceneGraphNode.
type ComponentFactory struct {
	*akara.ComponentFactory
}

// Add adds a SceneGraphNode component to the given entity and returns it
func (concrete *ComponentFactory) Add(id akara.EID) *SceneGraphNode {
	return concrete.ComponentFactory.Add(id).(*SceneGraphNode)
}

// Get returns the SceneGraphNode component for the given entity, and a bool for whether or not it exists
func (concrete *ComponentFactory) Get(id akara.EID) (*SceneGraphNode, bool) {
	component, found := concrete.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*SceneGraphNode), found
}
