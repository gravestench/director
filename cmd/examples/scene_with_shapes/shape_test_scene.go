package main

import (
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/systems/scene"
	"image/color"
)

const (
	sceneName = "shape test scene"
)

type ShapeTestScene struct {
	scene.Scene
}

func (s *ShapeTestScene) Init(_ *akara.World) {
	yellow := color.RGBA{R: 255, G: 255, A:255}
	pink := color.RGBA{R: 255, B:255, A:255}

	ww, wh := s.Width, s.Height

	s.Add.Rectangle(ww/2, wh/2, 100, 100, yellow, pink)
	s.Add.Circle(ww/4, wh*3/4, ww/8, nil, pink)
}

func (s *ShapeTestScene) IsInitialized() bool {
	return true
}

func (s *ShapeTestScene) Key() string {
	return sceneName
}

