package main

import (
	"fmt"
	director "github.com/gravestench/director/pkg"
	"math/rand"
)

const (
	sceneName = "lua shape test scene"
	maxRectangles = 500
	minWidth = 10
	maxWidth = 200
	minHeight = 10
	maxHeight = 200
)

type shapeTestFromLua struct {
	director.Scene
	numRectangles int
}

func (s *shapeTestFromLua) Key() string {
	return sceneName
}

func (s *shapeTestFromLua) Update() {
	if s.numRectangles >= maxRectangles {
		return
	}

	ww, wh := s.Director.Window.Width, s.Director.Window.Height
	rx, ry := randRange(0, ww), randRange(0, wh)
	rw, rh := randRange(minWidth, maxWidth), randRange(minHeight, maxHeight)

	script := `
		v = rectangle.new(%v, %v, %v, %v, "#7f00f7", "#ffffff")
	`

	if err := s.Lua.DoString(fmt.Sprintf(script, rx, ry, rw, rh)); err != nil {
		fmt.Print(err)
	} else {
		s.numRectangles++
	}
}

func randRange(min, max int) int {
	return rand.Intn(max - min) + min
}

