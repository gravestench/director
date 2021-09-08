package main

import (
	"github.com/gravestench/director/pkg/easing"
	"github.com/gravestench/director/pkg/systems/tween"
	"math/rand"
	"time"

	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
	"github.com/gravestench/mathlib"

	. "github.com/gravestench/director/pkg/common"
	"github.com/gravestench/director/pkg/systems/scene"
)

const (
	newImageInterval = time.Millisecond * 100
	imgUrl = "https://cdn.betterttv.net/emote/5e9c6c187e090362f8b0b9e8/3x"
)

type testScene struct {
	scene.Scene
	images []Entity
	elapsed time.Duration
}

func (scene *testScene) Init(_ *akara.World) {
	img := scene.Add.Image(imgUrl, 0, 0)
	scene.setRandomImagePosition(img)
	scene.images = append(scene.images, img)
}

func (scene *testScene) IsInitialized() bool {
	return true
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

	scene.Tweens.New(t)
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

	trs.Translation.Add(mathlib.NewVector3(float64(rand.Intn(3)-1), float64(rand.Intn(3)-1), 0))
	if trs.Translation.X > float64(scene.Window.Width + int(tw/2)) {
		trs.Translation.X = float64(-tw/2)
	}

	if trs.Translation.Y > float64(scene.Window.Height + int(th/2)) {
		trs.Translation.Y = float64(-th/2)
	}
}

func (scene *testScene) resizeCameraWithWindow() {
	for _, e := range scene.Viewports {
		rt, found := scene.Components.RenderTexture2D.Get(e)
		if !found {
			continue
		}

		if int(rt.Texture.Width) != scene.Window.Width || int(rt.Texture.Height) != scene.Window.Height {
			t := rl.LoadRenderTexture(int32(scene.Window.Width), int32(scene.Window.Height))
			rt.RenderTexture2D = &t

			for _, e := range scene.images {
				scene.setRandomImagePosition(e)
			}
		}
	}
}

func (scene *testScene) setRandomImagePosition(e Entity) {
	x, y := rand.Intn(scene.Window.Width), rand.Intn(scene.Window.Height)
	trs, _ := scene.Components.Transform.Get(e)

	trs.Translation.Set(float64(x), float64(y), 0)
}
