package color

import (
	"image/color"

	"github.com/gravestench/director/pkg/common/components"

	"github.com/gravestench/akara"
)

// Color is a wrapper component for a color.Color interface
type Color struct {
	color.Color
}

// New creates a new color component
func (*Color) New() akara.Component {
	return &Color{}
}

// static checks, should fail to compile if
// these interfaces can't be implemented
var (
	_ components.Component = &Component{}
	_ components.LuaExport = &ComponentFactory{}
)

type Component = Color // Component is an alias to Color
