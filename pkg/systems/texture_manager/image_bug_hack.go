package texture_manager

import (
	"image"
	"image/color"
)

type imageBugHack struct {
	img image.Image
}

func (h imageBugHack) ColorModel() color.Model {
	return h.img.ColorModel()
}

func (h imageBugHack) Bounds() image.Rectangle {
	return h.img.Bounds()
}

func (h imageBugHack) At(x, y int) color.Color {
	b := h.img.Bounds()

	return h.img.At(x, b.Dy()-y)
}