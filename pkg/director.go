package pkg

import (
	"github.com/gravestench/director/pkg/systems/animation"
	"github.com/gravestench/director/pkg/systems/renderer"
	"github.com/gravestench/director/pkg/systems/texture_manager"
	"time"

	"github.com/gen2brain/raylib-go/raylib"
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
	Sys directorSystems
}

// contains the base systems that are available when a director instance is created
type directorSystems struct {
	Events  *eventemitter.EventEmitter
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

	d.AddSystem(scene)
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

// Update iterates over all of the scenes, updating and rendering each.
// Because scenes are actually implementing the akara.System interface,
// this only calls all of tghe generic update methods, and akara calls the
// actual update methods at the end during world.Update
func (d *Director) Update(dt time.Duration) (err error) {
	d.updateScenes(dt)

	return d.World.Update(dt)
}

// updateScenes calls the generic scene update method for each scene
func (d *Director) updateScenes(dt time.Duration) {
	for _, ss := range d.scenes {
		ss.GenericUpdate(dt)
	}

	// this renders the scene objects to the scene's render texture
	// however, this will not actually display anything, that is done by the render system
	for idx := range d.scenes {
		d.scenes[idx].Render()
	}
}

// initDirectorSystems creates all of the systems that scenes will need.
// These systems are set up to do specific things during the update loop.
// For example, the tween system iterates over and processes all tweens, every update.
// likewise, input, file loading, texture management, etc are all functions that
// have been broken into their own systems.
func (d *Director) initDirectorSystems() {
	d.AddSystem(&screen_rendering.ScreenRenderingSystem{})

	d.Sys.Tweens = &tween.System{}
	d.AddSystem(d.Sys.Tweens)

	d.Sys.Renderer = &renderer.System{}
	d.AddSystem(d.Sys.Renderer)

	d.Sys.Input = &input.System{}
	d.AddSystem(d.Sys.Input)

	d.Sys.Load = &file_loader.System{}
	d.AddSystem(d.Sys.Load)

	d.Sys.Texture = &texture_manager.System{}
	d.AddSystem(d.Sys.Texture)

	d.AddSystem(&animation.System{})
}

// Run the director game loop. this is a blocking operation.
func (d *Director) Run() error {
	now := time.Now()
	last := now

	var delta time.Duration

	ww, wh := int32(d.Sys.Renderer.Window.Width), int32(d.Sys.Renderer.Window.Height)

	rl.SetTraceLog(rl.LogNone)
	rl.InitWindow(ww, wh, d.Sys.Renderer.Window.Title)
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
