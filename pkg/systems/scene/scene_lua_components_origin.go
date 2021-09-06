package scene

import (
	"github.com/gravestench/akara"
	"github.com/gravestench/director/pkg/components"
	lua "github.com/yuin/gopher-lua"
)

const (
	luaOriginComponentName = "origin"
)

func (s *Scene) luaExportComponentOrigin(mt *lua.LTable) {
	originTable := s.Lua.NewTypeMetatable(luaOriginComponentName)

	s.Lua.SetField(originTable, "add", s.Lua.NewFunction(s.luaOriginAdd()))
	s.Lua.SetField(originTable, "get", s.Lua.NewFunction(s.luaOriginGet()))

	s.Lua.SetField(mt, luaOriginComponentName, originTable)
}

func (s *Scene) luaOriginAdd() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		origin := s.Components.Origin.Add(*s.luaCheckEID())
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
		origin, found := s.Components.Origin.Get(akara.EID(id))

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

func (s *Scene) makeLuaTableComponentOrigin(origin *components.Origin) *lua.LTable {
	table := s.Lua.NewTable()

	s.Lua.SetField(table, "xyz", s.makeLuaTableVec3(origin.Vector3))

	return table
}

