package font

import (
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/common/components"
)

const (
	defaultFontFace = "Sans"
	defaultFontSize = 60 // px ?
)

// Font is a component that contains a font name as a string, and a font size in pixels.
//
// The font name is used for resolving a font file on the host machine.
type Font struct {
	Face string
	Size int
}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
func (*Font) New() akara.Component {
	return &Font{
		Face: defaultFontFace,
		Size: defaultFontSize,
	}
}

// static checks, should fail to compile if
// these interfaces can't be implemented
var (
	_ components.Component = &Component{}
	_ components.LuaExport = &ComponentFactory{}
)

type Component = Font // Component is an alias to Font
