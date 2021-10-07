package screen_rendering

import (
	"github.com/faiface/mainthread"
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
	sceneViewports *akara.Subscription
}

func (sys *ScreenRenderingSystem) Init(world *akara.World) {
	sys.components.Init(world)
	sys.initViewportSubscription()
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

func (sys *ScreenRenderingSystem) initViewportSubscription() {
	filter := sys.World.NewComponentFilter()

	filter.Require(
		&components.Viewport{},
		&components.Transform{},
		&components.RenderTexture2D{},
	)

	sys.sceneViewports = sys.World.AddSubscription(filter.Build())
}

func (sys *ScreenRenderingSystem) Update() {
	mainthread.Call(func() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Blank)
		rl.BeginBlendMode(rl.BlendAlpha)

		for _, e := range sys.sceneViewports.GetEntities() {
			sys.renderViewport(e)
		}

		rl.EndBlendMode()
		rl.EndDrawing()
	})
}

func (sys *ScreenRenderingSystem) renderViewport(e common.Entity) {
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

	var alpha uint8

	opacity, found := sys.components.Opacity.Get(e)
	if found {
		if opacity.Value > 1 {
			opacity.Value = 1
		} else if opacity.Value < 0 {
			opacity.Value = 0
		}

		alpha = uint8(opacity.Value*255)
	}

	position := rl.Vector2{
		X: float32(trs.Translation.X),
		Y: float32(trs.Translation.Y),
	}

	rotation := float32(trs.Rotation.Y)
	scale := float32(trs.Scale.X)

	rl.DrawTextureEx(rt.Texture, position, rotation, scale, rl.NewColor(0xff, 0xff, 0xff, alpha))
}
