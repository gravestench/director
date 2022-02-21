package main

import (
	"image/color"
	"time"

	. "github.com/gravestench/director"

	"github.com/gravestench/director"
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
	Scene
}

func (scene *TweenTest) Key() string {
	return "tween test"
}

func (scene *TweenTest) Update() {
	// noop
}

func (scene *TweenTest) Init(_ *World) {
	x, y := 1024/2, 768/2
	size := 100
	font := "" // stub, doesnt work right now
	red := color.RGBA{R: 255, A: 255}

	eid := scene.Add.Label("Director", x, y, size, font, red)

	scene.makeTween(eid)
}

func (scene *TweenTest) makeTween(eid Entity) {
	t := scene.Sys.Tweens.New()

	t.Time(time.Second * 5)
	t.Ease(easing.ElasticOut, []float64{0.5, 0.85, 0.5})
	t.Repeat(tween.RepeatForever)

	trs, found := scene.Components.Transform.Get(eid)
	if !found {
		return
	}

	t.OnUpdate(func(complete float64) {
		trs.Rotation.Y = complete * 360
		trs.Scale.Set(complete, complete, complete)
	})

	scene.Sys.Tweens.Add(t)
}
