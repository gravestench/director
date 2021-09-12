package components

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"image"
	"time"

	"github.com/gravestench/akara"
)

var _ akara.Component = &Animation{}

// Animation is a component that contains multiple images+textures,
// as well as frame durations and the current frame index that is being displayed
type Animation struct {
	FrameImages    []image.Image
	FrameTextures  []*rl.Texture2D
	FrameDurations []time.Duration
	CurrentFrame   int
	UntilNextFrame time.Duration
}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
func (*Animation) New() akara.Component {
	return &Animation{
		FrameImages:    make([]image.Image, 0),
		FrameTextures:  make([]*rl.Texture2D, 0),
		FrameDurations: make([]time.Duration, 0),
	}
}

// AnimationFactory is a wrapper for the generic component factory that returns Animation component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Animation.
type AnimationFactory struct {
	*akara.ComponentFactory
}

// Add adds a Animation component to the given entity and returns it
func (m *AnimationFactory) Add(id akara.EID) *Animation {
	return m.ComponentFactory.Add(id).(*Animation)
}

// Get returns the Animation component for the given entity, and a bool for whether or not it exists
func (m *AnimationFactory) Get(id akara.EID) (*Animation, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Animation), found
}
