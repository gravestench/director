package screen_rendering

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/common"
	"github.com/gravestench/director/pkg/components"
)

type ScreenRenderingSystem struct {
	akara.BaseSubscriberSystem
	components struct {
		common.BasicComponents
	}
	sceneCameras *akara.Subscription
}

func (sys *ScreenRenderingSystem) Init(world *akara.World) {
	sys.World = world

	sys.components.Init(world)
	sys.initCameraSubscription()
}

func (sys *ScreenRenderingSystem) IsInitialized() bool {
	if sys.World == nil {
		return false
	}

	if !sys.components.IsInit() {
		return false
	}

	return true
}

func (sys *ScreenRenderingSystem) initCameraSubscription() {
	filter := sys.World.NewComponentFilter()

	filter.Require(
		&components.Viewport{},
		&components.Transform{},
		&components.RenderTexture2D{},
		)

	sys.sceneCameras = sys.World.AddSubscription(filter.Build())
}

func (sys *ScreenRenderingSystem) Update() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Black)
	rl.BeginBlendMode(rl.BlendAlpha)

	for _, e := range sys.sceneCameras.GetEntities() {
		sys.renderCamera(e)
		break
	}

	rl.EndBlendMode()
	rl.EndDrawing()
}

func (sys *ScreenRenderingSystem) renderCamera(e common.Entity) {
	// we use the camera in the filter merely to tag the transform + rendertexture
	// for rendering here
	trs, found := sys.components.Transform.Get(e)
	if !found {
		return
	}

	rt, found := sys.components.RenderTexture2D.Get(e)
	if !found {
		return
	}

	alpha := 1.0
	opacity, found := sys.components.Opacity.Get(e)
	if found {
		alpha = opacity.Value
	}

	position := rl.Vector2{
		X: float32(trs.Translation.X),
		Y: float32(trs.Translation.Y),
	}

	rotation := float32(trs.Rotation.Y)
	scale := float32(trs.Scale.X)

	rl.DrawTextureEx(rt.Texture, position, rotation, scale, rl.NewColor(0xff, 0xff, 0xff, uint8(alpha * 255)))
}

