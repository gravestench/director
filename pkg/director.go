package pkg

import (
	"flag"
	"log"
	"os"
	"runtime/trace"
	"time"

	"github.com/gravestench/director/pkg/components/audio"

	"github.com/gravestench/director/pkg/components/viewport"

	"runtime/pprof"

	"github.com/faiface/mainthread"
	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/gravestench/akara"
	"github.com/gravestench/eventemitter"

	"github.com/gravestench/director/pkg/components/animation"
	"github.com/gravestench/director/pkg/components/camera"
	"github.com/gravestench/director/pkg/components/color"
	"github.com/gravestench/director/pkg/components/debug"
	fileLoadRequest "github.com/gravestench/director/pkg/components/file_load_request"
	fileLoadResponse "github.com/gravestench/director/pkg/components/file_load_response"
	fileType "github.com/gravestench/director/pkg/components/file_type"
	"github.com/gravestench/director/pkg/components/fill"
	"github.com/gravestench/director/pkg/components/font"
	hasChildren "github.com/gravestench/director/pkg/components/has_children"
	"github.com/gravestench/director/pkg/components/interactive"
	"github.com/gravestench/director/pkg/components/opacity"
	"github.com/gravestench/director/pkg/components/origin"
	renderOrder "github.com/gravestench/director/pkg/components/render_order"
	"github.com/gravestench/director/pkg/components/render_texture"
	sceneGraphNode "github.com/gravestench/director/pkg/components/scene_graph_node"
	"github.com/gravestench/director/pkg/components/size"
	"github.com/gravestench/director/pkg/components/stroke"
	"github.com/gravestench/director/pkg/components/text"
	texture2D "github.com/gravestench/director/pkg/components/texture"
	"github.com/gravestench/director/pkg/components/transform"
	"github.com/gravestench/director/pkg/components/uuid"

	animationSystem "github.com/gravestench/director/pkg/systems/animation"
	audioSystem "github.com/gravestench/director/pkg/systems/audio"
	"github.com/gravestench/director/pkg/systems/file_loader"
	"github.com/gravestench/director/pkg/systems/input"
	"github.com/gravestench/director/pkg/systems/renderer"
	"github.com/gravestench/director/pkg/systems/screen_rendering"
	"github.com/gravestench/director/pkg/systems/texture_manager"
	"github.com/gravestench/director/pkg/systems/tween"
)

// Director provides a scene management abstraction, with
// supporting systems. scenes are basically a superset of
// the functionality provided by an akara.System, but with
// a bunch of object creation facilities provided for free.
type Director struct {
	*akara.World
	scenes     map[string]SceneInterface
	Sys        DirectorSystems
	Components DirectorComponents
}

// DirectorSystems contains the base systems that are available when a director instance is created
type DirectorSystems struct {
	Events   *eventemitter.EventEmitter
	Load     *file_loader.System
	Renderer *renderer.System
	Texture  *texture_manager.System
	Tweens   *tween.System
	Input    *input.System
	Audio    *audioSystem.System
}

// DirectorComponents contains all of the primitive components that come with director.
// These are INTENTIONALLY nil instances, which can be used when creating component filters.
// See akara's `World.NewComponentFilter` and `World.AddSubscription`
type DirectorComponents struct {
	audio            *audio.Component
	interactive      *interactive.Component
	viewport         *viewport.Component
	camera           *camera.Camera
	color            *color.Color
	debug            *debug.Component
	fileLoadRequest  *fileLoadRequest.Component
	fileLoadResponse *fileLoadResponse.Component
	fileType         *fileType.Component
	fill             *fill.Component
	hasChildren      *hasChildren.Component
	animation        *animation.Component
	stroke           *stroke.Component
	font             *font.Component
	opacity          *opacity.Component
	origin           *origin.Component
	render_texture   *render_texture.Component
	renderOrder      *renderOrder.Component
	size             *size.Component
	sceneGraphNode   *sceneGraphNode.Component
	text             *text.Component
	texture2D        *texture2D.Component
	transform        *transform.Component
	uuid             *uuid.Component
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

	d.AddSystem(scene)
	d.scenes[scene.Key()] = scene
}

// AddScene adds a scene
func (d *Director) AddSystem(sys akara.System) {
	if sceneGeneric, ok := sys.(isGeneric); ok {
		sceneGeneric.GenericSceneInit(d)
	}

	if luaInitializer, ok := sys.(initializesLua); ok {
		luaInitializer.InitializeLua()
	}

	d.World.AddSystem(sys, true)
}

// AddLuaScene creates and adds a scene using a lua script
func (d *Director) AddLuaScene(key, path string) {
	d.AddScene(NewLuaScene(key, path))
}

// AddLuaSystem creates and adds a scene using a lua script
func (d *Director) AddLuaSystem(path string) {
	d.AddSystem(NewLuaSystem(path))
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
	return d.World.Update()
}

// initDirectorSystems creates all of the systems that scenes will need.
// These systems are set up to do specific things during the update loop.
// For example, the tween system iterates over and processes all tweens, every update.
// likewise, input, file loading, texture management, etc are all functions that
// have been broken into their own systems.
func (d *Director) initDirectorSystems() {
	screenRendering := &screen_rendering.ScreenRenderingSystem{}
	d.AddSystem(screenRendering)
	screenRendering.SetTickFrequency(0) // unlimited FPS. TODO: set this from somewhere?

	d.Sys.Tweens = &tween.System{}
	d.AddSystem(d.Sys.Tweens)

	d.Sys.Renderer = &renderer.System{}
	d.AddSystem(d.Sys.Renderer)

	d.Sys.Input = &input.System{}
	d.AddSystem(d.Sys.Input)
	d.Sys.Input.SetTickFrequency(1000)

	d.Sys.Load = &file_loader.System{}
	d.AddSystem(d.Sys.Load)

	d.Sys.Texture = &texture_manager.System{}
	d.AddSystem(d.Sys.Texture)

	d.AddSystem(&animationSystem.System{})

	d.Sys.Audio = &audioSystem.System{}
	d.AddSystem(d.Sys.Audio)
}

func (d *Director) Run() error {
	if f := flag.Lookup(FlagNameProfileCPU); f != nil {
		if f.Value != nil {
			path := f.Value.(flag.Getter).Get().(string)

			f, err := os.Create(path)
			if err == nil {
				log.Printf("%s: %s\n", "begin cpu profile", path)

				_ = pprof.StartCPUProfile(f)
				defer pprof.StopCPUProfile()
			}
		}
	}

	if f := flag.Lookup(FlagNameTrace); f != nil {
		if f.Value != nil && f.Value.(flag.Getter).Get().(bool) {
			f, err := os.Create("trace.out")
			if err != nil {
				log.Fatalf("failed to create trace output file: %v", err)
			}

			defer func() {
				if err := f.Close(); err != nil {
					log.Fatalf("failed to close trace file: %v", err)
				}
			}()

			if err := trace.Start(f); err != nil {
				log.Fatalf("failed to start trace: %v", err)
			}

			defer trace.Stop()
		}
	}

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
