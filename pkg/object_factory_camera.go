package pkg

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
	"time"
)

type cameraFactory struct {
	*basicComponents
}

func (*cameraFactory) New(s *Scene, x, y, w, h int) akara.EID {
	e := s.Add.generic.visibleEntity(s)

	cam := s.Camera2D.Add(e)
	cam.Camera2D = rl.NewCamera2D(rl.Vector2{}, rl.Vector2{}, 0, 1)

	trs, _ := s.Transform.Get(e) // this is a component all visible entities have
	rt := s.RenderTexture2D.Add(e)

	newRT := rl.LoadRenderTexture(int32(w), int32(h))
	rt.RenderTexture2D = &newRT

	trs.Translation.Set(float64(x), float64(y), trs.Translation.Z)

	return e
}

func (factory *cameraFactory) update(s *Scene, dt time.Duration) {

}

