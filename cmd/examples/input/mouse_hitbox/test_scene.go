package main

import (
	"fmt"
	"image"
	"image/color"

	"github.com/gravestench/director/pkg/systems/input/constants"

	"github.com/gravestench/director/pkg"

	"github.com/gravestench/akara"
)

type TestScene struct {
	pkg.Scene
	square common.entity
	label  common.entity
}

func (scene *TestScene) Key() string {
	return "test"
}

func (scene *TestScene) Update() {
	scene.updateLabel()
	scene.resetSquare()
}

func (scene *TestScene) updateLabel() {
	text, found := scene.Components.Text.Get(scene.label)
	if !found {
		return
	}

	const (
		fmtMouse = "Mouse (%.2f, %.2f)"
	)

	text.String = fmt.Sprintf(fmtMouse, scene.Sys.Input.MousePosition.X, scene.Sys.Input.MousePosition.Y)
}

func (scene *TestScene) Init(world *akara.World) {
	scene.makeSquare()
	scene.makeLabel()
	scene.bindInput()
}

func (scene *TestScene) makeSquare() {
	blue := color.RGBA{B: 255, A: 255}
	scene.square = scene.Add.Rectangle(100, 100, 30, 30, blue, nil)
}

func (scene *TestScene) resetSquare() {
	fill, found := scene.Components.Fill.Get(scene.square)
	if !found {
		return
	}

	r, g, b, a := fill.Color.RGBA()
	if g > 0 {
		g -= 1
		fill.Color = color.RGBA{
			R: uint8(r),
			G: uint8(g),
			B: uint8(b),
			A: uint8(a),
		}
	}
}

func (scene *TestScene) setSquareColor(c color.Color) {
	fill, found := scene.Components.Fill.Get(scene.square)
	if !found {
		return
	}

	fill.Color = c
}

func (scene *TestScene) makeLabel() {
	red := color.RGBA{R: 255, A: 255}
	scene.label = scene.Add.Label("", 400, 400, 24, "", red)
}

func (scene *TestScene) bindInput() {
	i := scene.Components.Interactive.Add(scene.square)

	i.Callback = func() (preventPropogation bool) {
		yellow := color.RGBA{R: 255, G: 255, A: 255}

		scene.setSquareColor(yellow)

		return false
	}

	i.Vector.SetMouseButton(constants.MouseButtonLeft)

	size, found := scene.Components.Size.Get(scene.square)
	if !found {
		return
	}

	trs, found := scene.Components.Transform.Get(scene.square)
	if !found {
		return
	}

	rHeight := scene.Sys.Renderer.Window.Height

	i.Hitbox = &image.Rectangle{
		Min: image.Point{
			X: int(trs.Translation.X) - size.Dx()/2,
			Y: rHeight - (int(trs.Translation.Y) + size.Dy()/2),
		},
		Max: image.Point{
			X: int(trs.Translation.X) + size.Dx()/2,
			Y: rHeight - (int(trs.Translation.Y) - size.Dy()/2),
		},
	}
}

func (scene *TestScene) IsInitialized() bool {
	return scene.square != 0
}
