package render_texture

import (
	"github.com/gravestench/akara"
)

// ComponentFactory is a wrapper for the generic component factory.
type ComponentFactory struct {
	*akara.ComponentFactory
}

// Add adds a RenderTexture2D component to the given entity and returns it
func (concrete *ComponentFactory) Add(id akara.EID) *RenderTexture2D {
	return concrete.ComponentFactory.Add(id).(*RenderTexture2D)
}

// Get returns the RenderTexture2D component for the given entity, and a bool for whether or not it exists
func (concrete *ComponentFactory) Get(id akara.EID) (*RenderTexture2D, bool) {
	component, found := concrete.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*RenderTexture2D), found
}
