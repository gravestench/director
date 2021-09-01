package common

import (
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/components"
)

type BasicComponents struct {
	Camera2D        components.Camera2DFactory
	Color           components.ColorFactory
	Debug           components.DebugFactory
	Fill            components.FillFactory
	Stroke          components.StrokeFactory
	Font            components.FontFactory
	Origin          components.OriginFactory
	RenderTexture2D components.RenderTexture2DFactory
	Size            components.SizeFactory
	SceneGraphNode  components.SceneGraphNodeFactory
	Text            components.TextFactory
	Texture2D       components.Texture2DFactory
	Transform       components.TransformFactory
	UUID            components.UUIDFactory
}

func (bc *BasicComponents) Init(w *akara.World) {
	injectComponent(w, &components.Camera2D{}, &bc.Camera2D.ComponentFactory)
	injectComponent(w, &components.Color{}, &bc.Color.ComponentFactory)
	injectComponent(w, &components.Debug{}, &bc.Debug.ComponentFactory)
	injectComponent(w, &components.Fill{}, &bc.Fill.ComponentFactory)
	injectComponent(w, &components.Origin{}, &bc.Origin.ComponentFactory)
	injectComponent(w, &components.Stroke{}, &bc.Stroke.ComponentFactory)
	injectComponent(w, &components.Font{}, &bc.Font.ComponentFactory)
	injectComponent(w, &components.SceneGraphNode{}, &bc.SceneGraphNode.ComponentFactory)
	injectComponent(w, &components.Text{}, &bc.Text.ComponentFactory)
	injectComponent(w, &components.RenderTexture2D{}, &bc.RenderTexture2D.ComponentFactory)
	injectComponent(w, &components.Size{}, &bc.Size.ComponentFactory)
	injectComponent(w, &components.Texture2D{}, &bc.Texture2D.ComponentFactory)
	injectComponent(w, &components.Transform{}, &bc.Transform.ComponentFactory)
	injectComponent(w, &components.UUID{}, &bc.UUID.ComponentFactory)
}

func (bc *BasicComponents) IsInit() bool {
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

	if bc.Fill.ComponentFactory == nil {
		return false
	}

	if bc.Origin.ComponentFactory == nil {
		return false
	}

	if bc.Debug.ComponentFactory == nil {
		return false
	}

	if bc.Stroke.ComponentFactory == nil {
		return false
	}

	if bc.Font.ComponentFactory == nil {
		return false
	}

	if bc.Size.ComponentFactory == nil {
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
