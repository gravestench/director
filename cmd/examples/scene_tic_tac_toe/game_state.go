package main

import (
	"github.com/gravestench/akara"
)

const (
	gridOrder = 3
	numCells = gridOrder * gridOrder
)

type state struct {
	turn Player
	cells []Player
}

// static check that GameState implements Component
var _ akara.Component = &GameState{}

// GameState is a component that contains normalized alpha transparency (0.0 ... 1.0)
type GameState struct {
	state
}

// New creates a new alpha component instance. The default alpha is opaque with value 1.0
func (*GameState) New() akara.Component {
	s := state{}
	s.cells = make([]Player, numCells)

	return &GameState{
		state: s,
	}
}

// GameStateFactory is a wrapper for the generic component factory that returns GameState component instances.
// This can be embedded inside of a system to give them the methods for adding, retrieving, and removing a GameState.
type GameStateFactory struct {
	*akara.ComponentFactory
}

// Add adds a GameState component to the given entity and returns it
func (m *GameStateFactory) Add(id akara.EID) *GameState {
	return m.ComponentFactory.Add(id).(*GameState)
}

// Get returns the GameState component for the given entity, and a bool for whether or not it exists
func (m *GameStateFactory) Get(id akara.EID) (*GameState, bool) {
	component, found := m.ComponentFactory.Get(id)
	if !found {
		return nil, found
	}

	return component.(*GameState), found
}

