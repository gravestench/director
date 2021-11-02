package pkg

import (
	lua "github.com/yuin/gopher-lua"

	"github.com/gravestench/director/pkg/common"
	"github.com/gravestench/director/pkg/components"
)

const (
	luaFontComponentName = "font"
)

/*
example lua:
	font = components.font.add(eid)

	fontFace = font.face()
	font.face("Comic Sans") -- lul

	fontSize = font.size()
	font.size(24) -- pixels
*/

func (s *Scene) luaExportComponentFont(mt *lua.LTable) {
	fontTable := s.Lua.NewTypeMetatable(luaFontComponentName)

	s.Lua.SetField(fontTable, "add", s.Lua.NewFunction(s.luaFontAdd()))
	s.Lua.SetField(fontTable, "get", s.Lua.NewFunction(s.luaFontGet()))
	s.Lua.SetField(fontTable, "remove", s.Lua.NewFunction(s.luaFontRemove()))

	s.Lua.SetField(mt, luaFontComponentName, fontTable)
}

func (s *Scene) luaFontAdd() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := common.Entity(s.Lua.CheckNumber(1))

		font := s.Components.Font.Add(e)
		L.Push(s.makeLuaTableComponentFont(font))
		return 1
	}

	return fn
}

func (s *Scene) luaFontGet() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		id := L.CheckNumber(1)
		font, found := s.Components.Font.Get(common.Entity(id))

		truthy := lua.LFalse
		if !found {
			L.Push(lua.LNil)
			L.Push(truthy)
			return 2
		} else {
			truthy = lua.LTrue
		}

		table := s.makeLuaTableComponentFont(font)

		L.SetMetatable(table, L.GetTypeMetatable(luaFontComponentName))

		L.Push(table)
		L.Push(truthy)

		return 2
	}

	return fn
}

func (s *Scene) luaFontRemove() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := common.Entity(s.Lua.CheckNumber(1))

		s.Components.Font.Remove(e)

		return 0
	}

	return fn
}

func (s *Scene) makeLuaTableComponentFont(font *components.Font) *lua.LTable {
	table := s.Lua.NewTable()

	s.Lua.SetField(table, "face", s.Lua.NewFunction(func(L *lua.LState) int {
		if L.GetTop() == 0 {
			L.Push(lua.LString(font.Face))

			return 1
		}

		face := L.CheckString(1)
		font.Face = face

		return 0
	}))

	s.Lua.SetField(table, "size", s.Lua.NewFunction(func(L *lua.LState) int {
		if L.GetTop() == 0 {
			L.Push(lua.LNumber(font.Size))

			return 1
		}

		size := int(L.CheckNumber(1))
		font.Size = size

		return 0
	}))

	return table
}
