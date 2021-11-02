package pkg

import (
	"fmt"
	"os"
	"time"

	lua "github.com/yuin/gopher-lua"

	"github.com/gravestench/akara"
)

// LuaSystem is a non-graphical system which is created purely from a lua script.
// The Lua script requires an init and update function be declared.
type LuaSystem struct {
	Base
	LuaScriptPath string
	scriptLoaded  bool
	initCalled    bool
}

func NewLuaSystem(name, scriptPath string) *LuaSystem {
	var scene LuaSystem

	scene.LuaScriptPath = scriptPath

	return &scene
}

func (scene *LuaSystem) IsInitialized() bool {
	return scene.scriptLoaded
}

func (scene *LuaSystem) Init(_ *akara.World) {
	scene.loadScript()
	scene.callLuaInitFn()
}

func (scene *LuaSystem) loadScript() {
	if _, err := os.Stat(scene.LuaScriptPath); err != nil {
		return
	}

	if err := scene.Lua.DoFile(scene.LuaScriptPath); err != nil {
		fmt.Printf("Lua script failed to execute: %s\n", err.Error())
	}

	scene.scriptLoaded = true
}

func (scene *LuaSystem) callLuaInitFn() {
	err := scene.Lua.CallByParam(lua.P{
		Fn:      scene.Lua.GetGlobal(luaFnInit),
		NRet:    0,
		Protect: true,
	})

	if err != nil {
		return
	}

	scene.initCalled = true
}

func (scene *LuaSystem) Update() {
	scene.callLuaUpdateFn(scene.TimeDelta)
}

func (scene *LuaSystem) callLuaUpdateFn(dt time.Duration) {
	err := scene.Lua.CallByParam(lua.P{
		Fn:      scene.Lua.GetGlobal(luaFnUpdate),
		NRet:    0,
		Protect: true,
	}, lua.LNumber(dt))

	if err != nil {
		fmt.Println(err)
	}
}

// a static check to verify that a `LuaSystem` implements the `akara.System` interface
var _ akara.System = &LuaSystem{}
