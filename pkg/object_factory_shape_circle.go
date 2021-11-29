package pkg

import (
	"image/color"
	"sync"
	"time"

	"github.com/gravestench/akara"

	"github.com/faiface/mainthread"
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/gravestench/director/pkg/common"
)

type circleFactory struct {
	common.EntityManager
	cache      map[akara.EID]*circleParameters
	cacheMutex sync.Mutex
}

type circleParameters struct {
	width, height int
	fill, stroke  color.Color
}

func (factory *circleFactory) putInCache(e akara.EID, width, height int, fill, stroke color.Color) {
	entry := &circleParameters{
		width:  width,
		height: height,
		fill:   fill,
		stroke: stroke,
	}

	factory.cacheMutex.Lock()
	factory.cache[e] = entry
	factory.cacheMutex.Unlock()
}

func (factory *circleFactory) New(s *Scene, x, y, radius int, fill, stroke color.Color) akara.EID {
	e := s.Add.generic.visibleEntity(s)

	size := s.Components.Size.Add(e)
	size.Max.X, size.Max.Y = radius*2, radius*2

	trs, _ := s.Components.Transform.Get(e)
	trs.Translation.X, trs.Translation.Y = float64(x), float64(y)

	if fill != nil {
		s.Components.Fill.Add(e).Color = fill
	}

	if stroke != nil {
		s.Components.Stroke.Add(e).Color = stroke
	}

	factory.EntityManager.AddEntity(e)

	return e
}

func (factory *circleFactory) update(s *Scene, dt time.Duration) {
	if !factory.EntityManager.IsInit() {
		factory.EntityManager.Init()
	}

	if factory.cache == nil {
		factory.cache = make(map[akara.EID]*circleParameters)
	}

	factory.EntitiesMutex.Lock()
	for e := range factory.EntityManager.Entities {
		if !factory.needsToGenerateTexture(s, e) {
			return
		}

		factory.generateNewTexture(s, e)
	}
	factory.EntitiesMutex.Unlock()

	factory.EntityManager.ProcessRemovalQueue()
}

func (factory *circleFactory) needsToGenerateTexture(s *Scene, e akara.EID) bool {
	factory.cacheMutex.Lock()
	entry, found := factory.cache[e]
	factory.cacheMutex.Unlock()
	if !found {
		return true
	}

	_, rtFound := s.Components.RenderTexture2D.Get(e)
	if !rtFound {
		return true
	}

	fill, fillFound := s.Components.Fill.Get(e)
	stroke, strokeFound := s.Components.Stroke.Get(e)
	col, colorFound := s.Components.Color.Get(e)

	size, sizeFound := s.Components.Size.Get(e)
	if !sizeFound {
		return true
	}

	if fillFound {
		if !colorsEqual(entry.fill, fill.Color) {
			return true
		}
	}

	if !fillFound && colorFound {
		if !colorsEqual(entry.fill, col.Color) {
			return true
		}
	}

	if strokeFound {
		if !colorsEqual(entry.fill, stroke.Color) {
			return true
		}
	}

	if entry.width != size.Dx() || entry.height != size.Dy() {
		return true
	}

	return false
}

func (factory *circleFactory) generateNewTexture(s *Scene, e akara.EID) {
	fill, fillFound := s.Components.Fill.Get(e)
	stroke, strokeFound := s.Components.Stroke.Get(e)
	col, colorFound := s.Components.Color.Get(e)

	size, sizeFound := s.Components.Size.Get(e)
	if !sizeFound {
		return
	}

	w, h := int32(size.Max.X), int32(size.Max.Y)

	if w < 1 || h < 1 {
		return
	}

	var fc, sc color.Color

	rt, rtFound := s.Components.RenderTexture2D.Get(e)
	mainthread.Call(func() {
		if !rtFound {
			rt = s.Components.RenderTexture2D.Add(e)
			newRT := rl.LoadRenderTexture(w, h)
			rt.RenderTexture2D = &newRT
		}

		rl.BeginTextureMode(*rt.RenderTexture2D)
		rl.ClearBackground(rl.Blank)

		if fillFound {
			fc = fill
			r, g, b, a := fill.RGBA()
			rl.DrawCircle(w/2, h/2, float32(w/2), rl.NewColor(uint8(r), uint8(g), uint8(b), uint8(a)))
		}

		if !fillFound && colorFound {
			fc = col
			r, g, b, a := col.RGBA()
			rl.DrawCircle(w/2, h/2, float32(w/2), rl.NewColor(uint8(r), uint8(g), uint8(b), uint8(a)))
		}

		if strokeFound {
			sc = stroke
			r, g, b, a := stroke.RGBA()
			rl.DrawCircleLines(w/2, h/2, float32(w/2), rl.NewColor(uint8(r), uint8(g), uint8(b), uint8(a)))
		}

		rl.EndTextureMode()
	})

	factory.putInCache(e, int(w), int(h), fc, sc)
}
