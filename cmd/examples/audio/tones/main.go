package main

import (
	"embed"

	"github.com/gravestench/director"
)

//go:embed data
var data embed.FS

func main() {
	d := director.New()

	d.AddScene(&AudioTonePlayerScene{})
	d.Sys.Load.AddProvider(&provider{data: &data})

	if err := d.Run(); err != nil {
		panic(err)
	}
}
