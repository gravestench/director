package components

import (
	"github.com/gravestench/akara"
)

var _ akara.Component = &RenderOrder{}

// RenderOrder is a component that contains normalized alpha transparency (0.0 ... 1.0)
type RenderOrder struct {
	Value int
}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
func (*RenderOrder) New() akara.Component {
	return &RenderOrder{}
}

// RenderOrderFactory is a wrapper for the generic component factory that returns RenderOrder component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a RenderOrder.
type RenderOrderFactory struct {
	*akara.ComponentFactory
}

// Add adds a RenderOrder component to the given entity and returns it
func (m *RenderOrderFactory) Add(id akara.EID) *RenderOrder {
	return m.ComponentFactory.Add(id).(*RenderOrder)
}

// Get returns the RenderOrder component for the given entity, and a bool for whether or not it exists
func (m *RenderOrderFactory) Get(id akara.EID) (*RenderOrder, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*RenderOrder), found
}
