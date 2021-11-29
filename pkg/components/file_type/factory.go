package file_type

import "github.com/gravestench/akara"

// ComponentFactory is a wrapper for the generic component factory.
type ComponentFactory struct {
	*akara.ComponentFactory
}

// Add adds a FileType component to the given entity and returns it
func (concrete *ComponentFactory) Add(id akara.EID) *FileType {
	return concrete.ComponentFactory.Add(id).(*FileType)
}

// Get returns the FileType component for the given entity, and a bool for whether or not it exists
func (concrete *ComponentFactory) Get(id akara.EID) (*FileType, bool) {
	component, found := concrete.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*FileType), found
}
