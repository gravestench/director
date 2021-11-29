package pkg

import (
	"time"

	"github.com/gravestench/akara"

	"github.com/faiface/mainthread"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/director/pkg/common"
)

type cameraFactory struct {
	common.EntityManager
	*common.SceneComponents
}

func (*cameraFactory) New(s *Scene, x, y, w, h int) akara.EID {
	ce := s.Add.generic.visibleEntity(s)

	cam := s.Components.Camera.Add(ce)
	var newRT rl.RenderTexture2D
	mainthread.Call(func() {
		cam.Camera2D = rl.NewCamera2D(rl.Vector2{}, rl.Vector2{}, 0, 1)

		newRT = rl.LoadRenderTexture(int32(w), int32(h))
	})

	rt := s.Components.RenderTexture2D.Add(ce)
	rt.RenderTexture2D = &newRT

	trs, _ := s.Components.Transform.Get(ce) // this is a component all visible entities have
	trs.Translation.Set(float64(x), float64(y), trs.Translation.Z)

	return ce
}

func (factory *cameraFactory) update(s *Scene, _ time.Duration) {
	if !factory.EntityManager.IsInit() {
		factory.EntityManager.Init()
	}

	factory.EntityManager.ProcessRemovalQueue()
}
