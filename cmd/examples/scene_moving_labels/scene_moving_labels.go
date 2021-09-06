package main

import (
	"fmt"
	"github.com/gravestench/director/pkg/systems/scene"
	"github.com/gravestench/mathlib"
	"image/color"
	"math"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/gravestench/akara"
)

const (
	key              = "Director Example - Moving Text"
	numTextObjects   = 1000
	maxVelocity      = 150
	maxVelocityDelta = maxVelocity / 10
)

type MovingLabelsScene struct {
	scene.Scene
	textObjects       [numTextObjects]akara.EID
	Velocity          VelocityFactory
	lastMousePosition mathlib.Vector2
}

func (scene *MovingLabelsScene) Key() string {
	return key
}

func (scene *MovingLabelsScene) IsInitialized() bool {
	if scene.Director.World == nil {
		return false
	}

	return true
}

func (scene *MovingLabelsScene) Init(w *akara.World) {
	fmt.Println("moving text scene init")

	rand.Seed(time.Now().UnixNano())

	scene.makeLabels()

	// TODO: how do we do this now
	//scene.Director.Window.Title = scene.Key()

	scene.InjectComponent(&Velocity{}, &scene.Velocity.ComponentFactory)
}

func (scene *MovingLabelsScene) makeLabels() {
	ww, wh := scene.Width, scene.Height

	fontSize := wh / 25

	for idx := range scene.textObjects {
		rx, ry := rand.Intn(ww), rand.Intn(wh)
		scene.textObjects[idx] = scene.Add.Label("", rx, ry, fontSize, "", randColor())
	}
}

func (scene *MovingLabelsScene) Update(dt time.Duration) {
	scene.updateString()

	for _, eid := range scene.textObjects {
		scene.updateVelocity(eid)
		scene.updatePosition(eid, dt)
	}

	scene.resizeCameraWithWindow()

	mp := rl.GetMousePosition()
	scene.lastMousePosition = mathlib.Vector2{
		X: float64(mp.X),
		Y: float64(mp.Y),
	}
}

func (scene *MovingLabelsScene) updateString() {
	for _, e := range scene.textObjects {
		text, found := scene.Components.Text.Get(e)
		if !found {
			continue
		}

		uuid, found := scene.Components.UUID.Get(e)
		if !found {
			continue
		}

		text.String = uuid.String()[:4]
	}
}

func (scene *MovingLabelsScene) updatePosition(eid akara.EID, dt time.Duration) {
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

	ww, wh := float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight())

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

func (scene *MovingLabelsScene) updateVelocity(eid akara.EID) {
	velocity, found := scene.Velocity.Get(eid)
	if !found {
		velocity = scene.Velocity.Add(eid)
	}

	mp := rl.GetMousePosition()
	mp2 := &mathlib.Vector2{
		X: float64(mp.X),
		Y: float64(mp.Y),
	}

	velocity.X += (rand.Float32() * maxVelocityDelta * 2) - maxVelocityDelta
	velocity.Y += (rand.Float32() * maxVelocityDelta * 2) - maxVelocityDelta

	mv := mp2.Subtract(&scene.lastMousePosition)
	velocity.X += float32(mv.X) / 2
	velocity.Y -= float32(mv.Y) / 2
}

func (scene *MovingLabelsScene) resizeCameraWithWindow() {
	for _, e := range scene.Cameras {
		rt, found := scene.Components.RenderTexture2D.Get(e)
		if !found {
			continue
		}

		if int(rt.Texture.Width) != scene.Window.Width || int(rt.Texture.Height) != scene.Window.Height {
			t := rl.LoadRenderTexture(int32(scene.Window.Width), int32(scene.Window.Height))
			rt.RenderTexture2D = &t
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
	return float32(mathlib.WrapInt(int(v - min), int(max - min))) + min
}

func randColor() color.Color {
	return &color.RGBA{
		R: uint8(rand.Intn(math.MaxUint8)),
		G: uint8(rand.Intn(math.MaxUint8)),
		B: uint8(rand.Intn(math.MaxUint8)),
		A: math.MaxUint8 - uint8(rand.Intn(math.MaxUint8>>2)),
	}
}
