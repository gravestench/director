package main

import (
	"image/color"
	"math"
	"math/rand"
	"time"

	"github.com/faiface/mainthread"
	"github.com/gravestench/director/pkg"

	rl "github.com/gen2brain/raylib-go/raylib"
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
	key = "Director Example - Origin Test"
)

type LabelTestScene struct {
	pkg.Scene
	singleLabel common.entity
	originPoint common.entity
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

	red := color.RGBA{R: 255, A: 255}

	scene.singleLabel = scene.Add.Label("Hello World!", ww/2, wh/2, fontSize, "", randColor())
	scene.originPoint = scene.Add.Rectangle(ww/2, wh/2, 10, 10, red, nil)

	scene.Components.Debug.Add(scene.singleLabel)
}

func (scene *LabelTestScene) Update() {
	scene.updateLabel()
	scene.resizeCameraWithWindow()
}

func (scene *LabelTestScene) updateLabel() {
	ww, wh := scene.Sys.Renderer.Window.Width, scene.Sys.Renderer.Window.Height

	o, found := scene.Components.Origin.Get(scene.singleLabel)
	if !found {
		return
	}

	n := float64(time.Now().UnixNano()) / 1e9
	o.X, o.Y = (math.Sin(n)/2)+0.5, (math.Cos(n)/2)+0.5

	trs, found := scene.Components.Transform.Get(scene.singleLabel)
	if !found {
		return
	}

	trs2, found := scene.Components.Transform.Get(scene.originPoint)
	if !found {
		return
	}

	trs.Rotation.Y += 0.1

	col, found := scene.Components.Color.Get(scene.singleLabel)
	if !found {
		return
	}

	col.Color = randColor()

	trs.Translation.X = float64(ww) / 2
	trs.Translation.Y = float64(wh) / 2

	trs2.Translation.X = float64(ww) / 2
	trs2.Translation.Y = float64(wh) / 2

	//text, found := scene.Components.Text.Get(scene.singleLabel)
	//if !found {
	//	return
	//}
	//
	//t := time.Now()
	//formatted := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
	//	t.Year(), t.Month(), t.Day(),
	//	t.Hour(), t.Minute(), t.Second())
	//
	//text.String = formatted
}

func (scene *LabelTestScene) resizeCameraWithWindow() {
	ww, wh := scene.Sys.Renderer.Window.Width, scene.Sys.Renderer.Window.Height

	for _, e := range scene.Viewports {
		rt, found := scene.Components.RenderTexture2D.Get(e)
		if !found {
			continue
		}

		if int(rt.Texture.Width) != ww || int(rt.Texture.Height) != wh {
			mainthread.Call(func() {
				t := rl.LoadRenderTexture(int32(ww), int32(wh))
				rt.RenderTexture2D = &t
			})
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
