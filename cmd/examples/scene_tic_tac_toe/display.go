package main

import (
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/systems/scene"
	"image/color"
)

type Display struct {
	scene.Scene
	ui struct {
		prompt    akara.EID
		highlight akara.EID
		x         []akara.EID
		o         []akara.EID
	}
}

func (scene *Display) Update() {

}

func (scene *Display) Init(_ *akara.World) {
	scene.makeGraphics()
}

func (scene *Display) IsInitialized() bool {
	return scene.ui.prompt != 0
}

func (scene *Display) makeGraphics() {
	const (
		ww, wh           = 1024, 768
		centerX, centerY = ww >> 1, wh >> 1
		gridSize         = wh
		cellSize         = gridSize / gridOrder
		margin           = cellSize / 10
		innerCellSize    = cellSize - (margin * 2)
		w, h             = innerCellSize, innerCellSize
		startX, startY   = centerX - cellSize, centerY + cellSize
	)

	highlight := color.RGBA{R: 255, G: 255, A: 128}

	scene.ui.prompt = scene.Add.Label("Testing!", 1024>>1, 768>>1, wh/6, "", color.White)
	scene.ui.highlight = scene.Add.Rectangle(-cellSize, -cellSize, w, h, highlight, nil)

	scene.ui.x = make([]akara.EID, numCells)
	scene.ui.o = make([]akara.EID, numCells)

	x, y := startX, startY
	for idx := range scene.ui.x {
		ox := (idx % 3) * cellSize
		oy := (idx / 3) * -cellSize
		scene.ui.x[idx] = scene.Add.Label(PlayerX.String(), x+ox, y+oy, innerCellSize, "", color.White)
		scene.ui.o[idx] = scene.Add.Label(PlayerO.String(), x+ox, y+oy, innerCellSize, "", color.White)
	}
}
