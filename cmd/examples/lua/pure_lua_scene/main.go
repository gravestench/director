package main

import (
	"github.com/gravestench/director"
)

func main() {
	d := director.New()
	d.AddScene(director.Scene.NewLuaScene("Lua Test Scene", "main.lua"))

	if err := d.Run(); err != nil {
		panic(err)
	}
}
