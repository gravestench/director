package file_type

import (
	"net/http"

	"github.com/gravestench/director/pkg/common/components"

	"github.com/gravestench/akara"
)

// FileType is a component that contains a string descriptor of the detected file type
type FileType struct {
	Type string
}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
func (*FileType) New() akara.Component {
	return &FileType{
		Type: http.DetectContentType(nil),
	}
}

// static checks, should fail to compile if
// these interfaces can't be implemented
var (
	_ components.Component = &Component{}
	_ components.LuaExport = &ComponentFactory{}
)

type Component = FileType // Component is an alias to FileType
