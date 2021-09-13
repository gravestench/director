package scene

import (
	lua "github.com/yuin/gopher-lua"
	"image/color"

	"github.com/gravestench/director/pkg/common"
	"github.com/gravestench/director/pkg/components"
)

const (
	luaFillComponentName = "fill"
)

/*
example lua:
	fill = scene.components.fill.add(eid)

	fill.rgba(255, 255, 255, 255)
	r, g, b, a = fill.rgba()
*/

func (s *Scene) luaExportComponentFill(mt *lua.LTable) {
	fillTable := s.Lua.NewTypeMetatable(luaFillComponentName)

	s.Lua.SetField(fillTable, "add", s.Lua.NewFunction(s.luaFillAdd()))
	s.Lua.SetField(fillTable, "get", s.Lua.NewFunction(s.luaFillGet()))
	s.Lua.SetField(fillTable, "remove", s.Lua.NewFunction(s.luaFillRemove()))

	s.Lua.SetField(mt, luaFillComponentName, fillTable)
}

func (s *Scene) luaFillAdd() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := common.Entity(s.Lua.CheckNumber(1))

		c := s.Components.Fill.Add(e)
		L.Push(s.makeLuaTableComponentFill(c))
		return 1
	}

	return fn
}

func (s *Scene) luaFillGet() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		id := L.CheckNumber(1)
		fill, found := s.Components.Fill.Get(common.Entity(id))

		truthy := lua.LFalse
		if !found {
			L.Push(lua.LNil)
			L.Push(truthy)
			return 2
		} else {
			truthy = lua.LTrue
		}

		table := s.makeLuaTableComponentFill(fill)

		L.SetMetatable(table, L.GetTypeMetatable(luaFillComponentName))

		L.Push(table)
		L.Push(truthy)

		return 2
	}

	return fn
}

func (s *Scene) luaFillRemove() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := common.Entity(s.Lua.CheckNumber(1))

		s.Components.Fill.Remove(e)

		return 0
	}

	return fn
}

func (s *Scene) makeLuaTableComponentFill(fill *components.Fill) *lua.LTable {
	table := s.Lua.NewTable()

	s.Lua.SetField(table, "rgba", s.Lua.NewFunction(func(L *lua.LState) int {
		if L.GetTop() == 4 {
			r := uint8(L.CheckNumber(1))
			g := uint8(L.CheckNumber(1))
			b := uint8(L.CheckNumber(1))
			a := uint8(L.CheckNumber(1))

			fill.Color = color.RGBA{
				R: r,
				G: g,
				B: b,
				A: a,
			}

			return 0
		}

		r, g, b, a := fill.RGBA()

		L.Push(lua.LNumber(r))
		L.Push(lua.LNumber(g))
		L.Push(lua.LNumber(b))
		L.Push(lua.LNumber(a))

		return 4
	}))

	return table
}
