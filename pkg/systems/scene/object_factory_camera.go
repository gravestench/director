package scene

import (
	"github.com/faiface/mainthread"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/director/pkg/common"
	"time"
)

type cameraFactory struct {
	common.EntityManager
	*common.BasicComponents
}

func (*cameraFactory) New(s *Scene, x, y, w, h int) common.Entity {
	ce := s.Add.generic.visibleEntity(s)

	cam := s.Components.Camera.Add(ce)
	mainthread.Call(func() {
		cam.Camera2D = rl.NewCamera2D(rl.Vector2{}, rl.Vector2{}, 0, 1)
	})

	trs, _ := s.Components.Transform.Get(ce) // this is a component all visible entities have
	rt := s.Components.RenderTexture2D.Add(ce)

	var newRT rl.RenderTexture2D
	mainthread.Call(func() {
		newRT = rl.LoadRenderTexture(int32(w), int32(h))
	})
	rt.RenderTexture2D = &newRT

	trs.Translation.Set(float64(x), float64(y), trs.Translation.Z)

	return ce
}

func (factory *cameraFactory) update(s *Scene, _ time.Duration) {
	if !factory.EntityManager.IsInit() {
		factory.EntityManager.Init()
	}

	factory.EntityManager.ProcessRemovalQueue()
}
