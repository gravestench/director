package scene

import (
	"github.com/gravestench/director/pkg/common"
	"github.com/gravestench/director/pkg/components"
	lua "github.com/yuin/gopher-lua"
	"image/color"
)

const (
	luaColorComponentName = "color"
)

/*
example lua:
	color = scene.components.color.add(eid)

	color.rgba(255, 255, 255, 255)
	r, g, b, a = color.rgba()
*/

func (s *Scene) luaExportComponentColor(mt *lua.LTable) {
	colorTable := s.Lua.NewTypeMetatable(luaColorComponentName)

	s.Lua.SetField(colorTable, "add", s.Lua.NewFunction(s.luaColorAdd()))
	s.Lua.SetField(colorTable, "get", s.Lua.NewFunction(s.luaColorGet()))
	s.Lua.SetField(colorTable, "remove", s.Lua.NewFunction(s.luaColorRemove()))

	s.Lua.SetField(mt, luaColorComponentName, colorTable)
}

func (s *Scene) luaColorAdd() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := common.Entity(s.Lua.CheckNumber(1))

		c := s.Components.Color.Add(e)
		L.Push(s.makeLuaTableComponentColor(c))
		return 1
	}

	return fn
}

func (s *Scene) luaColorGet() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		id := L.CheckNumber(1)
		color, found := s.Components.Color.Get(common.Entity(id))

		truthy := lua.LFalse
		if !found {
			L.Push(lua.LNil)
			L.Push(truthy)
			return 2
		} else {
			truthy = lua.LTrue
		}

		table := s.makeLuaTableComponentColor(color)

		L.SetMetatable(table, L.GetTypeMetatable(luaColorComponentName))

		L.Push(table)
		L.Push(truthy)

		return 2
	}

	return fn
}

func (s *Scene) luaColorRemove() lua.LGFunction {
	fn := func(L *lua.LState) int {
		if L.GetTop() != 1 {
			return 0
		}

		e := common.Entity(s.Lua.CheckNumber(1))

		s.Components.Color.Remove(e)

		return 0
	}

	return fn
}

func (s *Scene) makeLuaTableComponentColor(c *components.Color) *lua.LTable {
	table := s.Lua.NewTable()

	s.Lua.SetField(table, "rgba", s.Lua.NewFunction(func(L *lua.LState) int {
		if L.GetTop() == 4 {
			r := uint8(L.CheckNumber(1))
			g := uint8(L.CheckNumber(1))
			b := uint8(L.CheckNumber(1))
			a := uint8(L.CheckNumber(1))

			c.Color = color.RGBA{
				R: r,
				G: g,
				B: b,
				A: a,
			}

			return 0
		}

		r, g, b, a := c.RGBA()

		L.Push(lua.LNumber(r))
		L.Push(lua.LNumber(g))
		L.Push(lua.LNumber(b))
		L.Push(lua.LNumber(a))

		return 4
	}))

	return table
}
