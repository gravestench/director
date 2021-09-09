package main

import director "github.com/gravestench/director/pkg"

func main() {
	d := director.New()

	d.Window.Width = 1024
	d.Window.Height = 768
	d.TargetFPS = 60

	d.AddScene(&ShapeTestScene{})

	if err := d.Run(); err != nil {
		panic(err)
	}
}
