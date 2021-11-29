package uuid

import "github.com/gravestench/akara"

// ComponentFactory is a wrapper for the generic component factory that returns UUID component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a UUID.
type ComponentFactory struct {
	*akara.ComponentFactory
}

// Add adds a UUID component to the given entity and returns it
func (concrete *ComponentFactory) Add(id akara.EID) *UUID {
	return concrete.ComponentFactory.Add(id).(*UUID)
}

// Get returns the UUID component for the given entity, and a bool for whether or not it exists
func (concrete *ComponentFactory) Get(id akara.EID) (*UUID, bool) {
	component, found := concrete.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*UUID), found
}
