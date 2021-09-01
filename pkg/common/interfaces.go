package common

import (
	"github.com/gravestench/akara"
	"time"
)

type SceneFace interface {
	akara.System
	HasKey
	InitializesLua
	Update(duration time.Duration)
	GenericUpdate(duration time.Duration)
	Render()
	Initialize(width int, height int, world *akara.World, renderablesSubscription *akara.Subscription)
}

type HasKey interface {
	Key() string
}

type InitializesLua interface {
	LuaInitialized() bool
	InitializeLua()
	UninitializeLua()
}

type BasicObject interface {

}
