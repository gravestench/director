package pkg

import (
	"github.com/gravestench/director/pkg/systems/animation"
	"github.com/gravestench/director/pkg/systems/texture_manager"
	"time"

	"github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
	"github.com/gravestench/eventemitter"
	go_lua "github.com/yuin/gopher-lua"

	"github.com/gravestench/director/pkg/systems/file_loader"
	"github.com/gravestench/director/pkg/systems/input"
	"github.com/gravestench/director/pkg/systems/screen_rendering"
	"github.com/gravestench/director/pkg/systems/tween"
)

type Director struct {
	*akara.World
	Lua     *go_lua.LState
	Events  *eventemitter.EventEmitter
	Scenes  map[string]Scene
	Load    *file_loader.System
	Texture *texture_manager.System
	Tweens  *tween.System
	Input   *input.System
	Window  struct {
		Width, Height int // pixels
		Title         string
		ScaleFactor   float64
	}
	TargetFPS int
}

func New() *Director {
	director := Director{}
	director.World = akara.NewWorld(akara.NewWorldConfig())
	director.Events = eventemitter.New()

	director.Scenes = make(map[string]Scene)

	director.initDirectorSystems()

	director.Window.Width = defaultWidth
	director.Window.Height = defaultHeight
	director.Window.ScaleFactor = defaultScaleFactor
	director.Window.Title = defaultTitle

	return &director
}

func (d *Director) AddScene(scene Scene) {
	scene.GenericSceneInit(d, d.Window.Width, d.Window.Height)
	scene.InitializeLua()

	d.AddSystem(scene)
	d.Scenes[scene.Key()] = scene
}

func (d *Director) RemoveScene(key string) *Director {
	if ss, found := d.Scenes[key]; found {
		d.RemoveSystem(ss)
		delete(d.Scenes, key)
	}

	return d
}

func (d *Director) Update(dt time.Duration) (err error) {
	d.updateState()

	d.updateScenes(dt)

	// this renders the scene objects to the scene's render texture
	// however, this will not actually display anything, that is done by the render system
	d.renderScenes()

	return d.World.Update(dt)
}

func (d *Director) updateState() {
	w := rl.GetScreenWidth()
	h := rl.GetScreenHeight()
	d.Window.Width, d.Window.Height = w, h
}

func (d *Director) updateScenes(dt time.Duration) {
	for _, ss := range d.Scenes {
		ss.GenericUpdate(dt)
	}
}

func (d *Director) renderScenes() {
	for idx := range d.Scenes {
		d.Scenes[idx].Render()
	}
}

func (d *Director) initDirectorSystems() {
	d.AddSystem(&screen_rendering.ScreenRenderingSystem{})

	d.Tweens = &tween.System{}
	d.AddSystem(d.Tweens)

	d.Input = &input.System{}
	d.AddSystem(d.Input)

	d.Load = &file_loader.System{}
	d.AddSystem(d.Load)

	d.Texture = &texture_manager.System{}
	d.AddSystem(d.Texture)

	d.AddSystem(&animation.System{})
}

const (
	defaultTitle       = "Director"
	defaultWidth       = 1028
	defaultHeight      = 768
	defaultScaleFactor = 1.0
)

func (d *Director) Run() error {
	now := time.Now()
	last := now

	var delta time.Duration

	ww, wh := int32(d.Window.Width), int32(d.Window.Height)

	if ww <= 1 {
		ww = defaultWidth
	}

	if wh <= 1 {
		wh = defaultHeight
	}

	rl.SetTargetFPS(int32(d.TargetFPS))

	rl.InitWindow(ww, wh, d.Window.Title)
	defer rl.CloseWindow()

	for !rl.WindowShouldClose() {
		now = time.Now()
		delta = now.Sub(last)
		last = now

		if err := d.Update(delta); err != nil {
			return err
		}
	}

	return nil
}
