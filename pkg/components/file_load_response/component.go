package file_load_response

import (
	"io"

	"github.com/gravestench/director/pkg/common/components"

	"github.com/gravestench/akara"
)

// FileLoadResponse represents the response to a request for loading a file
type FileLoadResponse struct {
	Stream io.ReadSeeker
}

// New creates a new FileLoadResponse
func (*FileLoadResponse) New() akara.Component {
	return &FileLoadResponse{}
}

// static checks, should fail to compile if
// these interfaces can't be implemented
var (
	_ components.Component = &Component{}
	_ components.LuaExport = &ComponentFactory{}
)

type Component = FileLoadResponse // Component is an alias to FileLoadResponse
