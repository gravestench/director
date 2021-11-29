package viewport

import (
	"image/color"

	"github.com/gravestench/director/pkg/common/components"

	"github.com/gravestench/akara"
)

var _ akara.Component = &Viewport{}

// Viewport represents a rendering viewport within a scene. This consists of a camera and background color.
type Viewport struct {
	CameraEntity akara.EID
	Background   color.Color
}

// New creates a new viewport
func (*Viewport) New() akara.Component {
	return &Viewport{
		Background: color.Transparent,
	}
}

// static checks, should fail to compile if
// these interfaces can't be implemented
var (
	_ components.Component = &Component{}
	_ components.LuaExport = &ComponentFactory{}
)

type Component = Viewport // Component is an alias to Viewport
