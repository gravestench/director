package main

import (
	"github.com/gravestench/director"
)

func main() {
	d := director.New()

	d.AddScene(&LabelScene{})

	if err := d.Run(); err != nil {
		panic(err)
	}
}
