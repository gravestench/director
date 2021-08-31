package main

import (
	"github.com/gravestench/akara"
	director "github.com/gravestench/director/pkg"
	"image/color"
)

const sceneName = "Label Scene"

type LabelScene struct {
	director.Scene
	label akara.EID
}

func (scene *LabelScene) Init(world *akara.World) {
	scene.Add.Label("Hello World!", 100, 100, 20, "", color.White)
}

func (scene *LabelScene) IsInitialized() bool {
	return true
}

func (scene *LabelScene) Key() string {
	return sceneName
}

