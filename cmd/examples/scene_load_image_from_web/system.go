package main

import (
	director "github.com/gravestench/director/pkg"
	"github.com/gravestench/director/pkg/easing"
	"github.com/gravestench/director/pkg/systems/tween"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/mathlib"

	. "github.com/gravestench/director/pkg/common"
	"github.com/gravestench/director/pkg/systems/scene"
)

const (
	newImageInterval = time.Millisecond * 100
	imgUrl           = "https://cdn.betterttv.net/emote/5e9c6c187e090362f8b0b9e8/3x"
)

var _ director.SceneInterface = &testScene{}

type testScene struct {
	scene.Scene
	images  []Entity
	elapsed time.Duration
}

func (scene *testScene) Key() string {
	return "test"
}

func (scene *testScene) Update(dt time.Duration) {
	scene.elapsed += dt

	scene.handleNewImage()
	scene.resizeCameraWithWindow()

	for _, e := range scene.images {
		scene.updatePosition(e)
	}
}

func (scene *testScene) handleNewImage() {
	if scene.elapsed < newImageInterval {
		return
	}

	scene.elapsed = 0

	newImage := scene.Add.Image(imgUrl, 0, 0)

	scene.setRandomImagePosition(newImage)
	scene.images = append(scene.images, newImage)

	scene.fadeIn(newImage)
}

func (scene *testScene) fadeIn(e Entity) {
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

func (scene *testScene) updatePosition(e Entity) {
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

func (scene *testScene) resizeCameraWithWindow() {
	rWidth := scene.Sys.Renderer.Window.Width
	rHeight := scene.Sys.Renderer.Window.Height

	for _, e := range scene.Viewports {
		rt, found := scene.Components.RenderTexture2D.Get(e)
		if !found {
			continue
		}

		if int(rt.Texture.Width) != rWidth || int(rt.Texture.Height) != rHeight {
			t := rl.LoadRenderTexture(int32(rWidth), int32(rHeight))
			rt.RenderTexture2D = &t

			for _, e := range scene.images {
				scene.setRandomImagePosition(e)
			}
		}
	}
}

func (scene *testScene) setRandomImagePosition(e Entity) {
	rWidth := scene.Sys.Renderer.Window.Width
	rHeight := scene.Sys.Renderer.Window.Height

	x, y := rand.Intn(rWidth), rand.Intn(rHeight)
	trs, _ := scene.Components.Transform.Get(e)

	trs.Translation.Set(float64(x), float64(y), 0)
}
