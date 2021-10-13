package pkg

import (
	"github.com/faiface/mainthread"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/director/pkg/systems/animation"
	"github.com/gravestench/director/pkg/systems/renderer"
	"github.com/gravestench/director/pkg/systems/texture_manager"
	"time"

	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/systems/file_loader"
	"github.com/gravestench/director/pkg/systems/input"
	"github.com/gravestench/director/pkg/systems/screen_rendering"
	"github.com/gravestench/director/pkg/systems/tween"
	"github.com/gravestench/eventemitter"
)

// Director provides a scene management abstraction, with
// supporting systems. scenes are basically a superset of
// the functionality provided by an akara.System, but with
// a bunch of object creation facilities provided for free.
type Director struct {
	*akara.World
	scenes map[string]SceneInterface
	Sys    DirectorSystems
}

// DirectorSystems contains the base systems that are available when a director instance is created
type DirectorSystems struct {
	Events   *eventemitter.EventEmitter
	Load     *file_loader.System
	Renderer *renderer.System
	Texture  *texture_manager.System
	Tweens   *tween.System
	Input    *input.System
}

// New creates a new director instance, with default settings
func New() *Director {
	director := Director{}
	director.World = akara.NewWorld(akara.NewWorldConfig())
	director.Sys.Events = eventemitter.New()

	director.scenes = make(map[string]SceneInterface)

	director.initDirectorSystems()

	return &director
}

// AddScene adds a scene
func (d *Director) AddScene(scene SceneInterface) {
	scene.GenericSceneInit(d)
	scene.InitializeLua()

	d.AddSystem(scene, true)
	d.scenes[scene.Key()] = scene
}

// RemoveScene queues a scene for removal
func (d *Director) RemoveScene(key string) *Director {
	if ss, found := d.scenes[key]; found {
		d.RemoveSystem(ss)
		delete(d.scenes, key)
	}

	return d
}

// Update calls World.Update()
func (d *Director) Update(dt time.Duration) (err error) {
	return d.World.Update(dt)
}

// initDirectorSystems creates all of the systems that scenes will need.
// These systems are set up to do specific things during the update loop.
// For example, the tween system iterates over and processes all tweens, every update.
// likewise, input, file loading, texture management, etc are all functions that
// have been broken into their own systems.
func (d *Director) initDirectorSystems() {
	screenRendering := &screen_rendering.ScreenRenderingSystem{}
	d.AddSystem(screenRendering, true)

	d.Sys.Tweens = &tween.System{}
	d.AddSystem(d.Sys.Tweens, true)

	d.Sys.Renderer = &renderer.System{}
	d.AddSystem(d.Sys.Renderer, true)

	d.Sys.Input = &input.System{}
	d.AddSystem(d.Sys.Input, true)

	d.Sys.Load = &file_loader.System{}
	d.AddSystem(d.Sys.Load, true)

	d.Sys.Texture = &texture_manager.System{}
	d.AddSystem(d.Sys.Texture, true)

	d.AddSystem(&animation.System{}, true)
}

func (d *Director) Run() error {
	// mainthread.CallQueueCap = 16
	mainthread.Run(d.run)

	return nil
}

// run the director game loop. this is a blocking operation.
func (d *Director) run() {
	defer d.Stop()

	// open the raylib window. It'd be nice to have this in the renderer system, but not currently possible due to the
	// order that things are Init'd in.
	ww, wh := int32(d.Sys.Renderer.Window.Width), int32(d.Sys.Renderer.Window.Height)
	mainthread.Call(func() {
		rl.SetTraceLog(rl.LogNone)
		rl.InitWindow(ww, wh, d.Sys.Renderer.Window.Title)
	})

	now := time.Now()
	last := now

	var delta time.Duration

	ticker := time.NewTicker(time.Millisecond)
	defer ticker.Stop()
	for now = range ticker.C {
		if !d.Sys.Renderer.Window.IsOpen {
			break
		}

		delta = now.Sub(last)
		last = now

		if err := d.Update(delta); err != nil {
			panic(err)
		}
	}
}

// Stop deactivates all the Director's systems
func (d *Director) Stop() {
	for _, system := range d.Systems {
		system.Deactivate()
	}
}