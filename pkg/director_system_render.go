package pkg

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/components"
)

type screenRenderingSystem struct {
	akara.BaseSubscriberSystem
	components struct {
		basicComponents
	}
	sceneCameras *akara.Subscription
}

func (sys *screenRenderingSystem) Init(world *akara.World) {
	sys.World = world

	sys.components.init(world)
	sys.initCameraSubscription()
}

func (sys *screenRenderingSystem) IsInitialized() bool {
	if sys.World == nil {
		return false
	}

	if !sys.components.isInit() {
		return false
	}

	return true
}

func (sys *screenRenderingSystem) initCameraSubscription() {
	filter := sys.World.NewComponentFilter()

	filter.Require(
		&components.Camera2D{},
		&components.Transform{},
		&components.RenderTexture2D{},
		)

	sys.sceneCameras = sys.World.AddSubscription(filter.Build())
}

func (sys *screenRenderingSystem) Update() {
	rl.BeginDrawing()

	rl.ClearBackground(rl.Black)

	for _, e := range sys.sceneCameras.GetEntities() {
		sys.renderCamera(e)
	}

	rl.EndDrawing()
}

func (sys *screenRenderingSystem) renderCamera(e akara.EID) {
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

	position := rl.Vector2{
		X: float32(trs.Translation.X),
		Y: float32(trs.Translation.Y),
	}

	rotation := float32(trs.Rotation.Y)

	scale := float32(trs.Scale.X)

	rl.DrawTextureEx(rt.Texture, position, rotation, scale, rl.White)
}

