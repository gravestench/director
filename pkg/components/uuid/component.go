package uuid

import (
	"github.com/google/uuid"
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/common/components"
)

// UUID is a unique identifier
type UUID struct {
	uuid.UUID
}

// New creates a new uuid instance
func (*UUID) New() akara.Component {
	return &UUID{
		UUID: uuid.New(),
	}
}

// static checks, should fail to compile if
// these interfaces can't be implemented
var (
	_ components.Component = &Component{}
	_ components.LuaExport = &ComponentFactory{}
)

type Component = UUID // Component is an alias to UUID
