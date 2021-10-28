package main

import (
	"github.com/gravestench/director"
	"time"
)

func main() {
	d := director.New()

	d.AddScene(&testScene{
		name: "a",
		x:    260,
		y:    100,
		w:    200,
		h:    100,
		interval: time.Millisecond * 20,
	})

	d.AddScene(&testScene{
		name: "b",
		x:    250,
		y:    150,
		w:    520,
		h:    320,
		interval: time.Second/10,
	})

	if err := d.Run(); err != nil {
		panic(err)
	}
}
