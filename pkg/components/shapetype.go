package components

import (
"github.com/gravestench/akara"
)

// static check that Shape implements Component
var _ akara.Component = &Shape{}

type ShapeType int

// Shape is a component that contains normalized alpha transparency (0.0 ... 1.0)
type Shape struct {
	Value float64
}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
func (*Shape) New() akara.Component {
	const defaultShape = 1.0

	return &Shape{
		Value: defaultShape,
	}
}

// Factory is a wrapper for the generic component factory that returns Shape component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Shape.
type Factory struct {
	*akara.ComponentFactory
}

// Add adds a Shape component to the given entity and returns it
func (m *Factory) Add(id akara.EID) *Shape {
	return m.ComponentFactory.Add(id).(*Shape)
}

// Get returns the Shape component for the given entity, and a bool for whether or not it exists
func (m *Factory) Get(id akara.EID) (*Shape, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Shape), found
}

