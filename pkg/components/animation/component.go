package animation

import (
	"image"
	"time"

	"github.com/gravestench/director/pkg/common/components"

	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/gravestench/akara"
)

// Animation is a component that contains multiple images+textures,
// as well as frame durations and the current frame index that is being displayed
type Animation struct {
	FrameImages    []image.Image
	FrameTextures  []*rl.Texture2D
	FrameDurations []time.Duration
	CurrentFrame   int
	UntilNextFrame time.Duration
}

// New creates a new animation component instance.
func (*Animation) New() akara.Component {
	return &Animation{
		FrameImages:    make([]image.Image, 0),
		FrameTextures:  make([]*rl.Texture2D, 0),
		FrameDurations: make([]time.Duration, 0),
	}
}

// static checks, should fail to compile if
// these interfaces can't be implemented
var (
	_ components.Component = &Component{}
	_ components.LuaExport = &ComponentFactory{}
)

type Component = Animation // Component is an alias to Animation
