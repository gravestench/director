package components

import (
	"github.com/gravestench/akara"
)

// static check that Text implements Component
var _ akara.Component = &Text{}

// Text is a component that contains a text string
type Text struct {
	String string
}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
func (*Text) New() akara.Component {
	return &Text{}
}

// TextFactory is a wrapper for the generic component factory that returns Text component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Text.
type TextFactory struct {
	*akara.ComponentFactory
}

// Add adds a Text component to the given entity and returns it
func (m *TextFactory) Add(id akara.EID) *Text {
	return m.ComponentFactory.Add(id).(*Text)
}

// Get returns the Text component for the given entity, and a bool for whether or not it exists
func (m *TextFactory) Get(id akara.EID) (*Text, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Text), found
}
