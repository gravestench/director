package input

import (
	"github.com/gravestench/akara"
	"image"
)

// static check that Interactive implements Component
var _ akara.Component = &Interactive{}

type InputCallback = func() (preventPropogation bool)

func noop() bool {
	return false
}

// Interactive is used to define an input state and a callback function to execute when that state is reached
type Interactive struct {
	Enabled bool
	*Vector
	Hitbox   *image.Rectangle
	Callback InputCallback
	// TODO: better name for this? allows us to temporarily ignore this vector (for debouncing)
	UsedRecently bool
	// TODO: better name for this? disables the debouncing functionality so the input callback fires as fast as possible
	RapidFire bool
}

// New returns a Interactive component. By default, it contains a nil instance.
func (*Interactive) New() akara.Component {
	return &Interactive{
		Enabled:  true,
		Vector:   NewInputVector(),
		Hitbox:   nil,
		Callback: noop,
	}
}

// InteractiveFactory is a wrapper for the generic component factory that returns Interactive component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Interactive.
type InteractiveFactory struct {
	*akara.ComponentFactory
}

// Add adds a Interactive component to the given entity and returns it
func (m *InteractiveFactory) Add(id akara.EID) *Interactive {
	return m.ComponentFactory.Add(id).(*Interactive)
}

// Get returns the Interactive component for the given entity, and a bool for whether or not it exists
func (m *InteractiveFactory) Get(id akara.EID) (*Interactive, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Interactive), found
}
