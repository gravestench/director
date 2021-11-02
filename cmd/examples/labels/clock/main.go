package main

import (
	"fmt"
	"github.com/gravestench/director/pkg"
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/gravestench/akara"
	"github.com/gravestench/director"
)

func main() {
	d := director.New()

	d.AddScene(&LabelTestScene{})

	if err := d.Run(); err != nil {
		panic(err)
	}
}

const (
	key = "Director Example - Label Test"
)

type LabelTestScene struct {
	pkg.Scene
	singleLabel akara.EID
}

func (scene *LabelTestScene) Key() string {
	return key
}

func (scene *LabelTestScene) IsInitialized() bool {
	return true
}

func (scene *LabelTestScene) Init(w *akara.World) {
	scene.makeLabels()
}

func (scene *LabelTestScene) makeLabels() {
	ww, wh := scene.Sys.Renderer.Window.Width, scene.Sys.Renderer.Window.Height
	fontSize := wh / 10

	scene.singleLabel = scene.Add.Label("", ww/2, wh/2, fontSize, "", randColor())
}

func (scene *LabelTestScene) Update() {
	scene.updateLabel()
}

func (scene *LabelTestScene) updateLabel() {
	ww, wh := scene.Sys.Renderer.Window.Width, scene.Sys.Renderer.Window.Height

	trs, found := scene.Components.Transform.Get(scene.singleLabel)
	if !found {
		return
	}

	col, found := scene.Components.Color.Get(scene.singleLabel)
	if !found {
		return
	}

	col.Color = randColor()

	trs.Translation.X = float64(ww) / 2
	trs.Translation.Y = float64(wh) / 2

	text, found := scene.Components.Text.Get(scene.singleLabel)
	if !found {
		return
	}

	t := time.Now()
	formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())

	text.String = formatted
}

func randColor() color.Color {
	return &color.RGBA{
		R: uint8(rand.Intn(math.MaxUint8)),
		G: uint8(rand.Intn(math.MaxUint8)),
		B: uint8(rand.Intn(math.MaxUint8)),
		A: math.MaxUint8 - uint8(rand.Intn(math.MaxUint8>>2)),
	}
}
