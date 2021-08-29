package pkg

import "time"

type imageFactory struct {
	*basicComponents
}

func (factory *imageFactory) update(s *Scene, dt time.Duration)  {}
