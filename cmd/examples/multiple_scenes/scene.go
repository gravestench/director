package main

import (
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/common"
	"github.com/gravestench/director/pkg/easing"
	"github.com/gravestench/director/pkg/systems/tween"
	"github.com/gravestench/mathlib"
	"image/color"
	"math"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/director/pkg/systems/scene"
)

const (
	maxImages        = 20
	newImageInterval = time.Second
	imgUrl           = "https://cdn.betterttv.net/emote/5e9c6c187e090362f8b0b9e8/3x"
)

type testScene struct {
	scene.Scene
	x, y    int
	w, h    int
	bgColor color.Color
	name    string
	elapsed time.Duration
	stuff   []common.Entity
}

func (scene *testScene) IsInitialized() bool {
	return len(scene.stuff) > 0
}

func (scene *testScene) Key() string {
	return scene.name
}

func (scene *testScene) Init(_ *akara.World) {
	rand.Seed(time.Now().UnixNano())
	scene.stuff = append(scene.stuff, scene.Add.Circle(0, 0, 300, randColor(), nil))
}

func (scene *testScene) Update(dt time.Duration) {
	if scene.bgColor == nil {
		scene.bgColor = randColor()
	}

	scene.resizeViewport()
	scene.updateImages(dt)
}

func (scene *testScene) resizeViewport() {
	if len(scene.Viewports) < 1 {
		return
	}

	vp, found := scene.Components.Viewport.Get(scene.Viewports[0])
	if !found {
		return
	}

	vp.Background = scene.bgColor

	rt, found := scene.Components.RenderTexture2D.Get(scene.Viewports[0])
	if !found {
		return
	}

	trs, found := scene.Components.Transform.Get(scene.Viewports[0])
	if !found {
		return
	}

	if int(rt.Texture.Width) != scene.w || int(rt.Texture.Height) != scene.h {
		t := rl.LoadRenderTexture(int32(scene.w), int32(scene.h))
		rt.RenderTexture2D = &t
	}

	x, y := trs.Translation.XY()
	if int(x) != scene.x || int(y) != scene.y {
		trs.Translation.X, trs.Translation.Y = float64(scene.x), float64(scene.y)
	}
}

func (scene *testScene) updateImages(dt time.Duration) {
	scene.elapsed += dt

	scene.handleNewImage()

	for _, e := range scene.stuff {
		scene.updatePosition(e)
	}
}

func (scene *testScene) handleNewImage() {
	if scene.elapsed < newImageInterval {
		return
	}

	if len(scene.stuff) > maxImages {
		return
	}

	scene.elapsed = 0

	newImage := scene.Add.Image(imgUrl, 0, 0)

	scene.setRandomImagePosition(newImage)
	scene.stuff = append(scene.stuff, newImage)

	scene.fadeIn(newImage)
}

func (scene *testScene) fadeIn(e common.Entity) {
	t := tween.NewBuilder()
	t.Time(time.Second)
	t.Ease(easing.Sine)

	opacity, found := scene.Components.Opacity.Get(e)
	if !found {
		return
	}

	t.OnUpdate(func(progress float64) {
		opacity.Value = progress
	})

	t.OnUpdate(func(progress float64) {
		opacity.Value = progress
	})

	scene.Sys.Tweens.New(t)
}

func (scene *testScene) updatePosition(e common.Entity) {
	trs, found := scene.Components.Transform.Get(e)
	if !found {
		return
	}

	tex, found := scene.Components.Texture2D.Get(e)
	if !found {
		return
	}

	tw, th := tex.Width, tex.Height

	rWidth := scene.Sys.Renderer.Window.Width
	rHeight := scene.Sys.Renderer.Window.Height

	trs.Translation.Add(mathlib.NewVector3(float64(rand.Intn(3)-1), float64(rand.Intn(3)-1), 0))
	if trs.Translation.X > float64(rWidth+int(tw/2)) {
		trs.Translation.X = float64(-tw / 2)
	}

	if trs.Translation.Y > float64(rHeight+int(th/2)) {
		trs.Translation.Y = float64(-th / 2)
	}
}

func (scene *testScene) setRandomImagePosition(e common.Entity) {
	rWidth := scene.Sys.Renderer.Window.Width
	rHeight := scene.Sys.Renderer.Window.Height

	x, y := rand.Intn(rWidth), rand.Intn(rHeight)
	trs, _ := scene.Components.Transform.Get(e)

	trs.Translation.Set(float64(x), float64(y), 0)
}

func randColor() color.Color {
	return &color.RGBA{
		R: uint8(rand.Intn(math.MaxUint8)),
		G: uint8(rand.Intn(math.MaxUint8)),
		B: uint8(rand.Intn(math.MaxUint8)),
		A: math.MaxUint8 - uint8(rand.Intn(math.MaxUint8>>2)),
	}
}
