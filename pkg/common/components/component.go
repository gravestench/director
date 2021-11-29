package components

import (
	"github.com/gravestench/akara"
)

// Component to director is an akara component that can be exported to the lua state machine
type Component interface {
	akara.Component
	LuaExport
}
