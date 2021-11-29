package main

import (
	"github.com/gravestench/director"
)

func main() {
	d := director.New()

	//d.AddLuaScene("object spawn scene", "main.lua")
	d.AddLuaSystem("movement.lua")

	if err := d.Run(); err != nil {
		panic(err)
	}
}
