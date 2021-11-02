package pkg

import (
	lua "github.com/yuin/gopher-lua"

	"github.com/gravestench/director/pkg/common"
	"github.com/gravestench/director/pkg/components"
)

const (
	luaOriginComponentName = "origin"
)

/*
example lua:
	origin = components.origin.add(eid)

	ox, oy = origin.xy()
	anim.frame(0.5, 0.5) -- set the origin to the center of the entity
*/

func (s *Scene) luaExportComponentOrigin(mt *lua.LTable) {
	originTable := s.Lua.NewTypeMetatable(luaOriginComponentName)

	s.Lua.SetField(originTable, "add", s.Lua.NewFunction(s.luaOriginAdd()))
	s.Lua.SetField(originTable, "get", s.Lua.NewFunction(s.luaOriginGet()))
	s.Lua.SetField(originTable, "remove", s.Lua.NewFunction(s.luaOriginRemove()))

	s.Lua.SetField(mt, luaOriginComponentName, originTable)
}

func (s *Scene) luaOriginAdd() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := common.Entity(s.Lua.CheckNumber(1))

		origin := s.Components.Origin.Add(e)
		L.Push(s.makeLuaTableComponentOrigin(origin))
		return 1
	}

	return fn
}

func (s *Scene) luaOriginGet() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		id := L.CheckNumber(1)
		origin, found := s.Components.Origin.Get(common.Entity(id))

		truthy := lua.LFalse
		if !found {
			L.Push(lua.LNil)
			L.Push(truthy)
			return 2
		} else {
			truthy = lua.LTrue
		}

		table := s.makeLuaTableComponentOrigin(origin)

		L.SetMetatable(table, L.GetTypeMetatable(luaOriginComponentName))

		L.Push(table)
		L.Push(truthy)

		return 2
	}

	return fn
}

func (s *Scene) luaOriginRemove() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := common.Entity(s.Lua.CheckNumber(1))

		s.Components.Origin.Remove(e)

		return 0
	}

	return fn
}

func (s *Scene) makeLuaTableComponentOrigin(origin *components.Origin) *lua.LTable {
	table := s.Lua.NewTable()

	s.Lua.SetField(table, "xyz", s.makeLuaSetterGetterVec3(origin.Vector3))

	return table
}
