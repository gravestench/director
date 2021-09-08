package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/systems/scene"
	"github.com/gravestench/mathlib"
	"math/rand"
)

const (
	maxImages = 30
	imgUrl = "https://cdn.betterttv.net/emote/5e9c6c187e090362f8b0b9e8/3x"
)

type testScene struct {
	scene.Scene
	images []akara.EID
}

func (scene *testScene) Init(_ *akara.World) {
	for len(scene.images) < maxImages {
		scene.images = append(scene.images, scene.Add.Image(imgUrl, 0, 0))
	}

	scene.setRandomImagePositions()
}

func (scene *testScene) IsInitialized() bool {
	return true
}

func (scene *testScene) Update() {
	scene.resizeCameraWithWindow()

	for _, e := range scene.images {
		scene.updatePosition(e)
	}
}

func (scene *testScene) updatePosition(e akara.EID) {
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
	for _, e := range scene.Cameras {
		rt, found := scene.Components.RenderTexture2D.Get(e)
		if !found {
			continue
		}

		if int(rt.Texture.Width) != scene.Window.Width || int(rt.Texture.Height) != scene.Window.Height {
			t := rl.LoadRenderTexture(int32(scene.Window.Width), int32(scene.Window.Height))
			rt.RenderTexture2D = &t

			scene.setRandomImagePositions()
		}
	}
}

func (scene *testScene) setRandomImagePositions() {
	for _, e := range scene.images {
		x, y := rand.Intn(scene.Window.Width), rand.Intn(scene.Window.Height)
		trs, _ := scene.Components.Transform.Get(e)

		trs.Translation.Set(float64(x), float64(y), 0)
	}
}
