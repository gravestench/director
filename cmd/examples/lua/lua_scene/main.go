package main

import (
	"github.com/gravestench/director"
)

func main() {
	d := director.New()
	d.AddLuaScene("Lua Test Scene", "scene.lua")

	if err := d.Run(); err != nil {
		panic(err)
	}
}
