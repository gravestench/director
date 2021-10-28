package main

import "github.com/gravestench/director"

func main() {
	d := director.New()

	d.AddScene(&testScene{})

	if err := d.Run(); err != nil {
		panic(err)
	}
}
