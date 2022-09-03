package renderer

import (
	"fmt"

	"github.com/kbinani/screenshot"
)

func (s *System) updateScreenInfo() {
	n := screenshot.NumActiveDisplays()

	if len(s.Screens) > n {
		s.Screens = s.Screens[:n]
	} else if len(s.Screens) < n {
		s.Screens = make([]screenInfo, n)
	}

	for idx, screen := range s.Screens {
		screen.Name = fmt.Sprintf("Screen #%d", idx)
		bounds := screenshot.GetDisplayBounds(idx)

		s.Screens[idx].Width, s.Screens[idx].Height = bounds.Dx(), bounds.Dy()
	}
}

type screenInfo struct {
	Name          string
	Width, Height int
}
