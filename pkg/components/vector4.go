package components

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
)

// static check that Vector4 implements Component
var _ akara.Component = &Vector4{}

// Vector4 is a component that contains a 4-dimensional Vector
type Vector4 struct {
	rl.Vector4
}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
func (*Vector4) New() akara.Component {
	return &Vector4{}
}

// Vector4Factory is a wrapper for the generic component factory that returns Vector4 component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Vector4.
type Vector4Factory struct {
	*akara.ComponentFactory
}

// Add adds a Vector4 component to the given entity and returns it
func (m *Vector4Factory) Add(id akara.EID) *Vector4 {
	return m.ComponentFactory.Add(id).(*Vector4)
}

// Get returns the Vector4 component for the given entity, and a bool for whether or not it exists
func (m *Vector4Factory) Get(id akara.EID) (*Vector4, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Vector4), found
}

