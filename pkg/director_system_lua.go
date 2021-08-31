package pkg

import (
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/common"
	lua "github.com/yuin/gopher-lua"
)

type luaSystem struct {
	director *Director
	akara.BaseSystem
	bound map[string]*lua.LState
	//subscribed struct {
	//	unregisteredScripts *akara.Subscription
	//	apiEventHandlers *akara.Subscription
	//}
}

func (system *luaSystem) IsInitialized() bool {
	return true
}

func (system *luaSystem) Init(world *akara.World) {

}

func (system *luaSystem) registerLuaTypes(scene *Scene) {
	luaTypes := []common.LuaTypeExport{
		luaRectangleTypeExport(scene),
	}

	for _, luaType := range luaTypes {
		registerType(scene.Lua, luaType)
	}
}

func (system *luaSystem) Update() {
	for _, s := range system.director.scenes {
		scene := s.concrete()
		if scene.Lua == nil {
			scene.Lua = lua.NewState()
			system.registerLuaTypes(scene)
		}
	}

	//for _, eid := range system.subscribed.unregisteredScripts.GetEntities() {
	//	system.register(eid)
	//}

	//v3, err := vector3.FromLua(system.lua.GetGlobal("v").(*lua.LUserData))
	//if err != nil {
	//	panic(fmt.Sprintf("failed to retrieve dog from Lua state: %s", err.Error()))
	//}
	//
	//fmt.Printf("vector %v\n", v3)
	//
	//v3.X = 69
	//
	//src := `
	//	print(v:xyz())
	//`
	//
	//if err := system.lua.DoString(src); err != nil {
	//	fmt.Print(err)
	//}
	//
	//os.Exit(0)
}

func (system *luaSystem) register(e akara.EID) {
	system.bindEventHandlers(e)
}

func (system *luaSystem) bindEventHandlers(e akara.EID) {
	//panic("implement me")
}