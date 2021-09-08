package components

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"

	. "github.com/gravestench/director/pkg/common"
)

var _ akara.Component = &Viewport{}

// Viewport is a component that represents a camera within a scene. A camera can have have a background color
type Viewport struct {
	rl.Camera2D
	Background color.Color
}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
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
func (m *ViewportFactory) Add(id Entity) *Viewport {
	return m.ComponentFactory.Add(id).(*Viewport)
}

// Get returns the Viewport component for the given entity, and a bool for whether or not it exists
func (m *ViewportFactory) Get(id Entity) (*Viewport, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Viewport), found
}
