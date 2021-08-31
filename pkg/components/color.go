package components

import (
	"github.com/gravestench/akara"
	"image/color"
)

var _ akara.Component = &Color{}

// Color is a component that contains normalized alpha transparency (0.0 ... 1.0)
type Color struct {
	color.Color
}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
func (*Color) New() akara.Component {
	return &Color{}
}

// ColorFactory is a wrapper for the generic component factory that returns Color component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Color.
type ColorFactory struct {
	*akara.ComponentFactory
}

// Add adds a Color component to the given entity and returns it
func (m *ColorFactory) Add(id akara.EID) *Color {
	return m.ComponentFactory.Add(id).(*Color)
}

// Get returns the Color component for the given entity, and a bool for whether or not it exists
func (m *ColorFactory) Get(id akara.EID) (*Color, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Color), found
}
