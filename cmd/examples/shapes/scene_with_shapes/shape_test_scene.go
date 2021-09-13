package main

import (
	"image/color"

	"github.com/gravestench/akara"

	"github.com/gravestench/director/pkg/systems/scene"
)

const (
	sceneName = "shape test scene"
)

type ShapeTestScene struct {
	scene.Scene
}

func (scene *ShapeTestScene) Init(_ *akara.World) {
	yellow := color.RGBA{R: 255, G: 255, A: 255}
	pink := color.RGBA{R: 255, B: 255, A: 255}

	ww, wh := scene.Sys.Renderer.Window.Width, scene.Sys.Renderer.Window.Height

	scene.Add.Rectangle(ww/2, wh/2, 100, 100, yellow, pink)
	scene.Add.Circle(ww/4, wh*3/4, ww/8, nil, pink)
}

func (scene *ShapeTestScene) IsInitialized() bool {
	return true
}

func (scene *ShapeTestScene) Key() string {
	return sceneName
}
