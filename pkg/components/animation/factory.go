package animation

import (
	"github.com/gravestench/akara"
)

// ComponentFactory is a wrapper for the generic component factory that returns Component component instances.
type ComponentFactory struct {
	*akara.ComponentFactory
}

// Add adds a Component component to the given entity and returns it
func (concrete *ComponentFactory) Add(id akara.EID) *Component {
	return concrete.ComponentFactory.Add(id).(*Component)
}

// Get returns the Component component for the given entity, and a bool for whether or not it exists
func (concrete *ComponentFactory) Get(id akara.EID) (*Component, bool) {
	component, found := concrete.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Component), found
}
