package scene

import (
	"github.com/gravestench/director/pkg/common"
	"github.com/gravestench/director/pkg/components"
	lua "github.com/yuin/gopher-lua"
)

const (
	luaTextComponentName = "text"
)

/*
example lua:
	text = scene.components.text.add(eid)

	val = text.value()
	text.value(0.5) -- set the text to 50%
*/

func (s *Scene) luaExportComponentText(mt *lua.LTable) {
	textTable := s.Lua.NewTypeMetatable(luaTextComponentName)

	s.Lua.SetField(textTable, "add", s.Lua.NewFunction(s.luaTextAdd()))
	s.Lua.SetField(textTable, "get", s.Lua.NewFunction(s.luaTextGet()))
	s.Lua.SetField(textTable, "remove", s.Lua.NewFunction(s.luaTextRemove()))

	s.Lua.SetField(mt, luaTextComponentName, textTable)
}

func (s *Scene) luaTextAdd() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := common.Entity(s.Lua.CheckNumber(1))

		text := s.Components.Text.Add(e)
		L.Push(s.makeLuaTableComponentText(text))
		return 1
	}

	return fn
}

func (s *Scene) luaTextGet() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		id := L.CheckNumber(1)
		text, found := s.Components.Text.Get(common.Entity(id))

		truthy := lua.LFalse
		if !found {
			L.Push(lua.LNil)
			L.Push(truthy)
			return 2
		} else {
			truthy = lua.LTrue
		}

		table := s.makeLuaTableComponentText(text)

		L.SetMetatable(table, L.GetTypeMetatable(luaTextComponentName))

		L.Push(table)
		L.Push(truthy)

		return 2
	}

	return fn
}

func (s *Scene) luaTextRemove() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := common.Entity(s.Lua.CheckNumber(1))

		s.Components.Text.Remove(e)

		return 0
	}

	return fn
}

func (s *Scene) makeLuaTableComponentText(text *components.Text) *lua.LTable {
	table := s.Lua.NewTable()

	s.Lua.SetField(table, "string", s.Lua.NewFunction(func(L *lua.LState) int {
		if L.GetTop() == 1 {
			val := L.CheckString(1)

			text.String = val

			return 0
		}

		L.Push(lua.LString(text.String))

		return 1
	}))

	return table
}
