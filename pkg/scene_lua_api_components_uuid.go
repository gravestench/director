package pkg

import (
	lua "github.com/yuin/gopher-lua"

	"github.com/gravestench/director/pkg/common"
	"github.com/gravestench/director/pkg/components"
)

const (
	luaUUIDComponentName = "uuid"
)

/*
example lua:
	uuid = components.uuid.add(eid)

	val = uuid.value()
	uuid.value(0.5) -- set the uuid to 50%
*/

func (s *SceneSystem) luaExportComponentUUID(mt *lua.LTable) {
	uuidTable := s.Lua.NewTypeMetatable(luaUUIDComponentName)

	s.Lua.SetField(uuidTable, "add", s.Lua.NewFunction(s.luaUUIDAdd()))
	s.Lua.SetField(uuidTable, "get", s.Lua.NewFunction(s.luaUUIDGet()))
	s.Lua.SetField(uuidTable, "remove", s.Lua.NewFunction(s.luaUUIDRemove()))

	s.Lua.SetField(mt, luaUUIDComponentName, uuidTable)
}

func (s *SceneSystem) luaUUIDAdd() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := common.Entity(s.Lua.CheckNumber(1))

		uuid := s.Components.UUID.Add(e)
		L.Push(s.makeLuaTableComponentUUID(uuid))
		return 1
	}

	return fn
}

func (s *SceneSystem) luaUUIDGet() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		id := L.CheckNumber(1)
		uuid, found := s.Components.UUID.Get(common.Entity(id))

		truthy := lua.LFalse
		if !found {
			L.Push(lua.LNil)
			L.Push(truthy)
			return 2
		} else {
			truthy = lua.LTrue
		}

		table := s.makeLuaTableComponentUUID(uuid)

		L.SetMetatable(table, L.GetTypeMetatable(luaUUIDComponentName))

		L.Push(table)
		L.Push(truthy)

		return 2
	}

	return fn
}

func (s *SceneSystem) luaUUIDRemove() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := common.Entity(s.Lua.CheckNumber(1))

		s.Components.UUID.Remove(e)

		return 0
	}

	return fn
}

func (s *SceneSystem) makeLuaTableComponentUUID(uuid *components.UUID) *lua.LTable {
	table := s.Lua.NewTable()

	s.Lua.SetField(table, "string", s.Lua.NewFunction(func(L *lua.LState) int {
		if L.GetTop() != 1 {

			return 0
		}

		L.Push(lua.LString(uuid.String()))

		return 1
	}))

	return table
}
