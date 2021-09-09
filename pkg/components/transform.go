package components

import (
	"github.com/gravestench/akara"
	"github.com/gravestench/mathlib"
)

// static check that Transform implements Component
var _ akara.Component = &Transform{}

// Transform is a component that contains normalized alpha transparency (0.0 ... 1.0)
type Transform struct {
	Translation, Rotation, Scale *mathlib.Vector3
}

func (t *Transform) GetMatrix() *mathlib.Matrix4 {
	return mathlib.NewMatrix4(nil).
		Translate(t.Translation).
		RotateX(t.Rotation.X).
		RotateY(t.Rotation.Y).
		RotateZ(t.Rotation.Z).
		ScaleXYZ(t.Scale.XYZ())
}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
func (*Transform) New() akara.Component {
	return &Transform{
		Translation: mathlib.NewVector3(0, 0, 0),
		Rotation:    mathlib.NewVector3(0, 0, 0),
		Scale:       mathlib.NewVector3(1, 1, 1),
	}
}

// TransformFactory is a wrapper for the generic component factory that returns Transform component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Transform.
type TransformFactory struct {
	*akara.ComponentFactory
}

// Add adds a Transform component to the given entity and returns it
func (m *TransformFactory) Add(id akara.EID) *Transform {
	return m.ComponentFactory.Add(id).(*Transform)
}

// Get returns the Transform component for the given entity, and a bool for whether or not it exists
func (m *TransformFactory) Get(id akara.EID) (*Transform, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Transform), found
}
