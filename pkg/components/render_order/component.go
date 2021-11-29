package render_order

import (
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/common/components"
)

var _ akara.Component = &RenderOrder{}

// RenderOrder is a component that contains normalized alpha transparency (0.0 ... 1.0)
type RenderOrder struct {
	Value int
}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
func (*RenderOrder) New() akara.Component {
	return &RenderOrder{}
}

// static checks, should fail to compile if
// these interfaces can't be implemented
var (
	_ components.Component = &Component{}
	_ components.LuaExport = &ComponentFactory{}
)

type Component = RenderOrder // Component is an alias to RenderOrder
