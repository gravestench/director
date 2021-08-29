package pkg

import (
	"time"

	"github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
)

type Director struct {
	*akara.World
	scenes map[string]SceneFace
	Window struct {
		Width, Height int // pixels
		Title string
		ScaleFactor float64
	}
	TargetFPS int
}

func New() *Director {
	director := &Director{}
	director.World = akara.NewWorld(akara.NewWorldConfig())

	director.scenes = make(map[string]SceneFace)

	director.initSceneSystems()

	director.Window.Width = defaultWidth
	director.Window.Height = defaultHeight
	director.Window.ScaleFactor = defaultScaleFactor
	director.Window.Title = defaultTitle

	return director
}

func (d *Director) AddScene(scene SceneFace) *Director {
	scene.bind(d)

	d.AddSystem(scene)
	d.scenes[scene.Key()] = scene

	return d
}

func (d *Director) RemoveScene(key string) *Director {
	if scene, found := d.scenes[key]; found {
		d.RemoveSystem(scene)
		delete(d.scenes, key)
		scene.bind(nil)
	}

	return d
}

func (d *Director) Update(dt time.Duration) (err error) {
	err = d.updateSceneGraphs(dt)
	if err != nil {
		return err
	}

	err = d.updateSceneObjects(dt)
	if err != nil {
		return err
	}

	return d.renderScenes()
}

func (d *Director) updateSceneGraphs(dt time.Duration) error {
	for idx := range d.scenes {
		d.scenes[idx].updateSceneGraph()
	}

	return d.World.Update(dt)
}

func (d *Director) updateSceneObjects(dt time.Duration) error {
	for idx := range d.scenes {
		d.scenes[idx].updateSceneObjects(dt)
	}

	return d.World.Update(dt)
}

func (d *Director) renderScenes() error {
	for idx := range d.scenes {
		if err := d.scenes[idx].render(); err != nil {
			return err
		}
	}

	return nil
}

func (d *Director) initSceneSystems() {
	d.AddSystem(&renderSystem{})
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

		dpi:= rl.GetWindowScaleDPI()
		_ = dpi
	}

	return nil
}