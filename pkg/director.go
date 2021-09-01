package pkg

import (
	"github.com/gravestench/director/pkg/common"
	"github.com/gravestench/director/pkg/components"
	"github.com/gravestench/director/pkg/systems/screen_rendering"
	go_lua "github.com/yuin/gopher-lua"
	"time"

	"github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
)

type Director struct {
	*akara.World
	Lua    *go_lua.LState
	Scenes map[string]common.SceneFace
	Window struct {
		Width, Height int // pixels
		Title string
		ScaleFactor float64
	}
	TargetFPS int
}

func New() *Director {
	director := Director{}
	director.World = akara.NewWorld(akara.NewWorldConfig())

	director.Scenes = make(map[string]common.SceneFace)

	director.initDirectorSystems()

	director.Window.Width = defaultWidth
	director.Window.Height = defaultHeight
	director.Window.ScaleFactor = defaultScaleFactor
	director.Window.Title = defaultTitle

	return &director
}

func (d *Director) AddScene(scene common.SceneFace) {
	scene.Initialize(d.Window.Width, d.Window.Height, d.World, d.renderablesSubscription())
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
		ss.Update(dt)
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
}

const (
	defaultTitle = "Director"
	defaultWidth = 1028
	defaultHeight = 768
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

func (d *Director) renderablesSubscription() *akara.Subscription {
	f := d.NewComponentFilter()

	f.Require(&components.SceneGraphNode{})
	f.Require(&components.Transform{})
	f.RequireOne(&components.RenderTexture2D{}, &components.Texture2D{})

	return d.AddSubscription(f.Build())
}
