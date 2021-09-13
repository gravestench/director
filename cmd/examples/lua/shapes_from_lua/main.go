package main

import director "github.com/gravestench/director/pkg"

func main() {
	d := director.New()

	d.AddScene(&shapeTestFromLua{})

	if err := d.Run(); err != nil {
		panic(err)
	}
}