package origin

import (
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/common/components"
	"github.com/gravestench/mathlib"
)

// Origin is a component that contains normalized alpha transparency (0.0 ... 1.0)
type Origin struct {
	*mathlib.Vector3
}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
func (*Origin) New() akara.Component {
	const defaultOrigin = 0.5 // normalized, 0.5 is center

	return &Origin{
		Vector3: &mathlib.Vector3{
			X: defaultOrigin,
			Y: defaultOrigin,
		},
	}
}

// static checks, should fail to compile if
// these interfaces can't be implemented
var (
	_ components.Component = &Component{}
	_ components.LuaExport = &ComponentFactory{}
)

type Component = Origin // Component is an alias to Origin
