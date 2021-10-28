// the tick_graph scene is used for printing tick_graph information about all
// running scenes known to director
package tick_graph

import (
	"github.com/faiface/mainthread"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg"
	"github.com/gravestench/director/pkg/common"
	"github.com/gravestench/director/pkg/systems/scene"
	"image/color"
	"math"
	"math/rand"
	"time"
)

const (
	key                     = "director tick graph"
	graphWidth, graphHeight = 200, 60
)

type Scene struct {
	scene.Scene
	sceneColors          map[string]color.Color
	sceneSamples         map[string][]float64
	sceneSamplesPixels   map[string][]common.Entity
	sampleMax, sampleMin float64
	sampleIndex          int
}

func (s *Scene) Key() string {
	return key
}

func (s *Scene) Init(_ *akara.World) {
	s.sceneColors = make(map[string]color.Color)
	s.sceneSamples = make(map[string][]float64)
	s.sceneSamplesPixels = make(map[string][]common.Entity)
}

func (s *Scene) Update() {
	s.resizeViewportTexture(int32(graphWidth), int32(graphHeight))

	for idx := range s.Director.Systems {
		theScene, ok := s.Director.Systems[idx].(pkg.SceneInterface)
		if !ok {
			continue
		}

		if theScene.Key() == key {
			continue
		}

		s.updateSceneColors(theScene)
		s.updateSceneSamples(theScene)
	}

	s.getMinMaxFromSamples()
	s.updatePixels()
}

func (s *Scene) resizeViewportTexture(w, h int32) {
	for _, e := range s.Viewports {
		vp, found := s.Components.Viewport.Get(e)
		if !found {
			continue
		}

		vp.Background = color.RGBA{R: 0xFF, G:0xFF, B: 0xFF, A: 0x64}

		vprt, found := s.Components.RenderTexture2D.Get(e)
		if !found {
			continue
		}

		camrt, found := s.Components.RenderTexture2D.Get(vp.CameraEntity)
		if !found {
			continue
		}

		renderOrder, found := s.Components.RenderOrder.Get(e)
		if !found {
			renderOrder = s.Components.RenderOrder.Add(e)
		}

		renderOrder.Value = math.MaxInt32

		if vprt.Texture.Width != w || vprt.Texture.Height != h {
			mainthread.Call(func() {
				t := rl.LoadRenderTexture(w, h)
				vprt.RenderTexture2D = &t
			})
		}

		if camrt.Texture.Width != w || camrt.Texture.Height != h {
			mainthread.Call(func() {
				t := rl.LoadRenderTexture(w, h)
				camrt.RenderTexture2D = &t
			})
		}
	}
}

func (s *Scene) getMinMaxFromSamples() {
	s.sampleMin, s.sampleMax = 0, 0

	for _, samples := range s.sceneSamples {
		for idx := range samples {
			if samples[idx] < s.sampleMin {
				s.sampleMin = samples[idx]
			}

			if samples[idx] > s.sampleMax {
				s.sampleMax = samples[idx]
			}
		}
	}
}

func randUint8() uint8 {
	return uint8(rand.Intn(math.MaxUint8>>1) << 1)
}

func randomColor() color.Color {
	return &color.RGBA{
		R: randUint8(),
		G: randUint8(),
		B: randUint8(),
		A: math.MaxUint8,
	}
}

func (s *Scene) updateSceneColors(theScene pkg.SceneInterface) {
	key := theScene.Key()
	rand.Seed(time.Now().UnixNano())

	if _, found := s.sceneColors[key]; !found {
		s.sceneColors[key] = randomColor()
	}
}

func (s *Scene) updateSceneSamples(theScene pkg.SceneInterface) {
	key := theScene.Key()

	samples, found := s.sceneSamples[key]
	if !found {
		samples = make([]float64, graphWidth)
	}

	samples[s.sampleIndex] = float64(theScene.TickCount()) / theScene.Uptime().Seconds()
	s.sceneSamples[key] = samples
	s.sampleIndex = (s.sampleIndex + 1) % graphWidth
}

func (s *Scene) updatePixels() {
	for key := range s.sceneSamples {
		if _, found := s.sceneSamplesPixels[key]; !found {
			s.sceneSamplesPixels[key] = make([]common.Entity, graphWidth)

			sceneColor := s.sceneColors[key]

			for idx := range s.sceneSamplesPixels[key] {
				s.sceneSamplesPixels[key][idx] = s.Add.Rectangle(idx, 0, 1, 1, sceneColor, sceneColor)
			}
		}

		pixels := s.sceneSamplesPixels[key]

		for idx, pixel := range pixels {
			s.updatePixelPosition(key, idx, pixel)
		}
	}
}

func (s *Scene) updatePixelPosition(sceneKey string, sampleIndex int, pixelEID common.Entity) {
	trs, found := s.Components.Transform.Get(pixelEID)
	if !found {
		return
	}

	samples, found := s.sceneSamples[sceneKey]
	if !found {
		return
	}

	trs.Translation.Set(float64(sampleIndex), ((samples[sampleIndex] / s.sampleMax) + s.sampleMin) * graphHeight, 0)
}

// PrintDebugMessage prints a tick_graph message to stdout containing information about the running systems, including their
// target tick rates and their average tick rates.
//func (s *Scene) makeDebugMessage() fmt.Stringer {
//	buf := &bytes.Buffer{}
//	writer := tabwriter.NewWriter(buf, 0, 8, 1, '\t', tabwriter.AlignRight)
//	fmt.Fprintln(writer, "System\tActive?\tTarget Hz\tActual Hz")
//
//	for _, system := range s.Director.Systems {
//		fmt.Fprintln(writer, fmt.Sprintf("%s\t%v\t%.2f\t%.2f",
//			system.Name(),
//			system.Active(),
//			system.TickFrequency(),
//			float64(system.TickCount())/system.Uptime().Seconds()))
//	}
//
//	return buf
//}
