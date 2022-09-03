package main

import (
	"time"

	"github.com/gravestench/director"
)

func main() {
	d := director.New()

	d.AddScene(&testScene{
		name:     "a",
		x:        170,
		y:        100,
		w:        200,
		h:        100,
		interval: time.Second / 33 * 100,
	})

	d.AddScene(&testScene{
		name:     "b",
		x:        250,
		y:        150,
		w:        520,
		h:        320,
		interval: time.Second / 57 * 100,
	})

	d.AddScene(&testScene{
		name:     "b",
		x:        100,
		y:        300,
		w:        800,
		h:        160,
		interval: time.Second / 88 * 100,
	})

	if err := d.Run(); err != nil {
		panic(err)
	}
}
