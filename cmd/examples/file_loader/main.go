package main

import (
	"github.com/gravestench/director/cmd/examples/file_loader/loader/filesystemloader"
	director "github.com/gravestench/director/pkg"
	"os"
)

func main() {
	d := director.New()

	d.Window.Width = 1024
	d.Window.Height = 768
	d.Window.Title = "file loader test"
	d.TargetFPS = 60

	loader := &FileLoader{}
	fakeRequester := &FakeFileRequester{}

	d.AddSystem(loader) // handles loading file streams
	d.AddSystem(fakeRequester) // occasionally makes file load requests

	wd, _ := os.Getwd()
	loader.AddProvider(filesystemloader.New(wd))

	if err := d.Run(); err != nil {
		panic(err)
	}
}