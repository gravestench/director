package has_children

import (
	"github.com/gravestench/akara"
)

// ComponentFactory is a wrapper for the generic component factory.
type ComponentFactory struct {
	*akara.ComponentFactory
}

// Add a HasChildren component to the given entity and return it
func (concrete *ComponentFactory) Add(id akara.EID) *HasChildren {
	return concrete.ComponentFactory.Add(id).(*HasChildren)
}

// Get returns the HasChildren component for the given entity, and a bool for whether or not it exists
func (concrete *ComponentFactory) Get(id akara.EID) (*HasChildren, bool) {
	component, found := concrete.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*HasChildren), found
}
