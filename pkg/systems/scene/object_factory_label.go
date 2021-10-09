package scene

import (
	"github.com/faiface/mainthread"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/director/pkg/common"
	"image/color"
	"math"
	"math/rand"
	"sync"
	"time"
)

type labelFactory struct {
	common.EntityManager
	*common.BasicComponents
	cache map[common.Entity]*labelParameters
	cacheMutex sync.Mutex
}

type labelParameters struct {
	text     string
	fontSize int
	fontFace string
	color    color.Color
	debug    bool
}

func (factory *labelFactory) New(s *Scene, str string, x, y, fontsize int, fontName string, c color.Color) common.Entity {
	e := s.Add.generic.visibleEntity(s)

	trs, _ := s.Components.Transform.Get(e) // this is a component all visible entities have

	text := s.Components.Text.Add(e)
	font := s.Components.Font.Add(e)
	col := s.Components.Color.Add(e)
	r, g, b, a := c.RGBA()

	text.String = str
	trs.Translation.Set(float64(x), float64(y), trs.Translation.Z)
	col.Color = color.RGBA{
		R: uint8(r),
		G: uint8(g),
		B: uint8(b),
		A: uint8(a),
	}

	font.Face = fontName
	font.Size = fontsize

	//rl.MeasureTextEx()

	factory.EntityManager.AddEntity(e)

	return e
}

func (factory *labelFactory) update(s *Scene, _ time.Duration) {
	if !factory.EntityManager.IsInit() {
		factory.EntityManager.Init()
	}

	if factory.cache == nil {
		factory.cache = make(map[common.Entity]*labelParameters)
	}

	factory.EntitiesMutex.Lock()
	for e := range factory.Entities {
		if !factory.needsToGenerateTexture(s, e) {
			continue
		}

		factory.generateNewTexture(s, e)
	}
	factory.EntitiesMutex.Unlock()

	factory.EntityManager.ProcessRemovalQueue()
}

func (factory *labelFactory) putInCache(_ *Scene, e common.Entity, str, font string, size int, c color.Color) {
	r, g, b, a := c.RGBA() // need to make a copy

	entry := &labelParameters{
		text:     str,
		fontSize: size,
		fontFace: font,
		color: color.RGBA{
			R: uint8(r),
			G: uint8(g),
			B: uint8(b),
			A: uint8(a),
		},
	}

	factory.cacheMutex.Lock()
	factory.cache[e] = entry
	factory.cacheMutex.Unlock()
}

func (factory *labelFactory) needsToGenerateTexture(s *Scene, e common.Entity) bool {
	factory.cacheMutex.Lock()
	entry, found := factory.cache[e]
	factory.cacheMutex.Unlock()
	if !found || entry == nil {
		return true
	}

	text, found := s.Components.Text.Get(e)
	if !found {
		return true
	}

	font, found := s.Components.Font.Get(e)
	if !found {
		return true
	}

	c, found := s.Components.Color.Get(e)
	if !found {
		return true
	}

	_, debugFound := s.Components.Debug.Get(e)

	er, eg, eb, ea := entry.color.RGBA()
	cr, cg, cb, ca := c.RGBA()

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

	if debugFound != entry.debug {
		entry.debug = debugFound
		return true
	}

	return false
}

func (factory *labelFactory) generateNewTexture(s *Scene, e common.Entity) {
	text, textFound := s.Components.Text.Get(e)
	font, fontFound := s.Components.Font.Get(e)
	c, colorFound := s.Components.Color.Get(e)

	if !(textFound || colorFound || fontFound) {
		factory.RemoveEntity(e)
		return
	}

	rt, rtFound := s.Components.RenderTexture2D.Get(e)
	if !rtFound {
		rt = s.Components.RenderTexture2D.Add(e)
	}

	cr, cg, cb, ca := c.RGBA()
	rlc := rl.Color{
		R: uint8(cr),
		G: uint8(cg),
		B: uint8(cb),
		A: uint8(ca),
	}

	str := text.String

	_, debugFound := s.Components.Debug.Get(e)

	mainthread.Call(func() {
		w, h := factory.getTextureSize(text.String, font.Face, font.Size)
		if rt.RenderTexture2D == nil || rt.Texture.Width != int32(w) || rt.Texture.Height != int32(h) {
			newRT := rl.LoadRenderTexture(int32(w), int32(h))
			rt.RenderTexture2D = &newRT
		}

		rl.BeginTextureMode(*rt.RenderTexture2D)
		rl.ClearBackground(rl.Blank)
		rl.DrawText(str, 0, 0, int32(font.Size), rlc)

		if debugFound {
			rl.DrawRectangleLines(0, 0, int32(w), int32(h), rl.NewColor(randRGBA()))
		}

		rl.EndTextureMode()
	})

	factory.putInCache(s, e, str, "", font.Size, c)
}

func (factory *labelFactory) getTextureSize(text, fontName string, size int) (w, h int) {
	//font := rl.LoadFont(fontFace)
	font := rl.GetFontDefault()
	v := rl.MeasureTextEx(font, text, float32(size), 8)

	return int(v.X), int(v.Y)
}

func rand8() uint8 {
	return uint8(rand.Intn(math.MaxUint8))
}

func randRGBA() (r, g, b, a uint8) {
	return rand8(), rand8(), rand8(), math.MaxUint8
}
