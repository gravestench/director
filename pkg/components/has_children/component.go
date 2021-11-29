package has_children

import (
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/common/components"
)

// HasChildren represents an entity's ability to reference children in a parent-child relationship
type HasChildren struct {
	Children []akara.EID
}

// New creates a new color component
func (*HasChildren) New() akara.Component {
	return &HasChildren{
		Children: make([]akara.EID, 0),
	}
}

// static checks, should fail to compile if
// these interfaces can't be implemented
var (
	_ components.Component = &Component{}
	_ components.LuaExport = &ComponentFactory{}
)

type Component = HasChildren // Component is an alias to HasChildren
