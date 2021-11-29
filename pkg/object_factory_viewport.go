package pkg

import (
	"time"

	"github.com/gravestench/akara"

	"github.com/faiface/mainthread"
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/gravestench/director/pkg/common"
)

type viewportFactory struct {
	common.EntityManager
	*common.SceneComponents
}

func (*viewportFactory) New(s *Scene, x, y, w, h int) akara.EID {
	e := s.Add.generic.visibleEntity(s)

	viewport := s.Components.Viewport.Add(e)
	viewport.CameraEntity = s.Add.Camera(x, y, w, h)

	trs, _ := s.Components.Transform.Get(e)
	rt := s.Components.RenderTexture2D.Add(e)

	mainthread.Call(func() {
		newRT := rl.LoadRenderTexture(int32(w), int32(h))
		rt.RenderTexture2D = &newRT
	})

	trs.Translation.Set(float64(x), float64(y), trs.Translation.Z)

	return e
}

func (factory *viewportFactory) update(s *Scene, _ time.Duration) {
	if !factory.EntityManager.IsInit() {
		factory.EntityManager.Init()
	}

	factory.EntityManager.ProcessRemovalQueue()
}
