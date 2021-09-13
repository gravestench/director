package scene

import (
	lua "github.com/yuin/gopher-lua"

	"github.com/gravestench/director/pkg/common"
)

const (
	luaDebugComponentName = "debug"
)

/*
example lua:
	scene.components.debug.add(eid)

	-- this is a tag component, so has no other functionality

	-- it is used merely to flag something for debug, systems/scenes
	-- can check if an entity has this debug flag
*/

func (s *Scene) luaExportComponentDebug(mt *lua.LTable) {
	debugTable := s.Lua.NewTypeMetatable(luaDebugComponentName)

	s.Lua.SetField(debugTable, "add", s.Lua.NewFunction(s.luaDebugAdd()))
	s.Lua.SetField(debugTable, "get", s.Lua.NewFunction(s.luaDebugGet()))
	s.Lua.SetField(debugTable, "remove", s.Lua.NewFunction(s.luaDebugRemove()))

	s.Lua.SetField(mt, luaDebugComponentName, debugTable)
}

func (s *Scene) luaDebugAdd() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := common.Entity(s.Lua.CheckNumber(1))

		s.Components.Debug.Add(e)

		L.Push(s.Lua.NewTable()) // intentionally blank, this is a tag component

		return 1
	}

	return fn
}

func (s *Scene) luaDebugGet() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		id := L.CheckNumber(1)
		_, found := s.Components.Debug.Get(common.Entity(id))

		truthy := lua.LFalse
		if !found {
			L.Push(lua.LNil)
			L.Push(truthy)
			return 2
		} else {
			truthy = lua.LTrue
		}

		// intentionally left blank
		// this component is just used for tagging an entity for debug purposes
		table := s.Lua.NewTable()

		L.SetMetatable(table, L.GetTypeMetatable(luaDebugComponentName))

		L.Push(table)
		L.Push(truthy)

		return 2
	}

	return fn
}

func (s *Scene) luaDebugRemove() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := common.Entity(s.Lua.CheckNumber(1))

		s.Components.Debug.Remove(e)

		return 0
	}

	return fn
}
