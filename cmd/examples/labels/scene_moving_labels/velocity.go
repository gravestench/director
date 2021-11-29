package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/gravestench/akara"
)

// static check that Velocity implements Camera
var _ akara.Component = &Velocity{}

// Velocity is a component that contains a 2-dimensional Vector
type Velocity struct {
	rl.Vector2
}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
func (*Velocity) New() akara.Component {
	return &Velocity{}
}

// VelocityFactory is a wrapper for the generic component factory that returns Velocity component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a Velocity.
type VelocityFactory struct {
	*akara.ComponentFactory
}

// Add adds a Velocity component to the given entity and returns it
func (m *VelocityFactory) Add(id akara.EID) *Velocity {
	return m.ComponentFactory.Add(id).(*Velocity)
}

// Get returns the Velocity component for the given entity, and a bool for whether or not it exists
func (m *VelocityFactory) Get(id akara.EID) (*Velocity, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*Velocity), found
}
