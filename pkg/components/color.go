package components

import (
	"image/color"

	"github.com/gravestench/akara"
)

var _ akara.Component = &Color{}

// Color is a wrapper component for a color.Color interface
type Color struct {
	color.Color
}

// New creates a new color component
func (*Color) New() akara.Component {
	return &Color{}
}

// ColorFactory is a wrapper for the generic component factory.
type ColorFactory struct {
	*akara.ComponentFactory
}

// Add a Color component to the given entity and return it
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
