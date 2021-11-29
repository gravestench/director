package opacity

import (
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/common/components"
)

// Opacity is a component that contains normalized alpha transparency (0.0 ... 1.0)
type Opacity struct {
	Value float64
}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
func (*Opacity) New() akara.Component {
	const defaultAlpha = 1.0

	return &Opacity{
		Value: defaultAlpha,
	}
}

// static checks, should fail to compile if
// these interfaces can't be implemented
var (
	_ components.Component = &Component{}
	_ components.LuaExport = &ComponentFactory{}
)

type Component = Opacity // Component is an alias to Opacity
