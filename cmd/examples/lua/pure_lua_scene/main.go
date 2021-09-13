package main

import (
	director "github.com/gravestench/director/pkg"
	"github.com/gravestench/director/pkg/systems/scene"
)

func main() {
	d := director.New()

	d.AddScene(scene.NewLuaScene("Lua Test Scene", "main.lua"))

	if err := d.Run(); err != nil {
		panic(err)
	}
}
