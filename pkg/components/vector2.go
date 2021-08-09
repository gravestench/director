package components

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
)

// static check that Vector2 implements Component
var _ akara.Component = &Vector2{}

// Vector2 is a component that contains a 2-dimensional Vector
type Vector2 struct {
	rl.Vector2
}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
func (*Vector2) New() akara.Component {
	return &Vector2{}
}

// Vector2Factory is a wrapper for the generic component factory that returns Vector2 component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Vector2.
type Vector2Factory struct {
	*akara.ComponentFactory
}

// Add adds a Vector2 component to the given entity and returns it
func (m *Vector2Factory) Add(id akara.EID) *Vector2 {
	return m.ComponentFactory.Add(id).(*Vector2)
}

// Get returns the Vector2 component for the given entity, and a bool for whether or not it exists
func (m *Vector2Factory) Get(id akara.EID) (*Vector2, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Vector2), found
}

