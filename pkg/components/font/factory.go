package font

import (
	"github.com/gravestench/akara"
)

// ComponentFactory is a wrapper for the generic component factory that returns Font component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Font.
type ComponentFactory struct {
	*akara.ComponentFactory
}

// Add adds a Font component to the given entity and returns it
func (concrete *ComponentFactory) Add(id akara.EID) *Font {
	return concrete.ComponentFactory.Add(id).(*Font)
}

// Get returns the Font component for the given entity, and a bool for whether or not it exists
func (concrete *ComponentFactory) Get(id akara.EID) (*Font, bool) {
	component, found := concrete.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Font), found
}
