package transform

import (
	"github.com/gravestench/akara"
)

// ComponentFactory is a wrapper for the generic component factory.
type ComponentFactory struct {
	*akara.ComponentFactory
}

// Add adds a Transform component to the given entity and returns it
func (concrete *ComponentFactory) Add(id akara.EID) *Transform {
	return concrete.ComponentFactory.Add(id).(*Transform)
}

// Get returns the Transform component for the given entity, and a bool for whether or not it exists
func (concrete *ComponentFactory) Get(id akara.EID) (*Transform, bool) {
	component, found := concrete.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Transform), found
}
