package texture

import "github.com/gravestench/akara"

// ComponentFactory is a wrapper for the generic component factory.
type ComponentFactory struct {
	*akara.ComponentFactory
}

// Add adds a Texture2D component to the given entity and returns it
func (m *ComponentFactory) Add(id akara.EID) *Texture2D {
	return m.ComponentFactory.Add(id).(*Texture2D)
}

// Get returns the Texture2D component for the given entity, and a bool for whether or not it exists
func (m *ComponentFactory) Get(id akara.EID) (*Texture2D, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Texture2D), found
}
