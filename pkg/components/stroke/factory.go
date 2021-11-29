package stroke

import (
	"github.com/gravestench/akara"
)

// ComponentFactory is a wrapper for the generic component factory that returns Stroke component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Stroke.
type ComponentFactory struct {
	*akara.ComponentFactory
}

// Add adds a Stroke component to the given entity and returns it
func (concrete *ComponentFactory) Add(id akara.EID) *Stroke {
	return concrete.ComponentFactory.Add(id).(*Stroke)
}

// Get returns the Stroke component for the given entity, and a bool for whether or not it exists
func (concrete *ComponentFactory) Get(id akara.EID) (*Stroke, bool) {
	component, found := concrete.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Stroke), found
}
