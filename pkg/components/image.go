package components

import (
	"github.com/gravestench/akara"
	"image"
)

var _ akara.Component = &Image{}

// Image is a component that contains normalized alpha transparency (0.0 ... 1.0)
type Image struct {
	image.Image
}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
func (*Image) New() akara.Component {
	return &Image{}
}

// ImageFactory is a wrapper for the generic component factory that returns Image component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Image.
type ImageFactory struct {
	*akara.ComponentFactory
}

// Add adds a Image component to the given entity and returns it
func (m *ImageFactory) Add(id akara.EID) *Image {
	return m.ComponentFactory.Add(id).(*Image)
}

// Get returns the Image component for the given entity, and a bool for whether or not it exists
func (m *ImageFactory) Get(id akara.EID) (*Image, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Image), found
}



