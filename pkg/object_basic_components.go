package pkg

import (
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/components"
)

type basicComponents struct {
	Camera2D        components.Camera2DFactory
	Color           components.ColorFactory
	Font            components.FontFactory
	RenderTexture2D components.RenderTexture2DFactory
	SceneGraphNode  components.SceneGraphNodeFactory
	Text            components.TextFactory
	Texture2D       components.Texture2DFactory
	Transform       components.TransformFactory
	UUID            components.UUIDFactory
	Vector2         components.Vector2Factory
	Vector3         components.Vector3Factory
	Vector4         components.Vector4Factory
}

func (bc *basicComponents) init(w *akara.World) {
	injectComponent(w, &components.Camera2D{}, &bc.Camera2D.ComponentFactory)
	injectComponent(w, &components.Color{}, &bc.Color.ComponentFactory)
	injectComponent(w, &components.Font{}, &bc.Font.ComponentFactory)
	injectComponent(w, &components.Vector2{}, &bc.Vector2.ComponentFactory)
	injectComponent(w, &components.Vector3{}, &bc.Vector3.ComponentFactory)
	injectComponent(w, &components.Vector4{}, &bc.Vector4.ComponentFactory)
	injectComponent(w, &components.SceneGraphNode{}, &bc.SceneGraphNode.ComponentFactory)
	injectComponent(w, &components.Text{}, &bc.Text.ComponentFactory)
	injectComponent(w, &components.RenderTexture2D{}, &bc.RenderTexture2D.ComponentFactory)
	injectComponent(w, &components.Texture2D{}, &bc.Texture2D.ComponentFactory)
	injectComponent(w, &components.Transform{}, &bc.Transform.ComponentFactory)
	injectComponent(w, &components.UUID{}, &bc.UUID.ComponentFactory)
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

	if bc.Texture2D.ComponentFactory == nil {
		return false
	}

	if bc.RenderTexture2D.ComponentFactory == nil {
		return false
	}

	if bc.Transform.ComponentFactory == nil {
		return false
	}

	if bc.Color.ComponentFactory == nil {
		return false
	}

	if bc.Font.ComponentFactory == nil {
		return false
	}

	if bc.Camera2D.ComponentFactory == nil {
		return false
	}

	if bc.UUID.ComponentFactory == nil {
		return false
	}

	return true
}

func injectComponent(w *akara.World, c akara.Component, cf **akara.ComponentFactory) {
	cid := w.RegisterComponent(c)
	*cf = w.GetComponentFactory(cid)
}
