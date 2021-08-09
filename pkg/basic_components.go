package pkg

import (
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/components"
)

type basicComponents struct {
	Color components.ColorFactory
	Vector2 components.Vector2Factory
	Vector3 components.Vector3Factory
	Vector4 components.Vector4Factory
	Text components.TextFactory
}

func (bc *basicComponents) init(w *akara.World) {
	injectComponent(w, &components.Vector2{}, &bc.Vector2.ComponentFactory)
	injectComponent(w, &components.Vector3{}, &bc.Vector3.ComponentFactory)
	injectComponent(w, &components.Vector4{}, &bc.Vector4.ComponentFactory)
	injectComponent(w, &components.Text{}, &bc.Text.ComponentFactory)
	injectComponent(w, &components.Color{}, &bc.Color.ComponentFactory)
}

func (bc *basicComponents) isInit() bool {
	if bc.Vector2.ComponentFactory == nil {
		return false
	}

	if bc.Vector3.ComponentFactory == nil {
		return false
	}

	if bc.Vector4.ComponentFactory == nil {
		return false
	}

	if bc.Text.ComponentFactory == nil {
		return false
	}

	if bc.Color.ComponentFactory == nil {
		return false
	}

	return true
}

func injectComponent(w *akara.World, c akara.Component, cf **akara.ComponentFactory) {
	cid := w.RegisterComponent(c)
	*cf = w.GetComponentFactory(cid)
}
