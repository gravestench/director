package main

import (
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/systems/scene"
	"image/color"
)

const sceneName = "Label Scene"

type LabelScene struct {
	scene.Scene
	label akara.EID
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
