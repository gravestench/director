package file_load_response

import "github.com/gravestench/akara"

// ComponentFactory is a wrapper for the generic component factory
type ComponentFactory struct {
	*akara.ComponentFactory
}

func (concrete *ComponentFactory) Add(id akara.EID) *FileLoadResponse {
	return concrete.ComponentFactory.Add(id).(*FileLoadResponse)
}

func (concrete *ComponentFactory) Get(id akara.EID) (*FileLoadResponse, bool) {
	component, found := concrete.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*FileLoadResponse), found
}
