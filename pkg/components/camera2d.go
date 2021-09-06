package components

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
	"image/color"
)

var _ akara.Component = &Camera2D{}

// Camera2D is a component that contains normalized alpha transparency (0.0 ... 1.0)
type Camera2D struct {
	rl.Camera2D
	Background color.Color
}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
func (*Camera2D) New() akara.Component {
	return &Camera2D{
		Camera2D: rl.NewCamera2D(rl.Vector2{}, rl.Vector2{}, 0, 1),
		Background: color.Transparent,
	}
}

// Camera2DFactory is a wrapper for the generic component factory that returns
// Camera2D component instances. This can be embedded inside of a system to give
// them the methods for adding, retrieving, and removing a Camera2D.
type Camera2DFactory struct {
	*akara.ComponentFactory
}

// Add adds a Camera2D component to the given entity and returns it
func (m *Camera2DFactory) Add(id akara.EID) *Camera2D {
	return m.ComponentFactory.Add(id).(*Camera2D)
}

// Get returns the Camera2D component for the given entity, and a bool for whether or not it exists
func (m *Camera2DFactory) Get(id akara.EID) (*Camera2D, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Camera2D), found
}
