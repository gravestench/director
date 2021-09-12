package main

import (
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/systems/scene"
	"image"
	"image/color"
)

type testScene struct {
	scene.Scene
	highest int
}

const (
	numObjects = 10
)

func (scene *testScene) Key() string {
	return "test"
}

func (scene *testScene) IsInitialized() bool {
	return true
}

func (scene *testScene) Init(_ *akara.World) {
	startX, startY, step := 100, 100, 30
	w, h := 80, 80

	c := color.RGBA{R: 255, A: 255}
	white := color.RGBA{R: 255, G: 255, B: 255, A: 255}

	for idx := 0; idx < numObjects; idx++ {
		x := startX + (idx * step)
		y := startY + (idx * step)

		c.R -= uint8(step)
		c.G += uint8(step*2)
		c.B = uint8(step*16)
		e := scene.Add.Rectangle(x, y, w, h, c, white)

		ro, found := scene.Components.RenderOrder.Get(e)
		if !found {
			continue
		}

		in := scene.Components.Interactive.Add(e)

		size, found := scene.Components.Size.Get(e)
		if !found {
			return
		}

		trs, found := scene.Components.Transform.Get(e)
		if !found {
			return
		}

		rHeight := scene.Sys.Renderer.Window.Height

		in.Hitbox = &image.Rectangle{
			Min: image.Point{
				X: int(trs.Translation.X) - size.Dx()/2,
				Y: rHeight - (int(trs.Translation.Y) + size.Dy()/2),
			},
			Max: image.Point{
				X: int(trs.Translation.X) + size.Dx()/2,
				Y: rHeight - (int(trs.Translation.Y) - size.Dy()/2),
			},
		}

		in.Callback = func() (preventPropogation bool) {
			scene.highest++
			ro.Value = scene.highest

			return false
		}

		ro.Value = numObjects - idx
	}

	scene.highest = numObjects
}

func (scene *testScene) Update() { /* noop */ }