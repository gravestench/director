package main

import (
	director "github.com/gravestench/director/pkg"
	"github.com/gravestench/director/pkg/systems/scene"
)

func main() {
	d := director.New()

	d.Window.Width = 1024
	d.Window.Height = 768
	d.Window.Title = "lua test"
	d.TargetFPS = 60

	d.AddScene(scene.NewLuaScene("Lua Test Scene", "main.lua"))

	if err := d.Run(); err != nil {
		panic(err)
	}
}
