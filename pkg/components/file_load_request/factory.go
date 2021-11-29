package file_load_request

import "github.com/gravestench/akara"

// ComponentFactory is a wrapper for the generic component factory
type ComponentFactory struct {
	*akara.ComponentFactory
}

// Add a FileLoadRequest for the given entity
func (f *ComponentFactory) Add(id akara.EID) *FileLoadRequest {
	return f.ComponentFactory.Add(id).(*FileLoadRequest)
}

// Get a FileLoadRequest for the given entity (can be nil), and a bool for whether it was found
func (f *ComponentFactory) Get(id akara.EID) (*FileLoadRequest, bool) {
	component, found := f.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*FileLoadRequest), found
}
