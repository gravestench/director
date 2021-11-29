package render_order

import (
	"github.com/gravestench/akara"
)

// ComponentFactory is a wrapper for the generic component factory.
type ComponentFactory struct {
	*akara.ComponentFactory
}

// Add adds a RenderOrder component to the given entity and returns it
func (concrete *ComponentFactory) Add(id akara.EID) *RenderOrder {
	return concrete.ComponentFactory.Add(id).(*RenderOrder)
}

// Get returns the RenderOrder component for the given entity, and a bool for whether or not it exists
func (concrete *ComponentFactory) Get(id akara.EID) (*RenderOrder, bool) {
	component, found := concrete.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*RenderOrder), found
}
