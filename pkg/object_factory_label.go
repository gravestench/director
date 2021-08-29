package pkg

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
	"image/color"
	"time"
)

type labelFactory struct {
	entityManager
	*basicComponents
}

func (factory *labelFactory) update(s *Scene, _ time.Duration) {
	if !factory.entityManagerIsInit() {
		factory.entityManagerInit()
	}

	factory.generateNewTextures(s)
}

func (factory *labelFactory) New(s *Scene, str string, x, y, fontsize int, fontName string, c color.Color) akara.EID {
	e := s.Add.generic.visibleEntity(s)

	trs, _ := s.Transform.Get(e) // this is a component all visible entities have

	text := s.Text.Add(e)
	font := s.Font.Add(e)
	col := s.Color.Add(e)
	r, g, b, a := c.RGBA()

	text.String = str
	trs.Translation.Set(float64(x), float64(y), trs.Translation.Z)
	col.R, col.G, col.B, col.A = uint8(r), uint8(g), uint8(b), uint8(a)
	font.Face = fontName
	font.Size = fontsize

	//rl.MeasureTextEx()

	factory.entities = append(factory.entities, e)

	return e
}

func (factory *labelFactory) generateNewTextures(s *Scene) {
	for idx, e := range factory.entities {
		text, found := s.Text.Get(e)
		if !found {
			factory.removalQueue = append(factory.removalQueue, idx)
			continue
		}

		c, found := s.Color.Get(e)
		if !found {
			factory.removalQueue = append(factory.removalQueue, idx)
			continue
		}

		font, found := s.Font.Get(e)
		if !found {
			factory.removalQueue = append(factory.removalQueue, idx)
			continue
		}

		rt, found := s.RenderTexture2D.Get(e)
		if !found {
			rt = s.RenderTexture2D.Add(e)
		}

		w, h := factory.getTextureSize(text.String, font.Face, font.Size)
		if rt.RenderTexture2D == nil || rt.Texture.Width != int32(w) || rt.Texture.Height != int32(h) {
			newRT := rl.LoadRenderTexture(int32(w), int32(h))
			rt.RenderTexture2D = &newRT
		}

		rlc := rl.Color{
			R: c.R,
			G: c.G,
			B: c.B,
			A: c.A,
		}

		str := text.String

		rl.BeginTextureMode(*rt.RenderTexture2D)
		rl.DrawText(str, 0, 0, int32(font.Size), rlc)
		rl.EndTextureMode()
	}
}

func (factory *labelFactory) getTextureSize(text, font string, size int) (w, h int) {
	v := rl.MeasureTextEx(rl.LoadFont(font), text, float32(size), 0)

	return int(v.X), int(v.Y)
}