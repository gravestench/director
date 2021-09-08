package components

import (
	"github.com/gravestench/akara"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// static check that Texture2D implements Component
var _ akara.Component = &Texture2D{}

// Texture2D is a wrapper for a raylib texture
type Texture2D struct {
	*rl.Texture2D
}

// New creates a new raylib texture
func (*Texture2D) New() akara.Component {
	return &Texture2D{}
}

// Texture2DFactory is a wrapper for the generic component factory that returns Texture2D component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Texture2D.
type Texture2DFactory struct {
	*akara.ComponentFactory
}

// Add adds a Texture2D component to the given entity and returns it
func (m *Texture2DFactory) Add(id akara.EID) *Texture2D {
	return m.ComponentFactory.Add(id).(*Texture2D)
}

// Get returns the Texture2D component for the given entity, and a bool for whether or not it exists
func (m *Texture2DFactory) Get(id akara.EID) (*Texture2D, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Texture2D), found
}
