package scene

import (
	"fmt"
	"github.com/gravestench/akara"
	lua "github.com/yuin/gopher-lua"
	"os"
	"time"
)

const (
	luaFnInit   = "init"
	luaFnUpdate = "update"
)

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
	scene.callLuaUpdateFn(scene.Director.TimeDelta)
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
