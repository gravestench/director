package pkg

import (
	"fmt"
	"os"
	"time"

	lua "github.com/yuin/gopher-lua"

	"github.com/gravestench/akara"
)

// Lua scene API
//
// A lua script that contains an `init` and `update` function can be
// used to create a graphical scene in director. Part of the initialization
// for the lua environment BEFORE the script is loaded includes declaring
// director/scene stuff for use inside of the lua script
//
// TODO: need to elaborate on how director and scene stuff is used inside of lua.
const (
	luaFnInit   = "init"
	luaFnUpdate = "update"
)

// LuaScene is a graphical scene which is created purely from a lua script.
// The Lua script requires an init and update function be declared.
type LuaScene struct {
	Scene
	LuaScriptPath string
	scriptLoaded  bool
	initCalled    bool
}

func NewLuaScene(name, scriptPath string) *LuaScene {
	var scene LuaScene

	scene.key = name
	scene.LuaScriptPath = scriptPath

	return &scene
}

func (scene *LuaScene) Key() string {
	return scene.key
}

func (scene *LuaScene) IsInitialized() bool {
	return scene.scriptLoaded
}

func (scene *LuaScene) Init(_ *akara.World) {
	scene.InitializeLua()
	scene.loadScript()
	scene.callLuaInitFn()
}

func (scene *LuaScene) loadScript() {
	if _, err := os.Stat(scene.LuaScriptPath); err != nil {
		return
	}

	if err := scene.Lua.DoFile(scene.LuaScriptPath); err != nil {
		fmt.Printf("Lua script failed to execute: %s\n", err.Error())
	}

	scene.scriptLoaded = true
}

func (scene *LuaScene) callLuaInitFn() {
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

func (scene *LuaScene) Update() {
	go scene.callLuaUpdateFn(scene.TimeDelta)
}

func (scene *LuaScene) callLuaUpdateFn(dt time.Duration) {
	err := scene.Lua.CallByParam(lua.P{
		Fn:      scene.Lua.GetGlobal(luaFnUpdate),
		NRet:    0,
		Protect: true,
	}, lua.LNumber(dt))

	if err != nil {
		fmt.Println(err)
	}
}
