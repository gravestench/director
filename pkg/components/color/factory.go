package color

import "github.com/gravestench/akara"

// ComponentFactory is a wrapper for the generic component factory.
type ComponentFactory struct {
	*akara.ComponentFactory
}

// Add a Color component to the given entity and return it
func (concrete *ComponentFactory) Add(id akara.EID) *Color {
	return concrete.ComponentFactory.Add(id).(*Color)
}

// Get returns the Color component for the given entity, and a bool for whether or not it exists
func (concrete *ComponentFactory) Get(id akara.EID) (*Color, bool) {
	component, found := concrete.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Color), found
}
