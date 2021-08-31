package components

import (
	"github.com/gravestench/akara"
	"github.com/gravestench/mathlib"
)

var _ akara.Component = &Origin{}

// Origin is a component that contains normalized alpha transparency (0.0 ... 1.0)
type Origin struct {
	*mathlib.Vector3
}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
func (*Origin) New() akara.Component {
	const defaultOrigin = 0.5 // normalized, 0.5 is center

	return &Origin{
		Vector3: &mathlib.Vector3{
			X: defaultOrigin,
			Y: defaultOrigin,
		},
	}
}

// OriginFactory is a wrapper for the generic component factory that returns Origin component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Origin.
type OriginFactory struct {
	*akara.ComponentFactory
}

// Add adds a Origin component to the given entity and returns it
func (m *OriginFactory) Add(id akara.EID) *Origin {
	return m.ComponentFactory.Add(id).(*Origin)
}

// Get returns the Origin component for the given entity, and a bool for whether or not it exists
func (m *OriginFactory) Get(id akara.EID) (*Origin, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Origin), found
}
