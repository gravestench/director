package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gravestench/akara"
	lua "github.com/yuin/gopher-lua"

	"github.com/gravestench/director/cmd/examples/lua_api/vector3"
)

type LuaSystem struct {
	akara.BaseSystem
	lua *lua.LState
	//subscribed struct {
	//	unregisteredScripts *akara.Subscription
	//	apiEventHandlers *akara.Subscription
	//}
}

func (sys *LuaSystem) IsInitialized() bool {
	return sys.lua != nil
}

func (sys *LuaSystem) Init(world *akara.World) {
	sys.World = world
	sys.lua = lua.NewState()

	for _, luaType := range luaTypes {
		registerType(sys.lua, luaType)
	}

	src := `
		v = vector3.new(1, 2, 3)
		print(v:xyz())

		v:xyz(3, 2, 1)
		print(v:xyz())
	`

	if err := sys.lua.DoString(src); err != nil {
		fmt.Print(err)
	}
}

func (sys *LuaSystem) Update(duration time.Duration) {
	//for _, eid := range sys.subscribed.unregisteredScripts.GetEntities() {
	//	sys.register(eid)
	//}

	v3, err := vector3.FromLua(sys.lua.GetGlobal("v").(*lua.LUserData))
	if err != nil {
		panic(fmt.Sprintf("failed to retrieve dog from Lua state: %s", err.Error()))
	}

	fmt.Printf("vector %v\n", v3)

	v3.X = 69

	src := `
		print(v:xyz())
	`

	if err := sys.lua.DoString(src); err != nil {
		fmt.Print(err)
	}

	os.Exit(0)
}

func (sys *LuaSystem) register(e akara.EID) {
	sys.bindEventHandlers(e)
}

func (sys *LuaSystem) bindEventHandlers(e akara.EID) {
	panic("implement me")
}

