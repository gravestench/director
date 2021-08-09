package pkg

import (
	"image/color"

	"github.com/gravestench/akara"
)

type SceneObjectFactory struct {
	scene *Scene
	basicComponents
	shape shapeFactory
	image imageFactory
	text textFactory
}

func (of *SceneObjectFactory) Text(str string, x, y int, c color.Color) akara.EID {
	return of.text.New(of.scene, str, x, y, c)
}

type shapeFactory struct {
	*Director
	components struct {
		*basicComponents
	}
}

type imageFactory struct {
	*basicComponents
}

type textFactory struct {
	*basicComponents
}

func (tf *textFactory) New(s *Scene, str string, x, y int, c color.Color) akara.EID {
	e := s.Director.NewEntity()

	text := s.Text.Add(e)
	text.String = str

	vec2 := s.Vector2.Add(e)
	vec2.X, vec2.Y = float32(x), float32(y)

	col := s.Color.Add(e)
	r, g, b, a := c.RGBA()
	col.R, col.G, col.B, col.A = uint8(r), uint8(g), uint8(b), uint8(a)

	return e
}