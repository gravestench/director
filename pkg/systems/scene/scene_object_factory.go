package scene

import (
	"image/color"
	"time"

	"github.com/gravestench/director/pkg/common"
)

type SceneObjectFactory struct {
	common.BasicComponents
	scene    *Scene
	generic  genericFactory
	shape    shapeFactory
	image    imageFactory
	label    labelFactory
	viewport viewportFactory
	camera   cameraFactory
}

func (factory *SceneObjectFactory) update(dt time.Duration) {
	factory.generic.update(factory.scene, dt)
	factory.viewport.update(factory.scene, dt)
	factory.camera.update(factory.scene, dt)
	factory.shape.update(factory.scene, dt)
	factory.label.update(factory.scene, dt)
	factory.image.update(factory.scene, dt)
}

func (factory *SceneObjectFactory) Label(str string, x, y, size int, fontName string, c color.Color) common.Entity {
	return factory.label.New(factory.scene, str, x, y, size, fontName, c)
}

func (factory *SceneObjectFactory) Viewport(x, y, w, h int) common.Entity {
	return factory.viewport.New(factory.scene, x, y, w, h)
}

func (factory *SceneObjectFactory) Camera(x, y, w, h int) common.Entity {
	return factory.camera.New(factory.scene, x, y, w, h)
}

func (factory *SceneObjectFactory) Rectangle(x, y, w, h int, fill, stroke color.Color) common.Entity {
	return factory.shape.rectangle.New(factory.scene, x, y, w, h, fill, stroke)
}

func (factory *SceneObjectFactory) Circle(x, y, radius int, fill, stroke color.Color) common.Entity {
	return factory.shape.circle.New(factory.scene, x, y, radius, fill, stroke)
}

func (factory *SceneObjectFactory) Image(uri string, x, y int) common.Entity {
	return factory.image.New(factory.scene, uri, x, y)
}
