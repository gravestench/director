package pkg

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
	"time"
)

type cameraFactory struct {
	entityManager
	*basicComponents
}

func (*cameraFactory) New(s *Scene, x, y, w, h int) akara.EID {
	e := s.Add.generic.visibleEntity(s)

	cam := s.Components.Camera2D.Add(e)
	cam.Camera2D = rl.NewCamera2D(rl.Vector2{}, rl.Vector2{}, 0, 0.2)

	trs, _ := s.Components.Transform.Get(e) // this is a component all visible entities have
	rt := s.Components.RenderTexture2D.Add(e)

	newRT := rl.LoadRenderTexture(int32(w), int32(h))
	rt.RenderTexture2D = &newRT

	trs.Translation.Set(float64(x), float64(y), trs.Translation.Z)

	return e
}

func (factory *cameraFactory) update(s *Scene, _ time.Duration) {
	if !factory.entityManagerIsInit() {
		factory.entityManagerInit()
	}

	factory.applyTransformToCamera(s)
	factory.processRemovalQueue()
}

func (factory *cameraFactory) applyTransformToCamera(s *Scene) {
	for e := range factory.entities {
		cam, found := s.Components.Camera2D.Get(e)
		if !found {
			continue
		}

		trs, found := s.Components.Transform.Get(e)
		if !found {
			continue
		}

		cam.Rotation = float32(trs.Rotation.Y)
		cam.Offset = rl.Vector2{
			X: float32(trs.Translation.X),
			Y: float32(trs.Translation.Y),
		}
	}
}
