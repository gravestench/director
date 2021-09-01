package main

import (
	"fmt"
	"github.com/gravestench/director/pkg/systems/scene"
	"image/color"
	"math"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
	director "github.com/gravestench/director/pkg"
)

func main() {
	d := director.New()

	d.Window.Width = 1024
	d.Window.Height = 768
	d.TargetFPS = 60

	d.AddScene(&LabelTestScene{})

	if err := d.Run(); err != nil {
		panic(err)
	}
}

const (
	key              = "Director Example - Label Test"
)

type LabelTestScene struct {
	scene.Scene
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
	fontSize := scene.Height / 10

	scene.singleLabel = scene.Add.Label("", scene.Width/2, scene.Height/2, fontSize, "", randColor())
}

func (scene *LabelTestScene) Update(dt time.Duration) {
	scene.updateLabel()
	scene.resizeCameraWithWindow()
}

func (scene *LabelTestScene) updateLabel() {
	ww, wh := scene.Width, scene.Height

	trs, found := scene.Components.Transform.Get(scene.singleLabel)
	if !found {
		return
	}

	col, found := scene.Components.Color.Get(scene.singleLabel)
	if !found {
		return
	}

	col.Color = randColor()

	trs.Translation.X = float64(ww)/2
	trs.Translation.Y = float64(wh)/2

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

func (scene *LabelTestScene) resizeCameraWithWindow() {
	for _, e := range scene.Cameras {
		rt, found := scene.Components.RenderTexture2D.Get(e)
		if !found {
			continue
		}

		if int(rt.Texture.Width) != scene.Width || int(rt.Texture.Height) != scene.Height {
			t := rl.LoadRenderTexture(int32(scene.Width), int32(scene.Height))
			rt.RenderTexture2D = &t
		}
	}
}

func randColor() color.Color {
	return &color.RGBA{
		R: uint8(rand.Intn(math.MaxUint8)),
		G: uint8(rand.Intn(math.MaxUint8)),
		B: uint8(rand.Intn(math.MaxUint8)),
		A: math.MaxUint8 - uint8(rand.Intn(math.MaxUint8>>2)),
	}
}
