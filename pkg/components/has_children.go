package components

import (
	"github.com/gravestench/akara"
)

var _ akara.Component = &HasChildren{}

// HasChildren represents an entity's ability to reference children in a parent-child relationship
type HasChildren struct {
	Children []akara.EID
}

// New creates a new color component
func (*HasChildren) New() akara.Component {
	return &HasChildren{
		Children: make([]akara.EID, 0),
	}
}

// HasChildrenFactory is a wrapper for the generic component factory.
type HasChildrenFactory struct {
	*akara.ComponentFactory
}

// Add a HasChildren component to the given entity and return it
func (m *HasChildrenFactory) Add(id akara.EID) *HasChildren {
	return m.ComponentFactory.Add(id).(*HasChildren)
}

// Get returns the HasChildren component for the given entity, and a bool for whether or not it exists
func (m *HasChildrenFactory) Get(id akara.EID) (*HasChildren, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*HasChildren), found
}
