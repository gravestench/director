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

func (l *LabelScene) Init(world *akara.World) {
	l.Add.Label("Hello World!", 100, 100, 20, "", color.White)
}

func (l *LabelScene) IsInitialized() bool {
	return true
}

func (l *LabelScene) Key() string {
	return sceneName
}

