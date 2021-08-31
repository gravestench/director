package components

import (
	"github.com/gravestench/akara"
	"image"
)

// static check that Size implements Component
var _ akara.Component = &Size{}

// Size is a component that contains normalized alpha transparency (0.0 ... 1.0)
type Size struct {
	image.Rectangle
}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
func (*Size) New() akara.Component {
	return &Size{}
}

// SizeFactory is a wrapper for the generic component factory that returns Size component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Size.
type SizeFactory struct {
	*akara.ComponentFactory
}

// Add adds a Size component to the given entity and returns it
func (m *SizeFactory) Add(id akara.EID) *Size {
	return m.ComponentFactory.Add(id).(*Size)
}

// Get returns the Size component for the given entity, and a bool for whether or not it exists
func (m *SizeFactory) Get(id akara.EID) (*Size, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Size), found
}
