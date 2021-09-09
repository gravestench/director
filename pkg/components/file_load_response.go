package components

import (
	"github.com/gravestench/akara"
	"io"
)

var _ akara.Component = &FileLoadResponse{}

type FileLoadResponse struct {
	Stream io.ReadSeeker
}

func (*FileLoadResponse) New() akara.Component {
	return &FileLoadResponse{}
}

type FileLoadResponseFactory struct {
	*akara.ComponentFactory
}

func (m *FileLoadResponseFactory) Add(id akara.EID) *FileLoadResponse {
	return m.ComponentFactory.Add(id).(*FileLoadResponse)
}

func (m *FileLoadResponseFactory) Get(id akara.EID) (*FileLoadResponse, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*FileLoadResponse), found
}
