package main

import (
	director "github.com/gravestench/director/pkg"
)

func main() {
	d := director.New()
	d.SetDebug(true)

	d.AddScene(&AudioTonePlayerScene{})

	if err := d.Run(); err != nil {
		panic(err)
	}
}
