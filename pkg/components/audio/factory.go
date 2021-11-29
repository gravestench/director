package audio

import (
	"github.com/gravestench/akara"
)

// ComponentFactory is a wrapper for the generic component factory.
type ComponentFactory struct {
	*akara.ComponentFactory
}

// Add adds an Audio component to the given entity and returns it
func (m *ComponentFactory) Add(id akara.EID) *Audio {
	return m.ComponentFactory.Add(id).(*Audio)
}

// Get returns the Audio component for the given entity, and a bool for whether not it exists
func (m *ComponentFactory) Get(id akara.EID) (*Audio, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Audio), found
}
