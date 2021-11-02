package main

import (
	"fmt"
	"github.com/gravestench/director/pkg"
	"math/rand"
)

const (
	sceneName     = "lua shape test scene"
	maxRectangles = 500
	minWidth      = 10
	maxWidth      = 200
	minHeight     = 10
	maxHeight     = 200
)

type shapeTestFromLua struct {
	pkg.Scene
	numRectangles int
}

func (scene *shapeTestFromLua) Key() string {
	return sceneName
}

func (scene *shapeTestFromLua) Update() {
	if scene.numRectangles >= maxRectangles {
		return
	}

	ww, wh := scene.Sys.Renderer.Window.Width, scene.Sys.Renderer.Window.Height
	rx, ry := randRange(0, ww), randRange(0, wh)
	rw, rh := randRange(minWidth, maxWidth), randRange(minHeight, maxHeight)

	script := `
		v = scene.add.rectangle(%v, %v, %v, %v, "#7f00f7", "#ffffff")
	`

	if err := scene.Lua.DoString(fmt.Sprintf(script, rx, ry, rw, rh)); err != nil {
		fmt.Print(err)
	} else {
		scene.numRectangles++
	}
}

func randRange(min, max int) int {
	return rand.Intn(max-min) + min
}
