package size

import (
	"image"

	"github.com/gravestench/director/pkg/common/components"

	"github.com/gravestench/akara"
)

// Size is a component that contains normalized alpha transparency (0.0 ... 1.0)
type Size struct {
	image.Rectangle
}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
func (*Size) New() akara.Component {
	return &Size{}
}

// static checks, should fail to compile if
// these interfaces can't be implemented
var (
	_ components.Component = &Component{}
	_ components.LuaExport = &ComponentFactory{}
)

type Component = Size // Component is an alias to Size
