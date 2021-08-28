package main

import (
	"github.com/gravestench/akara"
	lua "github.com/yuin/gopher-lua"
	"time"
)

type LuaSystem struct {
	akara.BaseSystem
	state *lua.LState
	subscribed struct {
		unregisteredScripts *akara.Subscription
		apiEventHandlers *akara.Subscription
	}
}

func (sys *LuaSystem) IsInitialized() bool {
	return sys.state != nil
}

func (sys *LuaSystem) Init(world *akara.World) {
	sys.World = world
	sys.state = lua.NewState()

	for _, luaType := range luaTypes {
		registerType(sys.state, luaType)
	}
}

func (sys *LuaSystem) Update(duration time.Duration) {
	for _, eid := range sys.subscribed.unregisteredScripts.GetEntities() {
		sys.register(eid)
	}
}

func (sys *LuaSystem) register(e akara.EID) {
	sys.bindEventHandlers(e)
}

func (sys *LuaSystem) bindEventHandlers(e akara.EID) {
	panic("implement me")
}

