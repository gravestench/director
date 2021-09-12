package main

import (
	director "github.com/gravestench/director/pkg"
)

func main() {
	d := director.New()

	d.AddScene(&testScene{
		name: "a",
		x: 260,
		y: 100,
		w: 200,
		h: 100,
	})

	d.AddScene(&testScene{
		name: "b",
		x: 250,
		y: 150,
		w: 120,
		h: 320,
	})

	if err := d.Run(); err != nil {
		panic(err)
	}
}
