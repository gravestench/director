package pkg

import (
	lua "github.com/yuin/gopher-lua"

	"github.com/gravestench/director/pkg/common"
	"github.com/gravestench/director/pkg/components"
)

const (
	luaSizeComponentName = "size"
)

/*
example lua:
	s = components.size.add(eid)

	w, h = s.size()
	s.size(10, 10) -- the the size, in pixels
*/

func (s *Scene) luaExportComponentSize(mt *lua.LTable) {
	sizeTable := s.Lua.NewTypeMetatable(luaSizeComponentName)

	s.Lua.SetField(sizeTable, "add", s.Lua.NewFunction(s.luaSizeAdd()))
	s.Lua.SetField(sizeTable, "get", s.Lua.NewFunction(s.luaSizeGet()))
	s.Lua.SetField(sizeTable, "remove", s.Lua.NewFunction(s.luaSizeRemove()))

	s.Lua.SetField(mt, luaSizeComponentName, sizeTable)
}

func (s *Scene) luaSizeAdd() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := common.Entity(s.Lua.CheckNumber(1))

		size := s.Components.Size.Add(e)
		L.Push(s.makeLuaTableComponentSize(size))
		return 1
	}

	return fn
}

func (s *Scene) luaSizeGet() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		id := L.CheckNumber(1)
		size, found := s.Components.Size.Get(common.Entity(id))

		truthy := lua.LFalse
		if !found {
			L.Push(lua.LNil)
			L.Push(truthy)
			return 2
		} else {
			truthy = lua.LTrue
		}

		table := s.makeLuaTableComponentSize(size)

		L.SetMetatable(table, L.GetTypeMetatable(luaSizeComponentName))

		L.Push(table)
		L.Push(truthy)

		return 2
	}

	return fn
}

func (s *Scene) luaSizeRemove() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := common.Entity(s.Lua.CheckNumber(1))

		s.Components.Size.Remove(e)

		return 0
	}

	return fn
}

func (s *Scene) makeLuaTableComponentSize(size *components.Size) *lua.LTable {
	table := s.Lua.NewTable()

	s.Lua.SetField(table, "size", s.Lua.NewFunction(func(L *lua.LState) int {
		if L.GetTop() >= 2 {

			return 0
		}

		w, h := size.Max.X, size.Max.Y
		L.Push(lua.LNumber(w))
		L.Push(lua.LNumber(h))

		return 2
	}))

	return table
}
