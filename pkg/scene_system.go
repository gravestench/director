package pkg

import (
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/common"
	lua "github.com/yuin/gopher-lua"
)

// SceneSystem is the generic non-graphical stuff for a scene.
// This is kept separate so that non-graphical scenes can be created,
// such as for headless servers. This type of "scene" can still create objects,
// but is much more like a system in that it just inits and then does an update
// every tick.
//
// Example 1: You could have many modal ui menus, and a system that manages state
// for all of these ui's.
//
// Example 2: A scene that spawns objects with velocities, and another system that
// only uses the velocity to update their position every tick.
type SceneSystem struct {
	*Director
	akara.BaseSystem
	Lua        *lua.LState
	Components common.SceneComponents
	Add        ObjectFactory
}

func (s *SceneSystem) IsInitialized() bool {
	return s.Director.World != nil
}

func (s *SceneSystem) InitializeLua() {
	if s.LuaInitialized() {
		return
	}

	s.Lua = lua.NewState()
	if err := s.Lua.DoString(common.LuaLibSTD); err != nil {
		panic(err)
	}

	s.initLuaConstantsTable()
	s.initLuaSceneTable()
}

func (s *SceneSystem) UninitializeLua() {
	s.Lua = nil
}

func (s *SceneSystem) LuaInitialized() bool {
	return s.Lua != nil
}
