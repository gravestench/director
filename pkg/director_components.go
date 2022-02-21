package pkg

import (
	"image/color"

	"github.com/gravestench/director/pkg/components/animation"
	"github.com/gravestench/director/pkg/components/audio"
	"github.com/gravestench/director/pkg/components/camera"
	"github.com/gravestench/director/pkg/components/debug"
	fileLoadRequest "github.com/gravestench/director/pkg/components/file_load_request"
	fileLoadResponse "github.com/gravestench/director/pkg/components/file_load_response"
	fileType "github.com/gravestench/director/pkg/components/file_type"
	"github.com/gravestench/director/pkg/components/fill"
	"github.com/gravestench/director/pkg/components/font"
	hasChildren "github.com/gravestench/director/pkg/components/has_children"
	"github.com/gravestench/director/pkg/components/interactive"
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
	"github.com/gravestench/director/pkg/components/uuid"
	"github.com/gravestench/director/pkg/components/viewport"
)

// DirectorComponents contains all of the primitive components that come with director.
// These are INTENTIONALLY nil instances, which can be used when creating component filters.
// See akara's `World.NewComponentFilter` and `World.AddSubscription`
type DirectorComponents struct {
	Audio            *audio.Component
	Interactive      *interactive.Component
	Viewport         *viewport.Component
	Camera           *camera.Camera
	Color            *color.Color
	Debug            *debug.Component
	FileLoadRequest  *fileLoadRequest.Component
	FileLoadResponse *fileLoadResponse.Component
	FileType         *fileType.Component
	Fill             *fill.Component
	HasChildren      *hasChildren.Component
	Animation        *animation.Component
	Stroke           *stroke.Component
	Font             *font.Component
	Opacity          *opacity.Component
	Origin           *origin.Component
	RenderTexture    *renderTexture.Component
	RenderOrder      *renderOrder.Component
	Size             *size.Component
	SceneGraphNode   *sceneGraphNode.Component
	Text             *text.Component
	Texture2D        *texture2D.Component
	Transform        *transform.Component
	Uuid             *uuid.Component
}
