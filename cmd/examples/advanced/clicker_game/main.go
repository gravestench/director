package main

import director "github.com/gravestench/director/pkg"

func main() {
	d := director.New()
	d.SetDebug(true)

	d.AddScene(&GameScene{})

	if err := d.Run(); err != nil {
		panic(err)
	}
}
