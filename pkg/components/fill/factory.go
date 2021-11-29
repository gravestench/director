package fill

import "github.com/gravestench/akara"

// ComponentFactory is a wrapper for the generic component factory that returns Fill component instances.
type ComponentFactory struct {
	*akara.ComponentFactory
}

// Add adds a Fill component to the given entity and returns it
func (concrete *ComponentFactory) Add(id akara.EID) *Fill {
	return concrete.ComponentFactory.Add(id).(*Fill)
}

// Get returns the Fill component for the given entity, and a bool for whether or not it exists
func (concrete *ComponentFactory) Get(id akara.EID) (*Fill, bool) {
	component, found := concrete.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Fill), found
}
