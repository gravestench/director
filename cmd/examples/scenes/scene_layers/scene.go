package main

import (
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg"
	"github.com/gravestench/director/pkg/common"
	"image/color"
)

// NOTE: this is currently not working

type layerTest struct {
	pkg.Scene
	layer common.Entity
}

func (scene *layerTest) Key() string {
	return "layer test"
}

func (scene *layerTest) Update() {
	ltrs, found := scene.Components.Transform.Get(scene.layer)
	if !found {
		return
	}

	ltrs.Rotation.Y += .01
	ltrs.Scale.Set(1, 1, 1)
}

func (scene *layerTest) Init(_ *akara.World) {
	l := scene.Add.Layer(100, 100)
	lc, found := scene.Components.HasChildren.Get(l)
	if !found {
		return
	}

	r1 := scene.Add.Rectangle(10, 10, 20, 20, color.RGBA{R: 255, A: 255}, nil)
	r2 := scene.Add.Rectangle(15, 15, 20, 20, color.RGBA{G: 255, A: 255}, nil)
	r3 := scene.Add.Rectangle(30, 20, 20, 20, color.RGBA{B: 255, A: 255}, nil)

	scene.Components.Debug.Add(r1)

	lc.Children = append(lc.Children, r1, r2, r3)

	scene.layer = l
}

func (scene *layerTest) IsInitialized() bool {
	return true
}
