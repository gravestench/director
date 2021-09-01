package scene

import (
	"github.com/gravestench/director/pkg/common"
	"time"
)

type imageFactory struct {
	*common.BasicComponents
}

func (factory *imageFactory) update(s *Scene, dt time.Duration) {}
