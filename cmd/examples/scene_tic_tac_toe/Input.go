package main

import (
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/systems/scene"
)

type Input struct {
	scene.Scene
}

func (scene *Input) Update() {

}

func (scene *Input) Init(_ *akara.World) {
}

func (scene *Input) IsInitialized() bool {
	return true
}
