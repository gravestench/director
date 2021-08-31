package pkg

import "time"

type shapeFactory struct {
	*Director
	components struct {
		*basicComponents
	}

	rectangle rectangleFactory
	circle    circleFactory
}

func (factory *shapeFactory) update(s *Scene, dt time.Duration) {
	factory.rectangle.update(s, dt)
	factory.circle.update(s, dt)
}
