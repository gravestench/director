package interactive

import (
	"image"

	"github.com/gravestench/director/pkg/systems/input/vector"

	"github.com/gravestench/director/pkg/common/components"

	"github.com/gravestench/akara"
)

// static check that Interactive implements Camera
var _ akara.Component = &Interactive{}

type InputCallback = func() (preventPropogation bool)

func noop() bool {
	return false
}

// Interactive is used to define an input state and a callback function to execute when that state is reached
type Interactive struct {
	Enabled bool
	*vector.Vector
	Hitbox   *image.Rectangle
	Callback InputCallback
	// TODO: better componentName for this? allows us to temporarily ignore this vector (for debouncing)
	UsedRecently bool
	// TODO: better componentName for this? disables the debouncing functionality so the input callback fires as fast as possible
	RapidFire bool
}

// New returns a Interactive component. By default, it contains a nil instance.
func (*Interactive) New() akara.Component {
	return &Interactive{
		Enabled:  true,
		Vector:   vector.NewInputVector(),
		Hitbox:   nil,
		Callback: noop,
	}
}

// static checks, should fail to compile if
// these interfaces can't be implemented
var (
	_ components.Component = &Component{}
	_ components.LuaExport = &ComponentFactory{}
)

type Component = Interactive // Component is an alias to Interactive
