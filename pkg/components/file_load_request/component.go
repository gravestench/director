package file_load_request

import (
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/common/components"
)

// FileLoadRequest represents a request for a file to be loaded.
// When the file is loaded, a FileLoadResponse should be created for the entity (by a system...)
type FileLoadRequest struct {
	Path     string
	Attempts int
}

// New creates a new FileLoadRequest component
func (*FileLoadRequest) New() akara.Component {
	return &FileLoadRequest{}
}

// static checks, should fail to compile if
// these interfaces can't be implemented
var (
	_ components.Component = &Component{}
	_ components.LuaExport = &ComponentFactory{}
)

type Component = FileLoadRequest // Component is an alias to FileLoadRequest
