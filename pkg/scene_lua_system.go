package pkg

import (
	"fmt"
	"os"

	lua "github.com/yuin/gopher-lua"

	"github.com/gravestench/akara"
)

// LuaSystem is a non-graphical system which is created purely from a lua script.
// The Lua script requires an init and update function be declared.
type LuaSystem struct {
	SceneSystem
	LuaScriptPath string
	scriptLoaded  bool
	initCalled    bool
}

func NewLuaSystem(scriptPath string) *LuaSystem {
	var sys LuaSystem

	sys.LuaScriptPath = scriptPath

	return &sys
}

func (sys *LuaSystem) IsInitialized() bool {
	return sys.scriptLoaded
}

func (sys *LuaSystem) Init(_ *akara.World) {
	sys.InitializeLua()
	sys.loadScript()
	sys.callLuaInitFn()
}

func (sys *LuaSystem) loadScript() {
	if _, err := os.Stat(sys.LuaScriptPath); err != nil {
		return
	}

	if err := sys.Lua.DoFile(sys.LuaScriptPath); err != nil {
		fmt.Printf("Lua script failed to execute: %s\n", err.Error())
	}

	sys.scriptLoaded = true
}

func (sys *LuaSystem) callLuaInitFn() {
	err := sys.Lua.CallByParam(lua.P{
		Fn:      sys.Lua.GetGlobal(luaFnInit),
		NRet:    0,
		Protect: true,
	})

	if err != nil {
		return
	}

	sys.initCalled = true
}

func (sys *LuaSystem) Update() {
	go sys.callLuaUpdateFn()
}

func (sys *LuaSystem) callLuaUpdateFn() {
	err := sys.Lua.CallByParam(lua.P{
		Fn:      sys.Lua.GetGlobal(luaFnUpdate),
		NRet:    0,
		Protect: true,
	}, lua.LNumber(sys.TimeDelta))

	if err != nil {
		fmt.Println(err)
	}
}

// a static check to verify that a `LuaSystem` implements the `akara.System` interface
var _ akara.System = &LuaSystem{}
