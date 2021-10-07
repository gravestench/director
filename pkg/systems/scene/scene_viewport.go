package scene

import (
	"github.com/faiface/mainthread"
	"sort"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/director/pkg/common"
)

func (s *Scene) initViewport() {
	s.Viewports = make([]common.Entity, 0)
	rw, rh := s.Sys.Renderer.Window.Width, s.Sys.Renderer.Window.Height
	s.Viewports = append(s.Viewports, s.Add.Viewport(0, 0, rw, rh))
}

func (s *Scene) renderCameraForViewport(viewport common.Entity) {
	vp, found := s.Components.Viewport.Get(viewport)
	if !found {
		return
	}

	cam, found := s.Components.Camera.Get(vp.CameraEntity)
	if !found {
		return
	}

	camrt, found := s.Components.RenderTexture2D.Get(vp.CameraEntity)
	if !found {
		return
	}

	mainthread.Call(func() {
		rl.BeginTextureMode(*camrt.RenderTexture2D)
		defer rl.EndTextureMode()

		rl.BeginMode2D(cam.Camera2D)
		defer rl.EndMode2D()

		r, g, b, a := vp.Background.RGBA()
		rl.ClearBackground(rl.NewColor(uint8(r), uint8(g), uint8(b), uint8(a)))

		sort.Slice(s.renderList, func(i, j int) bool {
			a, b := s.renderList[i], s.renderList[j]
			roA, foundA := s.Components.RenderOrder.Get(a)
			roB, foundB := s.Components.RenderOrder.Get(b)

			if !foundA || !foundB {
				return a < b
			}

			return roA.Value < roB.Value
		})

		for _, entity := range s.renderList {
			if entity == vp.CameraEntity || entity == viewport {
				continue
			}

			s.renderEntity(entity)
		}
	})
}

func (s *Scene) renderCameraToViewport(viewport common.Entity) {
	vp, found := s.Components.Viewport.Get(viewport)
	if !found {
		return
	}

	vprt, found := s.Components.RenderTexture2D.Get(viewport)
	if !found {
		return
	}

	camrt, found := s.Components.RenderTexture2D.Get(vp.CameraEntity)
	if !found {
		return
	}

	mainthread.Call(func() {
		rl.BeginTextureMode(*vprt.RenderTexture2D)
		defer rl.EndTextureMode()

		rl.ClearBackground(rl.Blank)

		src := rl.Rectangle{
			X:      0,
			Y:      float32(camrt.Texture.Height),
			Width:  float32(camrt.Texture.Width),
			Height: -float32(camrt.Texture.Height),
		}

		dst := rl.Rectangle{
			X:      0,
			Y:      0,
			Width:  float32(vprt.Texture.Width),
			Height: float32(vprt.Texture.Height),
		}

		rl.DrawTexturePro(camrt.Texture, src, dst, rl.Vector2{}, 0, rl.White)
	})
}
