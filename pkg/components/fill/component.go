package fill

import (
	"image/color"

	"github.com/gravestench/director/pkg/common/components"

	"github.com/gravestench/akara"
)

// Fill is a component that contains a fill color
type Fill struct {
	color.Color
}

// New creates a new alpha component instance. The default color is fully opaque black.
func (*Fill) New() akara.Component {
	return &Fill{
		Color: color.RGBA{A: 255},
	}
}

// static checks, should fail to compile if
// these interfaces can't be implemented
var (
	_ components.Component = &Component{}
	_ components.LuaExport = &ComponentFactory{}
)

type Component = Fill // Component is an alias to Fill
