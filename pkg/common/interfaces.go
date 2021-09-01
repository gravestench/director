package common

import (
	"github.com/gravestench/akara"
	"time"
)

type Scene interface {
	akara.System
	HasKey
	initializesLua
	GenericUpdate(duration time.Duration)
	Render()
	Initialize(width int, height int, world *akara.World, renderablesSubscription *akara.Subscription)
}

type Updater interface {
	Update()
}

type UpdaterTimed interface {
	Update(duration time.Duration)
}

type LuaScene interface {
	Scene
	RunLuaScripts()
}

type HasKey interface {
	Key() string
}

type initializesLua interface {
	LuaInitialized() bool
	InitializeLua()
	UninitializeLua()
}

type BasicObject interface {

}
