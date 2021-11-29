package common

import (
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/common/components"
	"github.com/gravestench/director/pkg/components/audio"
	"github.com/gravestench/director/pkg/components/interactive"

	"github.com/gravestench/director/pkg/components/animation"
	"github.com/gravestench/director/pkg/components/camera"
	"github.com/gravestench/director/pkg/components/color"
	debugging "github.com/gravestench/director/pkg/components/debug"
	fileLoadRequest "github.com/gravestench/director/pkg/components/file_load_request"
	fileLoadResponse "github.com/gravestench/director/pkg/components/file_load_response"
	fileType "github.com/gravestench/director/pkg/components/file_type"
	"github.com/gravestench/director/pkg/components/fill"
	"github.com/gravestench/director/pkg/components/font"
	fontFace "github.com/gravestench/director/pkg/components/font"
	hasChildren "github.com/gravestench/director/pkg/components/has_children"
	"github.com/gravestench/director/pkg/components/opacity"
	"github.com/gravestench/director/pkg/components/origin"
	renderOrder "github.com/gravestench/director/pkg/components/render_order"
	renderTexture "github.com/gravestench/director/pkg/components/render_texture"
	sceneGraphNode "github.com/gravestench/director/pkg/components/scene_graph_node"
	"github.com/gravestench/director/pkg/components/size"
	"github.com/gravestench/director/pkg/components/stroke"
	"github.com/gravestench/director/pkg/components/text"
	texture2D "github.com/gravestench/director/pkg/components/texture"
	"github.com/gravestench/director/pkg/components/transform"
	UUID "github.com/gravestench/director/pkg/components/uuid"
	"github.com/gravestench/director/pkg/components/viewport"
)

// SceneComponents represents components that every scene has available
type SceneComponents struct {
	Animation        animation.ComponentFactory
	Camera           camera.ComponentFactory
	Color            color.ComponentFactory
	Debug            debugging.ComponentFactory
	FileLoadRequest  fileLoadRequest.ComponentFactory
	FileLoadResponse fileLoadResponse.ComponentFactory
	FileType         fileType.ComponentFactory
	Fill             fill.ComponentFactory
	Font             fontFace.ComponentFactory
	HasChildren      hasChildren.ComponentFactory
	Opacity          opacity.ComponentFactory
	Stroke           stroke.ComponentFactory
	Origin           origin.ComponentFactory
	RenderTexture2D  renderTexture.ComponentFactory
	RenderOrder      renderOrder.ComponentFactory
	Size             size.ComponentFactory
	SceneGraphNode   sceneGraphNode.ComponentFactory
	Text             text.ComponentFactory
	Texture2D        texture2D.ComponentFactory
	Transform        transform.ComponentFactory
	UUID             UUID.ComponentFactory
	Viewport         viewport.ComponentFactory
	Audible          audio.ComponentFactory
	Interactive      interactive.ComponentFactory
}

// Init initializes each component factory for the given world,
// putting the generic component factory inside of the concrete component factory
func (bc *SceneComponents) Init(w *akara.World) {
	injectComponent(w, &animation.Component{}, &bc.Animation.ComponentFactory)
	injectComponent(w, &camera.Component{}, &bc.Camera.ComponentFactory)
	injectComponent(w, &color.Component{}, &bc.Color.ComponentFactory)
	injectComponent(w, &debugging.Component{}, &bc.Debug.ComponentFactory)
	injectComponent(w, &fileLoadRequest.Component{}, &bc.FileLoadRequest.ComponentFactory)
	injectComponent(w, &fileLoadResponse.Component{}, &bc.FileLoadResponse.ComponentFactory)
	injectComponent(w, &fileType.Component{}, &bc.FileType.ComponentFactory)
	injectComponent(w, &fill.Component{}, &bc.Fill.ComponentFactory)
	injectComponent(w, &viewport.Component{}, &bc.Viewport.ComponentFactory)
	injectComponent(w, &hasChildren.Component{}, &bc.HasChildren.ComponentFactory)
	injectComponent(w, &origin.Component{}, &bc.Origin.ComponentFactory)
	injectComponent(w, &interactive.Component{}, &bc.Interactive.ComponentFactory)
	injectComponent(w, &opacity.Component{}, &bc.Opacity.ComponentFactory)
	injectComponent(w, &stroke.Component{}, &bc.Stroke.ComponentFactory)
	injectComponent(w, &font.Component{}, &bc.Font.ComponentFactory)
	injectComponent(w, &sceneGraphNode.Component{}, &bc.SceneGraphNode.ComponentFactory)
	injectComponent(w, &text.Component{}, &bc.Text.ComponentFactory)
	injectComponent(w, &renderOrder.Component{}, &bc.RenderOrder.ComponentFactory)
	injectComponent(w, &renderTexture.Component{}, &bc.RenderTexture2D.ComponentFactory)
	injectComponent(w, &size.Component{}, &bc.Size.ComponentFactory)
	injectComponent(w, &texture2D.Component{}, &bc.Texture2D.ComponentFactory)
	injectComponent(w, &transform.Component{}, &bc.Transform.ComponentFactory)
	injectComponent(w, &UUID.Component{}, &bc.UUID.ComponentFactory)
	injectComponent(w, &audio.Component{}, &bc.Audible.ComponentFactory)
}

// IsInit returns whether or not all of the basic component factories have been initialized
func (bc *SceneComponents) IsInit() bool {
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
func injectComponent(w *akara.World, c components.Component, cf **akara.ComponentFactory) {
	cid := w.RegisterComponent(c)
	*cf = w.GetComponentFactory(cid)
}
