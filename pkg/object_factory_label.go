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
	cache map[akara.EID]*labelParameters
}

type labelParameters struct {
	text     string
	fontSize int
	fontFace string
	color    color.RGBA
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

func (factory *labelFactory) update(s *Scene, _ time.Duration) {
	if !factory.entityManagerIsInit() {
		factory.entityManagerInit()
	}

	if factory.cache == nil {
		factory.cache = make(map[akara.EID]*labelParameters)
	}

	for _, e := range factory.entities {
		if !factory.needsToGenerateTexture(s, e) {
			continue
		}

		factory.generateNewTexture(s, e)
	}
}

func (factory *labelFactory) putInCache(s *Scene, e akara.EID, str, font string, size int, c color.RGBA) {
	entry := &labelParameters{
		text:     str,
		fontSize: size,
		fontFace: font,
		color:    c,
	}

	factory.cache[e] = entry
}

func (factory *labelFactory) needsToGenerateTexture(s *Scene, e akara.EID) bool {
	entry, found := factory.cache[e]
	if !found || entry == nil {
		return true
	}

	text, found := s.Text.Get(e)
	if !found {
		return true
	}

	font, found := s.Font.Get(e)
	if !found {
		return true
	}


	c, found := s.Color.Get(e)
	if !found {
		return true
	}

	er, eg, eb, ea := entry.color.RGBA()
	cr, cg, cb, ca := c.RGBA.RGBA()

	if er != cr || eg != cg || eb != cb || ea != ca {
		return true
	}

	if entry.text != text.String {
		return true
	}

	if entry.fontFace != font.Face {
		return true
	}

	if entry.fontSize != font.Size {
		return true
	}

	return false
}

func (factory *labelFactory) generateNewTexture(s *Scene, e akara.EID) {
	text, textFound := s.Text.Get(e)
	font, fontFound := s.Font.Get(e)
	c, colorFound := s.Color.Get(e)

	if !(textFound || colorFound || fontFound) {
		factory.removalQueue = append(factory.removalQueue, e)
		return
	}

	rt, rtFound := s.RenderTexture2D.Get(e)
	if !rtFound {
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
	rl.ClearBackground(rl.Blank)
	rl.DrawText(str, 0, 0, int32(font.Size), rlc)
	rl.EndTextureMode()

	factory.putInCache(s, e, str, "", font.Size, c.RGBA)
}

func (factory *labelFactory) getTextureSize(text, fontName string, size int) (w, h int) {
	//font := rl.LoadFont(fontFace)
	font := rl.GetFontDefault()
	v := rl.MeasureTextEx(font, text, float32(size), 8)

	return int(v.X), int(v.Y)
}