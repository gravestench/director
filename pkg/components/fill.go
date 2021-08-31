package components

import (
	"github.com/gravestench/akara"
	"image/color"
)

var _ akara.Component = &Fill{}

// Fill is a component that contains normalized alpha transparency (0.0 ... 1.0)
type Fill struct {
	color.Color
}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
func (*Fill) New() akara.Component {
	return &Fill{}
}

// FillFactory is a wrapper for the generic component factory that returns Fill component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Fill.
type FillFactory struct {
	*akara.ComponentFactory
}

// Add adds a Fill component to the given entity and returns it
func (m *FillFactory) Add(id akara.EID) *Fill {
	return m.ComponentFactory.Add(id).(*Fill)
}

// Get returns the Fill component for the given entity, and a bool for whether or not it exists
func (m *FillFactory) Get(id akara.EID) (*Fill, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Fill), found
}
