package pkg

import (
	"image/color"
	"time"

	"github.com/gravestench/akara"
)

type SceneObjectFactory struct {
	scene *Scene
	basicComponents
	generic genericFactory
	shape   shapeFactory
	image   imageFactory
	label   labelFactory
	camera  cameraFactory
}

func (factory *SceneObjectFactory) update(dt time.Duration) {
	factory.generic.update(factory.scene, dt)
	factory.camera.update(factory.scene, dt)
	factory.shape.update(factory.scene, dt)
	factory.label.update(factory.scene, dt)
	factory.image.update(factory.scene, dt)
}

func (factory *SceneObjectFactory) Label(str string, x, y, size int, fontName string, c color.Color) akara.EID {
	return factory.label.New(factory.scene, str, x, y, size, fontName, c)
}

func (factory *SceneObjectFactory) Camera(x, y, w, h int) akara.EID {
	return factory.camera.New(factory.scene, x, y, w, h)
}

func (factory *SceneObjectFactory) Rectangle(x, y, w, h int, fill, stroke color.Color) akara.EID {
	return factory.shape.rectangle.New(factory.scene, x, y, w, h, fill, stroke)
}
