package scene

import (
	"github.com/faiface/mainthread"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/director/pkg/common"
	"github.com/gravestench/mathlib"
	"math"
	"sort"
)

func (s *Scene) initViewport() {
	s.Viewports = make([]common.Entity, 1)
	rw, rh := s.Sys.Renderer.Window.Width, s.Sys.Renderer.Window.Height
	s.Viewports[0] = s.Add.Viewport(0, 0, rw, rh)
	vp, _ := s.Components.Viewport.Get(s.Viewports[0])

	// remove viewport and camera from the scene's renderList so they don't get rendered to themselves later
	newRenderList := make([]common.Entity, 0, len(s.renderList))
	for _, entity := range s.renderList {
		if entity != s.Viewports[0] && entity != vp.CameraEntity {
			newRenderList = append(newRenderList, entity)
		}
	}
	s.renderList = newRenderList
}

type entityRenderRequest struct {
	Texture  rl.Texture2D
	Position rl.Vector2
	Rotation float32
	Scale    float32
	Tint     rl.Color
}

func (s *Scene) generateEntityRenderBatch(entities []common.Entity) []entityRenderRequest {
	entityRenderRequests := make([]entityRenderRequest, 0, len(entities))

	for _, e := range entities {
		texture, textureFound := s.Components.Texture2D.Get(e)
		rt, rtFound := s.Components.RenderTexture2D.Get(e)
		if !textureFound && !rtFound || (textureFound && texture.Texture2D == nil) {
			continue
		}

		var t *rl.Texture2D

		if !rtFound {
			t = texture.Texture2D
		} else {
			t = &rt.Texture
		}

		trs, found := s.Components.Transform.Get(e)
		if !found {
			continue
		}

		origin, found := s.Components.Origin.Get(e)
		if !found {
			continue
		}

		tint := rl.White

		opacity, found := s.Components.Opacity.Get(e)
		if found {
			if opacity.Value > 1 {
				opacity.Value = 1
			} else if opacity.Value < 0 {
				opacity.Value = 0
			}

			tint.A = uint8(float64(math.MaxUint8) * opacity.Value)
		}

		if tint.A == 0 {
			continue
		}

		// this is rotating around the origin point from the origin component
		tmpVect.Set(float64(t.Width), float64(t.Height), 1)
		yRad := trs.Rotation.Y * mathlib.DegreesToRadians
		ov2 := mathlib.NewVector2(origin.Clone().Multiply(&tmpVect).XY()).Rotate(yRad).Negate()
		ov3 := mathlib.NewVector3(ov2.X, ov2.Y, 0)
		x, y := trs.Translation.Clone().Add(ov3).XY()
		v2 := mathlib.NewVector2(x, y)

		position := rl.Vector2{X: float32(v2.X), Y: float32(v2.Y)}
		rotation := float32(trs.Rotation.Y)
		scale := float32(trs.Scale.X)

		entityRenderRequests = append(entityRenderRequests, entityRenderRequest{
			Texture:  *t,
			Position: position,
			Rotation: rotation,
			Scale:    scale,
			Tint:     tint,
		})
	}

	return entityRenderRequests
}

func (s *Scene) drawEntitiesAndRender(viewport common.Entity) {
	vp, found := s.Components.Viewport.Get(viewport)
	if !found {
		return
	}

	vprt, found := s.Components.RenderTexture2D.Get(viewport)
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

	sort.Slice(s.renderList, func(i, j int) bool {
		a, b := s.renderList[i], s.renderList[j]
		roA, foundA := s.Components.RenderOrder.Get(a)
		roB, foundB := s.Components.RenderOrder.Get(b)

		if !foundA || !foundB {
			return a < b
		}

		return roA.Value < roB.Value
	})

	// prepare a batch of entities to render
	renderBatch := s.generateEntityRenderBatch(s.renderList)

	mainthread.Call(func() {
		// render all the entities to the camera's render texture
		rl.BeginTextureMode(*camrt.RenderTexture2D)
		rl.BeginMode2D(cam.Camera2D)

		r, g, b, a := vp.Background.RGBA()
		rl.ClearBackground(rl.NewColor(uint8(r), uint8(g), uint8(b), uint8(a)))

		for _, e := range renderBatch {
			rl.DrawTextureEx(e.Texture, e.Position, e.Rotation, e.Scale, e.Tint)
		}

		rl.EndMode2D()
		rl.EndTextureMode()

		// render the camera
		rl.BeginTextureMode(*vprt.RenderTexture2D)

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

		rl.EndTextureMode()
	})
}
