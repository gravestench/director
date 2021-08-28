package main

import (
	"github.com/gravestench/akara"
)

var _ akara.Component = &FileLoadRequest{}

type FileLoadRequest struct {
	Path string
	Attempts int
}

func (*FileLoadRequest) New() akara.Component {
	return &FileLoadRequest{}
}

type LoadRequestFactory struct {
	*akara.ComponentFactory
}

func (m *LoadRequestFactory) Add(id akara.EID) *FileLoadRequest {
	return m.ComponentFactory.Add(id).(*FileLoadRequest)
}

func (m *LoadRequestFactory) Get(id akara.EID) (*FileLoadRequest, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*FileLoadRequest), found
}


