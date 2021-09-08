package components

import (
	"github.com/gravestench/akara"
)

var _ akara.Component = &Debug{}

// Debug is a tag component, used as a flag for debugging an entity.
type Debug struct {}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
func (*Debug) New() akara.Component {
	return &Debug{}
}

// DebugFactory is a wrapper for the generic component factory that returns Debug component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Debug.
type DebugFactory struct {
	*akara.ComponentFactory
}

// Add adds a Debug component to the given entity and returns it
func (m *DebugFactory) Add(id akara.EID) *Debug {
	return m.ComponentFactory.Add(id).(*Debug)
}

// Get returns the Debug component for the given entity, and a bool for whether or not it exists
func (m *DebugFactory) Get(id akara.EID) (*Debug, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Debug), found
}
