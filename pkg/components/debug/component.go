package debug

import (
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/common/components"
)

// Debug is a tag component, used as a flag for debugging an entity.
type Debug struct{}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
func (*Debug) New() akara.Component {
	return &Debug{}
}

// static checks, should fail to compile if
// these interfaces can't be implemented
var (
	_ components.Component = &Component{}
	_ components.LuaExport = &ComponentFactory{}
)

type Component = Debug // Component is an alias to Debug
