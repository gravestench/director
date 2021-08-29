package components

import (
	"github.com/gravestench/akara"
)

// static check that Font implements Component
var _ akara.Component = &Font{}

const (
	defaultFontFace = "Sans"
	defaultFontSize = 60 // px ?
)

// Font is a component that contains normalized alpha transparency (0.0 ... 1.0)
type Font struct {
	Face string
	Size int
}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
func (*Font) New() akara.Component {
	return &Font{
		Face: defaultFontFace,
		Size: defaultFontSize,
	}
}

// FontFactory is a wrapper for the generic component factory that returns Font component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Font.
type FontFactory struct {
	*akara.ComponentFactory
}

// Add adds a Font component to the given entity and returns it
func (m *FontFactory) Add(id akara.EID) *Font {
	return m.ComponentFactory.Add(id).(*Font)
}

// Get returns the Font component for the given entity, and a bool for whether or not it exists
func (m *FontFactory) Get(id akara.EID) (*Font, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Font), found
}
