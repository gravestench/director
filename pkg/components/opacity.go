package components

import (
	"github.com/gravestench/akara"
)

// static check that Opacity implements Component
var _ akara.Component = &Opacity{}

// Opacity is a component that contains normalized alpha transparency (0.0 ... 1.0)
type Opacity struct {
	Value float64
}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
func (*Opacity) New() akara.Component {
	const defaultAlpha = 1.0

	return &Opacity{
		Value: defaultAlpha,
	}
}

// OpacityFactory is a wrapper for the generic component factory that returns Opacity component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Opacity.
type OpacityFactory struct {
	*akara.ComponentFactory
}

// Add adds a Opacity component to the given entity and returns it
func (m *OpacityFactory) Add(id akara.EID) *Opacity {
	return m.ComponentFactory.Add(id).(*Opacity)
}

// Get returns the Opacity component for the given entity, and a bool for whether or not it exists
func (m *OpacityFactory) Get(id akara.EID) (*Opacity, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Opacity), found
}
