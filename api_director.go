package director

import (
	"flag"
	director "github.com/gravestench/director/pkg"
	"github.com/gravestench/director/primitive_scenes/tick_graph"
)

// Director is the scene manager
type Director = director.Director

// New returns a new director instance
func New() *Director {
	flag.Parse()

	instance := director.New()

	if flag.Lookup(director.FlagNameDebug).Value.(flag.Getter).Get().(bool) {
		instance.AddScene(&tick_graph.Scene{})
	}

	return instance
}
