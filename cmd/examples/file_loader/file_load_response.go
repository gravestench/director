package main

import (
	"github.com/gravestench/akara"
	"io"
)

var _ akara.Component = &FileLoadResponse{}

type FileLoadResponse struct {
	io.ReadSeeker
}

func (*FileLoadResponse) New() akara.Component {
	return &FileLoadResponse{}
}

type FileStreamFactory struct {
	*akara.ComponentFactory
}

func (m *FileStreamFactory) Add(id akara.EID) *FileLoadResponse {
	return m.ComponentFactory.Add(id).(*FileLoadResponse)
}

func (m *FileStreamFactory) Get(id akara.EID) (*FileLoadResponse, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*FileLoadResponse), found
}


