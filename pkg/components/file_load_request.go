package components

import (
	"github.com/gravestench/akara"
)

var _ akara.Component = &FileLoadRequest{}

type FileLoadRequest struct {
	Path     string
	Attempts int
}

func (*FileLoadRequest) New() akara.Component {
	return &FileLoadRequest{}
}

type FileLoadRequestFactory struct {
	*akara.ComponentFactory
}

func (m *FileLoadRequestFactory) Add(id akara.EID) *FileLoadRequest {
	return m.ComponentFactory.Add(id).(*FileLoadRequest)
}

func (m *FileLoadRequestFactory) Get(id akara.EID) (*FileLoadRequest, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*FileLoadRequest), found
}
