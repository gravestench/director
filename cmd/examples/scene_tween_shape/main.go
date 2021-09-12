package main

import (
	"image/color"
	"time"

	"github.com/gravestench/akara"

	director "github.com/gravestench/director/pkg"
	. "github.com/gravestench/director/pkg/common"
	"github.com/gravestench/director/pkg/easing"
	"github.com/gravestench/director/pkg/systems/scene"
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
	scene.Scene
	object      Entity
	squareTween *tween.Tween
}

func (t *TweenTest) Key() string {
	return "tween test"
}

func (t *TweenTest) Update() {
	// noop
}

func (t *TweenTest) Init(_ *akara.World) {
	red := color.RGBA{R: 255, A: 255}

	t.object = t.Add.Label("LOLWUT", 1024/2, 768/2, 100, "", red)

	t.makeTween()
}

func (t *TweenTest) IsInitialized() bool {
	return t.object != 0
}

func (t *TweenTest) makeTween() {
	builder := tween.NewBuilder()

	builder.Time(time.Second * 4)
	builder.Ease(easing.ElasticOut, []float64{0.5, 0.85, 0.5})
	builder.Repeat(tween.RepeatForever)
	builder.Delay(time.Second * 3)

	trs, found := t.Components.Transform.Get(t.object)
	if !found {
		return
	}

	builder.OnUpdate(func(complete float64) {
		trs.Rotation.Y = complete * 360
		trs.Scale.Set(complete, complete, complete)
	})

	t.squareTween = builder.Build()

	t.Sys.Tweens.New(t.squareTween)
}
