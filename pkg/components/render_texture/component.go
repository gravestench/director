package render_texture

import (
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/common/components"

	rl "github.com/gen2brain/raylib-go/raylib"
)

// static check that RenderTexture2D implements Camera
var _ akara.Component = &RenderTexture2D{}

// RenderTexture2D is a wrapper for a raylib texture
type RenderTexture2D struct {
	*rl.RenderTexture2D
}

// New creates a new raylib texture
func (*RenderTexture2D) New() akara.Component {
	return &RenderTexture2D{}
}

// static checks, should fail to compile if
// these interfaces can't be implemented
var (
	_ components.Component = &Component{}
	_ components.LuaExport = &ComponentFactory{}
)

type Component = RenderTexture2D // Component is an alias to RenderTexture2D
