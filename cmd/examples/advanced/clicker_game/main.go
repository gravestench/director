package main

import "github.com/gravestench/director"

func main() {
	d := director.New()

	d.AddScene(&GameScene{})

	if err := d.Run(); err != nil {
		panic(err)
	}
}
