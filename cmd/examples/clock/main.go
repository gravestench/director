package main

import (
	"fmt"
	"image/color"
	"math"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"

	"github.com/gravestench/akara"
	director "github.com/gravestench/director/pkg"
)

const key = "Clock Scene"

const numClocks = 50

type clockScene struct {
	director.Scene
	elapsed time.Duration
	seconds int
	updates int
	textObjects [numClocks]akara.EID
	Velocity VelocityFactory
}

func (c *clockScene) Key() string {
	return key
}

func (c *clockScene) IsInitialized() bool {
	return c.Director.World != nil
}

func (c *clockScene) Init(w *akara.World) {
	fmt.Println("clock scene init")

	rand.Seed(time.Now().UnixNano())

	ww, wh := rl.GetScreenWidth(), rl.GetScreenHeight()

	for idx := range c.textObjects {
		c.textObjects[idx] = c.Add.Text("", ww/2, wh/2, randColor())
	}

	c.Director.World = w
	c.BaseSystem.World = w
	c.InjectComponent(&Velocity{}, &c.Velocity.ComponentFactory)
}

func (c *clockScene) Update(delta time.Duration) {
	c.elapsed += delta
	c.updates++

	c.updateString()

	for _, eid := range c.textObjects {
		c.updatePosition(eid)
	}
}

func (c *clockScene) updateString() {
	mes := int(c.elapsed.Milliseconds())
	es := int(c.elapsed.Seconds())

	c.seconds = es
	ms, s, m, h := mes % 1000 , es % 60, (es / 60) % 60, (es / 60 / 60) % 24

	fps := fmt.Sprintf("%.02f FPS", rl.GetFPS())
	str := fmt.Sprintf("%02d:%02d:%02d:%04d \n(%v)", h, m, s, ms, fps)

	for _, eid := range c.textObjects {
		if text, found := c.Text.Get(eid); found {
			text.String = str
		}
	}

	c.updates = 0
}

func (c *clockScene) updatePosition(eid akara.EID) {
	position, found := c.Vector2.Get(eid)
	if !found  {
		return
	}

	velocity, found := c.Velocity.Get(eid)
	if !found  {
		velocity = c.Velocity.Add(eid)
	}

	const max = 8

	position.X += velocity.X
	position.Y += velocity.Y

	velocity.X += (rand.Float32() * 2) - 1
	velocity.Y += (rand.Float32() * 2) - 1

	velocity.X = float32(clamp(velocity.X, -max, max))
	velocity.Y = float32(clamp(velocity.Y, -max, max))

	ww, wh := float32(rl.GetScreenWidth()), float32(rl.GetScreenHeight())
	tw, th := float32(400), float32(140)

	position.X = float32(wrap(position.X, -tw, ww))
	position.Y = float32(wrap(position.Y, -th, wh))
}

func main() {
	d := director.New()

	d.Window.Width = 1024
	d.Window.Height = 768
	d.Window.Title = "clock test"
	d.TargetFPS = 60

	d.AddScene(&clockScene{})

	if err := d.Run(); err != nil {
		panic(err)
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
	if v > max {
		v = min
	} else if v < min {
		v = max
	}

	return v
}

func randColor() color.Color {
	return &color.RGBA{
		R: uint8(rand.Intn(math.MaxUint8)),
		G: uint8(rand.Intn(math.MaxUint8)),
		B: uint8(rand.Intn(math.MaxUint8)),
		A: uint8(rand.Intn(math.MaxUint8)),
	}
}

