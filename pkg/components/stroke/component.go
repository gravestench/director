package stroke

import (
	"image/color"

	"github.com/gravestench/director/pkg/common/components"

	"github.com/gravestench/akara"
)

var _ akara.Component = &Stroke{}

// Stroke is a component that contains normalized alpha transparency (0.0 ... 1.0)
type Stroke struct {
	LineWidth int
	color.Color
}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
func (*Stroke) New() akara.Component {
	return &Stroke{}
}

// static checks, should fail to compile if
// these interfaces can't be implemented
var (
	_ components.Component = &Component{}
	_ components.LuaExport = &ComponentFactory{}
)

type Component = Stroke // Component is an alias to Stroke
