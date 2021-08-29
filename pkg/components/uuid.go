package components

import (
	"github.com/google/uuid"
	"github.com/gravestench/akara"
)

// static check that UUID implements Component
var _ akara.Component = &UUID{}

// UUID is a unique identifier
type UUID struct {
	uuid.UUID
}

// New creates a new uuid instance
func (*UUID) New() akara.Component {
	return &UUID{
		UUID: uuid.New(),
	}
}

// UUIDFactory is a wrapper for the generic component factory that returns UUID component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a UUID.
type UUIDFactory struct {
	*akara.ComponentFactory
}

// Add adds a UUID component to the given entity and returns it
func (m *UUIDFactory) Add(id akara.EID) *UUID {
	return m.ComponentFactory.Add(id).(*UUID)
}

// Get returns the UUID component for the given entity, and a bool for whether or not it exists
func (m *UUIDFactory) Get(id akara.EID) (*UUID, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*UUID), found
}
