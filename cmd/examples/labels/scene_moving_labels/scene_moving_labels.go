package main

import (
	"fmt"
	"github.com/faiface/mainthread"
	"github.com/gravestench/director/pkg/systems/scene"
	"github.com/gravestench/mathlib"
	"image/color"
	"math"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"

	"github.com/gravestench/director/pkg/common"
)

const (
	key              = "Director Example - Moving Text"
	numTextObjects   = 1000
	maxVelocity      = 150
	maxVelocityDelta = maxVelocity / 10
)

type MovingLabelsScene struct {
	scene.Scene
	textObjects          [numTextObjects]common.Entity
	Velocity             VelocityFactory
	lastMousePosition    mathlib.Vector2
	currentMousePosition mathlib.Vector2
}

func (scene *MovingLabelsScene) Key() string {
	return key
}

func (scene *MovingLabelsScene) IsInitialized() bool {
	return scene.Director.World != nil
}

func (scene *MovingLabelsScene) Init(w *akara.World) {
	fmt.Println("moving text scene init")

	rand.Seed(time.Now().UnixNano())

	go scene.makeLabels()

	scene.Sys.Renderer.Window.Title = scene.Key()

	scene.InjectComponent(&Velocity{}, &scene.Velocity.ComponentFactory)
}

var messages = []string{
	"BRB",
	"be",
	"right",
	"back",
	"I'll be back in a moment",
	"I'm just running to the store",
	"please wait",
	"getting booze",
	"went to the store",
	"enjoy the music",
	"fork me on github!",
}

func (scene *MovingLabelsScene) makeLabels() {
	ww, wh := scene.Sys.Renderer.Window.Width, scene.Sys.Renderer.Window.Height

	fontSize := wh / 25

	for idx := range scene.textObjects {
		rx, ry := rand.Intn(ww), rand.Intn(wh)
		ri := rand.Intn(len(messages))
		scene.textObjects[idx] = scene.Add.Label(messages[ri], rx, ry, fontSize, "", randColor())
	}
}

func (scene *MovingLabelsScene) Update() {
	scene.lastMousePosition = scene.currentMousePosition
	var mp rl.Vector2
	mainthread.Call(func() {
		mp = rl.GetMousePosition()
	})
	scene.currentMousePosition = mathlib.Vector2{
		X: float64(mp.X),
		Y: float64(mp.Y),
	}

	for _, eid := range scene.textObjects {
		scene.updateVelocity(eid)
		scene.updatePosition(eid, scene.Director.TimeDelta)
	}

	scene.resizeCameraWithWindow()
}

func (scene *MovingLabelsScene) updatePosition(eid common.Entity, dt time.Duration) {
	trs, found := scene.Components.Transform.Get(eid)
	if !found {
		return
	}

	position := trs.Translation

	velocity, found := scene.Velocity.Get(eid)
	if !found {
		velocity = scene.Velocity.Add(eid)
		velocity.X = (rand.Float32() - 0.5) * maxVelocity
		velocity.Y = (rand.Float32() - 0.5) * maxVelocity
	}

	velocity.X = clamp(velocity.X, -maxVelocity, maxVelocity)
	velocity.Y = clamp(velocity.Y, -maxVelocity, maxVelocity)

	scalar := float64(dt.Seconds())

	position.X += float64(velocity.X) * scalar
	position.Y += float64(velocity.Y) * scalar

	ww, wh := float32(scene.Director.Sys.Renderer.Window.Width), float32(scene.Director.Sys.Renderer.Window.Width)

	var tw, th int32

	t, tFound := scene.Components.Texture2D.Get(eid)
	rt, rtFound := scene.Components.RenderTexture2D.Get(eid)
	if tFound {
		tw, th = t.Texture2D.Width, t.Texture2D.Height
	} else if rtFound {
		tw, th = rt.Texture.Width, rt.Texture.Height
	}

	position.X = float64(wrap(float32(position.X), float32(-tw), ww+float32(tw)))
	position.Y = float64(wrap(float32(position.Y), float32(-th), wh+float32(th)))
}

func (scene *MovingLabelsScene) updateVelocity(eid common.Entity) {
	velocity, found := scene.Velocity.Get(eid)
	if !found {
		velocity = scene.Velocity.Add(eid)
	}

	velocity.X += (rand.Float32() * maxVelocityDelta * 2) - maxVelocityDelta
	velocity.Y += (rand.Float32() * maxVelocityDelta * 2) - maxVelocityDelta

	// copy these vectors because Subtract() mutates them
	currentMousePos := scene.currentMousePosition
	lastMousePos := scene.lastMousePosition
	mv := currentMousePos.Subtract(&lastMousePos)
	velocity.X += float32(mv.X) / 2
	velocity.Y -= float32(mv.Y) / 2
}

func (scene *MovingLabelsScene) resizeCameraWithWindow() {
	for _, e := range scene.Viewports {
		rt, found := scene.Components.RenderTexture2D.Get(e)
		if !found {
			continue
		}

		rWidth := scene.Sys.Renderer.Window.Width
		rHeight := scene.Sys.Renderer.Window.Height

		if int(rt.Texture.Width) != rWidth || int(rt.Texture.Height) != rHeight {
			mainthread.Call(func() {
				t := rl.LoadRenderTexture(int32(rWidth), int32(rHeight))
				rt.RenderTexture2D = &t
			})
		}
	}
}

func clamp(v, min, max float32) float32 {
	if v > max {
		v = max
	} else if v < min {
		v = min
	}

	return v
}

func wrap(v, min, max float32) float32 {
	return float32(mathlib.WrapInt(int(v-min), int(max-min))) + min
}

func randColor() color.Color {
	return &color.RGBA{
		R: uint8(rand.Intn(math.MaxUint8)),
		G: uint8(rand.Intn(math.MaxUint8)),
		B: uint8(rand.Intn(math.MaxUint8)),
		A: math.MaxUint8 - uint8(rand.Intn(math.MaxUint8>>2)),
	}
}
