package components

import (
	"github.com/gravestench/akara"
	"image/color"
)

var _ akara.Component = &Stroke{}

// Stroke is a component that contains normalized alpha transparency (0.0 ... 1.0)
type Stroke struct {
	LineWidth int
	color.Color
}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
func (*Stroke) New() akara.Component {
	return &Stroke{}
}

// StrokeFactory is a wrapper for the generic component factory that returns Stroke component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Stroke.
type StrokeFactory struct {
	*akara.ComponentFactory
}

// Add adds a Stroke component to the given entity and returns it
func (m *StrokeFactory) Add(id akara.EID) *Stroke {
	return m.ComponentFactory.Add(id).(*Stroke)
}

// Get returns the Stroke component for the given entity, and a bool for whether or not it exists
func (m *StrokeFactory) Get(id akara.EID) (*Stroke, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Stroke), found
}
