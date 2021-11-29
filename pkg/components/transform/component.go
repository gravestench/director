package transform

import (
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/common/components"
	"github.com/gravestench/mathlib"
)

// Transform is a component that contains a matrix4 that expresses
// translation, rotation, and scale.
type Transform struct {
	Translation, Rotation, Scale *mathlib.Vector3
}

// GetMatrix returns a matrix4, composed from the translation, rotation, and scale of this transform
func (component *Transform) GetMatrix() *mathlib.Matrix4 {
	return mathlib.NewMatrix4(nil).
		Translate(component.Translation).
		RotateX(component.Rotation.X).
		RotateY(component.Rotation.Y).
		RotateZ(component.Rotation.Z).
		ScaleXYZ(component.Scale.XYZ())
}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
func (*Transform) New() akara.Component {
	return &Transform{
		Translation: mathlib.NewVector3(0, 0, 0),
		Rotation:    mathlib.NewVector3(0, 0, 0),
		Scale:       mathlib.NewVector3(1, 1, 1),
	}
}

// static checks, should fail to compile if
// these interfaces can't be implemented
var (
	_ components.Component = &Component{}
	_ components.LuaExport = &ComponentFactory{}
)

type Component = Transform // Component is an alias to Transform
