package components

import (
	"github.com/gravestench/akara"
	"net/http"
)

var _ akara.Component = &FileType{}

// FileType is a component that contains normalized alpha transparency (0.0 ... 1.0)
type FileType struct {
	Type string
}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
func (*FileType) New() akara.Component {
	return &FileType{
		Type: http.DetectContentType(nil),
	}
}

// FileTypeFactory is a wrapper for the generic component factory that returns FileType component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a FileType.
type FileTypeFactory struct {
	*akara.ComponentFactory
}

// Add adds a FileType component to the given entity and returns it
func (m *FileTypeFactory) Add(id akara.EID) *FileType {
	return m.ComponentFactory.Add(id).(*FileType)
}

// Get returns the FileType component for the given entity, and a bool for whether or not it exists
func (m *FileTypeFactory) Get(id akara.EID) (*FileType, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*FileType), found
}

