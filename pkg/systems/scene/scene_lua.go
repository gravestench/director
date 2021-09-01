package scene

import (
	"fmt"
	"os"
)

type LuaScene struct {
	Scene
	LuaScriptPath string
	scriptRan     bool
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
	return scene.scriptRan
}

func (scene *LuaScene) Update() {
	if scene.scriptRan {
		return
	}

	if _, err := os.Stat(scene.LuaScriptPath); err != nil {
		return
	}

	if err := scene.Lua.DoFile(scene.LuaScriptPath); err != nil {
		fmt.Printf("Lua script failed to execute: %s\n", err.Error())
		return
	}

	scene.scriptRan = true
}




