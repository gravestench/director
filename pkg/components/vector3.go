package components

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
)

// static check that Vector3 implements Component
var _ akara.Component = &Vector3{}

// Vector3 is a component that contains a 3-dimensional Vector
type Vector3 struct {
	rl.Vector3
}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
func (*Vector3) New() akara.Component {
	return &Vector3{}
}

// Vector3Factory is a wrapper for the generic component factory that returns Vector3 component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Vector3.
type Vector3Factory struct {
	*akara.ComponentFactory
}

// Add adds a Vector3 component to the given entity and returns it
func (m *Vector3Factory) Add(id akara.EID) *Vector3 {
	return m.ComponentFactory.Add(id).(*Vector3)
}

// Get returns the Vector3 component for the given entity, and a bool for whether or not it exists
func (m *Vector3Factory) Get(id akara.EID) (*Vector3, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Vector3), found
}

