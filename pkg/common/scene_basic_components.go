package common

import (
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/components"
	"github.com/gravestench/director/pkg/systems/audio"
	"github.com/gravestench/director/pkg/systems/input"
)

// BasicComponents represents components that every scene has available
type BasicComponents struct {
	Viewport         components.ViewportFactory
	Camera           components.CameraFactory
	Color            components.ColorFactory
	Debug            components.DebugFactory
	FileLoadRequest  components.FileLoadRequestFactory
	FileLoadResponse components.FileLoadResponseFactory
	FileType         components.FileTypeFactory
	Fill             components.FillFactory
	HasChildren      components.HasChildrenFactory
	Animation        components.AnimationFactory
	Stroke           components.StrokeFactory
	Font             components.FontFactory
	Interactive      input.InteractiveFactory
	Opacity          components.OpacityFactory
	Origin           components.OriginFactory
	RenderTexture2D  components.RenderTexture2DFactory
	RenderOrder      components.RenderOrderFactory
	Size             components.SizeFactory
	SceneGraphNode   components.SceneGraphNodeFactory
	Text             components.TextFactory
	Texture2D        components.Texture2DFactory
	Transform        components.TransformFactory
	UUID             components.UUIDFactory
	Audible          audio.AudibleFactory
}

// Init initializes each component factory for the given world,
// putting the generic component factory inside of the concrete component factory
func (bc *BasicComponents) Init(w *akara.World) {
	injectComponent(w, &components.Viewport{}, &bc.Viewport.ComponentFactory)
	injectComponent(w, &components.Camera{}, &bc.Camera.ComponentFactory)
	injectComponent(w, &components.Color{}, &bc.Color.ComponentFactory)
	injectComponent(w, &components.Debug{}, &bc.Debug.ComponentFactory)
	injectComponent(w, &components.FileLoadRequest{}, &bc.FileLoadRequest.ComponentFactory)
	injectComponent(w, &components.FileLoadResponse{}, &bc.FileLoadResponse.ComponentFactory)
	injectComponent(w, &components.FileType{}, &bc.FileType.ComponentFactory)
	injectComponent(w, &components.Fill{}, &bc.Fill.ComponentFactory)
	injectComponent(w, &components.HasChildren{}, &bc.HasChildren.ComponentFactory)
	injectComponent(w, &components.Animation{}, &bc.Animation.ComponentFactory)
	injectComponent(w, &components.Origin{}, &bc.Origin.ComponentFactory)
	injectComponent(w, &input.Interactive{}, &bc.Interactive.ComponentFactory)
	injectComponent(w, &components.Opacity{}, &bc.Opacity.ComponentFactory)
	injectComponent(w, &components.Stroke{}, &bc.Stroke.ComponentFactory)
	injectComponent(w, &components.Font{}, &bc.Font.ComponentFactory)
	injectComponent(w, &components.SceneGraphNode{}, &bc.SceneGraphNode.ComponentFactory)
	injectComponent(w, &components.Text{}, &bc.Text.ComponentFactory)
	injectComponent(w, &components.RenderOrder{}, &bc.RenderOrder.ComponentFactory)
	injectComponent(w, &components.RenderTexture2D{}, &bc.RenderTexture2D.ComponentFactory)
	injectComponent(w, &components.Size{}, &bc.Size.ComponentFactory)
	injectComponent(w, &components.Texture2D{}, &bc.Texture2D.ComponentFactory)
	injectComponent(w, &components.Transform{}, &bc.Transform.ComponentFactory)
	injectComponent(w, &components.UUID{}, &bc.UUID.ComponentFactory)
	injectComponent(w, &audio.Audible{}, &bc.Audible.ComponentFactory)
}

// IsInit returns whether or not all of the basic component factories have been initialized
func (bc *BasicComponents) IsInit() bool {
	if bc.Text.ComponentFactory == nil ||
		bc.Texture2D.ComponentFactory == nil ||
		bc.RenderTexture2D.ComponentFactory == nil ||
		bc.Transform.ComponentFactory == nil ||
		bc.Color.ComponentFactory == nil ||
		bc.FileType.ComponentFactory == nil ||
		bc.FileLoadRequest.ComponentFactory == nil ||
		bc.FileLoadResponse.ComponentFactory == nil ||
		bc.Fill.ComponentFactory == nil ||
		bc.Animation.ComponentFactory == nil ||
		bc.Origin.ComponentFactory == nil ||
		bc.Opacity.ComponentFactory == nil ||
		bc.Debug.ComponentFactory == nil ||
		bc.Stroke.ComponentFactory == nil ||
		bc.Font.ComponentFactory == nil ||
		bc.HasChildren.ComponentFactory == nil ||
		bc.Size.ComponentFactory == nil ||
		bc.RenderOrder.ComponentFactory == nil ||
		bc.Viewport.ComponentFactory == nil ||
		bc.UUID.ComponentFactory == nil {
		return false
	}

	return true
}

// injectComponent uses a component to retrieve a singleton component factory for that
// component type from the given world, assigning the generic factory to the given **ComponentFactory
func injectComponent(w *akara.World, c akara.Component, cf **akara.ComponentFactory) {
	cid := w.RegisterComponent(c)
	*cf = w.GetComponentFactory(cid)
}
