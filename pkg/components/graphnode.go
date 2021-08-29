package components

import (
	"github.com/gravestench/akara"
	"github.com/gravestench/scenegraph"
)

var _ akara.Component = &SceneGraphNode{}

// SceneGraphNode is a component that contains normalized alpha transparency (0.0 ... 1.0)
type SceneGraphNode struct {
	*scenegraph.Node
}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
func (*SceneGraphNode) New() akara.Component {
	return &SceneGraphNode{
		Node: scenegraph.NewNode(),
	}
}

// SceneGraphNodeFactory is a wrapper for the generic component factory that returns SceneGraphNode component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a SceneGraphNode.
type SceneGraphNodeFactory struct {
	*akara.ComponentFactory
}

// Add adds a SceneGraphNode component to the given entity and returns it
func (m *SceneGraphNodeFactory) Add(id akara.EID) *SceneGraphNode {
	return m.ComponentFactory.Add(id).(*SceneGraphNode)
}

// Get returns the SceneGraphNode component for the given entity, and a bool for whether or not it exists
func (m *SceneGraphNodeFactory) Get(id akara.EID) (*SceneGraphNode, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*SceneGraphNode), found
}
