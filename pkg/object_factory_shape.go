package pkg

import "time"

type shapeFactory struct {
	*Director
	components struct {
		*basicComponents
	}

	rectangle rectangleFactory
}

func (factory *shapeFactory) update(s *Scene, dt time.Duration) {
	factory.rectangle.update(s, dt)
}
