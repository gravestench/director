package main

import (
	"github.com/gravestench/director/pkg"
	"image/color"

	"github.com/gravestench/akara"
)

const sceneName = "Label Scene"

type LabelScene struct {
	pkg.Scene
}

func (scene *LabelScene) Key() string {
	return sceneName
}

func (scene *LabelScene) IsInitialized() bool {
	return true
}

func (scene *LabelScene) Init(_ *akara.World) {
	scene.Add.Label("Hello World!", 200, 200, 20, "", color.White)
}

func (scene *LabelScene) Update() {

}
