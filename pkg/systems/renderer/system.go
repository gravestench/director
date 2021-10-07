package renderer

import (
	"github.com/faiface/mainthread"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
)

const (
	defaultTitle       = "Director"
	defaultWidth       = 1024
	defaultHeight      = 768
	defaultFPS         = 768
	defaultScaleFactor = 1.0
	defaultLogging     = rl.LogNone
)

type System struct {
	akara.BaseSystem
	Window struct {
		Width, Height int // pixels
		Title         string
		ScaleFactor   float64
	}
	Logging   int
	TargetFPS int
}

func (s *System) New(args ...interface{}) {
	if !s.IsInitialized() {
		s.Init(nil)
	}
}

func (s *System) Update() {
	mainthread.Call(func() {
		s.Window.Width, s.Window.Height = rl.GetScreenWidth(), rl.GetScreenHeight()

		if s.Window.Width <= 1 {
			s.Window.Width = defaultWidth
		}

		if s.Window.Height <= 1 {
			s.Window.Height = defaultHeight
		}

		// rl.SetTargetFPS(int32(s.TargetFPS))
		rl.SetTraceLog(rl.LogNone)
	})
}

func (s *System) Init(_ *akara.World) {
	s.Window.Width = defaultWidth
	s.Window.Height = defaultHeight
	s.Window.ScaleFactor = defaultScaleFactor
	s.Window.Title = defaultTitle
	s.TargetFPS = defaultFPS
	s.Logging = defaultLogging
}

func (s *System) IsInitialized() bool {
	return true
}
