package components

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
)

// Camera represents a rendering viewport within a scene. This consists of a camera and background color.
type Camera struct {
	rl.Camera2D
}

// New creates a new viewport
func (*Camera) New() akara.Component {
	return &Camera{}
}

// CameraFactory is a wrapper for the generic viewport component factory.
type CameraFactory struct {
	*akara.ComponentFactory
}

// Add adds a Camera component to the given entity and returns it
func (m *CameraFactory) Add(id akara.EID) *Camera {
	return m.ComponentFactory.Add(id).(*Camera)
}

// Get returns the Camera component for the given entity, and a bool for whether or not it exists
func (m *CameraFactory) Get(id akara.EID) (*Camera, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Camera), found
}
