package texture

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/common/components"
)

// Texture2D is a wrapper for a raylib texture
type Texture2D struct {
	*rl.Texture2D
}

// New creates a new raylib texture
func (*Texture2D) New() akara.Component {
	return &Texture2D{}
}

// static checks, should fail to compile if
// these interfaces can't be implemented
var (
	_ components.Component = &Component{}
	_ components.LuaExport = &ComponentFactory{}
)

type Component = Texture2D // Component is an alias to Texture2D
