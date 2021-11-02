package pkg

import (
	lua "github.com/yuin/gopher-lua"

	"github.com/gravestench/director/pkg/common"
	"github.com/gravestench/director/pkg/components"
)

const (
	luaOpacityComponentName = "opacity"
)

/*
example lua:
	opacity = components.opacity.add(eid)

	val = opacity.value()
	opacity.value(0.5) -- set the opacity to 50%
*/

func (s *Scene) luaExportComponentOpacity(mt *lua.LTable) {
	opacityTable := s.Lua.NewTypeMetatable(luaOpacityComponentName)

	s.Lua.SetField(opacityTable, "add", s.Lua.NewFunction(s.luaOpacityAdd()))
	s.Lua.SetField(opacityTable, "get", s.Lua.NewFunction(s.luaOpacityGet()))
	s.Lua.SetField(opacityTable, "remove", s.Lua.NewFunction(s.luaOpacityRemove()))

	s.Lua.SetField(mt, luaOpacityComponentName, opacityTable)
}

func (s *Scene) luaOpacityAdd() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := common.Entity(s.Lua.CheckNumber(1))

		opacity := s.Components.Opacity.Add(e)
		L.Push(s.makeLuaTableComponentOpacity(opacity))
		return 1
	}

	return fn
}

func (s *Scene) luaOpacityGet() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		id := L.CheckNumber(1)
		opacity, found := s.Components.Opacity.Get(common.Entity(id))

		truthy := lua.LFalse
		if !found {
			L.Push(lua.LNil)
			L.Push(truthy)
			return 2
		} else {
			truthy = lua.LTrue
		}

		table := s.makeLuaTableComponentOpacity(opacity)

		L.SetMetatable(table, L.GetTypeMetatable(luaOpacityComponentName))

		L.Push(table)
		L.Push(truthy)

		return 2
	}

	return fn
}

func (s *Scene) luaOpacityRemove() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := common.Entity(s.Lua.CheckNumber(1))

		s.Components.Opacity.Remove(e)

		return 0
	}

	return fn
}

func (s *Scene) makeLuaTableComponentOpacity(opacity *components.Opacity) *lua.LTable {
	table := s.Lua.NewTable()

	s.Lua.SetField(table, "value", s.Lua.NewFunction(func(L *lua.LState) int {
		if L.GetTop() == 1 {
			val := float64(L.CheckNumber(1))

			opacity.Value = val

			return 0
		}

		L.Push(lua.LNumber(opacity.Value))

		return 1
	}))

	return table
}
