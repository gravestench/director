package main

import (
	"github.com/gravestench/director/pkg"
	"image/color"
	"time"

	"github.com/gravestench/akara"

	"github.com/gravestench/director"
	"github.com/gravestench/director/pkg/common"
	"github.com/gravestench/director/pkg/easing"
	"github.com/gravestench/director/pkg/systems/tween"
)

func main() {
	d := director.New()

	d.AddScene(&TweenTest{})

	if err := d.Run(); err != nil {
		panic(err)
	}
}

type TweenTest struct {
	pkg.Scene
	object common.Entity
}

func (scene *TweenTest) Key() string {
	return "tween test"
}

func (scene *TweenTest) Update() {
	// noop
}

func (scene *TweenTest) Init(_ *akara.World) {
	red := color.RGBA{R: 255, A: 255}

	scene.object = scene.Add.Label("Director", 1024/2, 768/2, 100, "", red)

	scene.makeTween()
}

func (scene *TweenTest) IsInitialized() bool {
	return scene.object != 0
}

func (scene *TweenTest) makeTween() {
	t := scene.Sys.Tweens.New()

	t.Time(time.Second * 5)
	t.Ease(easing.ElasticOut, []float64{0.5, 0.85, 0.5})
	t.Repeat(tween.RepeatForever)

	trs, found := scene.Components.Transform.Get(scene.object)
	if !found {
		return
	}

	t.OnUpdate(func(complete float64) {
		trs.Rotation.Y = complete * 360
		trs.Scale.Set(complete, complete, complete)
	})

	scene.Sys.Tweens.Add(t)
}
