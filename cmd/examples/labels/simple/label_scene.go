package main

import (
	"image/color"

	. "github.com/gravestench/director"
)

const sceneName = "Label Scene"

type LabelScene struct {
	Scene
}

func (scene *LabelScene) Key() string {
	return sceneName
}

func (scene *LabelScene) IsInitialized() bool {
	return true
}

func (scene *LabelScene) Init(_ *World) {
	scene.Add.Label("Hello World!", 200, 200, 20, "", color.White)
}

func (scene *LabelScene) Update() {

}
