package size

import (
	"github.com/gravestench/akara"
)

// ComponentFactory is a wrapper for the generic component factory.
type ComponentFactory struct {
	*akara.ComponentFactory
}

// Add adds a Size component to the given entity and returns it
func (concrete *ComponentFactory) Add(id akara.EID) *Size {
	return concrete.ComponentFactory.Add(id).(*Size)
}

// Get returns the Size component for the given entity, and a bool for whether or not it exists
func (concrete *ComponentFactory) Get(id akara.EID) (*Size, bool) {
	component, found := concrete.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Size), found
}
