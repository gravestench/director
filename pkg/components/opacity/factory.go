package opacity

import "github.com/gravestench/akara"

// ComponentFactory is a wrapper for the generic component factory that returns Opacity component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Opacity.
type ComponentFactory struct {
	*akara.ComponentFactory
}

// Add adds a Opacity component to the given entity and returns it
func (concrete *ComponentFactory) Add(id akara.EID) *Opacity {
	return concrete.ComponentFactory.Add(id).(*Opacity)
}

// Get returns the Opacity component for the given entity, and a bool for whether or not it exists
func (concrete *ComponentFactory) Get(id akara.EID) (*Opacity, bool) {
	component, found := concrete.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Opacity), found
}
