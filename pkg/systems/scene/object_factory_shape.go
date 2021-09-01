package scene

import (
	"github.com/gravestench/director/pkg/common"
	"time"
)

type shapeFactory struct {
	components struct {
		*common.BasicComponents
	}

	rectangle rectangleFactory
	circle    circleFactory
}

func (factory *shapeFactory) update(s *Scene, dt time.Duration) {
	factory.rectangle.update(s, dt)
	factory.circle.update(s, dt)
}
