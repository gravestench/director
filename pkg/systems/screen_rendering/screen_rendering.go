package screen_rendering

import (
	"sort"

	renderTexture "github.com/gravestench/director/pkg/components/render_texture"

	"github.com/gravestench/director/pkg/components/transform"
	"github.com/gravestench/director/pkg/components/viewport"

	"github.com/faiface/mainthread"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"

	"github.com/gravestench/director/pkg/common"
)

type ScreenRenderingSystem struct {
	akara.BaseSystem
	components struct {
		common.SceneComponents
	}
	sceneViewports *akara.Subscription
}

func (sys *ScreenRenderingSystem) Name() string {
	return "ScreenRendering"
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
	filter := sys.World.NewComponentFilter().Require(
		&viewport.Component{},
		&transform.Component{},
		&renderTexture.Component{},
	).Build()

	sys.sceneViewports = sys.World.AddSubscription(filter)
}

func (sys *ScreenRenderingSystem) Update() {
	mainthread.Call(func() {
		rl.BeginDrawing()
		defer rl.EndDrawing()

		rl.ClearBackground(rl.Blank)

		rl.BeginBlendMode(rl.BlendAlpha)
		defer rl.EndBlendMode()

		for _, e := range sys.getSortedViewports() {
			sys.renderViewport(e)
		}
	})
}

func (sys *ScreenRenderingSystem) getSortedViewports() []akara.EID {
	renderList := sys.sceneViewports.GetEntities()

	// sort viewports by their render order, if they have it
	sort.Slice(renderList, func(i, j int) bool {
		a, b := renderList[i], renderList[j]
		roA, foundA := sys.components.RenderOrder.Get(a)
		roB, foundB := sys.components.RenderOrder.Get(b)

		if !foundA || !foundB {
			return a < b
		}

		return roA.Value < roB.Value
	})

	return renderList
}

func (sys *ScreenRenderingSystem) renderViewport(e akara.EID) {
	// we use the camera in the filter merely to tag the transform + rendertexture
	// for rendering here
	trs, found := sys.components.Transform.Get(e)
	if !found {
		return
	}

	rt, found := sys.components.RenderTexture2D.Get(e)
	if !found || rt.RenderTexture2D == nil {
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

		alpha = uint8(opacity.Value * 255)
	}

	position := rl.Vector2{
		X: float32(trs.Translation.X),
		Y: float32(trs.Translation.Y),
	}

	rotation := float32(trs.Rotation.Y)
	scale := float32(trs.Scale.X)

	rl.DrawTextureEx(rt.Texture, position, rotation, scale, rl.NewColor(0xff, 0xff, 0xff, alpha))
}
