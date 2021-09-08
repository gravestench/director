package components

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
)

var _ akara.Component = &Viewport{}

// Viewport represents a rendering viewport within a scene. This consists of a camera and background color.
type Viewport struct {
	rl.Camera2D
	Background color.Color
}

// New creates a new viewport
func (*Viewport) New() akara.Component {
	return &Viewport{
		Camera2D: rl.NewCamera2D(rl.Vector2{}, rl.Vector2{}, 0, 1),
		Background: color.Transparent,
	}
}

// ViewportFactory is a wrapper for the generic viewport component factory.
type ViewportFactory struct {
	*akara.ComponentFactory
}

// Add adds a Viewport component to the given entity and returns it
func (m *ViewportFactory) Add(id akara.EID) *Viewport {
	return m.ComponentFactory.Add(id).(*Viewport)
}

// Get returns the Viewport component for the given entity, and a bool for whether or not it exists
func (m *ViewportFactory) Get(id akara.EID) (*Viewport, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Viewport), found
}
