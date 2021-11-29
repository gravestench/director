package viewport

import "github.com/gravestench/akara"

// ComponentFactory is a wrapper for the generic viewport component factory.
type ComponentFactory struct {
	*akara.ComponentFactory
}

// Add adds a Viewport component to the given entity and returns it
func (concrete *ComponentFactory) Add(id akara.EID) *Viewport {
	return concrete.ComponentFactory.Add(id).(*Viewport)
}

// Get returns the Viewport component for the given entity, and a bool for whether or not it exists
func (concrete *ComponentFactory) Get(id akara.EID) (*Viewport, bool) {
	component, found := concrete.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Viewport), found
}
