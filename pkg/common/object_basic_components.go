package common

import (
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/components"
	"github.com/gravestench/director/pkg/systems/input"
)

type BasicComponents struct {
	Viewport         components.ViewportFactory
	Color            components.ColorFactory
	Debug            components.DebugFactory
	FileLoadRequest  components.FileLoadRequestFactory
	FileLoadResponse components.FileLoadResponseFactory
	FileType         components.FileTypeFactory
	Fill             components.FillFactory
	Animation        components.AnimationFactory
	Stroke           components.StrokeFactory
	Font             components.FontFactory
	Interactive      input.InteractiveFactory
	Opacity          components.OpacityFactory
	Origin           components.OriginFactory
	RenderTexture2D  components.RenderTexture2DFactory
	Size             components.SizeFactory
	SceneGraphNode   components.SceneGraphNodeFactory
	Text             components.TextFactory
	Texture2D        components.Texture2DFactory
	Transform        components.TransformFactory
	UUID             components.UUIDFactory
}

func (bc *BasicComponents) Init(w *akara.World) {
	injectComponent(w, &components.Viewport{}, &bc.Viewport.ComponentFactory)
	injectComponent(w, &components.Color{}, &bc.Color.ComponentFactory)
	injectComponent(w, &components.Debug{}, &bc.Debug.ComponentFactory)
	injectComponent(w, &components.FileLoadRequest{}, &bc.FileLoadRequest.ComponentFactory)
	injectComponent(w, &components.FileLoadResponse{}, &bc.FileLoadResponse.ComponentFactory)
	injectComponent(w, &components.FileType{}, &bc.FileType.ComponentFactory)
	injectComponent(w, &components.Fill{}, &bc.Fill.ComponentFactory)
	injectComponent(w, &components.Animation{}, &bc.Animation.ComponentFactory)
	injectComponent(w, &components.Origin{}, &bc.Origin.ComponentFactory)
	injectComponent(w, &input.Interactive{}, &bc.Interactive.ComponentFactory)
	injectComponent(w, &components.Opacity{}, &bc.Opacity.ComponentFactory)
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

	if bc.FileType.ComponentFactory == nil {
		return false
	}

	if bc.FileLoadRequest.ComponentFactory == nil {
		return false
	}

	if bc.FileLoadResponse.ComponentFactory == nil {
		return false
	}

	if bc.Fill.ComponentFactory == nil {
		return false
	}

	if bc.Animation.ComponentFactory == nil {
		return false
	}

	if bc.Origin.ComponentFactory == nil {
		return false
	}

	if bc.Opacity.ComponentFactory == nil {
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

	if bc.Viewport.ComponentFactory == nil {
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
