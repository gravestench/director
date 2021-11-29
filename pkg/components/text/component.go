package text

import (
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/common/components"
)

// static check that Text implements Camera
var _ akara.Component = &Text{}

// Text is a component that contains a text string
type Text struct {
	String string
}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
func (*Text) New() akara.Component {
	return &Text{}
}

// static checks, should fail to compile if
// these interfaces can't be implemented
var (
	_ components.Component = &Component{}
	_ components.LuaExport = &ComponentFactory{}
)

type Component = Text // Component is an alias to Text
