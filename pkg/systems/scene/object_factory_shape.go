package scene

import (
	"time"
)

type shapeFactory struct {
	rectangle rectangleFactory
	circle    circleFactory
}

func (factory *shapeFactory) update(s *Scene, dt time.Duration) {
	factory.rectangle.update(s, dt)
	factory.circle.update(s, dt)
}
