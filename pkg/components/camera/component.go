package camera

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/common/components"
)

// Camera represents a camera used for rendering within a scene.
type Camera struct {
	rl.Camera2D
}

// New creates a new viewport
func (*Camera) New() akara.Component {
	return &Camera{}
}

// static checks, should fail to compile if
// these interfaces can't be implemented
var (
	_ components.Component = &Component{}
	_ components.LuaExport = &ComponentFactory{}
)

type Component = Camera // Component is an alias to Camera
