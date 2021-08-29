package components

import (
	"github.com/gravestench/akara"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// static check that RenderTexture2D implements Component
var _ akara.Component = &RenderTexture2D{}

// RenderTexture2D is a wrapper for a raylib texture
type RenderTexture2D struct {
	*rl.RenderTexture2D
}

// New creates a new raylib texture
func (*RenderTexture2D) New() akara.Component {
	return &RenderTexture2D{}
}

// RenderTexture2DFactory is a wrapper for the generic component factory that returns RenderTexture2D component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a RenderTexture2D.
type RenderTexture2DFactory struct {
	*akara.ComponentFactory
}

// Add adds a RenderTexture2D component to the given entity and returns it
func (m *RenderTexture2DFactory) Add(id akara.EID) *RenderTexture2D {
	return m.ComponentFactory.Add(id).(*RenderTexture2D)
}

// Get returns the RenderTexture2D component for the given entity, and a bool for whether or not it exists
func (m *RenderTexture2DFactory) Get(id akara.EID) (*RenderTexture2D, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*RenderTexture2D), found
}
